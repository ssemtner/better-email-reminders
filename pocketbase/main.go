package main

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"

	"github.com/robfig/cron/v3"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func testEmail() {
	from := mail.NewEmail("Better Email Reminders", "sjsemtner@gmail.com")
	subject := "Testing"
	to := mail.NewEmail("Jimmy", "ssemtner2023@hightechhigh.org")

	message := mail.NewSingleEmail(from, subject, to, "Hello", "<strong>Hello World</strong>")
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}

func main() {
	// testEmail()

	c := cron.New()
	c.AddFunc("*/10 * * * *", func() {
		log.Println("cron: every 10 seconds")
	})
	c.Start()

	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
