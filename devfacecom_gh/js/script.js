let interval;
let isModalClosed = false; // Переменная для отслеживания состояния модалки (авторизован/не авторизован)
let coefficient; // Коэффициент для расчета суммы

// Запуск анимации прогресс бара
function startProgress() {
    if (!isModalClosed) {
        showModal(); // Открываем модалку, если пользователь не авторизован
        return;
    }

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
    if (!isModalClosed) { // Проверяем состояние авторизации
        document.getElementById('authModal').style.display = 'block'; // Показать модальное окно
        document.body.classList.add('modal-open'); // Добавить класс, чтобы скрыть кнопку Start
    }
}

// Закрытие модалки
function closeModal() {
    const username = document.getElementById('username');
    const password = document.getElementById('password');

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
    document.body.classList.remove('modal-open'); // Убираем блокировку кнопки Start
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

// Добавляем обработчик для наведения на кнопку "Start"
document.querySelector('.play-button').addEventListener('mouseover', () => {
    if (!isModalClosed) { // Если пользователь не авторизован
        showModal(); // Показать модальное окно при наведении
    }
});

// Добавляем обработчик событий для кнопки "Start"
document.querySelector('.play-button').addEventListener('click', startProgress);