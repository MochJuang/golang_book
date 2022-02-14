package models

type Book struct {
	ID          uint64 `gorm:"primaryKey:autoIncrement" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// UserID      uint64
	// User        User `gorm:"constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
