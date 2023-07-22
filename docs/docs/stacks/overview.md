---
title: Overview
slug: /stacks
sidebar_position: 1
---

# Stacks Overview

A stack is a representation of the software delivery tooling available during action plan execution.

Stacks contain "out-of-band" application configuration parameters that cannot be securely or reliably derived from a source code repository.

## Input Fields

Stacks parameters are provided as input fields in the [Action Plan](/action-plans). The fields required by the stack file are under the `inputs` key as key/value pairs.

:::info
Parameters defined in the stack configuration are only used when needed by a given action plan.
:::

## Stack Configuration File

The `stack.json` is a JSON formatted configuration file that contains all of the stack inputs for a given environment.

### Encryption

[Stackfiles](#stackfile) contain secrets such as kuebconfigs, docker credentials, and other service credentials. The Stackfile must be encrypted at rest using [sops](https://github.com/getsops/sops) with an [age](https://github.com/FiloSottile/age) key. 

Use the links below to install sops and age:

- sops: https://github.com/getsops/sops/releases
- age: https://github.com/FiloSottile/age/releases


Run the below actions to create an encrypted stack configuration.

1. Create the age private and public key.

```bash
age-keygen -o key.txt
```

2. Encrypt your stack.json file using sops with age.

```bash
sops -e --age <age-public-key> stack.json > stack.enc.json
```

You should have a file named `stack.enc.json` on your filesyste with the encrypted contents of your stack.json file.

:::tip
You can name your encrypted file anything that you want, but you should avoid overwriting your stack file using a command such as:

```
sops -e --age <age-public-key> stack.json > stack.json
```

If you want to use the name stack.json you can add the `-i` argument to sops to encrypt your stack configuration in place.
:::

:::caution
*If you encrypt your stack.json in place, and you lose your age key, then you won't be able to recover your stack configuration.*
:::
