FROM golang:1.9 AS build
COPY k8s-vault-tester.go /go/src/app/app.go

WORKDIR /go/src/app
RUN go get -d ./... && \
 CGO_ENABLED=0 go build -o app .

# copy the binary from the build stage to the final stage
FROM alpine:3.6
RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*
    
COPY --from=build /go/src/app/app /k8s-vault-tester
CMD ["/k8s-vault-tester"]
