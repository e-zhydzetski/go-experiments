FROM golang:1.19.5-alpine3.17
RUN apk add --update tcpdump && rm -rf /var/cache/apk/*
ENV CGO_ENABLED=0
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENTRYPOINT ["./test.sh"]