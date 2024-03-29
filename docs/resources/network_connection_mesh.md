---
layout: ""
page_title: "Network Connection Mesh"
description: |-
---
# syntropystack_network_connection_mesh ( Resource )

Creates [Mesh Connections](https://docs.syntropystack.com/docs/network-as-code-topologies#creating-complex-topologies).
To establish a mesh connection between any number of agents, you need to provide an array of `agent IDs` (minimum of two).

## Example Usage
 ```terraform
data "syntropystack_agent_search" "results" {
  filter = {
    type             = ["LINUX"]
    location_country = ["US", "UK"]
  }
}

resource "syntropystack_network_connection_mesh" "test_connection_mesh" {
  agent_ids   = data.syntropystack_agent_search.results.agents.*.id
  sdn_enabled = true
}
```

 <!-- schema generated by tfplugindocs -->
## Schema

### Required

- `agent_ids` (Set of Number) List of agent IDs for network connection mesh

### Optional

- `sdn_enabled` (Boolean) Should SDN be enabled?

### Read-Only

- `connections` (Attributes List) List of network connections created by mesh resource (see [below for nested schema](#nestedatt--connections))
- `id` (String) Network connection mesh ID randomly generated

<a id="nestedatt--connections"></a>
### Nested Schema for `connections`

Read-Only:

- `agent_1_id` (Number) Agent 1 ID
- `agent_2_id` (Number) Agent 2 ID
- `connection_group_id` (Number) Unique identifier for the connection
- `services` (Attributes List) List of services inside in network connection mesh (see [below for nested schema](#nestedatt--connections--services))

<a id="nestedatt--connections--services"></a>
### Nested Schema for `connections.services`

Read-Only:

- `agent_id` (Number) Network connection agent ID that service is created
- `enabled` (Boolean) Is network connection service enabled?
- `id` (Number) Network connection service ID
- `ip` (String) Network connection service IP
- `name` (String) Network connection service name
- `type` (String) Network connection service type (Kubernetes, Docker, etc.)



