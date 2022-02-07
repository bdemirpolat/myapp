FROM golang:1.16

WORKDIR /go/src/app
COPY main.go .

RUN go build main.go
CMD ["./main"]