FROM alpine:latest

WORKDIR /app

COPY --chmod=0755 daily-todo.sh .

ENTRYPOINT [ "./daily-todo.sh" ]