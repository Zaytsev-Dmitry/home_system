FROM golang:1.23.5-alpine
WORKDIR /app
COPY /cmd/main.go .
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /app/cmd/main.go
EXPOSE 8081
# Set environment variable for Gin mode
ENV GIN_MODE=release
# Run the executable
CMD ["/app/cmd/main.go"]