package limiter

import (
	"RateLimiter/pkg/datetime"
	"sync"
	"time"
)

type client struct {
	quota       int
	lastUpdated time.Time
	mutex       *sync.Mutex
}

func newClient() *client {
	var m = sync.Mutex{}
	return &client{
		quota:       1,
		lastUpdated: datetime.Clock.Now(),
		mutex:       &m,
	}
}

func (client *client) weight(maxQuota int, periodInSeconds float64) bool {
	timeNow := datetime.Clock.Now()
	client.mutex.Lock()
	secondsElapsed := timeNow.Sub(client.lastUpdated)

	if secondsElapsed.Seconds() > periodInSeconds {
		client.quota = 1
		client.lastUpdated = timeNow
		client.mutex.Unlock()
		return true
	}

	if maxQuota <= client.quota {
		client.mutex.Unlock()
		return false
	}

	client.quota += 1
	client.lastUpdated = timeNow
	client.mutex.Unlock()
	return true
}
