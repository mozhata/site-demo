package common

import (
	"encoding"
	"time"
)

// Date can be used in parameter schema to parse date.
type Date time.Time

var _ encoding.TextUnmarshaler = (*Date)(nil)

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (d *Date) UnmarshalText(data []byte) error {
	t, err := time.Parse("2006-01-02", string(data))
	*d = Date(t)
	return err
}

func (d *Date) ValuePtr() *time.Time {
	return (*time.Time)(d)
}

func (d Date) Value() time.Time {
	return (time.Time)(d)
}

// CommaSeparated can be used in parameter schema to parse comma separated string list.
type CommaSeparated []string

var _ encoding.TextUnmarshaler = (*CommaSeparated)(nil)

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (c *CommaSeparated) UnmarshalText(data []byte) error {
	*c = SplitAndTrim(string(data), ",")
	return nil
}

func (c CommaSeparated) Value() []string {
	return ([]string)(c)
}
