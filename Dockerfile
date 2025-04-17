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

# Copy the executable from the builder stage
COPY --from=builder /run-app /usr/local/bin/

# Copy template files to the final image
COPY --from=builder /usr/src/app/templates /templates
COPY --from=builder /usr/src/app/static /static

# Set working directory where templates might be expected
WORKDIR /

CMD ["run-app"]
