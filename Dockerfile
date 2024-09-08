# Этап сборки
FROM golang:latest AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем все файлы в контейнер
COPY . .

# Загружаем зависимости и компилируем приложение
RUN go mod download
RUN go build -o myapp

# Финальный этап
FROM alpine:latest

# Устанавливаем необходимые зависимости
RUN apk add --no-cache ca-certificates

WORKDIR /

# Копируем скомпилированное приложение из фазы сборки
COPY --from=builder /app/myapp /myapp

# Указываем команду для запуска приложения
CMD ["myapp"]
