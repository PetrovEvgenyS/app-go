# Генератор случайных цитат

Этот проект представляет собой веб-приложение на Go, которое отображает случайные цитаты. Приложение использует HTML-шаблоны для отображения данных и предоставляет возможность обновлять цитаты с помощью кнопки.

## Структура проекта

- **`main.go`**: Основной файл приложения, содержащий обработчики HTTP-запросов и логику генерации случайных цитат.
- **`templates/index.html`**: HTML-шаблон для отображения страницы с цитатой.
- **`static/styles.css`**: CSS-файл для стилизации веб-страницы.
- **`go.mod`**: Файл, определяющий модуль Go и его зависимости.

## Установка и запуск

1. Запустите приложение:
   ```bash
   go run main.go
   ```
2. Откройте браузер и перейдите по адресу [http://localhost:8080](http://localhost:8080).

## Функциональность

- Отображение случайной цитаты.
- Кнопка для генерации новой цитаты.
- Отображение версии приложения и IP-адреса сервера.
