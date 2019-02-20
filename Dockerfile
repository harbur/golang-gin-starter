# Build Image
FROM golang:1.11.5-alpine AS build
RUN apk add --no-cache g++
COPY . /go/src/api-gateway
WORKDIR /go/src/api-gateway
RUN go test ./... && go build .

# Runtime Image
FROM alpine:3.9
RUN apk --no-cache add ca-certificates
COPY --from=build /go/src/api-gateway/api-gateway /
EXPOSE 8080
CMD ["/api-gateway"]
