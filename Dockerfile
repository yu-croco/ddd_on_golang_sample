FROM golang:1.15 AS build

RUN mkdir /ddd_on_golang
WORKDIR /ddd_on_golang

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN cd cmd/ddd_on_golang && \
    CGO_ENABLED=0 go build -o /bin/ddd_on_golang

FROM scratch

COPY --from=build /bin/ddd_on_golang /bin/ddd_on_golang

CMD ["bin/ddd_on_golang"]
