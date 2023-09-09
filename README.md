# TruStacks Software Delivery Engine

## What is TruStacks?

TruStacks is a pipeless software delivery engine that removes the need to build pipelines.

#### Pipeless Sofware Delivery

*Intermediary Wiring* is the pipeline code that sits between the source and the delivered product that is typically implemented with yaml or other no-code DSL's.

Pipeless software delivery (PSD) works without intermediary wiring. TruStacks accomplishes PSD through action plans, bundles of actions determined by rules applied to source facts, so that developers can go from source to deployment with no intermediary wiring.

## Usage

### Installation

Download the latest release [here](https://github.com/TruStacks/trustacks/releases)

### Plan

Run the following command to build an action plan

```bash
tsctl plan --name <plan_name> --source <path_to_source>
```

This will generate a plan file in the working directory named `<plan_name>.plan`

<small>**Run `tsctl explain <plan_name>.plan` to get a detailed description of the action plan*.</small>

### Input

Run the following command to generate stack inputs from the plan. The inputs contain credentials and parameters for your software delivery stack that will be used by actions in the action plan.

```bash
tsctl stack init --from-plan <plan_name>.plan
```

This will generate an input file in the working directory named `inputs.yaml` with the keys and null values for the required inputs.

| <small>⚠ *DO NOT MODIFY INPUT KEYS*</small>

##### Encryption

TruStacks uses [sops](https://github.com/getsops/sops) and [age](https://github.com/FiloSottile/age) for input encryption.

Run the following commands to encrypt your inputs

1. Generate an age key

```bash
age-keygen -o key.txt
```

2. Encrypt your stack inputs file using the public key from the *key.txt* file

```bash
sops -e -i --age <age_public_key> inputs.yaml
```

## Run

Action plans require [docker](https://www.docker.com/).

Run the following in a docker environment (local or CI/CD) from the root of your source.

```bash
SOPS_AGE_KEY=<age_secret_key> tsctl run <plan_file> --source <path_to_source> --inputs <path_to_encrypted_inputs> --stages feedback
```

<small>**Remove the `--stages` option to run the complete plan*.</small>

## Project Catalyst

[Catalyst](https://github.com/orgs/TruStacks/projects/2) is the current project the ecompasses the core engine features of TruStacks. The status of current and planned features is below.

#### Frameworks

|Framework|Implemented|
|-|-|
|React|✅|
|Angular|❌|
|Vue|❌|
|Gin|❌|
|FastAPI|❌|
|.Net|❌|
|Flutter|❌|

#### Actions

|Actions|Implemented|
|-|-|
|React Build|✅|
|React Test|✅|
|GolangCI Lint|❌|
|Go Test|❌|
|Container Build|✅|
|Container Deploy|✅|
|ESLint|✅|
|Helm Release (Stage)|✅|
|Helm Release (Release)|❌|
|SonarQube Scan|✅|
|Trivy Container Scan|✅|
|Cypress|❌|
|ArgoCD Sync|❌|

### Contributing

If you are interested in contributing to TruStacks please join our [Community](https://discord.gg/usgjQj7QTd) to get involved. 

| <small>⚠ TruStacks is early alpha software and is not recommended for use in production systems. We provide provide no guarantees of API stability until we reach a stable v1.0+.</small>