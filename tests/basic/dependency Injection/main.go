package main

import "fmt"

type MessageService interface {
	sendMessage(message string, number int) error
}

type SMSSendMessage struct {
}

func (sms *SMSSendMessage) sendMessage(message string, number int) error {
	fmt.Println("SMS send message::", message, "number is::", number)
	return nil
}

type EmailSendMessage struct {
}

func (e *EmailSendMessage) sendMessage(message string, number int) error {
	fmt.Println("Email send message::", message, "number is::", number)
	return nil
}

type MyApplication struct {
	messageService MessageService
}

func (a *MyApplication) processMessage(message string, number int) error {
	err := a.messageService.sendMessage(message, number)
	if err != nil {
		return err
	}
	return nil
}

type OTPSendMessage struct {
}

func (otp *OTPSendMessage) sendMessage(message string, number int) error {
	fmt.Println("Email send message::", message, "number is::", number)
	return nil
}

func main() {
	sms := &SMSSendMessage{}
	app := &MyApplication{messageService: sms}
	app.processMessage("hello sms", 1)

	fmt.Println("------------")

	email := &EmailSendMessage{}
	app.messageService = email
	app.processMessage("hello email", 2)

	fmt.Println("------------")

	otp := &OTPSendMessage{}
	app.messageService = otp
	app.processMessage("hello otp", 3)
}
