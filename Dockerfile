FROM golang:1.11.5

ADD . /go/src/gin-book-mementos/

WORKDIR /go/src/gin-book-mementos/

RUN go get -u github.com/gin-gonic/gin && go get github.com/jinzhu/gorm && go get github.com/go-sql-driver/mysql

CMD go run main.go