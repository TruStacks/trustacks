---
title: Image
slug: /actions/trivy/image
---

# Trivy - Image

The image action runs the [trivy](https://trivy.dev/) container security scanner.

:::tip
This action uses the [trivy.yaml](https://aquasecurity.github.io/trivy/v0.33/docs/references/customization/config-file/) configuration file in the project root.
:::

### Artifacts

#### Inputs:

|Name|Type|Description|
|-|-|-|
|image.tar|image|OCI compliant container image tar|
