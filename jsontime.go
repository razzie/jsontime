package jsontime

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var (
	_ json.Marshaler     = (*JSONTime[RFC3339])(nil)
	_ fmt.Stringer       = (*JSONTime[RFC3339])(nil)
	_ TimeFormatProvider = (*RFC3339)(nil)
)

type JSONTimeRFC3339 JSONTime[RFC3339]

type TimeFormatProviderConstraint interface {
	~struct{}
	TimeFormatProvider
}

type TimeFormatProvider interface {
	TimeFormat() string
}

type RFC3339 struct{}

func (RFC3339) TimeFormat() string {
	return time.RFC3339
}

type JSONTime[F TimeFormatProviderConstraint] struct {
	time.Time
	format F
}

func (t JSONTime[F]) MarshalJSON() ([]byte, error) {
	stamp := `"` + t.Time.Format(t.format.TimeFormat()) + `"`
	return []byte(stamp), nil
}

func (t *JSONTime[F]) UnmarshalJSON(data []byte) error {
	ts := string(data)
	ts = strings.TrimFunc(ts, func(r rune) bool { return r == '"' })
	stamp, err := time.Parse(t.format.TimeFormat(), ts)
	if err != nil {
		return err
	}
	t.Time = stamp
	return nil
}

func (t JSONTime[F]) String() string {
	return t.Time.Format(t.format.TimeFormat())
}
