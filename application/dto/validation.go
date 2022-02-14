package dto

type BookCreateRuleValidation struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
	// UserID      uint64 `validate:"required"`
}

type BookUpdateRuleValidation struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
}
