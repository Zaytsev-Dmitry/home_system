FROM golang:1.24.0 AS builder
WORKDIR /app

# Копируем только необходимые файлы
COPY /cmd/main.go .
COPY go.mod .
RUN go mod download

# Копируем весь проект
COPY . .

# Убеждаемся, что папка `config/` есть
RUN ls -la /app/config/

# Убеждаемся, что папка `migrations/` есть
RUN ls -la /app/migrations/

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o /migrator main.go

# Минимизируем финальный образ
FROM scratch
COPY --from=builder /migrator /bin/migrator
COPY --from=builder /app/config /app/config
COPY --from=builder /app/migrations /migrator/migrations

# Объявляем аргументы
ARG CONFIG_PATH_ARG
ARG APP_PROFILE_ARG

# Проверяем, что аргументы переданы (иначе билд прервётся)
RUN test -n "$CONFIG_PATH_ARG" || (echo "CONFIG_PATH_ARG is empty!")
RUN test -n "$APP_PROFILE_ARG" || (echo "APP_PROFILE_ARG is empty!")

# Устанавливаем переменные среды
ENV CONFIG_PATH=$CONFIG_PATH_ARG
ENV APP_PROFILE=$APP_PROFILE_ARG

ENTRYPOINT ["/bin/migrator"]
