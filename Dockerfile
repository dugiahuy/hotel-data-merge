FROM golang:1.13
WORKDIR /go/src/github.com/dugiahuy/hotel-data-merge
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server cmd/*.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=0 /go/src/github.com/dugiahuy/hotel-data-merge/server /server
EXPOSE 8080
CMD ["/server"]
