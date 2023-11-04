package notifications

import (
	"RateLimiter/pkg/datetime"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var notificationService = NewService()
var timeMock = datetime.InitializeMockTime()

func beforeTest() {
	timeMock = datetime.InitializeMockTime()
	notificationService = NewService()
}

func TestSendNotificationTypeNotConfiguredShouldReturnError(t *testing.T) {
	beforeTest()
	err := notificationService.Send("not-a-type", "user1", "My notification message 2!")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "type not configured")

}

func TestSendNotificationNewsShouldAllowOnlyOneRequestPerDay(t *testing.T) {
	beforeTest()

	err := notificationService.Send("news", "user1", "My notification message 1!")
	assert.Nil(t, err)

	err = notificationService.Send("news", "user1", "My notification message 2!")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "has exceeded the quota")

	timeMock.ModifyTime(timeMock.Now().AddDate(0, 0, 1).Add(time.Second))
	err = notificationService.Send("news", "user1", "My notification message 3!")
	assert.Nil(t, err)
}

func TestSendNotificationNewsShouldAllowOnlyOneRequestPerDayMoreThanOneClient(t *testing.T) {
	beforeTest()
	err := notificationService.Send("news", "user1", "My notification user 1 message 1!")
	assert.Nil(t, err)

	err = notificationService.Send("discount", "user2", "My notification user 2 message 1!")
	assert.Nil(t, err)

	err = notificationService.Send("discount", "user2", "My notification user 2 message 2!")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "has exceeded the quota")

	timeMock.ModifyTime(timeMock.Now().AddDate(0, 0, 7).Add(time.Second))
	err = notificationService.Send("news", "user1", "My notification user 1 message 3!")
	assert.Nil(t, err)

	err = notificationService.Send("discount", "user2", "My notification user 2 message 3!")
	assert.Nil(t, err)
}

func TestSendNotificationStatusShouldAllowOnly2RequestsPerMinute(t *testing.T) {
	beforeTest()
	err := notificationService.Send("status", "user1", "My notification message 1!")
	assert.Nil(t, err)

	err = notificationService.Send("status", "user1", "My notification message 2!")
	assert.Nil(t, err)

	err = notificationService.Send("status", "user1", "My notification message 3!")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "has exceeded the quota")

	timeMock.ModifyTime(timeMock.Now().Add(time.Minute).Add(time.Second))
	err = notificationService.Send("status", "user1", "My notification message 4!")
	assert.Nil(t, err)
}

func TestSendNotificationMarketingShouldAllowOnly3RequestsPerHour(t *testing.T) {
	beforeTest()
	for i := 1; i <= 3; i++ {
		err := notificationService.Send("marketing", "user1", fmt.Sprintf("My notification message %d!", i))
		assert.Nil(t, err)
	}

	err := notificationService.Send("marketing", "user1", "My notification message 4!")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "has exceeded the quota")

	timeMock.ModifyTime(timeMock.Now().Add(3 * time.Hour).Add(time.Second))
	err = notificationService.Send("marketing", "user1", "My notification message 5!")
	assert.Nil(t, err)
}
