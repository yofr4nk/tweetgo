FROM golang:latest

WORKDIR /go/src/github.com/yofr4nk/tweetgo

COPY ./ /go/src/github.com/yofr4nk/tweetgo

RUN go get -u github.com/golang/dep/cmd/dep

RUN go get github.com/githubnemo/CompileDaemon

CMD dep ensure && CompileDaemon --build="go build main.go" --command=./main