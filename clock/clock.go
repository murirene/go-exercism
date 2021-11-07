package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func normalizeMinute(min int) (int, int) {
	hour := min / 60
	min %= 60
	if min < 0 {
		hour = hour - 1
		min += 60
	}
	return hour, min
}
func normalizeHour(hour int) int {
	hour %= 24
	if hour < 0 {
		hour += 24
	}
	return hour
}
func New(h, m int) Clock {
	hDelta, min := normalizeMinute(m)
	hour := normalizeHour(h + hDelta)
	return Clock{
		hour:   hour,
		minute: min,
	}
}
func (c Clock) Add(m int) Clock {
	hDelta, min := normalizeMinute(c.minute + m)
	hour := normalizeHour(c.hour + hDelta)
	c.hour = hour
	c.minute = min

	return c
}
func (c Clock) Subtract(m int) Clock {
	hDelta, min := normalizeMinute(c.minute - m)
	hour := normalizeHour(c.hour + hDelta)
	c.hour = hour
	c.minute = min

	return c
}
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
