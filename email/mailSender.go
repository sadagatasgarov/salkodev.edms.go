package email

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

// Змінна оточення - системна адреса (email) від імені якої працює сервіс (робот)
const SystemEmailAddressEnv = "SALKODEV_EDMS_SYSTEM_EMAIL"

// Пасс до системн.адреси
const SystemEmailPswEnv = "SALKODEV_EDMS_SYSTEM_EMAIL_PSW"

// Хост SMTP
const SystemEmailSMTPHostEnv = "SALKODEV_EDMS_SYSTEM_EMAIL_SMTP_HOST"

// Порт SMTP
const SystemEmailSmtpPortEnv = "SALKODEV_EDMS_SYSTEM_EMAIL_SMTP_PORT"

func SendMail(toEmail string, subject string, body string) {

	emailAddr := os.Getenv(SystemEmailAddressEnv)
	emailPass := os.Getenv(SystemEmailPswEnv)
	smtpHost := os.Getenv(SystemEmailSMTPHostEnv)
	smtpPortStr := os.Getenv(SystemEmailSmtpPortEnv)

	smtpPort, errConv := strconv.Atoi(smtpPortStr)

	if errConv != nil {
		fmt.Println("Error during conversion smtpPortStr:", smtpPortStr, errConv.Error())
		return
	}

	message := gomail.NewMessage()

	message.SetHeader("From", emailAddr)
	message.SetHeader("To", toEmail)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	d := gomail.NewDialer(smtpHost, smtpPort, emailAddr, emailPass)

	//if SSL/TLS certificate is not valid on server (in prod.set to false)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	//Send
	err := d.DialAndSend(message)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
