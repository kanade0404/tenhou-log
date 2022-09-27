FROM golang:1.19
RUN mkdir /main
WORKDIR /main
COPY ../../go.* .
RUN go install github.com/go-task/task/v3/cmd/task@latest && \
    go install entgo.io/ent/cmd/ent@latest && \
    go install github.com/cosmtrek/air@latest && \
    go mod tidy

