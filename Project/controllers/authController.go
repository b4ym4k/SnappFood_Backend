package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"root/database"
	"root/models"
	"strconv"
	"time"
)

const SecretKey = "secret"

//==================Manager=====================

func ManagerRegister(c *fiber.Ctx) error {
	var data map[string]string

	//err:=c.BodyParser(&data)
	//if err != nil {
	//	return err
	//}
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	manager := models.Manager{
		Name:  data["name"],
		Email: data["email"],
		//Region:   data["region"],
		Address:  data["address"],
		Password: password, //data["password"],
	}

	database.DB.Create(&manager)

	return c.JSON(manager)
}

func ManagerLogin(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var manager models.Manager
	database.DB.Where("email = ?", data["email"]).First(&manager)

	//user doesn't exist
	if manager.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//if user exist -> user should go to pass check
	if err := bcrypt.CompareHashAndPassword(manager.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	//return jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(manager.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	//cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	//successfully entered email & pass
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func ManagerUser(c *fiber.Ctx) error {
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

	//claims := token.Claims
	//claims to pointer of standard claims
	claims := token.Claims.(*jwt.StandardClaims)

	var manager models.Manager

	database.DB.Where("id = ?", claims.Issuer).First(&manager)

	return c.JSON(manager)
	//return c.JSON(claims)
}

func ManagerLogout(c *fiber.Ctx) error {
	//just remove cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

//==================User=====================

func UserRegister(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:        data["name"],
		PhoneNumber: data["phoneNumber"],
		Address:     data["address"],
		//Credit:      int["Credit"],
		Password: password, //data["password"],
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func UserLogin(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user models.User
	database.DB.Where("phone_number = ?", data["phoneNumber"]).First(&user)

	//user doesn't exist
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//if user exist -> user should go to pass check
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	//return jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	//cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	//successfully entered email & pass
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UserUser(c *fiber.Ctx) error {
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

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
	//return c.JSON(claims)
}

func UserLogout(c *fiber.Ctx) error {
	//just remove cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
