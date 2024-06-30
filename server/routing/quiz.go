package routing

import (
	"github.com/gofiber/fiber/v2"
    "fmt"

	db "nuxxian/mol-clone/server/database"
	"nuxxian/mol-clone/server/models"
)

func Default(c *fiber.Ctx) error {
	return c.Status(200).
		JSON(fiber.Map{"status": "succes", "message": "Hello Server"})
}
func AddAnswer(c *fiber.Ctx) error {
	req := new(models.QuizEntry)
	if err := c.BodyParser(req); err != nil {
        fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "unsuccesful", "message": "parsing " + err.Error()})
	}

    fmt.Println(req)
	if err := db.AddEntry(*req, c.Params("id")); err != nil {
        fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "unsuccesful", "message": "db err adddrink: " + err.Error()})
	}
	return c.Status(200).
		JSON(fiber.Map{"status": "succes", "message": "entry added"})
}
//
// func GetDrinksUser(c *fiber.Ctx) error {
// 	drinks, err := db.GetDrinksUser()
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).
// 			JSON(fiber.Map{"status": "unsuccesful", "message": "db err adddrink: " + err.Error()})
// 	}
// 	return c.Status(200).
// 		JSON(fiber.Map{"status": "succes", "message": "Users drinks fetched", "data": fiber.Map{"drinks": drinks}})
// }
//
// type updateDrinkRequest struct {
// 	DrinkId   int `json:"drinkId"`
// 	Increment int `json:"increment"`
// }
//
// func UpdateDrink(c *fiber.Ctx) error {
// 	req := new(updateDrinkRequest)
// 	if err := c.BodyParser(req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).
// 			JSON(fiber.Map{"status": "unsuccesful", "message": "parsing " + err.Error()})
// 	}
// 	res, err := db.UpdateDrink(req.DrinkId, req.Increment)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).
// 			JSON(fiber.Map{"status": "unsuccesful", "message": err.Error()})
// 	}
// 	return c.Status(200).
// 		JSON(fiber.Map{"status": "succes", "message": "drink updated", "data": fiber.Map{"quantity": res}})
// }
//
// func GetAllDrinks(c *fiber.Ctx) error {
// 	drinks, err := db.GetAllDrinks(db.TokenUser.InstanceId)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).
// 			JSON(fiber.Map{"status": "unsuccesful", "message": err.Error()})
// 	}
// 	return c.Status(200).
// 		JSON(fiber.Map{"status": "succes", "message": "data follows", "data": fiber.Map{"allDrinks": drinks}})
// }
//
