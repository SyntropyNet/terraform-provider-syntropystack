---
layout: ""
page_title: "Agent"
description: |-
---

# syntropystack_agent ( Resource )

Syntropy Agent is an easy-to-use dependency to automatically encrypt and connect endpoints within a network. By using Terraform, you have an ability to create a `virtual agent`. Syntropy's user interface allows heightened visibility into network health and connectivity.
For more information about what else you can do with our *agent*, feel free to checkout out our [documentation](https://docs.syntropystack.com/docs/what-is-syntropy-agent).

## Example Usage
 ```terraform
resource "syntropystack_agent" "agent" {
  name  = "terraform-provider-syntropystack-agent"
  token = "<AGENT_TOKEN>"
}
```

 <!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Agent name
- `token` (String) Agent token

### Optional

- `tags` (Set of String) Agent tags

### Read-Only

- `id` (Number) Agent ID



## How to generate *Agent Token*?

First things first - to start using Syntropy Agent you need to set up an Agent token. Head to User section to create one.
Click on New Agent Token and create one by adding a name and its expiration date.
More details can be found in [documentation](https://docs.syntropystack.com/docs/get-your-agent-token).