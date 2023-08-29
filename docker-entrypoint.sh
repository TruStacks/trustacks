#!/bin/sh

export SSH_KNOWN_HOSTS=/tmp/known_hosts
eval `ssh-agent` > /dev/null
tsctl "$@"
