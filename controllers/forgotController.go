package controllers

import (
	"go-lang-react/database"
	"go-lang-react/models"
	"math/rand"

	"net/smtp"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	token := RandStringRunes(12)
	PasswordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}
	database.DB.Create(&PasswordReset)
	from := "admin@example.com"
	to := []string{
		data["email"],
	}
	url := "http://localhost:3000/reset/" + token
	message := []byte("click <a href=\"" + url + "\">here</a>to reset your password!")
	err := smtp.SendMail("0.0.0.0:1025", nil, from, to, message)

	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})

}
func Reset(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"]!=data["password_confirm"]{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Password do not match!",
		})
	}
	var passwordReset = models.PasswordReset{}
	if err := database.DB.Where("token =?",data["token"]).Last(&passwordReset); err.Error !=nil{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"invalid token!",
		})
	}
	password,_:= bcrypt.GenerateFromPassword([]byte(data["password"]),14)
	database.DB.Model(&models.User{}).Where("email =?",passwordReset.Email).Update("password",password)
	return c.JSON(fiber.Map{
		"message":"success",
	})
}
func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
