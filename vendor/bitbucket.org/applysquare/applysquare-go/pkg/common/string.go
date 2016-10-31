package common

import (
	"fmt"
	"regexp"
	"strings"
)

func SplitAndTrim(s string, d string) []string {
	return FilterStringSlice(
		MapStringSlice(strings.Split(s, d), strings.TrimSpace),
		func(s string) bool {
			return s != ""
		})
}

func FmtMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		if v, ok := msgAndArgs[0].(string); ok {
			return v
		}
		if v, ok := msgAndArgs[0].(error); ok {
			return v.Error()
		}
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}

func NewStrPtr(s string) *string {
	return &s
}

// EmailPattern email pattern
var EmailPattern = `^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})$`

// ValidEmail if the given email is valid, return true
func ValidEmail(email string) bool {
	valid, _ := regexp.MatchString(EmailPattern, email)
	return valid
}

// GetSourceByEmail figure out which source a email belongs to
func GetSourceByEmail(email string) string {
	var source string
	if ValidEmail(email) {
		switch strings.Split(strings.Split(email, "@")[1], ".")[0] {
		case "qq":
			source = "qq"
		case "weibo":
			source = "weibo"
		case "wechat":
			source = "wechat"
		case "trial":
			source = "visitor"
		default:
			source = "email"
		}
	} else {
		source = "Not an email"
	}
	return source
}
