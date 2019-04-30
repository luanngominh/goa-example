FROM golang:1.12.4-alpine3.9 as build

WORKDIR /go/src/github.com/luanngominh/goa-example
ADD . /go/src/github.com/luanngominh/goa-example
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /goa-example ext/cmd/goa-example/main.go

FROM alpine:3.8

WORKDIR /
COPY --from=build /goa-example /goa-example

EXPOSE 8888

ENTRYPOINT [ "/goa-example" ]
