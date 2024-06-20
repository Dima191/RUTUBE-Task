package notifications

import "time"

func (n *Notification) timeAfter() <-chan time.Time {
	now := time.Now()
	location := now.Location()
	tomorrow := now.AddDate(0, 0, 1)
	duration := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 8, 0, 0, 0, location).Sub(now)
	checkBirth := time.After(duration)

	return checkBirth
}
