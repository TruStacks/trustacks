---
slug: /config/python
title: Python
---

# Python Configuration

Table: `python`

|Name|Type|Description|Example|Native|
|-|-|-|-|-|
|version|string|the python version|"3.10"|✅|
|libs|array|C libraries that are required to install python dependencies|["libxmlsec1-dev"]|❌|
|dev_reqs|string|the path to a development requirements file|"dev.txt"|❌|

Usage Example: 

```toml
[python]
version = "3.10"
libs = ["libxmlsec1-dev"]
```