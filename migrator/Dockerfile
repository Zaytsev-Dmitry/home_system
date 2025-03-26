FROM golang:1.24.0 AS builder
WORKDIR /build

COPY /cmd/main.go .
COPY go.mod .
RUN go mod download
RUN CGO_ENABLED=0 go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go

FROM scratch
COPY --from=builder /main /bin/main

ENTRYPOINT ["/bin/main"]
