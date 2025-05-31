# Use latest Go with Debian base
FROM golang:1.23.2 as builder
WORKDIR /app
COPY . .
RUN go build -o app

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/app .
COPY static ./static
EXPOSE 8080
CMD ["./app"]
