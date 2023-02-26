package service

import (
	"context"
	"email_auth/initer"
	"fmt"
	"github.com/wneessen/go-mail"
)

func CreateNewMail(targetAddr string, code string) {
	mailMessage := mail.NewMsg()

	mailMessage.From(initer.AppConfig.SmtpUser)
	mailMessage.To(targetAddr)
	mailMessage.Subject("You got a verification code")
	mailMessage.SetBodyString(mail.TypeTextPlain, fmt.Sprintf("code: %s", code))

	mailClient, err := mail.NewClient(initer.AppConfig.SmtpHost, mail.WithPort(initer.AppConfig.SmtpPort), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithSSL(), mail.WithUsername(initer.AppConfig.SmtpUser), mail.WithPassword(initer.AppConfig.SmtpPass))
	if err != nil {
		fmt.Println(err)
	}

	mailClient.DialWithContext(context.TODO())

	if err := mailClient.Send(mailMessage); err != nil {
		fmt.Println(err)
	}

}
