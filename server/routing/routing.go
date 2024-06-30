package routing

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func routerHandlers(router *fiber.App) {
 //    drinks := router.Group("/drinks", AuthMiddleware())
	// drinks.Get("/get", GetDrinksUser)
	// drinks.Get("/getAll", GetAllDrinks)
	// drinks.Get("/add/:name", AddDrink)
	// drinks.Post("/update", UpdateDrink)
    router.Post("/answer/:id", AddAnswer)
    router.Get("/answers/:id", GetAnswers)
    router.Get("/", Default)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Listen(":" + port))
}

func Route() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3001, http://localhost:3000, https://nuxx.be, https://demol-clone-dashboard.web.app/, https://demol-clone.web.app/, https://demol-dashboard.nuxx.be/, https://demol.nuxx.be/",
		AllowMethods: "POST GET DELETE PUT",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	routerHandlers(app)
}

