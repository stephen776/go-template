FROM golang:1.16.0-alpine3.13 as builder
COPY go.mod go.sum /go/src/bitbucket.org/truefit/tf-manifest/
WORKDIR /go/src/bitbucket.org/truefit/tf-manifest
RUN go mod download
COPY . /go/src/bitbucket.org/truefit/tf-manifest
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/tf-manifest bitbucket.org/truefit/tf-manifest

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/bitbucket.org/truefit/tf-manifest/build/tf-manifest /usr/bin/tf-manifest
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/tf-manifest"]