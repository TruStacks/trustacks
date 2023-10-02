# TruStacks Software Delivery Engine

## What is TruStacks?

TruStacks is a generative software delivery engine that removes the need to build pipelines.

#### Generative Sofware Delivery

*Intermediary Wiring* is the pipeline code that sits between the source and the delivered product that is typically implemented with yaml or other no-code DSL's.

TruStacks uses a rule engine to generate actions plans that contain actions based on discovered facts in the application sources.

## Usage

### Installation

Download the latest release [here](https://github.com/trustacks/trustacks/pkg/releases)

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

### Contributing

If you are interested in contributing to TruStacks please join our [Community](https://discord.gg/usgjQj7QTd) to get involved. 

| <small>⚠ TruStacks is early alpha software and is not recommended for use in production systems. We provide no guarantees of API stability until stable v1.0.</small>