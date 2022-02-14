package controller

import (
	helper "github.com/MochJuang/golang_book/application/helpers"
	service "github.com/MochJuang/golang_book/application/services"
	"github.com/gofiber/fiber/v2"
)

type BookController interface {
	All(c *fiber.Ctx)
	FindById(c *fiber.Ctx)
	Insert(c *fiber.Ctx)
	Update(c *fiber.Ctx)
	Delete(c *fiber.Ctx)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(bs service.BookService) BookController {
	return &bookController{
		bookService: bs,
	}
}

func (b *bookController) All(c *fiber.Ctx) {
	books, err := b.bookService.All()
	if err != nil {
		res := helper.BuildErrorResponse(err)
		c.JSON()
	}

}
func (b *bookController) FindById(c *fiber.Ctx) {

}
func (b *bookController) Insert(c *fiber.Ctx) {

}
func (b *bookController) Update(c *fiber.Ctx) {

}
func (b *bookController) Delete(c *fiber.Ctx) {

}
