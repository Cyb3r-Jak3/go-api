FROM golang:1.17.3-alpine as build

WORKDIR /go/src/app
COPY . /go/src/app
COPY ./git .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN apk add make
RUN make build

FROM gcr.io/distroless/static
COPY --from=build /go/bin/app /
ENV PRODUCTION=TRUE
EXPOSE 5000
CMD ["/app"]