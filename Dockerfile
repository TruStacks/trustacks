FROM golang:1.20 AS builder
ARG version=0.1.0
COPY . /usr/src/app
WORKDIR /usr/src/app
RUN CGO_ENABLED=0 go build -o /tmp/main -ldflags="-X 'main.version=${version}'" ./cmd
RUN curl -Lo /tmp/sops https://github.com/getsops/sops/releases/download/v3.7.3/sops-v3.7.3.linux.amd64

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /tmp/main /main
COPY --from=builder /tmp/sops /usr/local/bin/sops
ENTRYPOINT ["/main"]