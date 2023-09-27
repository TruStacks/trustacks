---
title: Test
slug: /actions/npm/test
---

# Npm - Test

The test action runs the test suite using the [npm test](https://docs.npmjs.com/cli/v6/commands/npm-test) command.

:::tip
This action will utilize the command provided by the `test` script in the project's [package.json](https://docs.npmjs.com/cli/v10/configuring-npm/package-json).

`npm install` will be run before the test command to ensure that dependencies are installed beforehand.
:::