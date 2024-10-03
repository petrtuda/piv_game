#!/bin/bash

# Вывод текущей рабочей директории
echo "Текущая рабочая директория: $(pwd)"

# Проверяем существование папки Sec_Back
if [ ! -d "Sec_Back" ]; then
    echo "Папка Sec_Back не существует."
else
    echo "Папка Sec_Back существует."
fi