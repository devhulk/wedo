package api

// ListItem - a list item with associated Details
type ListItem struct {
	ID   int
	Name string
}

// CreateItem - Creates a Wedo Item
func (l *ListItem) CreateItem() {
	fmt.Println("I am creating a item")
}

// ReadItem - Reads a Wedo Item
func (l *ListItem) ReadItem() {
	fmt.Println("I am reading a item")
}

// UpdateItem - Updates a Wedo Item
func (l *ListItem) UpdateItem() {
	fmt.Println("I am updating a item")
}

// DeleteItem - Deletes a Wedo Item
func (l *ListItem) DeleteItem() {
	fmt.Println("I am deleting a item")
}
