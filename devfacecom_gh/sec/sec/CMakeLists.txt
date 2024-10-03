cmake_minimum_required(VERSION 3.10)
project(sec)

set(CMAKE_CXX_STANDARD 20)

# Указание путей к OpenSSL
include_directories(/opt/homebrew/opt/openssl@3/include)
link_directories(/opt/homebrew/opt/openssl@3/lib)

# Указание путей к libcurl
find_package(CURL REQUIRED)
include_directories(${CURL_INCLUDE_DIR})

add_executable(sec main.cpp)

# Линковка библиотек OpenSSL и libcurl
target_link_libraries(sec ssl crypto ${CURL_LIBRARIES})