package flaggy

import (
	"time"
)

const (
	roundDay    = 24 * time.Hour
	DateZFormat = "2006-01-02"
)

// DateZ represents a time.Time date parsed and formatted with UTC "2006-01-02"
// format.
type DateZ time.Time

// IsZero reports whether d represents the zero date, 0001-01-01 UTC.
func (d DateZ) IsZero() bool {
	return time.Time(d).Truncate(roundDay).IsZero()
}

// NewDateZ returns a new DateZ. If date is "" or "0000-00-00" a zero time
// instant is used to create the DateZ value represented by 0001-01-01 UTC.
func NewDateZ(date string) (DateZ, error) {
	var t time.Time
	var err error
	switch date {
	case "", "0000-00-00":
		t = time.Time{}
	default:
		t, err = time.Parse(DateZFormat, date)
		if err != nil {
			return DateZ{}, err
		}
	}
	return DateZ(t), nil
}

// EnvDecode complies with the envconfig.Decoder interface
func (d *DateZ) EnvDecode(value string) error {
	tmp, err := NewDateZ(value)
	if err != nil {
		return err
	}
	*d = tmp
	return nil
}

// Time return the underline time.Time
func (d *DateZ) Time() time.Time {
	if d == nil {
		return time.Time{}
	}
	return time.Time(*d).Truncate(roundDay)
}

// String returns a textual representation of the time value formatted as
// DateZFormat "2006-01-02".
func (d *DateZ) String() string {
	t := d.Time()
	return t.Format(DateZFormat)
}

// TimeZ represents a time.Time parsed and formatted with UTC RFC3339 format.
type TimeZ time.Time

// IsZero reports whether t represents the zero time.
func (t TimeZ) IsZero() bool {
	return time.Time(t).IsZero()
}

// NewTimeZ returns a new TimeZ. If date is "" a zero time instant is used to
// create the TimeZ value.
func NewTimeZ(value string) (TimeZ, error) {
	var t time.Time
	var err error
	switch value {
	case "":
		t = time.Time{}
	default:
		t, err = time.Parse(time.RFC3339, value)
		if err != nil {
			return TimeZ{}, err
		}
	}
	return TimeZ(t), nil
}

// EnvDecode complies with the envconfig.Decoder interface
func (t *TimeZ) EnvDecode(value string) error {
	tmp, err := NewTimeZ(value)
	if err != nil {
		return err
	}
	*t = tmp
	return nil
}

// Time return the underline time.Time
func (t *TimeZ) Time() time.Time {
	if t == nil {
		return time.Time{}
	}
	return time.Time(*t)
}

// String returns a textual representation of the time value formatted as
// time.RFC3339.
func (t *TimeZ) String() string {
	return t.Time().Format(time.RFC3339)
}
