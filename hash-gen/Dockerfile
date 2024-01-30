FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o hash-gen .

CMD [ "./hash-gen" ]
