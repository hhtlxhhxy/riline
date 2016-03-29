package main

import (
	"net/smtp"
	"strings"
	"fmt"
	"os"
)

const (
	Host = "smtp.163.com"
	Server_addr = "smtp.163.com:25"
	User = "lvxiaohan1415@163.com"
	Password = "cherish19921217"
)

type Email struct {
	to      string
	subject string
	msg     string
}

func NewEmail(to, subject, msg string) *Email {
	return &Email{to:to, subject:subject, msg:msg}
}

func (e *Email)SendEmail() (error) {
	sendTo := strings.Split(e.to, ",")
	auth := smtp.PlainAuth("", User, Password, Host)
	for _, val := range sendTo {
		str := strings.Replace("From: " + User + "~To: " + val + "~subject: " + e.subject + "~~", "~", "\r\n", -1) + e.msg
		err := smtp.SendMail(Server_addr, auth, User, []string{val}, []byte(str))
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	email := NewEmail("a605937391@vip.qq.com", "test", "haha")
	err := email.SendEmail()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
