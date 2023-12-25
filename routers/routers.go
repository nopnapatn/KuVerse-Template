// routers/routers.go
package routers

import (
	"kuverse/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	route := app.Group("/api/v1")
	report := route.Group("/report")
	report.Get("/", controllers.GetAllReport)

}
