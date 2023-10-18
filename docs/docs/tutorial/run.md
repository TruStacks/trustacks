---
title: Run
sidebar_position: 4
slug: /tutorial/run
---

# Run

## Running The Plan

Run the action plan using the following command:

```
tsctl run --stages feedback
```

:::tip
Removed the `--stages feedback` option to the run the entire pipeline if you have configured all necessary inputs in the previous step.
:::

This command will orchestrate the action plan into a runnable pipeline and runs them using [Dagger](https://dagger.io/)

## Troubleshooting

