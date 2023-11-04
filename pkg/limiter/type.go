package limiter

import (
	"errors"
	"sync"
)

type Type struct {
	MaxQuota        int
	PeriodInSeconds float64
	clients         *sync.Map
}

func NewType(maxQuota int, period float64) *Type {
	var cMap = sync.Map{}
	return &Type{
		MaxQuota:        maxQuota,
		PeriodInSeconds: period,
		clients:         &cMap,
	}
}

func (typ Type) Permit(userId string) (bool, error) {
	emptyClient := newClient()
	value, loaded := typ.clients.LoadOrStore(userId, emptyClient)
	if !loaded {
		//new client initialized on map
		return true, nil
	}
	cl, ok := value.(*client)
	if !ok {
		return false, errors.New("error retrieving client information")
	}
	return cl.weight(typ.MaxQuota, typ.PeriodInSeconds), nil
}
