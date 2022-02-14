package controller

import (
	"strconv"

	"github.com/MochJuang/golang_book/application/dto"
	helper "github.com/MochJuang/golang_book/application/helpers"
	"github.com/MochJuang/golang_book/application/models"
	service "github.com/MochJuang/golang_book/application/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BookController interface {
	All(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Insert(c *fiber.Ctx) error
	// Update(c *fiber.Ctx) error
	// Delete(c *fiber.Ctx) error
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(bs service.BookService) BookController {
	return &bookController{
		bookService: bs,
	}
}

func (b *bookController) All(c *fiber.Ctx) error {
	books, err := b.bookService.All()
	if err != nil {
		res := helper.BuildErrorResponse("Error Reponse", err.Error(), helper.EmptyObj{})
		return c.JSON(res)
	} else {
		res := helper.BuildResponse(true, "Success Get Books", books)
		return c.JSON(res)
	}

}
func (b *bookController) FindById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		return c.JSON(res)
	}
	book, err := b.bookService.FindByID(id)
	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), helper.EmptyObj{})
		return c.JSON(res)
	} else {
		res := helper.BuildResponse(true, "Success Get Book", book)
		return c.JSON(res)
	}

}

var validate *validator.Validate

func (b *bookController) Insert(c *fiber.Ctx) error {
	payload := struct {
		Title       string `json:"title" form:"title"`
		Description string `json:"description" form:"description"`
		// UserID      uint64 `josn:"user_id"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		res := helper.BuildErrorResponse("Error Parse Body", err.Error(), helper.EmptyObj{})
		return c.JSON(res)
	}

	validate = validator.New()
	bookData := dto.BookCreateRuleValidation{
		Title:       payload.Title,
		Description: payload.Description,
	}

	errValidate := validate.Struct(bookData)

	if errValidate != nil {
		res := helper.BuildErrorResponse("Error Parse Body", errValidate.Error(), helper.EmptyObj{})
		return c.JSON(res)
	}

	book, err := b.bookService.Insert(bookData)

	if err != nil {
		res := helper.BuildErrorResponse("Error Insert Data", err.Error(), helper.EmptyObj{})
		return c.JSON(res)
	} else {
		res := helper.BuildResponse(true, "Success Get Book", book)
		return c.JSON(res)
	}

}

func (b *bookController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		return c.JSON(res)
	}

	payload := struct {
		Title       string `json:"title" form:"title"`
		Description string `json:"description" form:"description"`
		// UserID      uint64 `josn:"user_id"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		res := helper.BuildErrorResponse("Error Validation", err.Error(), helper.EmptyObj{})
		return c.JSON(res)
	}

	validate = validator.New()
	bookData := dto.BookUpdateRuleValidation{
		Title:       payload.Title,
		Description: payload.Description,
	}

	errValidate := validate.Struct(bookData)

	if errValidate != nil {
		res := helper.BuildErrorResponse("Error Parse Body", errValidate.Error(), helper.EmptyObj{})
		return c.JSON(res)
	}

	book, err := b.bookService.Update(id, bookData)

	if err != nil {
		res := helper.BuildErrorResponse("Error Insert Data", err.Error(), helper.EmptyObj{})
		return c.JSON(res)
	} else {
		res := helper.BuildResponse(true, "Success Get Book", book)
		return c.JSON(res)
	}
}

func (b *bookController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		return c.JSON(res)
	}
	book, err := b.bookService.Delete(models.Book{ID: id})
	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), helper.EmptyObj{})
		return c.JSON(res)
	} else {
		res := helper.BuildResponse(true, "Success Delete Book", book)
		return c.JSON(res)
	}
}
