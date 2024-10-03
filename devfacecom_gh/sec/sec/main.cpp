#include <iostream>
#include <curl/curl.h>
#include <openssl/ssl.h>

void checkSSL(const char* url) {
    CURL* curl = curl_easy_init();
    if (curl) {
        curl_easy_setopt(curl, CURLOPT_URL, url);
        curl_easy_setopt(curl, CURLOPT_USE_SSL, CURLUSESSL_ALL);
        CURLcode res = curl_easy_perform(curl);
        if (res != CURLE_OK) {
            std::cerr << "Error: " << curl_easy_strerror(res) << std::endl;
        } else {
            long certinfo;
            res = curl_easy_getinfo(curl, CURLINFO_SSL_ENGINES, &certinfo);
            if (res == CURLE_OK) {
                std::cout << "SSL Certificate is valid.\n";
            } else {
                std::cerr << "SSL Certificate check failed.\n";
            }
        }
        curl_easy_cleanup(curl);
    } else {
        std::cerr << "Curl initialization failed." << std::endl;
    }
}

void getHeaders(const char* url) {
    CURL* curl = curl_easy_init();
    if (curl) {
        curl_easy_setopt(curl, CURLOPT_URL, url);
        curl_easy_setopt(curl, CURLOPT_NOBODY, 1L); // Only headers
        curl_easy_setopt(curl, CURLOPT_HEADER, 1L); // Include headers
        CURLcode res = curl_easy_perform(curl);
        if (res != CURLE_OK) {
            std::cerr << "Error: " << curl_easy_strerror(res) << std::endl;
        } else {
            std::cout << "Headers retrieved successfully.\n";
        }
        curl_easy_cleanup(curl);
    }
}

int main() {
    const char* url = "https://hh.ru";

    std::cout << "Checking SSL Certificate...\n";
    checkSSL(url);

    std::cout << "\nFetching headers...\n";
    getHeaders(url);

    return 0;
}