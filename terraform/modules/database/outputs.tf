output "cosmos_endpoint" {
  value = azurerm_cosmosdb_account.db.endpoint
}

output "cosmos_connection_strings" {
  value = azurerm_cosmosdb_account.db.connection_strings
}