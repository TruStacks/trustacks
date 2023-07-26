---
title: Action Plans
slug: /action-plans
---

# Action Plans

An action plan is an orchestrated collection of actions that perform software delivery.

### Phases

Actions plan activities are categorized into fixed phases. Each action within a respective phase must be resolved before moving on to the next phase. The four phases are [Feedback](#feedback), [Stage](#stage), [QA](#qa), and [Release](#release).


- ##### Feedback

    The feedback phase includes test and analysis actions that are performed early in the delivery process. Feedback actions must complete successfully before moving on to later states.

- ##### Stage

    The stage phase includes package and publish, and staging deployment actions. Completion of the staging phase will result in a staged application that is ready for manual or automated quality assurance.

- ##### QA

    The QA phase includes load, user acceptance, and other live application tests. QA is the final phase before releasing the application to production.

- ##### Release

    The release phases includes the production release actions. After the release phase the application will be at its final destination and live for user consumption.

## Actions

[Actions](/actions) are the individual steps that make make up an action plan. Actions are categorized into phases and execution order is determined by the orchestrator.

## Generation

Actions plans are generated using [Admission Criteria](/actions#admission-criteria) that are implemented as rules. Actions are admitted to the action plan based on naive matching. 

Matched rules include their associated inputs that must be populated with an appropriate stack configuration file.

:::tip
***Matching*** *makes no garauntees that an action will be* ***scheduled***. 

Read on to learn more about [scheduling](#scheduling) in the next section.
:::

## Scheduling

Actions are assigned to activity [Phases](#phases) in a "schedule". 

Scheduling places matched actions in the action plan in the phase defined in the action source.

#### On-Demand scheduling

Understanding on-demand actions is necessary when calculating which actions in the plan are scheduled.

On-Demand is a system managed phase that contains **"Producer Actions"**. Producer actions are run solely for the purpose of creating an output that will be consumed by an action later in the action plan.

Since producer actions are not associated with a finite activity phase, they are dynamically scheduled in the first matching phase that contains an action that requires their output.

If a producer action creates an output that has no schedule action to consume it, then the action will be omitted from the action plan.
