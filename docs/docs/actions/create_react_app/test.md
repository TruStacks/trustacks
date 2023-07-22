---
title: Test
slug: /actions/create-react-app/test
---

# Create React App - Test

The test action runs [Jest Tests](https://create-react-app.dev/docs/running-tests) for a [Create React App](https://create-react-app.dev) (CRA) application.

# Admission Criteria:

- A `package.json` configuration file exists in the root of the source directory.
- `package.json` contains a `scripts` key, which contains the `build` sub-key, with a value of `react-scripts build`.

### Artifacts

#### Inputs:

N/A

#### Outputs:

N/A

#### Reports:

|Name|Description|
|-|-|
test-report.html|An html report generated with [jest-html-reporter](https://www.npmjs.com/package/jest-html-reporter)|
coverage|The unit test coverage report|