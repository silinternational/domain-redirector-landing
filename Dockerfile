FROM golang:latest as builder

WORKDIR /src
ADD . .
RUN go get ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o redirector

FROM alpine:latest
WORKDIR /
COPY --from=builder /src/redirector .
CMD ["./redirector"]