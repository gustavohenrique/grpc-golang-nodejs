package customer

import "time"

type Customer struct {
	ID int32 `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	BirthDate time.Time `json:"birth_date" db:"birth_date"`
}
