package timex

import "time"

type DurationUnit string

var (
	DurationSecond DurationUnit = "second"
	DurationMinute DurationUnit = "minute"
	DurationHour   DurationUnit = "hour"
	DurationDay    DurationUnit = "day"
)

type Duration struct {
	Number int          `json:"number"`
	Unit   DurationUnit `json:"unit"`
}

func (d *Duration) Duration() time.Duration {
	switch d.Unit {
	case "second":
		return time.Duration(d.Number) * time.Second
	case "minute":
		return time.Duration(d.Number) * time.Minute
	case "hour":
		return time.Duration(d.Number) * time.Hour
	case "day":
		return time.Duration(d.Number) * time.Hour * 24
	default:
		return 0
	}
}
