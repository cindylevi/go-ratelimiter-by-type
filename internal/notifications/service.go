package notifications

import (
	"RateLimiter/pkg/gateaway"
	"RateLimiter/pkg/limiter"
	"errors"
	"fmt"
)

const (
	minute_to_seconds = 60
	hour_to_seconds   = 60 * minute_to_seconds
	day_to_seconds    = hour_to_seconds * 24
	week_to_seconds   = day_to_seconds * 7
)

type Service struct {
	limiter *limiter.Limiter
	gateway gateaway.Service
}

func NewService() *Service {
	var configs = map[string]*limiter.Type{
		"test":      limiter.NewType(2, 1),
		"status":    limiter.NewType(2, minute_to_seconds),
		"marketing": limiter.NewType(3, hour_to_seconds),
		"news":      limiter.NewType(1, day_to_seconds),
		"discount":  limiter.NewType(1, week_to_seconds),
	}

	lim := limiter.NewLimiter(configs)
	return &Service{
		limiter: &lim,
		gateway: gateaway.NewService(),
	}
}

func (s Service) Send(notificationType string, userId string, message string) error {
	allowed, err := s.limiter.Permit(notificationType, userId)
	if err != nil {
		return err
	}
	if !allowed {
		return errors.New(fmt.Sprintf("notification type %s for user %s has exceeded the quota", notificationType, userId))
	}
	s.gateway.Send(userId, message)
	return nil
}
