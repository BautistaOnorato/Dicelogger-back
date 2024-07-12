package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEmailVerificationLink(userEmail string, emailVerificationLink string) error {
	sender := "diceloggerdnd@gmail.com"
	recipient := userEmail
	smtpHost := "smtp.gmail.com"
	password := "exnf fyeg vsfw koda"
	smtpPort := 587
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "Verify your email")
	m.SetBody("text/html", fmt.Sprintf("<h1>Bienvenido a DiceLogger!</h1><p>Para verificar tu email haz click en el siguiente : <a href=\"%s\">LINK</a> </p>\r\n\r\n", emailVerificationLink))
	d := gomail.NewDialer(smtpHost, smtpPort, sender, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	fmt.Println("Email sent successfully")
	return nil

}
