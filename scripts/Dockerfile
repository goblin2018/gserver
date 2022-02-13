FROM golang:1.17-alpine as builder
WORKDIR /usr/src/app
# ENV GOPROXY=https://goproxy.cn
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN  go build -ldflags "-s -w" -o server

FROM scratch as runner
COPY --from=builder /usr/src/app/server /opt/app/
CMD ["/opt/app/server"]
