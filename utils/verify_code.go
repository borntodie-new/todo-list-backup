package utils

import (
	"bytes"
	"github.com/borntodie-new/todo-list-backup/config"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

type EmailCodeService struct {
	To      string `json:"to"`
	Code    string `json:"code"`
	Subject string `json:"subject"`
}

func NewEmailCodeService(to, code, subject string) *EmailCodeService {
	return &EmailCodeService{
		To:      to,
		Code:    code,
		Subject: subject,
	}
}

func (e *EmailCodeService) SendCode() error {
	conf := config.GetConfig()
	m := gomail.NewMessage()
	m.SetHeader("From", conf.EmailConfig.From)
	m.SetHeader("To", e.To)
	m.SetHeader("Subject", e.Subject)

	data := struct {
		Username string
		Code     string
	}{
		Username: e.To,
		Code:     e.Code,
	}

	t, err := template.ParseFiles("./template/email.html")
	if err != nil {
		log.Printf("parse file to template failed, and error is %s\n", err.Error())
		return err
	}
	buf := bytes.NewBufferString("")
	if err = t.Execute(buf, data); err != nil {
		log.Printf("execute template faile, and error is %s\n", err.Error())
		return err
	}
	m.SetBody("text/html", buf.String())
	d := gomail.NewDialer(conf.EmailConfig.Host, conf.EmailConfig.Port, conf.EmailConfig.Username, conf.EmailConfig.Password)
	if err = d.DialAndSend(m); err != nil {
		log.Printf("send email to %s error, and error is %s\n", e.To, err.Error())
		return err
	}
	return nil
}
