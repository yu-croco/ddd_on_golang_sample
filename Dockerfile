FROM golang:1.15

RUN mkdir /ddd_on_golang
WORKDIR /ddd_on_golang

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/go-app

CMD ["go", "run", "/ddd_on_golang/main.go"]