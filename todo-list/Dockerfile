FROM golang:alpine

WORKDIR /todo

COPY . .

ENV PORT="4000"

RUN go build -o todolist .

CMD [ "./todolist" ]