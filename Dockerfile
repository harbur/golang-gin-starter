# Build Image
FROM golang:1.23.5-alpine AS build
RUN apk add --no-cache g++ make git
WORKDIR /go/src/github.com/harbur/golang-gin-starter
COPY Makefile .
RUN make setup

COPY . .
ENV CGO_ENABLED=0
RUN make install

# Runtime Image
FROM golang:1.23.5-alpine
COPY --from=build /go/bin/golang-gin-starter /bin/

EXPOSE 8080
CMD ["golang-gin-starter", "-address", "0.0.0.0:8080"]
