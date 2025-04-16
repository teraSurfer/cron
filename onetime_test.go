package cron

import (
	"testing"
	"time"
)

func TestOneTimeSchedule_Next(t *testing.T) {

	now := time.Now()
	futureTime := now.Add(10 * time.Second)

	tests := []struct {
		name     string
		schedule OneTimeSchedule
		now      time.Time
		expected time.Time
	}{
		{
			name:     "Zero schedule",
			schedule: OneTimeSchedule{},
			now:      now,
			expected: time.Time{},
		},
		{
			name:     "Future time",
			schedule: OneTimeSchedule{Time: futureTime},
			now:      now,
			expected: futureTime,
		},
		{
			name:     "Past time",
			schedule: OneTimeSchedule{Time: time.Now().Add(-10 * time.Second)},
			now:      time.Now(),
			expected: time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.schedule.Next(tt.now)
			if !actual.Equal(tt.expected) {
				t.Errorf("Next() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
