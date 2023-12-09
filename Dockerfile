FROM quay.io/containers/podman

COPY .build /tmp/build

RUN mv /tmp/build/tsctl /usr/local/bin/tsctl && \
    chmod +x /usr/local/bin/tsctl && \
    rm -rf /tmp/build

RUN ssh-keyscan github.com >> /tmp/known_hosts && \
    ssh-keyscan gitlab.com >> /tmp/known_hosts && \
    ssh-keyscan bitbucket.org >> /tmp/known_hosts && \
    ssh-keyscan ssh.dev.azure.com >> /tmp/known_hosts

RUN curl https://cli.configu.com/install.sh | sh

RUN sudo ln -s /usr/bin/podman /usr/local/bin/docker

COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint
ENTRYPOINT ["docker-entrypoint"]
CMD ["_run"]
