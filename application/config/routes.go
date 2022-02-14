package config

import (
	controller "github.com/MochJuang/golang_book/application/controllers"
	repository "github.com/MochJuang/golang_book/application/repositories"
	service "github.com/MochJuang/golang_book/application/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = SetupDatabaseConnection()
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	bookService    service.BookService       = service.NewBookService(bookRepository)
	bookController controller.BookController = controller.NewBookController(bookService)
)

func Routes(app *fiber.App) {
	book := app.Group("/book")

	book.Get("/", bookController.All)
	book.Get("/:id", bookController.FindById)
	book.Post("/", bookController.Insert)
	book.Put("/:id", bookController.Update)
	book.Delete("/:id", bookController.Delete)

}
