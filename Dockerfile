FROM golang:1.17-alpine AS builder

RUN apk --no-cache add ca-certificates

WORKDIR /build

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o analyze_text .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder ["/build/analyze_text", "/"]

EXPOSE 3008

ENTRYPOINT ["/analyze_text"]