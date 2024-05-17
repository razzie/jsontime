package jsontime

type TimeFormatProvider interface {
	TimeFormat() string
}

type TimeFormatProviderConstraint interface {
	~struct{}
	TimeFormatProvider
}
