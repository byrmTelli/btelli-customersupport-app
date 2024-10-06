package models

type ComplaintDTO struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CustomerID  uint   `json:"customer_id"`
	CategoryID  uint   `json:"category_id"`
}
