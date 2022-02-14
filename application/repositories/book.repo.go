package repository

import (
	"errors"

	"github.com/MochJuang/golang_book/application/constants"
	"github.com/MochJuang/golang_book/application/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Insert(b models.Book) (models.Book, error)
	Update(id uint64, b models.Book) (models.Book, error)
	Delete(b models.Book) (bool, error)
	FindById(id uint64) (models.Book, error)
	GetAll() ([]models.Book, error)
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(dbCon *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbCon,
	}
}

func (db *bookConnection) Insert(b models.Book) (models.Book, error) {
	result := db.connection.Create(&b)
	if result.Error != nil {
		return b, errors.New(constants.ErrorAction)
	}
	db.connection.Find(&b)

	return b, nil
}

func (db *bookConnection) Update(id uint64, b models.Book) (models.Book, error) {
	result := db.connection.Model(&b).Where("id = ?", id).Updates(&b)
	if result.Error != nil {
		return b, errors.New(constants.ErrorAction)
	}
	db.connection.Find(&b)

	return b, nil
}

func (db *bookConnection) Delete(b models.Book) (bool, error) {
	result := db.connection.Delete(&b)
	if result.Error != nil {
		return false, errors.New(constants.ErrorAction)
	}
	return true, nil
}
func (db *bookConnection) FindById(id uint64) (models.Book, error) {
	var book models.Book
	result := db.connection.First(&book, "where id = ?", id)
	if result.Error != nil {
		return book, errors.New(constants.ErrorAction)
	}
	return book, nil
}
func (db *bookConnection) GetAll() ([]models.Book, error) {
	var book []models.Book
	result := db.connection.Find(&book)
	if result.Error != nil {
		return book, errors.New(constants.ErrorAction)
	}
	return book, nil
}
