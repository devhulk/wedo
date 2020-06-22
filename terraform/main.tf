provider "azurerm" {
  version = "=2.0.0"
  features {}
}

resource "azurerm_resource_group" "wedo-dev" {
  name     = var.resource_group_name
  location = var.resource_group_location
  tags = {
      Key = "DoNotDelete"
      Value = "True"
  }
}

module "database" {
  source = "./modules/database"
  resource_group_name = var.resource_group_name
  resource_group_location = var.resource_group_location
  failover_location = "West US"
}
