FROM golang:1.23.5 as builder
WORKDIR /build
COPY /cmd/main.go .
COPY go.mod .
RUN go mod download
RUN go mod tidy
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go
EXPOSE 8082
ENV GIN_MODE=release

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]
