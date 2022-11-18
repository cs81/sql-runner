FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.com.cn,direct
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM scratch
WORKDIR /app
COPY --from=builder /app/sql-runner .
ENTRYPOINT ["./sql-runner"]