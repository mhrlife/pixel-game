FROM golang:1.23-alpine AS build

WORKDIR /app

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN  CGO_ENABLED=0 go build -o main .

# Stage 2: Run the Go application
FROM alpine:latest

WORKDIR /root/

COPY --from=build /internal/app/main .

EXPOSE 8001

CMD ["./main","serve"]