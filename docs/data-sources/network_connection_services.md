---
layout: ""
page_title: "Network Connection Service Data Source"
description: |-
---

# syntropystack_network_connection_services ( Data Source )

Datasource retrieves list of services that were discovered in connection.

## Example Usage
 ```terraform
data "syntropystack_agent" "agent1" {
  name = "syntropy-agent-prod-1"
}

data "syntropystack_agent" "agent2" {
  name = "syntropy-agent-prod-2"
}

resource "syntropystack_network_connection" "p2p" {
  agent_peer  = [data.syntropystack_agent.agent1.id, data.syntropystack_agent.agent2.id]
  sdn_enabled = false
}

data "syntropystack_network_connection_services" "filtered_services" {
  connection_group_id = syntropystack_network_connection.p2p.id
  filter = {
    service_name_substring = "movie-service"
    service_type           = "DOCKER"
  }
}
```

 <!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `connection_group_id` (Number) Unique identifier for the connection.
- `filter` (Attributes) Network connection service filters (see [below for nested schema](#nestedatt--filter))

### Read-Only

- `id` (String) Network connection service ID randomly generated
- `services` (Attributes List) List of services inside in network connection (see [below for nested schema](#nestedatt--services))

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Optional:

- `agent_id` (Number) Filter service list by agent ID
- `service_id` (Number) Filter service list by subnet ID
- `service_name_substring` (String) Filter service list by connection service name substring that is running on agent
- `service_type` (String) Filter service list by connection service type that is running on agent


<a id="nestedatt--services"></a>
### Nested Schema for `services`

Read-Only:

- `agent_id` (Number) Network connection agent ID that service is created
- `enabled` (Boolean) Is network connection service enabled?
- `id` (Number) Network connection service ID
- `ip` (String) Network connection service IP
- `name` (String) Network connection service name
- `type` (String) Network connection service type (Kubernetes, Docker, etc.)


