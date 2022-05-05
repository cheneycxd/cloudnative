FROM golang:1.18-alpine AS build
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY go.mod go.sum main.go /go/src/httpserver/
WORKDIR /go/src/httpserver
RUN go env -w GO111MODULE=on 
RUN go env -w GOPROXY="https://goproxy.cn,direct" 
RUN go mod download 
RUN go build -o /bin/httpserver


FROM scratch
COPY --from=build /bin/httpserver /bin/httpserver
CMD ["/bin/httpserver","-log_dir=./"]

EXPOSE 8001

