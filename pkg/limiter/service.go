package limiter

import (
	"errors"
)

type Limiter struct {
	types map[string]*Type
}

func NewLimiter(configs map[string]*Type) Limiter {
	return Limiter{
		types: configs,
	}
}

// Permit returns true if the userId has quota for the requested type and false if not
func (limiter Limiter) Permit(requestedType string, userId string) (bool, error) {
	notificationType := limiter.types[requestedType]
	if notificationType == nil {
		return false, errors.New("type not configured")
	}

	return notificationType.Permit(userId)
}
