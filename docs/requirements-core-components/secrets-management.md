# Secrets Management

The secrets manager is responsible for managing access to sensitive application data.
Service Accounts

The secrets manager must provide a method for components to store service account credentials for use during CI execution. The secrets manager must also enable easy rotation of service account secrets.


## General Purpose Secrets

All other secrets in the secrets manager must be consumable by the workflow during CI execution via some form of value injection. Secrets must be injected into the CI executor’s dagger plan environment mounted at `/mnt/secrets/*`. 

A tool such as [external-secrets](https://github.com/external-secrets/external-secrets/) in combination with kubernetes volume mounts should be used for secret injection.


## Single Sign On
