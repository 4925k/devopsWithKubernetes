FROM golang:alpine

WORKDIR /todofe

COPY . .

ENV PORT="3000"

RUN go build -o todolistfe ./app

CMD [ "./todolistfe" ]