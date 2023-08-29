FROM golang:1.20 AS builder
ARG version=0.1.0
COPY . /usr/src/app
WORKDIR /usr/src/app
RUN CGO_ENABLED=0 go build -o tsctl -ldflags="-X 'main.version=${version}'" ./cmd

FROM alpine
COPY --from=builder /usr/src/app/tsctl /usr/local/bin/tsctl

RUN chmod +x /usr/local/bin/tsctl
RUN apk add --no-cache docker curl openssh-client

RUN ssh-keyscan github.com >> /tmp/known_hosts && \
    ssh-keyscan gitlab.com >> /tmp/known_hosts && \
    ssh-keyscan bitbucket.org >> /tmp/known_hosts && \
    ssh-keyscan ssh.dev.azure.com >> /tmp/known_hosts
RUN curl -Lo /usr/local/bin/sops https://github.com/getsops/sops/releases/download/v3.7.3/sops-v3.7.3.linux.amd64 && \
    chmod +x /usr/local/bin/sops

COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint
ENTRYPOINT ["docker-entrypoint"]
