# Build Image
FROM golang:1.24.6-bullseye AS build
RUN apt-get update && apt-get install -y gcc libc-dev sqlite3
WORKDIR /go/src/github.com/harbur/golang-gin-starter
COPY Makefile .
RUN make setup

COPY . .
RUN make install

# Runtime Image
FROM golang:1.24.6-bullseye
COPY --from=build /go/bin/golang-gin-starter /bin/

EXPOSE 8080
CMD ["golang-gin-starter", "-address", "0.0.0.0:8080"]
