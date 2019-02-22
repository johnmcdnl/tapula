FROM golang 

ENV	PROTOCOL    = "http"
ENV	DOMAIN      = "localhost"

WORKDIR /go/src/github.com/johnmcdnl/tapula 

ADD . .

RUN go build -o tapula ./cmd/main/main.go 

ENTRYPOINT [ "./tapula" ]

