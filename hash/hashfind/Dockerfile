FROM golang:alpine

WORKDIR /usr/src/app

COPY . .

RUN go build -o hashfind .

EXPOSE 4444

CMD [ "./hashfind" ]