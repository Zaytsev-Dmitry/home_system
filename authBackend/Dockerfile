FROM golang:1.23.5 as builder
WORKDIR /build
COPY /cmd/main.go .
COPY go.mod .
RUN go mod download
RUN go mod tidy
COPY . .

RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
RUN oapi-codegen -package=generatedApi -generate "types,spec,gin" /build/api/spec/auth-server-api.yml > api/spec/auth-server-api.gen.go

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go
EXPOSE 8081
ENV GIN_MODE=release

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]
