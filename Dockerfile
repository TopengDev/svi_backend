# build
FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# build main at root (.) — change if your main is elsewhere
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# run
FROM alpine:3.20
WORKDIR /app
RUN adduser -D -H appuser
USER appuser
COPY --from=builder /app/server /app/server
ENV GIN_MODE=release
EXPOSE 8080
CMD ["/app/server"]

