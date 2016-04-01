FROM golang:1.6-alpine
MAINTAINER mihail.samoylov@gmail.com

RUN apk add --update \
  git && \
  git clone https://github.com/teampact/breaker.git /go/src/breaker && \
  cd /go/src/breaker && go get -v && go build -v && cp ./breaker /usr/bin/

CMD ["/usr/bin/breaker"]
