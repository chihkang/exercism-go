package clock

import "fmt"

const testVersion = 4

type Clock struct {
	h int
	m int
}

func New(hour, minute int) Clock {
	hour, minute = convertTime(hour, minute)
	return Clock{h: hour, m: minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	c.h, c.m = convertTime(c.h, c.m+minutes)
	return Clock{c.h, c.m}
}
func convertTime(hour, minute int) (int, int) {
	for minute >= 60 {
		minute -= 60
		hour += 1
	}
	for minute < 0 {
		minute += 60
		hour -= 1
	}
	for hour >= 24 {
		hour -= 24
	}
	for hour < 0 {
		hour += 24
	}
	return hour, minute
}
