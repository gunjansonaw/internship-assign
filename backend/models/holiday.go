package models

// Holiday represents a holiday with its ID, name, date, month, and year
type Holiday struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date int    `json:"date"` // Date in the month
	Month int   `json:"month"`
	Year  int   `json:"year"`
}
