package datetime

import "time"

type MockDateTime struct {
	time time.Time
}

func (m *MockDateTime) Now() time.Time {
	return m.time
}

func (m *MockDateTime) ModifyTime(newTime time.Time) {
	m.time = newTime
}

func newMockDateTimeProvider() *MockDateTime {
	return &MockDateTime{
		time: time.Date(2023, time.November, 2, 12, 0, 0, 0, time.UTC),
	}
}

func InitializeMockTime() *MockDateTime {
	mock := newMockDateTimeProvider()
	Clock = mock
	return mock
}
