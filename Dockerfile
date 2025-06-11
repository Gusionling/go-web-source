FROM golang:1.20.1-alpine3.17 AS builder
WORKDIR /work
COPY . .

# 👉 리눅스/amd64 환경용으로 컴파일 (쿠버네티스 호환)
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o server main.go

FROM alpine:3.14
COPY --from=builder /work/server /work/server
ENTRYPOINT ["/work/server"]