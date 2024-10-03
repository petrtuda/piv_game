let interval;
let isModalClosed = false; // Переменная для отслеживания состояния модалки
let coefficient; // Коэффициент для расчета суммы


// Firebase Realtime Database
const database = firebase.database();

// Обработчик нажатия кнопки Start
document.getElementById("start-button").addEventListener("click", function() {
    // Создаем уникальный ID для запуска игры
    const gameId = database.ref().child('games').push().key;

    // Получаем текущее время
    const startTime = new Date().toISOString();

    // Получаем имя пользователя и сумму пива (пример)
    const username = document.getElementById("username-input").value;
    const wonAmount = calculateBeerAmount(); // Функция расчета суммы пива

    // Записываем данные в Firebase Realtime Database
    database.ref('games/' + gameId).set({
        gameId: gameId,
        startTime: startTime,
        username: username,
        wonAmount: wonAmount
    }).then(function() {
        console.log("Game data successfully saved!");
    }).catch(function(error) {
        console.error("Error saving game data: ", error);
    });
});

// Функция для записи данных о запуске игры
function writeGameData(gameId, username, beerAmount) {
  const startTime = new Date().toISOString(); // Получаем текущее время

  // Создаем запись в базе данных
  firebase.database().ref('games/' + gameId).set({
    username: username,
    startTime: startTime,
    beerAmount: beerAmount
  }, (error) => {
    if (error) {
      console.error("Ошибка при записи данных: ", error);
    } else {
      console.log("Данные успешно записаны в базу");
    }
  });
}

// Пример вызова функции
writeGameData('game123', 'username123', 10);

function calculateBeerAmount() {
    // Получаем значение, выбранное пользователем в dropdown
    const dropdownValue = document.getElementById("beer-amount-dropdown").value;

    // Задаем коэффициент для расчета суммы пива
    const coefficient = 1.5;  // Например, 1.5 литра за каждый выбор

    // Рассчитываем количество пива
    const wonAmount = dropdownValue * coefficient;

    return wonAmount;
}

// Запуск анимации прогресс бара
function startProgress() {
    const progressElement = document.getElementById('progress');
    let width = 0;

    clearInterval(interval);

    interval = setInterval(() => {
        if (width >= 100) {
            clearInterval(interval);
            // Рассчитываем сумму
            const beerSelect = document.getElementById('beerSelect');
            const selectedValue = parseInt(beerSelect.value);
            coefficient = Math.floor(Math.random() * 11); // Генерация коэффициента от 0 до 10
            const totalAmount = selectedValue * coefficient;

            // Выводим результат на экран
            displayResult(totalAmount);
        } else {
            width++;
            progressElement.style.width = width + '%';
        }
    }, 50);
}

document.querySelector('#startButton').addEventListener('click', function() {
    const gameId = 'game_' + Date.now();  // Создаем уникальный ID игры
    const username = 'Пользователь123';   // Можно заменить на сгенерированный username
    const beerAmount = 5;  // Сумма пива
  
    writeGameData(gameId, username, beerAmount);  // Записываем данные в Firebase
  });
// Функция для отображения результата
function displayResult(amount) {
    const resultDisplay = document.createElement('p');
    resultDisplay.textContent = `Сумма: ${amount} 💰 Yo! 🎉`; // Отображаем сумму и текст
    document.body.appendChild(resultDisplay);
}

// Обновление количества пива
function updateBeerCount() {
    const beerSelect = document.getElementById('beerSelect');
    const selectedValue = beerSelect.value;
    document.getElementById('beerCount').textContent = selectedValue;
}

// Показ модалки
function showModal() {
    if (!isModalClosed) { // Проверяем состояние модалки
        const modal = document.getElementById('authModal');
        modal.classList.add('show'); // Добавляем класс для анимации
        modal.style.display = 'block'; // Показать модальное окно
        document.body.classList.add('modal-open'); // Добавить класс, чтобы скрыть кнопку Start
    }
}

function closeModal() {
    const modal = document.getElementById('authModal');
    modal.classList.remove('show'); // Убираем класс для анимации

    // Удаляем стиль отображения после завершения анимации
    setTimeout(() => {
        modal.style.display = 'none'; // Закрываем модалку
    }, 300); // Задержка должна совпадать с длительностью анимации

    document.body.classList.remove('modal-open'); // Показать кнопку Start
}

// Закрытие модалки
function closeModal() {
    const modal = document.getElementById('authModal');
    modal.classList.remove('show'); // Убираем класс для анимации

    // Удаляем стиль отображения после завершения анимации
    setTimeout(() => {
        modal.style.display = 'none'; // Закрываем модалку
    }, 300); // Задержка должна совпадать с длительностью анимации

    document.body.classList.remove('modal-open'); // Показать кнопку Start
}
    // Убираем старые подсветки
    username.style.border = '';
    password.style.border = '';

    // Проверяем заполнены ли поля
    if (username.value.trim() === '' || password.value.trim() === '') {
        // Подсвечиваем пустые поля
        if (username.value.trim() === '') {
            username.style.border = '2px solid red'; // Подсветка поля
        }
        if (password.value.trim() === '') {
            password.style.border = '2px solid red'; // Подсветка поля
        }

        // Выводим уведомление
        alert("Введите данные");
        return; // Не закрываем модалку, если есть ошибки
    }

    // Если оба поля заполнены, закрываем модалку и выводим сообщение об успехе
    alert("Успех!");
    isModalClosed = true; // Устанавливаем состояние, что модалка успешно закрыта
    document.getElementById('authModal').style.display = 'none'; // Закрываем модал
    document.body.classList.remove('modal-open'); // Показать кнопку Start
}

// Генерация юзернейма
function generateUsername() {
    const username = 'Petr' + Math.floor(Math.random() * 10000);
    document.getElementById('username').value = username;
}

// Генерация пароля
function generatePassword() {
    const password = 'Hillo' + Math.floor(Math.random() * 10000);
    document.getElementById('password').value = password;
}

// Закрытие модального окна при клике вне его
window.onclick = function(event) {
    const modal = document.getElementById('authModal');
    if (event.target === modal) {
        closeModal();
    }
};

// Генерация юзернейма при клике на кнопку
document.getElementById('generateUsernameButton').addEventListener('click', generateUsername);

// Добавляем обработчик событий для кнопки "Start"
document.querySelector('.play-button').addEventListener('click', startProgress);