package common

import (
	"bytes"
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/golang/glog"
)

const (
	JSONTimeFormat  = "2006-01-02T15:04:05"
	JSONTimeZFormat = "2006-01-02T15:04:05Z"
)

type JSONTime time.Time
type JSONTimeZ time.Time

func JSONNow() JSONTime {
	return JSONTime(Now())
}

var _ encoding.TextUnmarshaler = (*JSONTime)(nil)

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *JSONTime) UnmarshalText(data []byte) error {
	tt, err := time.Parse(JSONTimeFormat, string(data))
	*t = JSONTime(tt)
	return err
}

func (t JSONTime) Time() time.Time {
	return time.Time(t)
}

func MarshalJSONOrDie(v interface{}) string {
	b, err := json.Marshal(v)
	Check(err)
	return string(b)
}

func UnmarshalJSONOrDie(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	Check(err)
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format(JSONTimeFormat))), nil
}

func (t *JSONTime) UnmarshalJSON(s []byte) error {
	if string(s) == "null" {
		*t = JSONTime(time.Time{})
		return nil
	}
	q, err := strconv.Unquote(string(s))
	if err != nil {
		return err
	}
	*(*time.Time)(t), err = time.Parse(JSONTimeFormat, q)
	return err
}

// Time return time.Time
func (t JSONTimeZ) Time() time.Time {
	return time.Time(t)
}

// Before implement time.Before
func (t JSONTimeZ) Before(u JSONTimeZ) bool {
	return t.Time().Before(u.Time())
}

// Format return foramted time
func (t JSONTimeZ) Format() string {
	return t.Time().Format(JSONTimeZFormat)
}

// return if a JSONTimeZ is after startTime and before endTime
func (t JSONTimeZ) Between(startTime, endTime time.Time) bool {
	tem := t.Time()
	return tem.After(startTime) && tem.Before(endTime)
}
func (t JSONTimeZ) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%sZ\"", time.Time(t).Format(JSONTimeFormat))), nil
}

// UnmarshalJSON offer the way to UnmarshalJSON
func (t *JSONTimeZ) UnmarshalJSON(s []byte) error {
	if string(s) == "null" {
		*t = JSONTimeZ(time.Time{})
		return nil
	}
	q, err := strconv.Unquote(string(s))
	if err != nil {
		return err
	}
	*(*time.Time)(t), err = time.Parse(JSONTimeZFormat, q)
	return err
}

func Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return ErrorfSkipFrames(1, "unmarshal failed: %v, source: %s", err, string(data))
	}
	return nil
}

func CleanJSON(i interface{}) cleanJSON {
	return cleanJSON{I: i}
}

// Object wrapped in this struct is marshaled into json
// where any fields with empty values are removed.
type cleanJSON struct {
	I interface{}
}

func (j cleanJSON) MarshalJSON() ([]byte, error) {
	b0, err := json.Marshal(j.I)
	if err != nil {
		return nil, err
	}
	return StripJsonBytes(b0), nil
}

func StripJsonBytes(b []byte) []byte {
	var r interface{}

	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&r)
	if err != nil {
		panic(err)
	}
	newR, _ := jsonStripInterface(r)
	newB, err := json.Marshal(newR)
	if err != nil {
		panic(err)
	}
	return newB
}

func jsonStripInterface(i interface{}) (interface{}, bool) {
	if i == nil {
		return i, true
	}
	switch v := i.(type) {
	case string:
		return v, v == ""
	case []interface{}:
		newi := make([]interface{}, 0)
		for _, ii := range v {
			newii, strip := jsonStripInterface(ii)
			if !strip {
				newi = append(newi, newii)
			}
		}
		return newi, len(newi) == 0
	case json.Number:
		f, err := v.Float64()
		if err != nil {
			panic(err)
		}
		return v, f == 0
	case bool:
		return v, v == false
	case map[string]interface{}:
		for key, value := range v {
			newValue, strip := jsonStripInterface(value)
			if strip {
				delete(v, key)
			} else {
				v[key] = newValue
			}
		}
		return v, len(v) == 0
	}
	panic(i)
}

func LogJSON(j interface{}) {
	b, err := json.Marshal(j)
	glog.Infof("%v err: %v", string(b), err)
}

func JSONMergeSample(sample interface{}, target interface{}) interface{} {
	if sample == nil {
		return target
	}
	switch s := sample.(type) {
	case string:
		if _, ok := target.(string); !ok {
			return s
		}
		return target
	case float64:
		if _, ok := target.(float64); !ok {
			return s
		}
		return target
	case bool:
		if _, ok := target.(bool); !ok {
			return s
		}
		return target
	case int:
		if _, ok := target.(int); !ok {
			return s
		}
		return target
	case []interface{}:
		if target, ok := target.([]interface{}); !ok {
			return s
		} else {
			for index, _ := range target {
				if len(s) == 0 {
					glog.Infof("JSONMergeSample Warning: Insufficient Sample Information")
				} else {
					target[index] = JSONMergeSample(s[0], target[index])
				}
			}
			return target
		}
	case map[string]interface{}:
		if target, ok := target.(map[string]interface{}); !ok {
			return s
		} else {
			for key, value := range s {
				if _, ok := target[key]; !ok {
					target[key] = value
				} else {
					target[key] = JSONMergeSample(value, target[key])
				}
			}
			return target
		}
	}
	panic(sample)
}
