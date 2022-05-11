FROM golang:1.18
RUN mkdir /main
WORKDIR /main
ADD . /main
RUN go install github.com/go-task/task/v3/cmd/task@latest
