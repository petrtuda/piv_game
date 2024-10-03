#include <iostream>
#include <curl/curl.h>  // Для выполнения HTTP-запросов
#include <string>

// Функция для отправки данных в Firebase
void sendDataToFirebase(const std::string& jsonData) {
    CURL* curl;
    CURLcode res;

    // URL твоей базы данных (не забудь изменить на свой URL)
    std::string firebaseUrl = "https://psec-bb70d.firebaseio.com/ports.json";

    curl = curl_easy_init();
    if (curl) {
        // Установка URL для запроса
        curl_easy_setopt(curl, CURLOPT_URL, firebaseUrl.c_str());
        // Указание, что мы отправляем POST-запрос
        curl_easy_setopt(curl, CURLOPT_POST, 1L);
        // Установка данных, которые будут отправлены (в формате JSON)
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, jsonData.c_str());

        // Выполнение запроса
        res = curl_easy_perform(curl);

        // Проверка на ошибки
        if (res != CURLE_OK) {
            std::cerr << "Ошибка отправки данных в Firebase: " << curl_easy_strerror(res) << std::endl;
        }

        // Освобождение ресурсов
        curl_easy_cleanup(curl);
    }
}

// Основная функция для проверки портов и отправки данных
void checkOpenPortsAndSend() {
    FILE* pipe = popen("ss -tuln", "r");
    if (!pipe) {
        std::cerr << "Ошибка открытия pipe!" << std::endl;
        return;
    }

    char buffer[128];
    std::string result = "";

    // Чтение вывода команды ss
    while (fgets(buffer, sizeof(buffer), pipe) != nullptr) {
        result += buffer;
    }

    // Закрытие pipe
    int status = pclose(pipe);
    if (status == -1) {
        std::cerr << "Ошибка закрытия pipe!" << std::endl;
    }

    // Создание JSON-данных для отправки в Firebase
    std::string jsonData = "{\"open_ports\": \"" + result + "\"}";

    // Отправка данных в Firebase
    sendDataToFirebase(jsonData);
}

int main() {
    std::cout << "Проверка открытых портов и отправка данных в Firebase...\n";
    checkOpenPortsAndSend();
    return 0;
}