FROM golang:1.23.5-alpine as builder
WORKDIR /app
COPY /cmd/main.go .
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
RUN oapi-codegen -package=generatedApi -generate "types,spec,gin" /app/api/spec/auth-server-api.yml > api/spec/auth-server-api.gen.go
RUN go build -o /app/cmd/main.go
RUN go build -o /main main.go
EXPOSE 8081
ENV GIN_MODE=release

# Финальный этап, копируем собранное приложение
FROM alpine:latest
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]
