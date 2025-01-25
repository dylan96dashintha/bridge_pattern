package main

import "fmt"

// MessageGateway - Implementation Interface
type MessageGateway interface {
	SendMessage(content string)
}

// HutchGateway - Concrete Implementation
type HutchGateway struct{}

func (h *HutchGateway) SendMessage(content string) {
	fmt.Println("Sending Message Through Hutch Gateway, Message:\n", content)
}

// DialogGateway - Concrete Implementation
type DialogGateway struct{}

func (d *DialogGateway) SendMessage(content string) {
	fmt.Println("Sending Message Through Dialog Gateway, Message:\n", content)
}

// Message - Abstraction
type Message interface {
	Send(content string)
}

// SMSMessage - Refined Abstraction
type SMSMessage struct {
	gateway MessageGateway
}

func (sms *SMSMessage) Send(content string) {
	fmt.Println("SMSMessage:")
	sms.gateway.SendMessage(content)
}

// EmailMessage - Refined Abstraction
type EmailMessage struct {
	gateway MessageGateway
}

func (email *EmailMessage) Send(content string) {
	fmt.Println("EmailMessage:")
	email.gateway.SendMessage(content)
}

func main() {
	// Use Hutch Gateway for SMS
	hutchGateway := &HutchGateway{}
	smsMessage := &SMSMessage{gateway: hutchGateway}
	smsMessage.Send("This is an SMS sent through Hutch.")

	// Use Dialog Gateway for Email
	dialogGateway := &DialogGateway{}
	emailMessage := &EmailMessage{gateway: dialogGateway}
	emailMessage.Send("This is an Email sent through Dialog.")

	// Switching Gateways
	fmt.Println("\nSwitching to Dialog Gateway for SMS:")
	smsMessage.gateway = dialogGateway
	smsMessage.Send("This is an SMS now sent through Dialog.")
}
