---
title: Stacks
slug: /stacks
sidebar_position: 1
---

# Stacks Overview

A stack is a representation of the software delivery tooling available during action plan execution.

Stacks contain external configuration parameters that cannot be securely or reliably discovered from a source code repository.

## Input Variables

Stacks parameters are provided as environment variables in the action plan. The inputs required by the action plan can be viewed in the `inputs` in the action plan. For a detailed view from `tsctl explain <plan>`.

:::info
Parameters defined in the stack configuration are only used when needed by a given action plan.
:::
