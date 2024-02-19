FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o be ./app

CMD ["./be"]