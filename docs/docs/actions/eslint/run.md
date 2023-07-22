---
title: Run
slug: /actions/eslint/run
---

# ESLint - Run

The run action runs the [ESLint](https://eslint.org/) linter against the javascript source code.

# Admission Criteria:

1. A `package.json` configuration file exists in the root of the source directory.
2. An ESLint configuration file exists in the root of the source directory with one of  `.eslintrc.js`, `.eslintrc.cjs`, `.eslintrc.yaml`, `.eslintrc.yml`, `.eslintrc.json`, `package.json` as the filename.
3. A compatible ESlint configuration exists and the configuration is not `package.json`, or `package.json` contains the `eslintConfig` key.

### Artifacts

#### Inputs:

N/A

#### Outputs:

|Name|Type|Description|
|-|-|-|
|dist|dir|The production react application build distribution|


#### Reports:

N/A