FROM golang:1.24.0 as builder
WORKDIR /build
COPY /cmd/main.go .
COPY go.mod .
RUN go mod download
RUN go mod tidy
COPY . .

RUN go build -o migrator main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go
EXPOSE 8083

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder main /bin/main

ENTRYPOINT ["/bin/main"]
