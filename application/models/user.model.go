package models

type User struct {
	ID       uint64  `gorm:"primaryKey:autoIncrement" json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Token    string  `gorm:"-" json:"token,omitempty"`
	Books    []*Book `json:"books,omitempty"`
}
