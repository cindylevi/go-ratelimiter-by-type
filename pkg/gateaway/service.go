package gateaway

import "log"

type Service interface {
	Send(userId string, message string)
}

type Gateway struct{}

func NewService() Gateway {
	return Gateway{}
}
func (s Gateway) Send(userId string, message string) {
	log.Printf("Hello %s! Our message is: %s", userId, message)
}
