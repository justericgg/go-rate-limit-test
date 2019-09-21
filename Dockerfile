FROM golang:1.13.0 as builder
WORKDIR /go-rate-limit-test
COPY go.mod .
#COPY go.sum .
RUN go mod download
COPY ./ ./
WORKDIR /go-rate-limit-test/cmd/http
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

FROM alpine:latest
RUN apk add --no-cache ca-certificates && apk add tzdata
ENV TZ=Asia/Taipei
WORKDIR /go/bin
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]