FROM golang:1.18.3-alpine as build

WORKDIR /go/src/app
COPY . ./
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go get ./...
RUN apk add make git
RUN make build

FROM gcr.io/distroless/static
COPY --from=build /go/src/app/app /
ENV PRODUCTION=TRUE
EXPOSE 5000
CMD ["/app"]