---
title: Container
slug: /inputs/container
---

# Container Inputs

Container inputs provide parameters for publishing and authenticating against container registries.

## Fields

| name | type | description |
| - | - |-|
| ContainerRegistry | string | The full url of the container registry, used for authentication and iamge publishing (ie. `docker.io/my-repo/my-app`). |
| ContainerRegistryUsername | string | the authentication username for the container registry. |
| ContainerRegistryPassword | string | the authentication password for the container registry. |