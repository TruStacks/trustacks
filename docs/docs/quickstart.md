---
title: Quickstart
sidebar_position: 2
---

### Installation

Download the cli [here](https://github.com/TruStacks/trustacks/releases)

### Plan

Run the following command to build an action plan

```bash
tsctl plan --name <plan_name> --source <path_to_source>
```

This will generate a plan file in the working directory named `<plan_name>.plan`

:::tip
Run `tsctl explain <plan_name>.plan` to get a detailed description of the action plan
:::

### Input

Run the following command to generate stack inputs from the plan. The inputs contain credentials and parameters for your software delivery stack that will be used by actions in the action plan.

```bash
tsctl stack init --from-plan <plan_name>.plan
```

This will generate an input file in the working directory named `inputs.yaml` with the keys and null values for the required inputs.

:::caution
DO NOT MODIFY INPUT KEYS
:::

##### Encryption

TruStacks uses [sops](https://github.com/getsops/sops/releases) and [age](https://github.com/FiloSottile/age/releases) for input encryption.

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

Action plans require [docker](https://docs.docker.com/get-docker/).

Run the following in a docker environment (local or CI/CD) from the root of your source.

```bash
SOPS_AGE_KEY=<age_secret_key> tsctl run <plan_file> --source <path_to_source> --inputs <path_to_encrypted_inputs> --stages feedback
```

:::tip
Remove the `--stages` option to run the complete plan
:::