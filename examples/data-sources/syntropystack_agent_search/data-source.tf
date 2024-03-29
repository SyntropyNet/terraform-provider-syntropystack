data "syntropystack_agent_search" "agent_1" {
  take = 1
  filter = {
    type             = ["LINUX"]
    location_country = ["US", "UK"]
  }
}

data "syntropystack_agent_search" "agent_2" {
  take = 1
  filter = {
    type   = ["LINUX"]
    status = ["CONNECTED"]
  }
}