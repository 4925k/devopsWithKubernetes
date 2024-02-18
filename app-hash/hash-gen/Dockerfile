FROM golang:alpine

WORKDIR /usr/src/app

COPY . .

RUN go build -o hash-gen .

CMD [ "./hash-gen" ]
