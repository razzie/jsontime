package jsontime

import "time"

var _ TimeFormatProvider = (*RFC3339)(nil)

type JSONTimeRFC3339 JSONTime[RFC3339]

type RFC3339 struct{}

func (RFC3339) TimeFormat() string {
	return time.RFC3339
}
