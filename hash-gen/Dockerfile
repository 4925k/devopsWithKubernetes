FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o hash-gen .

EXPOSE 4445

CMD [ "./hash-gen" ]
