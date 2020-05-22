FROM golang:latest

WORKDIR /tweetgo

COPY ./ /tweetgo

RUN go get github.com/githubnemo/CompileDaemon

CMD CompileDaemon --build="go build main.go" --command=./main