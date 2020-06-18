variable "resource_group_name" {
    type = "string"
}

variable "resource_group_location" {
    type = "string"
    description = "Azure region that resource group will be deployed to."
}

variable "failover_location" {
    type = "string"
    description = "The name of the Azure region to host replicated data."
}