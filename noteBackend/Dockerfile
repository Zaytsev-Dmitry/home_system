FROM golang:1.23.5-alpine
WORKDIR /app
COPY /cmd/main.go .
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
RUN oapi-codegen -package=generatedApi -generate "types,spec,gin" /app/api/spec/note-api.yml > api/spec/note-api.gen.go
RUN go build -o /app/cmd/main.go

EXPOSE 8080
ENV GIN_MODE=release
CMD ["/app/cmd/main.go"]