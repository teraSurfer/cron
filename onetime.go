package cron

import "time"

// OneTimeSchedule represents a one-time job that runs at a specific time.
type OneTimeSchedule struct {
	// Time is the time at which this job should run.
	Time time.Time
}

// Next returns the next time this job should run.
func (schedule OneTimeSchedule) Next(t time.Time) time.Time {
	if schedule.Time.IsZero() {
		return time.Time{}
	}
	if schedule.Time.After(t) {
		return schedule.Time
	}
	return time.Time{}
}

// IsZero returns true if the schedule is zero, meaning it has no time set.
func (schedule OneTimeSchedule) IsZero() bool {
	return schedule.Time.IsZero()
}

// Once creates a new OneTimeSchedule for the given duration from now.
func Once() OneTimeSchedule {
	return OneTimeSchedule{Time: time.Now().Add(time.Second)}
}
