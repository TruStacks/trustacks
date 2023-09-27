---
title: Overview
slug: /specification
sidebar_position: 1
---

<small>Draft v0.1</small>

## Purpose
The Open Software Delivery Specification (OSDS) defines terminology, activities, and utilities related to the implementation of software delivery processes.

## What is Software Delivery?
The OSDS defines software delivery as follows:

Software delivery is the process of testing, securing, and building software source code, resulting in a deliverable product deployed to a target destination.


State snapshots are immutable sources of software delivery process state. State Snapshots are satisfied by properly constructed source commits.

## The OSDS Golden Rule
> *Software delivery processes must not contain actions that modify the snapshot state such as generating semantic versions or any other artifacts that may change or result in the loss of contextual information required to complete said software delivery process.*

In other words, the snapshot must provide all that is needed to deliver the application beforehand. Scripts and other helpers such as git hooks should be used prior to shippable source commits in order to comply with this rule.

State snapshots and the resulting software delivery process execution must have the following traits:
- Snapshots are immutable.
- Snapshots contain all relevant context such as versions, files, and dependencies.
- Artifacts created during a software delivery process execution must be ephemeral and additive without modification to the snapshot state.

The rules defined above ensure that any snapshot of a given source will always result in the same outcome when executed via the same software delivery process. The implications of such snapshots go beyond general idempotence for the sake of consistency. The rules also ensure that sources and artifacts in regulated systems are verifiable and reproducible.

### Process Inputs

The golden rule does not imply that process inputs, such as credentials or other mutable parameters, must be defined in the source. Changing credentials does not result in significant changes to the outcome of process deliverables so these values are exempt from inclusion in the source. If credentials or other inputs without side effect must be provided to the software delivery process they can, and likely should, be provided via external means such as environment variables or files.

# Versioning

The most common versioning strategy is [semantic versioning](https://semver.org/) (or semver). Semver is the desired versioning strategy in OSDS compliant software delivery processes, however there are times when alternatives are appropriate such as commit hashes.

Semver should be preferred when:
- You are building an application that is distributed to end-users that require context for API changes such as breaks or bug fixes.
- You are building a library that is distributed to developers that will use it in their own applications.

Commit hashes can be used when:
- You are building an application that is deployed as a SaaS or other means where users are not exposed to implementation details.
- You require no context from the build version, other than the hash itself, in order to reference the source code of a deliverable for debugging.

If you are not sure of which versioning to use then semver is preferred with the exception being that semver requires more maintenance. Semver must be meticulously maintained in all relevant sources such as `package.json`, `pom.xml` and `helm` to ensure that build assets all match during a software delivery process execution.

Commit hashes are always an artifact of a source commit thus there is no effort on the part of the developer to generate them.

# Appendix A - Glossary

## Build & Versioning

### Conventional Commits

> Conventional commits provide a standard source commit syntax in order to generate semantic versions. Additional tools are needed to parse the commit messages in addition to hooks and other utilities to ensure that commit messages are compliant with the [Conventional Commits Specification](https://www.conventionalcommits.org/en/v1.0.0/#specification).

### Builds

> A build is source code that is converted into a deliverable. Builds are completed by framework or language specific tools and take on many forms. Binaries, containers, mobile applications, and FaaS code are all examples of builds.

## Testing

### White/Black Box Test

> A white box test is a test that is fully aware of, and given access to, internal knowledge about, or interfaces to the system under test (SUT). White box tests are “inside out”.
>
>A black box test is a test that is unaware of the internals and interfaces of the SUT. Black box tests are “outside in”.
>
>*SUT includes all testable systems including software, applications, and all other testable system types.

### Unit Test

> A unit test is a software test that executes with no external dependencies. Unit tests are generally small and granular, isolated to a single function or interface, and built to run fast to provide the developer with rapid feedback on the health of the application.

### Integration Test

> An integration test is a software test that verifies that inter-service interactions are healthy. Integration tests differ from e2e tests in that they include stubs or other mocking methodologies in order to isolate the desired service under test and eliminate the need to deploy the entire application and all associated dependencies.  
>
> An example of an integration test would be testing the interactions between a service and its associated database or service to service interaction with no requirement to verify the dependency service’s downstream integrations.


### API Contract Test

> An API contract test is a solution test that verifies the interactions between connected services. Contracts can be deployed on services or stubs. A successful API contract test creates confidence that inter-service operations are compliant with the required contract inputs and outputs.

### End-to-End Test

> An End-to-End (E2E) test is a solution test that tests the entire system and all associated services, database, queues and all other systems and subsystems in totality. The E2E test is the most comprehensive and complex test type. E2E tests are run in staging environments because they require the entire solution to be deployed which is unattainable in lower environments (namely development). 
> 
> E2E tests require a tool such as Cypress, Selenium, Cucumber, or Detox to drive the solution rather than a framework or language tool that is executed against the source code.

