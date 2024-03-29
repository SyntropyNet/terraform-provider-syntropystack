---
layout: ""
page_title: "Provider: Syntropy Stack"
description: |-
---

# Syntropy Stack Provider

The Syntropy Stack provider is used to configure your Syntropy Stack Platform infrastructure. See the [Getting Started page](https://docs.syntropystack.com/docs/getting-started) for more details and for to learn how you can use Syntropy Agents and connect them securely.

## Example Usage

  ```terraform
terraform {
  required_providers {
    syntropystack = {
      source  = "SyntropyNet/syntropystack"
      version = "~> 0.1"
    }
  }
}

provider "syntropystack" {
  access_token = "<ACCESS_TOKEN>"
}
```

## Additional Info

If you have configuration questions, or general questions about using the provider, try checking out:
- Syntropy Stack Devs community on Discord (`#stack-support` channel)
- Syntropy [Developer Portal](https://docs.syntropystack.com/)

## Releases
Interested in the provider's latest features, or want to make sure you're up to date? Check out the most recent releases [Releases · SyntropyNet/terraform-provider-syntropystack · GitHub](https://github.com/SyntropyNet/terraform-provider-syntropystack/releases)  for release notes and additional information.

## Feature and Bug Requests
At Syntropy we aim to provide the highest quality code to serve our developers, so we prepared a number of ways to report a bug or request a feature:
- on Github repo in [Issues](https://github.com/SyntropyNet/terraform-provider-syntropystack/issues) tab
- via [Discord](https://discord.gg/UYDyHwk5gN) community channel dedicated for `#stack-support` 
- via [Customer Portal](https://syntropy.atlassian.net/servicedesk/customer/portals)

 <!-- schema generated by tfplugindocs -->
## Schema

### Required

- `access_token` (String) Syntropy platform access token

### Optional

- `api_url` (String) Syntropy platform API URL

