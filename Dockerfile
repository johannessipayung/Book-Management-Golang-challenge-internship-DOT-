# ---------- Stage 1: Build ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o book-management .

# ---------- Stage 2: Run ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/book-management .

EXPOSE 8080

CMD ["./book-management"]