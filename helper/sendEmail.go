package helper

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(email, code string) (bool, error) {
	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_ADDRESS"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Verification Code")
	m.SetBody("text/plain", "Here is your verification code: \n"+code)

	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), emailPort, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}

func SendEmailActivedAccount(email, fullname, code, unique string) error {

	url := fmt.Sprintf("http://localhost:8080/api/v1/accounts/activate?e=%s&un=%s&unique=%s", email, fullname, unique)

	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_ADDRESS"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Activation  Account")
	m.SetBody("text/plain", "Click link to activated your account: \n"+url+"\n \nThis is information about your account for Login after activation: \nEmail: "+email+"\nUsername: "+fullname+"\nPassword: "+code)

	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), emailPort, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendEmailForgotPin(email, username, code, unique string) error {
	url := fmt.Sprintf("https://api.yusharwz.my.id/api/v1/auth/reset-pin?email=%s&username=%s&unique=%s&code=%s", email, username, unique, code)

	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_ADDRESS"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Reset PIN")
	m.SetBody("text/plain", "Click link to create new pin: \n"+url)

	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), emailPort, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
