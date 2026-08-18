package models

type Expense struct {
	Description string  `csv:"description"`
	Amount      float64 `csv:"amount"`
	CreatedAt   string  `csv:"created_at"`
}
