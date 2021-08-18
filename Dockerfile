FROM golang:1.16 AS build
WORKDIR /go/src/github.com/ch4nn0n/rpi-led-control/
COPY . ./
RUN go get
RUN GOOS=linux GOARCH=arm GOARM=5 go build

FROM arm32v7/alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/ch4nn0n/rpi-led-control/rpi-led-control ./
CMD ["./rpi-led-control"]
