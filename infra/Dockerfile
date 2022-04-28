# Build Stage
FROM golang:1.18-alpine3.15 AS builder

WORKDIR /app

COPY ../go.mod go.sum ./
RUN go mod download

COPY .. .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o profile-tracker cmd/*.go

# Run Stage
FROM alpine:3.15

WORKDIR /app

COPY --from=builder /app/profile-tracker .


## PORT NEEDS TO BE DEFINED ELSEWHERE

ENTRYPOINT ["/app/profile-tracker"]
