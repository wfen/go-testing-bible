package service

import "fmt"

// MessageService - interface for sending messages to customers
type MessageService interface {
	SendChargeNotification(int) bool
}

// SMSService - sms service
type SMSService struct{}

// MyService - my service
type MyService struct {
	MessageService MessageService
}

// SendChargeNotification - implementation for sending notifications
func (sms SMSService) SendChargeNotification(value int) bool {
	fmt.Println("Sending production charge notification")
	return true
}

// ChargeCustomer - charges a customer for our service
func (ms MyService) ChargeCustomer(value int) error {
	ms.MessageService.SendChargeNotification(value)
	fmt.Printf("Charging the customer for value %d\n", value)
	return nil
}
