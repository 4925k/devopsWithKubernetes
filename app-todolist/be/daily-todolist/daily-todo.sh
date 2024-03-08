#!/bin/sh

# install curl and psql
apk update
apk add curl
apk add postgresql-client

# get daily url to read
postgres_url="postgres://postgres:admin@postgres?sslmode=disable"
daily_url=$(curl -s -L -I -o /dev/null -w '%{url_effective}' "https://en.wikipedia.org/wiki/Special:Random")
STMT="INSERT INTO todolist (id, description) VALUES (0, 'READ $daily_url') ON CONFLICT (id) DO UPDATE SET description='READ $daily_url'"

psql $postgres_url -c "$STMT"


# curl -s -L -I -o /dev/null -w '%{url_effective}' "https://en.wikipedia.org/wiki/Special:Random"
