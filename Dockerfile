FROM golang:1.16
RUN go get github.com/cespare/reflex
WORKDIR /app
COPY . .
ENTRYPOINT ["reflex", "-c", "reflex.conf"]