FROM golang:1.20 AS builder
ARG version=0.1.0
COPY . /usr/src/app
WORKDIR /usr/src/app
RUN CGO_ENABLED=0 go build -o tsctl -ldflags="-X 'main.version=${version}'" ./cmd

FROM quay.io/containers/podman
COPY --from=builder /usr/src/app/tsctl /usr/local/bin/tsctl

RUN chmod +x /usr/local/bin/tsctl

RUN ssh-keyscan github.com >> /tmp/known_hosts && \
    ssh-keyscan gitlab.com >> /tmp/known_hosts && \
    ssh-keyscan bitbucket.org >> /tmp/known_hosts && \
    ssh-keyscan ssh.dev.azure.com >> /tmp/known_hosts

RUN curl https://cli.configu.com/install.sh | sh

RUN sudo ln -s /usr/bin/podman /usr/local/bin/docker

COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint
ENTRYPOINT ["docker-entrypoint"]
CMD ["_run"]
