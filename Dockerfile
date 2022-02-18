#FROM golang:alpine as builder
#
#WORKDIR /app
#
#COPY go.mod .
#COPY go.sum .
#
#RUN go env -w GOPROXY="https://goproxy.cn,goproxy.io,direct"
#RUN go mod download
#
#COPY main.go .
#COPY ./internal .
#COPY ./cmd .
#
#
#RUN go build -o /hermes

FROM alpine:latest

WORKDIR /app

#COPY --from=builder /hermes .
COPY  ./hermes .

EXPOSE 3000

ENTRYPOINT ["./hermes"]

CMD ["serve"]