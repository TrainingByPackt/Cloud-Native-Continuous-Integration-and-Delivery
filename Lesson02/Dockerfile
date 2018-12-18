FROM golang:alpine as builder
ADD . /go/src/gitlab.com/onuryilmaz/book-server
WORKDIR /go/src/gitlab.com/onuryilmaz/book-server/cmd
RUN go build -o book-server

FROM alpine as production
COPY --from=builder /go/src/gitlab.com/onuryilmaz/book-server/cmd/book-server /book-server
ENTRYPOINT ["/book-server"]
