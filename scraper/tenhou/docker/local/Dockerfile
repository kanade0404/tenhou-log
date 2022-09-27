FROM golang:1.19
RUN mkdir /main
WORKDIR /main
COPY ../../go.* .
COPY ../../.air.toml .
RUN go install github.com/go-task/task/v3/cmd/task@latest && \
    go install github.com/volatiletech/sqlboiler/v4@v4.10.2 && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.10.2 && \
    go install github.com/cosmtrek/air@latest && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    go mod tidy
CMD ["air", "-c", ".air.toml"]
