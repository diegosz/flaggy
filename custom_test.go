package flaggy

import (
	"reflect"
	"testing"
	"time"
)

func TestDateZ_IsZero(t *testing.T) {
	tests := []struct {
		name string
		d    DateZ
		want bool
	}{
		{name: "zero", d: DateZ(time.Time{}), want: true},
		{name: "zero_plus_one_day", d: DateZ(time.Time{}.Add(roundDay)), want: false},
		{name: "zero_plus_one_day_minus_one_second", d: DateZ(time.Time{}.Add(roundDay - 1*time.Second)), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IsZero(); got != tt.want {
				t.Errorf("DateZ.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDateZ(t *testing.T) {
	tests := []struct {
		name    string
		date    string
		want    time.Time
		wantErr bool
	}{
		{name: "empty", date: "", want: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
		{name: "0000-00-00", date: "0000-00-00", want: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
		{name: "0001-01-01", date: "0001-01-01", want: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
		{name: "2006-01-02", date: "2006-01-02", want: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC)},
		{name: "20060102", date: "20060102", wantErr: true},
		{name: "2006-03", date: "2006-03", wantErr: true},
		{name: "2006-01-13", date: "2006-01-13", want: time.Date(2006, 1, 13, 0, 0, 0, 0, time.UTC)},
		{name: "2006-13-02", date: "2006-13-02", wantErr: true},
		{name: "2006-01-32", date: "2006-01-32", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDateZ(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDateZ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotTime := got.Time()
			if !reflect.DeepEqual(gotTime, tt.want) {
				t.Errorf("NewDateZ() = %v, want %v", gotTime, tt.want)
			}
		})
	}
}

func TestDateZ_String(t *testing.T) {
	tests := []struct {
		name string
		d    DateZ
		want string
	}{
		{name: "zero", d: DateZ(time.Time{}), want: "0001-01-01"},
		{name: "zero_plus_one_day", d: DateZ(time.Time{}.Add(roundDay)), want: "0001-01-02"},
		{name: "zero_plus_one_day_minus_one_second", d: DateZ(time.Time{}.Add(roundDay - 1*time.Second)), want: "0001-01-01"},
		{name: "empty", d: mustNewDateZ(t, ""), want: "0001-01-01"},
		{name: "0000-00-00", d: mustNewDateZ(t, "0000-00-00"), want: "0001-01-01"},
		{name: "0001-01-01", d: mustNewDateZ(t, "0001-01-01"), want: "0001-01-01"},
		{name: "2006-01-02", d: mustNewDateZ(t, "2006-01-02"), want: "2006-01-02"},
		{name: "2006-01-13", d: mustNewDateZ(t, "2006-01-13"), want: "2006-01-13"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("DateZ.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateZ_Time(t *testing.T) {
	zero := DateZ(time.Time{})
	zeroPlusOneDay := DateZ(time.Time{}.Add(roundDay))
	zeroPlusOneDayMinusOneSecond := DateZ(time.Time{}.Add(roundDay - 1*time.Second))
	someDate := mustNewDateZ(t, "2006-01-02")
	tests := []struct {
		name string
		d    *DateZ
		want time.Time
	}{
		{name: "zero", d: &zero, want: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
		{name: "zero_plus_one_day", d: &zeroPlusOneDay, want: time.Date(1, 1, 2, 0, 0, 0, 0, time.UTC)},
		{name: "zero_plus_one_day_minus_one_second", d: &zeroPlusOneDayMinusOneSecond, want: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
		{name: "2006-01-02", d: &someDate, want: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC)},
		{name: "nil", d: nil, want: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateZ.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mustNewDateZ(t *testing.T, date string) DateZ {
	t.Helper()
	d, err := NewDateZ(date)
	if err != nil {
		t.Fatalf("mustNewDatez failed: %v", err)
	}
	return d
}

func TestDateZ_EnvDecode(t *testing.T) {
	zero := DateZ(time.Time{})
	tests := []struct {
		name    string
		value   string
		want    DateZ
		wantErr bool
	}{
		{name: "empty", value: "", want: zero},
		{name: "0000-00-00", value: "0000-00-00", want: zero},
		{name: "0001-01-01", value: "0001-01-01", want: zero},
		{name: "2006-01-02", value: "2006-01-02", want: DateZ(time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC))},
		{name: "20060102", value: "20060102", wantErr: true},
		{name: "2006-03", value: "2006-03", wantErr: true},
		{name: "2006-01-13", value: "2006-01-13", want: DateZ(time.Date(2006, 1, 13, 0, 0, 0, 0, time.UTC))},
		{name: "2006-13-02", value: "2006-13-02", wantErr: true},
		{name: "2006-01-32", value: "2006-01-32", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var assignmentVar interface{}
			var got DateZ
			assignmentVar = &got
			existing := assignmentVar.(*DateZ)
			if err := existing.EnvDecode(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("DateZ.EnvDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateZ.EnvDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeZ_IsZero(t *testing.T) {
	tests := []struct {
		name string
		d    TimeZ
		want bool
	}{
		{name: "zero", d: TimeZ(time.Time{}), want: true},
		{name: "zero_plus_one_second", d: TimeZ(time.Time{}.Add(1 * time.Second)), want: false},
		{name: "zero_plus_one_day", d: TimeZ(time.Time{}.Add(roundDay)), want: false},
		{name: "zero_plus_one_day_minus_one_second", d: TimeZ(time.Time{}.Add(roundDay - 1*time.Second)), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IsZero(); got != tt.want {
				t.Errorf("TimeZ.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTimeZ(t *testing.T) {
	tests := []struct {
		name    string
		date    string
		want    time.Time
		wantErr bool
	}{
		{name: "empty", date: "", want: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
		{name: "0000-00-00T00:00:00Z", date: "0000-00-00T00:00:00Z", wantErr: true},
		{name: "0001-01-01T00:00:00Z", date: "0001-01-01T00:00:00Z", want: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
		{name: "2006-01-02T00:00:00Z", date: "2006-01-02T00:00:00Z", want: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC)},
		{name: "2006-01-02T00:00:00Z", date: "2006-01-02T15:04:05-03:00", want: time.Date(2006, 1, 2, 15, 4, 5, 0, time.FixedZone("-03", -3*60*60))},
		{name: "2006-01-02T00:00:00Z", date: "2006-01-02T15:04:05+03:00", want: time.Date(2006, 1, 2, 15, 4, 5, 0, time.FixedZone("03", 3*60*60))},
		{name: "2006-01-02T00:00:00", date: "2006-01-02T00:00:00", wantErr: true},
		{name: "0001-01-01T00:00:01Z", date: "0001-01-01T00:00:01Z", want: time.Date(1, 1, 1, 0, 0, 1, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimeZ(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTimeZ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotTime := got.Time()
			if !sameTime(gotTime, tt.want) {
				t.Errorf("NewTimeZ() = %v, want %v", gotTime, tt.want)
			}
		})
	}
}

func TestTimeZ_String(t *testing.T) {
	tests := []struct {
		name string
		d    *TimeZ
		want string
	}{
		{name: "empty", d: mustNewTimeZ(t, ""), want: "0001-01-01T00:00:00Z"},
		{name: "2006-01-02T00:00:00Z", d: mustNewTimeZ(t, "2006-01-02T00:00:00Z"), want: "2006-01-02T00:00:00Z"},
		{name: "2006-01-02T00:00:00-03:00", d: mustNewTimeZ(t, "2006-01-02T00:00:00-03:00"), want: "2006-01-02T00:00:00-03:00"},
		{name: "2006-01-02T00:00:00+03:00", d: mustNewTimeZ(t, "2006-01-02T00:00:00+03:00"), want: "2006-01-02T00:00:00+03:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("TimeZ.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeZ_Time(t *testing.T) {
	zero := TimeZ(time.Time{})
	zeroPlusOneDay := TimeZ(time.Time{}.Add(roundDay))
	zeroPlusOneDayMinusOneSecond := TimeZ(time.Time{}.Add(roundDay - 1*time.Second))
	tests := []struct {
		name string
		d    *TimeZ
		want time.Time
	}{
		{name: "zero", d: &zero, want: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
		{name: "zero_plus_one_day", d: &zeroPlusOneDay, want: time.Date(1, 1, 2, 0, 0, 0, 0, time.UTC)},
		{name: "zero_plus_one_day_minus_one_second", d: &zeroPlusOneDayMinusOneSecond, want: time.Date(1, 1, 1, 23, 59, 59, 0, time.UTC)},
		{name: "2006-01-02T00:00:00Z", d: mustNewTimeZ(t, "2006-01-02T00:00:00Z"), want: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC)},
		{name: "2006-01-02T00:00:00-03:00", d: mustNewTimeZ(t, "2006-01-02T00:00:00-03:00"), want: time.Date(2006, 1, 2, 0, 0, 0, 0, time.FixedZone("-03", -3*60*60))},
		{name: "2006-01-02T00:00:00+03:00", d: mustNewTimeZ(t, "2006-01-02T00:00:00+03:00"), want: time.Date(2006, 1, 2, 0, 0, 0, 0, time.FixedZone("03", 3*60*60))},
		{name: "nil", d: nil, want: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Time(); !sameTime(got, tt.want) {
				t.Errorf("TimeZ.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mustNewTimeZ(t *testing.T, value string) *TimeZ {
	t.Helper()
	d, err := NewTimeZ(value)
	if err != nil {
		t.Fatalf("mustNewTimeZ failed: %v", err)
	}
	return &d
}

func TestTimeZ_EnvDecode(t *testing.T) {
	zero := TimeZ(time.Time{})
	tests := []struct {
		name    string
		value   string
		want    TimeZ
		wantErr bool
	}{
		{name: "empty", value: "", want: zero},
		{name: "zero", value: "0001-01-01T00:00:00Z", want: zero},
		{name: "2006-01-02T00:00:00Z", value: "2006-01-02T00:00:00Z", want: TimeZ(time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC))},
		{name: "2006-01-02T00:00:00-03:00", value: "2006-01-02T00:00:00-03:00", want: TimeZ(time.Date(2006, 1, 2, 0, 0, 0, 0, time.FixedZone("-03", -3*60*60)))},
		{name: "2006-01-02T00:00:00+03:00", value: "2006-01-02T00:00:00+03:00", want: TimeZ(time.Date(2006, 1, 2, 0, 0, 0, 0, time.FixedZone("03", 3*60*60)))},
		{name: "2006-01-02T00:00:00", value: "2006-01-02T00:00:00", wantErr: true},
		{name: "2006-01-32", value: "2006-01-32", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var assignmentVar interface{}
			var got TimeZ
			assignmentVar = &got
			existing := assignmentVar.(*TimeZ)
			if err := existing.EnvDecode(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("TimeZ.EnvDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !sameTime(got.Time(), tt.want.Time()) {
				t.Errorf("TimeZ.EnvDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sameTime(t time.Time, u time.Time) bool {
	// Check aggregates.
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	_, offset := t.Zone()
	uYear, uMonth, uDay := u.Date()
	uHour, uMin, uSec := u.Clock()
	_, uOffset := u.Zone()
	uNano := u.Nanosecond()
	uWeekday := u.Weekday()
	if year != uYear || month != uMonth || day != uDay ||
		hour != uHour || min != uMin || sec != uSec ||
		offset != uOffset {
		return false
	}
	// Check individual entries.
	return t.Year() == uYear &&
		t.Month() == uMonth &&
		t.Day() == uDay &&
		t.Hour() == uHour &&
		t.Minute() == uMin &&
		t.Second() == uSec &&
		t.Nanosecond() == uNano &&
		t.Weekday() == uWeekday
}
