FROM golang:1.11.5-alpine
RUN apk add --no-cache g++
COPY . /go/src/app
WORKDIR /go/src/app
RUN go test ./... && go build .
EXPOSE 8080
CMD ["/go/src/app/app"]
