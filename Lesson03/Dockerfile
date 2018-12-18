FROM golang:1.11.2-alpine3.8 as builder
ADD . /go/src/gitlab.com/onuryilmaz/book-server-cd
WORKDIR /go/src/gitlab.com/onuryilmaz/book-server-cd/cmd
ARG VERSION
RUN go build -ldflags "-X main.version=$VERSION" -o book-server

FROM alpine:3.8 as production
COPY --from=builder /go/src/gitlab.com/onuryilmaz/book-server-cd/cmd/book-server /book-server
ENTRYPOINT ["/book-server"]