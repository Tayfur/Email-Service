package main

import (
	"SendEmail-Service/logic"
	"SendEmail-Service/pkg/rabbitmq"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	WelcomeQueue := "WelcomeQueue"
	WeeklyReportQueue := "WeeklyReportQueue"
	FeatureNotificationQueue := "FeatureNotificationQueue"
	// consumelari dinliyoruz
	go rabbitmq.Consume(WeeklyReportQueue)
	go rabbitmq.Consume(FeatureNotificationQueue)
	go rabbitmq.Consume(WelcomeQueue)


	// Create a new Fiber instance.
	app := fiber.New()
	app.Use(
		logger.New(), // add simple logger
	)
	//bulk maili trigger etmek icin bu endpoint i kullanıyoruz
	app.Get("/FeatureNotification", func(c *fiber.Ctx) error {
		go logic.BulkMail(FeatureNotificationQueue)
		return c.SendString("Feature notification Sended")
	})
	//transactional maili trigger etmek icin bu endpoint i kullanıyoruz
	app.Get("/WelcomeNotification", func(c *fiber.Ctx) error {
		msg := []byte(c.Query("msg"))
		logic.Transactional(msg,WelcomeQueue)
		return c.SendString("Welcome Sended to "+string(msg))
	})
	log.Fatal(app.Listen(":3000"))

}
