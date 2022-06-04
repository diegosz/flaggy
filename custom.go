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
