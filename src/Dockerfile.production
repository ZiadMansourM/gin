FROM golang:1.18.2-alpine3.16 AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /server

FROM scratch
WORKDIR /app
COPY --from=builder /server .
EXPOSE 3000

ENTRYPOINT ["/app/server"]