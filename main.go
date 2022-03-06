package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@123@/goreact"), &gorm.Config{})
	if err != nil {
		panic("error db connetion")
	}
	fmt.Println(db)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
