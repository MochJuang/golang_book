package repository

import "github.com/MochJuang/golang_book/application/models"

type UserRepository interface {
	Insert(u models.User) (models.User, error)
	Update(id uint64, u models.User) (models.User, error)
}
