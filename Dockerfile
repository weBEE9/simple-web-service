FROM golang:1.17-alpine3.14

LABEL maintainer="Weeee9 <teletubby332@gmail.com>"

RUN apk add bash ca-certificates git gcc g++ libc-dev make
WORKDIR /app

COPY . .
RUN make

CMD ["./bin/app"]