FROM golang:1.16

RUN mkdir /todoapp
WORKDIR /todoapp

RUN go mod init github.com/obrkn/todoapp
RUN go get -u github.com/go-sql-driver/mysql

CMD go run cmd/todoapp/main.go