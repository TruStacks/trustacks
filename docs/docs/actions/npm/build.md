---
title: Build
slug: /actions/npm/build
---

# NPM - Build

The build action creates a package build from the [npm build](https://docs.npmjs.com/cli/v6/commands/npm-build) command.

:::tip
This action will utilize the command provided by the `build` script in the project's [package.json](https://docs.npmjs.com/cli/v10/configuring-npm/package-json).

`npm install` will be run before the build command to ensure that dependencies are installed beforehand.
:::

### Artifacts

#### Outputs:

|Name|Type|Description|
|-|-|-|
|dist|dir|The built application package|
