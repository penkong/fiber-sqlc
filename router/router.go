package apirouters

import (
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Repo *pgdb.Repo
}

func NewHandlers(repo *pgdb.Repo) *Handlers {
	return &Handlers{Repo: repo}
}

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")
	SetUpAuth(v1)
	SetUpUser(v1)
}
