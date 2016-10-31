package common

import "time"

var (
	fakeTime *time.Time
)

func SetFakeTime(t time.Time) {
	fakeTime = &t
}

func AdvanceFakeTime(d time.Duration) {
	t := fakeTime.Add(d)
	fakeTime = &t
}

func SetAFakeTime() {
	t, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41.123456Z")
	Check(e)
	SetFakeTime(t)
}

func UnSetFakeTime() {
	fakeTime = nil
}

func Now() time.Time {
	if fakeTime != nil {
		return *fakeTime
	}
	return time.Now().UTC()
}

func ParseTimeOrInvalidError(s string, format string) (time.Time, error) {
	value, err := time.Parse(format, s)
	if err != nil {
		return value, InvalidArgumentError("Failed to parse time-string %s against %s", s, format)
	}
	return value, nil
}
