package main

import (
	"RateLimiter/internal/notifications"
	"fmt"
	"log"
)

func main() {
	notificationsService := notifications.NewService()

	for i := 1; i <= 3; i++ {
		err := notificationsService.Send("news", "cindy", fmt.Sprintf("My notification message %d!", i))
		if err != nil {
			log.Printf("Too many notifications")
		}
	}

}
