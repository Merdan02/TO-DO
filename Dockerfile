FROM golang:1.23.0

WORKDIR /app

# Копируем весь проект, включая vendor
COPY . .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["go", "run", "cmd/main.go"]
