---
sidebar_position: 4
slug: /tutorial/input
title: Input
---

# Input

### What are inputs?

Inputs are parameters that are required by actions in an action plan. Inputs can be viewed by inspecting the contents of the `trustacks.plan`.

ex. `{"actions":[...],"inputs":["CONTAINER_REGISTRY", ...]}`

### Generating A Configu Schema (Optional)

TruStacks supports [Configu](https://configu.com/) schemas for managing action inputs. 

To generate an input schema run the following command:

```bash
tsctl config init --from-plan trustacks.plan
```

This command generates a [ConfigSchema](https://configu.com/docs/config-schema/) in the current directory named `trustacks.cfgu.json`.

Follow [this](https://configu.com/docs/get-started/) guide to configure your variables using your desired [ConfigStore](https://configu.com/docs/config-store/). 

:::tip
The `trustacks.cfgu.json` schema file contains a description of the inputs for reference.
:::

:::tip
Make sure to use the name `trustacks` for the desired ConfigStore in your .configu file like so:

```json
{
  "stores": {
    "trustacks": {
      "type": "<config-store-type>",
      "configuration": { 
        ... 
      }
    }
  }
}

```
:::

### Exporting Inputs

Once you have [upserted](https://configu.com/docs/commands/#upsert) your config to your config store, run the following command to export your config as environment variables.

```bash
for e in $(configu eval --store 'trustacks' --set '<config-set>' --schema ./trustacks.cfgu.json | configu export --format 'Dotenv'); do export $e; done
```

This command will validate your configuration against the input schema and export each variable into the environment.