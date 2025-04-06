FROM golang:1.24.0 AS builder

WORKDIR /migrator

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN ls -la /migrator/configs/
RUN ls -la /migrator/migrations/


RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM scratch

WORKDIR /migrator

COPY --from=builder /main /bin/main
COPY --from=builder /migrator/configs /migrator/configs
COPY --from=builder /migrator/migrations /migrator/migrations

ARG APP_PROFILE_ARG
ENV CONFIG_FILE=/app/configs/config-${APP_PROFILE_ARG}.yaml

EXPOSE 8085

ENTRYPOINT ["/bin/main"]

