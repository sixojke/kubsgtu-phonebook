package types

type Contact struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Position   string `json:"position" db:"position"`
	Phone      string `json:"phone" db:"phone"`
	Email      string `json:"email" db:"email"`
	Room       string `json:"room" db:"room"`
	Department string `json:"department" db:"department"`
}
