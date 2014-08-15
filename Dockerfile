FROM golang

ADD . /go/src/github.com/sysboy/looper

RUN go get github.com/sysboy/looper

ENTRYPOINT ["looper"]
