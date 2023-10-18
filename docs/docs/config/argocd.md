---
slug: /configuration/argocd
title: ArgoCD
---

# ArgoCD Configuration

Table: `argocd`

|Name|Type|Description|Example|
|-|-|-|-|
|insecure|boolean|skip server certificate and domain verification|false|
|grpcWeb|boolean|enables the gRPC-web protocol|false|

Usage Example:

```
[argocd]
insecure = false
grpcWeb = true
```