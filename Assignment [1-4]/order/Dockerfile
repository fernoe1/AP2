FROM golang

WORKDIR /usr/src/app

COPY . .

RUN go build -o order cmd/main.go

CMD ["./order"]