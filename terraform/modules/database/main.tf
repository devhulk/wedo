
resource "random_integer" "ri" {
  min = 10000
  max = 99999
}

resource "azurerm_cosmosdb_account" "db" {
  name                = "wedo-db-${random_integer.ri.result}"
  location            = var.resource_group_location
  resource_group_name = var.resource_group_name
  offer_type          = "Standard"
  kind                = "MongoDB"

  enable_automatic_failover = true

  consistency_policy {
    consistency_level       = "BoundedStaleness"
    max_interval_in_seconds = 301
    max_staleness_prefix    = 100001
  }

  geo_location {
    location          = var.failover_location
    failover_priority = 1
  }

  geo_location {
    prefix            = "wedo-db-${random_integer.ri.result}-customid"
    location          = var.resource_group_location
    failover_priority = 0
  }

  tags = {
      Key = "DoNotDelete"
  }
}