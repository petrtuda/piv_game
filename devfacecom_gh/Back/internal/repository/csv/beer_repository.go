// File: internal/repository/csv/beer_repository.go

package csv

import (
	"encoding/csv"
	"fmt"
	"myapp/internal/domain"
	"os"
	"strconv"
	"sync"
	"time"
)

type BeerRepository struct {
	depositFilename string
	balanceFilename string
	mutex           sync.RWMutex
}

func NewBeerRepository(depositFilename, balanceFilename string) *BeerRepository {
	return &BeerRepository{
		depositFilename: depositFilename,
		balanceFilename: balanceFilename,
	}
}

func (r *BeerRepository) SaveDeposit(deposit *domain.BeerDeposit) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	file, err := os.OpenFile(r.depositFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		deposit.Username,
		"dep",
		fmt.Sprintf("%.2f", deposit.Amount),
		"beer",
		deposit.Time.Format(time.RFC3339),
	}

	if err := writer.Write(record); err != nil {
		return err
	}

	return r.updateBalance(deposit.Username, deposit.Amount)
}

func (r *BeerRepository) GetBalance(username string) (float64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	file, err := os.OpenFile(r.balanceFilename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	for _, record := range records {
		if record[0] == username {
			return strconv.ParseFloat(record[1], 64)
		}
	}

	return 0, nil
}

func (r *BeerRepository) updateBalance(username string, amount float64) error {
	balances, err := r.readBalances()
	if err != nil {
		return err
	}

	updated := false
	for i, balance := range balances {
		if balance[0] == username {
			currentBalance, _ := strconv.ParseFloat(balance[1], 64)
			balances[i][1] = fmt.Sprintf("%.2f", currentBalance+amount)
			updated = true
			break
		}
	}

	if !updated {
		balances = append(balances, []string{username, fmt.Sprintf("%.2f", amount)})
	}

	return r.writeBalances(balances)
}

func (r *BeerRepository) readBalances() ([][]string, error) {
	file, err := os.OpenFile(r.balanceFilename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func (r *BeerRepository) writeBalances(balances [][]string) error {
	file, err := os.Create(r.balanceFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(balances)
}