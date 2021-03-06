package module1

// Budget stores budget information
type Budget struct {
	max float32
	items []Item
}

// Item stores item information
type Item struct {
	description string
	price float32
}