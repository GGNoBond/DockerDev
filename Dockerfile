FROM golang:1.22_dlv_vim

MAINTAINER aox aox@163.com

#ENV GO111MODULE=off
#
#ADD main.go /go/src/main.go
#
#EXPOSE 8080
#
#RUN cd bin && go build ../src/main.go
#
#ENTRYPOINT ["/go/bin/main"]

ADD . $GOPATH/src