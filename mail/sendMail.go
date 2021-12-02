package mail

import (
	"SendEmail-Service/pkg/config"
	"SendEmail-Service/pkg/parseHtml"
	"fmt"
	"net/smtp"
)
// maili gonderecek olan maile giris yapiyoruz
func auth(senderMail string, senderPassword string, smtpHost string) smtp.Auth {
	auth := smtp.PlainAuth("", senderMail, senderPassword, smtpHost)
	return auth

}
func SendMail(to []string, tamplate string) error {
	senderMail := config.Config("SENDER_MAIL")
	password := config.Config("SENDER_PASSWORD")
	smtpHost := config.Config("SMTP_HOST")
	smtpPort := config.Config("SMTP_PORT")
	var body string
	auth := auth(senderMail, password, smtpHost)

	//gonderilmek istenilen template gore parse ediyoruz
	switch tamplate {
	case "WelcomeQueue":
		body = parseHtml.WelcomeTemplate()
		break
	case "WeeklyReportQueue":
		body = parseHtml.WeeklyReportTemplate()
		break
	case "FeatureNotificationQueue":
		body = parseHtml.FeatureNotificationTemplate()
		break
	}
	// parse edilmis html templati ile maili gonderiyoruz
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderMail, to, []byte(body))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}
