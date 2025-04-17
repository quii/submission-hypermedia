ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod ./
# Only copy go.sum if it exists, otherwise create it with go mod tidy
COPY go.sum* ./
RUN if [ ! -f go.sum ]; then go mod tidy; fi
COPY . .
RUN go build -v -o /run-app .

FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
