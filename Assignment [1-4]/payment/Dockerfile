FROM golang

WORKDIR /usr/src/app

COPY . .

RUN go build -o payment cmd/main.go

CMD ["./payment"]