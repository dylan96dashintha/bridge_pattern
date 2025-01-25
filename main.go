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

// TextMessage - Refined Abstraction
type TextMessage struct {
	gateway MessageGateway
}

func (t *TextMessage) Send(content string) {
	fmt.Println("Text Message:")
	t.gateway.SendMessage(content)
}

// WhatsappMessage - Refined Abstraction
type WhatsappMessage struct {
	gateway MessageGateway
}

func (w *WhatsappMessage) Send(content string) {
	fmt.Println("Whatsapp Message:")
	w.gateway.SendMessage(content)
}

func main() {
	// Use Hutch Gateway for SMS
	hutchGateway := &HutchGateway{}
	smsMessage := &TextMessage{gateway: hutchGateway}
	smsMessage.Send("This is an Text Message sent through Hutch.")

	// Use Dialog Gateway for Email
	dialogGateway := &DialogGateway{}
	emailMessage := &WhatsappMessage{gateway: dialogGateway}
	emailMessage.Send("This is an whatsapp message sent through Dialog.")

	// Switching Gateways
	fmt.Println("\nSwitching to Dialog Gateway for SMS:")
	smsMessage.gateway = dialogGateway
	smsMessage.Send("This is an text message now sent through Dialog.")
}
