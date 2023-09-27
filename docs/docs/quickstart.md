---
title: Quickstart
sidebar_position: 2
---

### Installation

Download the cli [here](https://github.com/TruStacks/trustacks/releases)

### Plan

Run the following command to build an action plan from your project source.

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

This will generate an input file in the working directory named `inputs.env` with unix exports and empty variables.

:::caution
Do not modify the varaible names
:::


## Run

Action plans require [docker](https://docs.docker.com/get-docker/).

Run the following in a docker environment (local or CI/CD) from the root of your source.

```bash
tsctl run <plan_file> --source <path_to_source> --inputs <path_to_encrypted_inputs> --stages feedback
```

:::tip
Remove the `--stages` option to run the complete plan
:::