---
title: Build
slug: /actions/containerize/build
---

# Containerize - Build

The build action builds an [OCI compliant](https://opencontainers.org/) image from the application Dockerfile or Containerfile.

<!-- ### Build Artifact

The containerize action takes a build input artifact. The build artifact contains the built application assets. 

The Dockerfile must define a [build argument](https://docs.docker.com/build/guide/build-args/) named `app_dist`. The Copy directive should reference the build argument in the Dockerfile in order to allow flexibility between the TruStacks action plan and alternative container builds.

Example Dockerfile:

```dockerfile
FROM nginx:stable-alpine
ARG app_dist
COPY $app_dist/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
CMD ["nginx", "-g", "daemon off;"]
``` -->

### Stack Inputs

- [ContainerRegistry](/inputs/container#fields)
- [ContainerRegistryUsername](/inputs/container#fields)
- [ContainerRegistryPassword](/inputs/container#fields)

### Artifacts

#### Input:

|Name|Type|Description|
|-|-|-|
|dist|dir|Build artifacts from the application build action|

#### Output:

|Name|Type|Description|
|-|-|-|
|image.tar|file|OCI compliant container image tar|
