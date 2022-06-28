# Hooks

Components should be deployed via Helm. In order to support the myriad of deviations between components the component developer should take advantage of [Helm hooks](https://helm.sh/docs/topics/charts_hooks/#the-available-hooks).

Helm hooks enable the developer to inject custom scripts into helm actions to manage the component.


## Requirements

It is the responsibility of the developer to ensure the Day 1 readiness of their component. The TSSF has the following minimum requirements:

- Prerequisite tasks such as the creation of OIDC clients or one time execution tasks must use the pre-release hook.
- All API calls that configure the internals of the underlying component application must use the post-installation hook.
- All components backed by a persistent datastore must run a backup operation using the pre-upgrade hook. The backup must be pushed to the TSSF bucket storage for DR purposes.
- All artifacts created in the TSSF must be cleaned up in the post-delete hook.