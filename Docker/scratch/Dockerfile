FROM golang:1.11.4-alpine3.8 AS build

WORKDIR /go/src/test
COPY ./test.go .
RUN go build -o /bin/hello

FROM scratch
COPY --from=build /bin/hello /bin/hello
ENTRYPOINT ["/bin/hello"]
CMD []
