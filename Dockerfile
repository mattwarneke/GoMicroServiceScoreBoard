FROM golang:alpine

ADD src/app /go/src/app
WORKDIR /go/src/app

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get -d github.com/gorilla/mux

ENV PORT=8080
CMD ["go", "run", "main.go", "score.go"]