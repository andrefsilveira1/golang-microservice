package domain

type Item struct {
	Id			uint	`json:"id"`
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Price		float64	`json:"price"`
}