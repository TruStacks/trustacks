---
title: Build
slug: /actions/containerize/build
---

# Containerize - Build

The build action builds an [OCI compliant](https://opencontainers.org/) image from the application Dockerfile or Containerfile.

### Stack Inputs

- [ContainerRegistry](/inputs#container)
- [ContainerRegistryUsername](/inputs#container)
- [ContainerRegistryPassword](/inputs#container)

### Artifacts

#### Input:

|Name|Type|Description|
|-|-|-|
|dist|dir|Build artifacts from the application build action|

#### Output:

|Name|Type|Description|
|-|-|-|
|image.tar|file|OCI compliant container image tar|
