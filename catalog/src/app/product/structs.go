package product

type Product struct {
	ID int32 `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	PriceInCents int32 `json:"price_in_cents" db:"price_in_cents"`
}
