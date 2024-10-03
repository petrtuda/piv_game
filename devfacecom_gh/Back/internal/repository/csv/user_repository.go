// File: internal/repository/csv/user_repository.go

package csv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myapp/internal/domain"
	"os"
	"sync"
)

type UserRepository struct {
	filename string
	mutex    sync.RWMutex
}

func NewUserRepository(filename string) *UserRepository {
	return &UserRepository{
		filename: filename,
	}
}

func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users, err := r.readUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, domain.ErrUserNotFound
}

func (r *UserRepository) Save(user *domain.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	users, err := r.readUsers()
	if err != nil {
		return err
	}

	for i, u := range users {
		if u.Username == user.Username {
			users[i] = *user
			return r.writeUsers(users)
		}
	}

	users = append(users, *user)
	return r.writeUsers(users)
}

func (r *UserRepository) readUsers() ([]domain.User, error) {
	data, err := ioutil.ReadFile(r.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []domain.User{}, nil
		}
		return nil, err
	}

	var users []domain.User
	err = json.Unmarshal(data, &users)
	return users, err
}

func (r *UserRepository) writeUsers(users []domain.User) error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.filename, data, 0644)
}
