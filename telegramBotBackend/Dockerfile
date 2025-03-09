FROM golang:1.23.5-alpine
WORKDIR /app
COPY /cmd/main.go .
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /app/cmd/main.go
EXPOSE 8080
ENV GIN_MODE=release
CMD ["/app/cmd/main.go"]