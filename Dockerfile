FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app ./cmd/main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
