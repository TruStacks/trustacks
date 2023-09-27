---
title: Copy
slug: /actions/container/copy
---

# Containerize - Copy

The copy action publishes a container image to an image registry.

### Input Variables

- [CONTAINER_REGISTRY](/inputs#container)
- [CONTAINER_REGISTRY_USERNAME](/inputs#container)
- [CONTAINER_REGISTRY_PASSWORD](/inputs#container)

### Artifacts

#### Inputs:

|Name|Type|Description|
|-|-|-|
|image.tar|image|OCI compliant container image tar|
|version|file|The semantic version for the build that will be used as the container image tag
