FROM golang

WORKDIR /usr/src/app

COPY . .

RUN go build -o notification cmd/main.go

CMD ["./notification"]