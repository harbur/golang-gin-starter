# Build Image
FROM golang:1.14.3-alpine AS build
RUN apk add --no-cache g++ make git
WORKDIR /app
COPY Makefile .
RUN make setup

COPY . .
RUN make deps test install

# Runtime Image
FROM golang:1.14.3-alpine
COPY --from=build /go/bin/golang-gin-starter /bin/

EXPOSE 8080
CMD ["golang-gin-starter", "-address", "0.0.0.0:8080"]
