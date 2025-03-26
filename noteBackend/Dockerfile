FROM golang:1.24.0 as builder
WORKDIR /app
COPY /cmd/main.go .
COPY go.mod .
RUN go mod download
RUN go mod tidy
COPY . .

# Убеждаемся, что папка `config/` есть
RUN ls -la /app/config/

RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
RUN oapi-codegen -package=generatedApi -generate "types,spec,gin" /app/api/spec/openapi/note-api.yml > api/spec/openapi/note-api.gen.go

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go

ENV GIN_MODE=release

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder main /bin/main

# Объявляем аргументы
ARG CONFIG_PATH_ARG
ARG APP_PROFILE_ARG

# Проверяем, что аргументы переданы (иначе билд прервётся)
RUN test -n "$CONFIG_PATH_ARG" || (echo "CONFIG_PATH_ARG is empty!")
RUN test -n "$APP_PROFILE_ARG" || (echo "APP_PROFILE_ARG is empty!")

# Устанавливаем переменные среды
ENV CONFIG_PATH=$CONFIG_PATH_ARG
ENV APP_PROFILE=$APP_PROFILE_ARG

ENTRYPOINT ["/bin/main"]