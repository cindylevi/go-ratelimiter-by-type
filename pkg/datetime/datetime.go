package datetime

import "time"

type Provider interface {
	Now() time.Time
}

type DateTime struct{}

func (r DateTime) Now() time.Time {
	return time.Now()
}

var Clock Provider = DateTime{}
