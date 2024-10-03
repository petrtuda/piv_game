#include <iostream>
#include <cstdlib>   // Для system()
#include <cstdio>    // Для popen(), pclose()
#include <string>

void checkOpenPorts(const std::string& host) {
    std::string command = "nmap -p- " + host;
    FILE* pipe = popen(command.c_str(), "r");
    if (!pipe) {
        std::cerr << "Ошибка открытия pipe!" << std::endl;
        return;
    }

    char buffer[128];
    std::string result = "";

    while (fgets(buffer, sizeof(buffer), pipe) != nullptr) {
        result += buffer;
    }

    int status = pclose(pipe);
    if (status == -1) {
        std::cerr << "Ошибка закрытия pipe!" << std::endl;
    }

    std::cout << "Открытые порты для " << host << ":\n" << result << std::endl;
}

void runTests() {
    // Тест 1: Проверка выполнения команды nmap
    std::string testHost = "localhost";  // Используем локальный хост для тестирования
    std::cout << "Запуск теста для " << testHost << "...\n";

    FILE* pipe = popen(("nmap -p- " + testHost).c_str(), "r");
    if (!pipe) {
        std::cerr << "Ошибка выполнения команды nmap для теста!" << std::endl;
        return;
    }

    char buffer[128];
    bool success = false;

    // Чтение вывода команды nmap
    while (fgets(buffer, sizeof(buffer), pipe) != nullptr) {
        std::cout << buffer; // Выводим результат теста
        success = true;  // Если мы получили вывод, тест считается успешным
    }

    int status = pclose(pipe);
    if (status == -1) {
        std::cerr << "Ошибка закрытия pipe для теста!" << std::endl;
    }

    if (success) {
        std::cout << "Тест для " << testHost << " выполнен успешно.\n";
    } else {
        std::cout << "Тест для " << testHost << " завершился неудачей.\n";
    }
}

int main() {
    std::string host = "linkedin.com"; // Задаем хост
    std::cout << "Проверка открытых портов для " << host << "...\n";
    checkOpenPorts(host);
    
    // Запуск тестов
    std::cout << "\n--- Запуск тестов ---\n";
    runTests();

    return 0;
}