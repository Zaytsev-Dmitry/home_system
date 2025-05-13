FROM golang:1.24.0 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 && \
    oapi-codegen -package=http -generate "types,spec,gin" api/http/user-service-api.yml > api/http/user-service-api.gen.go && \
    go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /main /bin/main
COPY --from=builder /app/configs /app/configs
COPY --from=builder /app/third_party/swagger_ui /app/third_party/swagger_ui
COPY --from=builder /app/api/http/user-service-api.yml /app/api/http/user-service-api.yml
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ARG APP_PROFILE
ENV CONFIG_PATH=/app/configs/config-${APP_PROFILE}.yaml

EXPOSE 8080

ENTRYPOINT ["/bin/main"]
