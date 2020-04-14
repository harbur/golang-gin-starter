# Build Image
FROM golang:1.11.5-alpine AS build
RUN apk add --no-cache g++ make git \
  && go get github.com/ahmetb/govvv
COPY . /go/src/github.com/harbur/golang-starter
WORKDIR /go/src/github.com/harbur/golang-gin-starter/
RUN make test install

# Runtime Image
FROM golang:1.11.5-alpine
COPY --from=build /go/bin/golang-starter-server /bin/golang-starter-server
WORKDIR /

EXPOSE 8080
CMD ["golang-starter-server", "--host=0.0.0.0", "--port=8080"]
