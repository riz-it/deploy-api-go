FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
ENV GO111MODULE=on \
    GOPROXY=proxy.golang.org
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod tidy
COPY src/ .
RUN go build -o /app/hello

FROM scratch
COPY --from=builder /app/hello /app/hello
EXPOSE 3000

ENTRYPOINT ["/app/hello"]
