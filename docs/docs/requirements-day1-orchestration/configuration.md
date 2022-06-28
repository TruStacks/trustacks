# Configuration

All components must be configured for Day 1 consumption. Day 1 orchestration requirements are as follows:


1. The component is configured with OIDC or SAML authentication with the integrated SSO solution.
2. The component defines viewer, editor, and admin group roles.
3. If the component exposes service endpoint (ie. HTTP API or CLI API) for CI execution tasks then a service account must be configured with credentials stored in the secrets manager.
4. Any best practices or sane default should be included in the component orchestration (ie. configuring password policy and MFA for the SSO component).