package api

import "fmt"

// List - Wedo list with associated Items
type List struct {
	ID    int
	items []Item
}

// CreateList - Creates a Wedo List
func (l *List) CreateList() {
	fmt.Println("I am creating a list")
}

// ReadList - Reads a Wedo List
func (l *List) ReadList() {
	fmt.Println("I am reading a list")
}

// UpdateList - Updates a Wedo List
func (l *List) UpdateList() {
	fmt.Println("I am updating a list")
}

// DeleteList - Deletes a Wedo List
func (l *List) DeleteList() {
	fmt.Println("I am deleting a list")
}
