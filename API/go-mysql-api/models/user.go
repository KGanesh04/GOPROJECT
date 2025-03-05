package models

// ✅ User Struct
type User struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}
