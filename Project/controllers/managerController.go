package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"root/database"
	"root/models"
)

func ManagerCreateFood(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//get cookie
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	//user is not login
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	//claims to pointer of standard claims
	claims := token.Claims.(*jwt.StandardClaims)

	//var manager models.Manager

	restaurant := models.Restaurant{
		Food: []models.Food{
			models.Food{
				FoodName: data["foodName"],
				Price:    data["price"],
			},
		},
	}

	database.DB.Where("id = ?", claims.Issuer).Model(&restaurant).Update("food", restaurant.Food)
	database.DB.Create(&restaurant.Food)
	return c.JSON(restaurant)
}

func ManagerDeleteFood(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func ManagerUpdateFood(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func ManagerComment(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}
