#!/bin/sh

if [ "$1" == "_run" ]; then
  if [ -f .configu ]; then
    configu eval --store 'trustacks' --set $TS_CONFIG_SET --schema 'trustacks.cfgu.json' | configu export --format 'Dotenv' > .env && eval export $(cat .env)
  fi
  stages=""
  if [ ! -z "$TS_RUN_STAGES" ]; then
    stages="--stages $TS_RUN_STAGES"
  fi
  tsctl run $stages
else
  export SSH_KNOWN_HOSTS=/tmp/known_hosts
  eval `ssh-agent` > /dev/null
  tsctl "$@"
fi