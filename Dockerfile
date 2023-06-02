# Ensure this version tracks with go.mod and Dockerfile-dev
FROM golang:1.20 as builder

WORKDIR /src
ADD . .
RUN go get ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o redirector

FROM alpine:3
WORKDIR /
COPY --from=builder /src/redirector .
CMD ["./redirector"]
