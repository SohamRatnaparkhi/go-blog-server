package utils

import (
	"fmt"
	"os"

	"github.com/go-mail/mail"
)

func SendMail(to string, emailText string, subject string) {
	message := mail.NewMessage()
	// fmt.Println(os.Getenv("EMAIL_ADDRESS") + " - " + os.Getenv("EMAIL_ADDRESS_PASSWORD"))
	message.SetHeader("From", os.Getenv("EMAIL_ADDRESS"))
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", emailText)
	fmt.Println(message)
	dial := mail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_ADDRESS_PASSWORD"))
	if err := dial.DialAndSend(message); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Email sent")
	}
}
