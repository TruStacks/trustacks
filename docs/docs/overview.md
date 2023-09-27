---
title: Overview
sidebar_position: 3
---

## What is TruStacks?

TruStacks is a generative software delivery engine that removes the need to build pipelines.

### Generative Sofware Delivery

*Intermediary Wiring* is the pipeline code that sits between the source and the delivered product that is typically implemented with yaml or other no-code DSL's.

TruStacks uses a rule engine to generate actions plans that contain actions based on discovered facts in the application sources.

### Generation Flow

TruStacks uses the following flow to generate and execute action plans.

![Sonar Create Project](./assets/overview-diagram.svg)

#### Facts

Fact collection is the first step in the generation flow. During fact collection the engine sets attributes about sources such as language, frameworks, and granular facts such as multi-stage docker builds or test script discovery.

#### Rules & Actions

After collecting facts the engine applies matching rules against the fact set. If rules are matched then the appropriate actions will be admitted into the action plan. Actions contain common steps in a CI/CD pipeline such as linting or unit testing.

#### Action Plan

The action plan contains the list of matched actions and their associated inputs. Inputs are parameters and credentials that exists outside of the application source. 

Inputs must be populated before executing the action plan.

#### Schedule

Actions admitted into an action plan are naive with no specific order. The scheduler places rules in appropriate order based on action classification and artifacts. 

Rules can be generally classified in a fixed stage, or selected for execution in a stage at runtime by the scheduler as "feeder" actions. Feeder actions exist only to provide inputs to a downstream action such as a container build step that "feeds" the output image to a vulnerability scan or image publish step.

The scheduler ensures that actions between stages and inner stage are executed in the order of the required inputs. If no input is required by a given action the scheduler will run the action at whatever point it is introduced into the schedule.