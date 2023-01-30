FROM golang:1.19-buster

WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN go build -v

CMD ["./bcc_university"]