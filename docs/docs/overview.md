---
  sidebar_position: 1
---

# Overview

The TruStacks software factory (TSSF) is a consumable toolchain designed to ease the burden of software delivery.

The goal of this document is to describe the requirements for such toolchains to ensure consistency and repeatability for software teams leveraging TruStacks workflows.

## Implementation

The TSSF must be implemented as a kubernetes cluster. Kubernetes offers the essential scaffolding to support scaling components and CI execution workloads in addition to secret injection, configuration mapping, and environment variables.