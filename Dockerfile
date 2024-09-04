FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o ./app cmd/main.go

FROM gcr.io/distroless/static

COPY --from=builder /app/app /

EXPOSE 8080
CMD ["/app"]