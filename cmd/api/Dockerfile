FROM golang 

WORKDIR /go/src/github.com/johnmcdnl/

ADD . . 

RUN go build -o api 

EXPOSE 3333 

ENTRYPOINT [ "./api" ]


