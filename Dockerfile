FROM golang:alpine

WORKDIR /go/src/github.com/Devansh3712/portfolio
COPY . .
RUN apk update && apk add --no-cache git

RUN go get ./...
RUN go build
EXPOSE 8080
CMD ["./portfolio"]
