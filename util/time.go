package util

import (
	"time"
)

func GetSystemTime() time.Time {
	start := time.Time{}

	return start
}

func IncreaseTime(t time.Time, count int) time.Time {
	t = t.Add(time.Hour * time.Duration(count))

	return t
}

func ChangeTimeFormat(t time.Time) string {
	return t.Format("15:04")
}

func IncreaseTimeCommand(t time.Time) string {
	return "Time is " + ChangeTimeFormat(t)
}
