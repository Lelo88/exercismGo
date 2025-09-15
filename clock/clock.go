package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour, minute int
}

func New(h, m int) Clock {
	if m < 0 {
		h += (m - 59) / 60
		m = (m % 60 + 60) % 60
	} else {
		h += m / 60
		m = m % 60
	}
	h = h % 24
	if h < 0 {
		h += 24
	}
	return Clock{hour: h, minute: m}
}

func (c Clock) Add(m int) Clock {
	if m < 0 {
		return c.Subtract(-m)
	}
	h := c.hour + (m/60)%24
	min := c.minute + m%60
	if min >= 60 {
		h += min / 60
		min = min % 60
	}
	h = h % 24
	if h < 0 {
		h += 24
	}
	return Clock{hour: h, minute: min}
}

func (c Clock) Subtract(m int) Clock {
	if m < 0 {
		return c.Add(-m)
	}
	h := c.hour - (m/60)%24
	min := c.minute - m%60
	if min < 0 {
		h -= (1 + (-min)/60)
		min = (min + 60) % 60
	}
	h = h % 24
	if h < 0 {
		h += 24
	}
	return Clock{hour: h, minute: min}
}

func (c Clock) String() string {
	h := c.hour
	if h < 10 {
		h = 0*10 + h
	}
	m := c.minute
	if m < 10 {
		m = 0*10 + m
	}
	return fmt.Sprintf("%02d:%02d", h, m)
}
