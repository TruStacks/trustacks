---
id: ci-executor
title: CI Executor
---

# CI Executor

The CI executor is responsible for orchestrating dependencies and running [Dagger](https://dagger.io/) plans that are provided by workflows. The CI executor must have the following capabilities:

Run containers
- Pass environment variables to the dagger runtime
- Pull from version control
- Retry for intermittent failures

The CI executor must delegate two tasks to the dagger runtime

## Build

The build task is executed prior to the release of an application. The purpose of this task is to prepare the code for release. All analysis such as unit tests, static analysis, and security scans should be run in this task. Any build artifacts such as modules, binaries, or containers should also be created and stored in the artifact server.

Additional actions will be taken depending on the application type. For example a containerized application will also include a container security scan.

The analysis task will generate a number of report artifacts that are available to the product team for inspection.

Any failures encountered during the analysis step will cause the entire CI execution to fail until the issue is resolved.

## Release

The release task is executed after analysis. The release task will deliver code per the workflow. The delivery method will vary depending on the workflow. For example a Kubernetes delivery method would deploy the application container using Helm or Kustomize.

## Idempotence

Idempotence is required to ensure that duplication of artifacts and side effects are prevented in the case of a CI task retry.