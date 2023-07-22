---
sidebar_position: 1
title: Overview
slug: /actions
---

# Actions Overview

Actions are the individual steps that make make up an action plan. Actions are categorized into phases and their execution order is determined by the orchestrator.
<!-- 
## Source Injection

All actions are given a copy of the application source at the path `/src`. The `/src` directory is always set as the default working directory for all action sub-commands. 

:::caution
TruStacks strictly enforces a read-only methodology for git respositories. Pushing to git repositories in actions should only occur if explicitly required by an action.

Actions must be diligent to keep sources and artifacts flowing downstream without cycling back to the source.
::: -->

## Admission Criteria

Actions are admitted into an action plan based on conditions such as the presence of certain files or directories, or based on the contents of files.

These conditions are known as **Admission Criteria**.

#### Admission Tests

To build an action plan, the engine executes each action's registered **Admission Test**. An admission test contains the [Admission Criteria](#admission-criteria) associated with the action.

If the Aadmission Test passes then the action is added to the action plan

:::tip a word on bias
Admission criteria are designed to be highly biased to avoid the need for conflict resolution. It should be rare that two actions of the same classification would result in a successful admission test.
:::

#### Fact gathering

In addition to action plan admission, each action can generate zero or more facts during an [Admission Test](#admission-tests). 

Facts are added to the action plan in addition to unpopulated user-defined fields that will be used later during action plan orchestration and execution.

User-defined fields must be populated by the user before executing the action plan.