---
sidebar_position: 4
slug: /get-started/input
title: Input
---

# Input

This part of the guide will walk through setting up action inputs

### What are stack inputs?

Stack inputs are user-defined parameters and credentials that are used in action plans to complete the delivery of an application.

<!---

## Stacks

To get started with inputs click the stack icon beside the **New Action Plan** button.

You are now on the stacks page, and you should see a number of locked inputs.

### Using Age

In order to unlock and edit stack inputs, we need an age secret key.

Run the following command to create an age key pair:

```
age-keygen -o key.txt
```

:::tip
If you missed the age installation, go back to the [setup](/get-started/setup#age) page.
:::

This will create a key file with a public key and a secret key beginning with `AGE-SECRET-KEY-...`.

Copy the secret key, paste it in the **Age Secret Key** input on the option bar, and click the lock icon to the right of the input box.

The lock should change color and the icon should show an open lock. Input placeholders should be visible and input should no longer be disabled.

:::caution key management
TruStacks does not store keys of any kind. This includes age keys and ssh keys. 

Keep your keys handy because if you lose them we have no way to decrypt your inputs.
:::

### Input Management

Inputs can be managed in the stack context or in the application.

#### Stack context inputs

Inputs defined in the stack context will be available to every application.

Examples of stack context inputs are:
- container registry credentials
- kubeconfigs

#### Application inputs

Inputs defined in the application are limited to a single application.

Examples of application inputs are:
- container registry name
- kubernertes namespace

:::tip use cases
Use Stack inputs if you are defining inputs that are shared between applications and use application inputs if you defining inputs for a single application.
:::

-->

### Entering Stack Inputs

##### Using Age

Inputs must be encyrpted before being used in action plans. Run the following command to create an age key pair:

```
age-keygen -o key.txt
```

Keep this key handy for the next step.

#### Generate a stack input file

Before we populate stack inputs we need to generate the appropriate input keys.

TruStacks provides a helper command to generate a keyed input file from an existing plan. Run the following command to generate a new input file:

```
tsctl stack init --from-plan trustacks-react-sample.plan
```

A file named `inputs.yaml` is placed on the filesystem. You can now populate the keys with the appropriate values.

:::tip
Run `tsctl explain trutacks-react-sample.plan` to view a detailed description of the action plan. The inputs section will contain documentation links for all included inputs with more information.
:::

After populating the input values run the following command to encrypt the inputs file:

```
sops -e -i --age <age_public_key> inputs.yaml
```

Your inputs file should now be encrypted, and you are ready to move on to the next step.