FROM golang:alpine

WORKDIR /todo

COPY . .

ENV PORT="3000"

RUN go build -o todolist ./cmd/todolist

CMD [ "./todolist" ]