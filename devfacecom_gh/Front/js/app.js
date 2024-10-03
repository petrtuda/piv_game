// Импортируем необходимые модули
const express = require('express'); // Подключаем Express для создания сервера
const firebase = require('firebase'); // Подключаем Firebase

// Инициализация приложения Express
const app = express();

// Настройки Firebase
const firebaseConfig = {
    apiKey: process.env.FIREBASE_API_KEY,
    authDomain: process.env.FIREBASE_AUTH_DOMAIN,
    projectId: process.env.FIREBASE_PROJECT_ID,
    storageBucket: process.env.FIREBASE_STORAGE_BUCKET,
    messagingSenderId: process.env.FIREBASE_MESSAGING_SENDER_ID,
    appId: process.env.FIREBASE_APP_ID
};

// Инициализация Firebase
firebase.initializeApp(firebaseConfig);
const db = firebase.firestore();

// Middleware для обработки данных из форм
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Роут для главной страницы
app.get('/', (req, res) => {
    res.send('Привет, мир! Добро пожаловать на мой сайт.');
});

// Роут для получения всех пользователей из Firestore
app.get('/users', async (req, res) => {
    try {
        const usersSnapshot = await db.collection('users').get();
        const usersList = usersSnapshot.docs.map(doc => doc.data());
        res.json(usersList);
    } catch (error) {
        res.status(500).send('Ошибка при получении данных пользователей.');
    }
});

// Роут для добавления нового пользователя в Firestore
app.post('/add-user', async (req, res) => {
    const { username, password } = req.body;
    if (!username || !password) {
        return res.status(400).send('Необходимо указать username и password');
    }

    try {
        await db.collection('users').add({
            username,
            password,
            createdAt: firebase.firestore.FieldValue.serverTimestamp()
        });
        res.status(201).send('Пользователь успешно добавлен');
    } catch (error) {
        res.status(500).send('Ошибка при добавлении пользователя');
    }
});

// Настройка порта и запуск сервера
const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Сервер запущен на порту ${PORT}`);
});