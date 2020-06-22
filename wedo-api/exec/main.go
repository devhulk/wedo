package main

import (
	"fmt"

	"wedo.io/api"
)

func main() {
	wedo := api.List{
		ID: 1,
		Items: []api.ListItem{
			api.Item{
				ID:   1,
				Name: "This is wedo item 1",
			},
		},
	}

	wedo.CreateList()
}
