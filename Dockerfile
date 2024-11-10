FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o wallet ./cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/wallet /usr/local/bin/wallet
COPY --from=builder /app/config.env /app

ENTRYPOINT ["wallet"]
