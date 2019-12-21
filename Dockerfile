FROM golang:1.13
RUN mkdir -p /go/src/github.com/dugiahuy/hotel-data-merge
WORKDIR /go/src/github.com/dugiahuy/hotel-data-merge
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server cmd/*.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /

## config for timezone
COPY --from=0 /go/src/github.com/dugiahuy/hotel-data-merge server
COPY --from=0 /usr/share/zoneinfo /usr/share/zoneinfo
EXPOSE 8000
CMD ["/server"]
