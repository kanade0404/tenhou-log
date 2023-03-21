FROM golang:1.20 AS builder

WORKDIR /app
COPY . .
WORKDIR /app/services/scraper
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go

FROM gcr.io/distroless/static-debian11

WORKDIR /app
COPY --from=builder /app/services/scraper/main .
EXPOSE 8088
ENV ENV=dev
CMD ["./main"]
