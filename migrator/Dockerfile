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

ENTRYPOINT ["/bin/migrator"]
