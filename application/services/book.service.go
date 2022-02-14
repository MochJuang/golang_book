package service

import (
	"errors"

	"github.com/MochJuang/golang_book/application/constants"
	"github.com/MochJuang/golang_book/application/dto"
	"github.com/MochJuang/golang_book/application/models"
	repository "github.com/MochJuang/golang_book/application/repositories"
	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(b dto.BookCreateRuleValidation) (models.Book, error)
	Update(id uint64, b dto.BookUpdateRuleValidation) (models.Book, error)
	Delete(b models.Book) (bool, error)
	All() ([]models.Book, error)
	FindByID(bookID uint64) (models.Book, error)
	// IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

//NewBookService .....
func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (bs *bookService) Insert(bd dto.BookCreateRuleValidation) (models.Book, error) {
	var book models.Book

	err := smapping.FillStruct(&book, smapping.MapFields(&bd))

	if err != nil {
		return book, errors.New(constants.ErrorAction)
	}
	res, err := bs.bookRepository.Insert(book)

	if err != nil {
		return book, errors.New(err.Error())
	}
	return res, nil
}

func (bs *bookService) Update(id uint64, bd dto.BookUpdateRuleValidation) (models.Book, error) {
	var book models.Book
	err := smapping.FillStruct(&book, smapping.MapFields(&bd))

	if err != nil {
		return book, errors.New(constants.ErrorAction)
	}
	res, err := bs.bookRepository.Update(id, book)

	if err != nil {
		return book, errors.New(err.Error())
	}
	return res, nil
}

func (bs *bookService) Delete(book models.Book) (bool, error) {
	res, err := bs.bookRepository.Delete(book)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (bs *bookService) FindByID(id uint64) (models.Book, error) {
	res, err := bs.bookRepository.FindById(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (bs *bookService) All() ([]models.Book, error) {
	res, err := bs.bookRepository.GetAll()
	if err != nil {
		return res, err
	}
	return res, nil
}
