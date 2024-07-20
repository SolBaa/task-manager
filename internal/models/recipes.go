package models

type Recipe struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	OwnerID     string       `json:"owner_id"`
	Owner       User         `json:"owner"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}
