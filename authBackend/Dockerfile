FROM golang:1.23.5-alpine
WORKDIR /app
COPY /cmd/main.go .
COPY go.mod go.sum ./
RUN go mod download
RUN go get -u all
COPY . .

RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
RUN oapi-codegen -package=generatedApi -generate "types,spec,gin" /app/api/spec/auth-server-api.yml > api/spec/auth-server-api.gen.go
RUN go build -o /app/cmd/main.go

EXPOSE 8081
# Set environment variable for Gin mode
ENV GIN_MODE=release
# Run the executable
CMD ["/app/cmd/main.go"]
