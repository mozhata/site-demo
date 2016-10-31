package common

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
)

const defaultSkip = 2
const errorPrefixFormat string = "%s:%d: "

type A2Error struct {
	// The http status code is used here. But this error class could be used to pass canonical error that's not related to http.
	StatusCode int
	Message    string
}

func (e *A2Error) Error() string {
	return fmt.Sprintf("%v: %v", e.StatusCode, e.Message)
}

func UnavailableError(msgAndArgs ...interface{}) error {
	return &A2Error{http.StatusServiceUnavailable, FmtMsgAndArgs(msgAndArgs...)}
}

func IsUnavailableError(err error) bool {
	e, ok := err.(*A2Error)
	return ok && e.StatusCode == http.StatusServiceUnavailable
}

func NotFoundError(msgAndArgs ...interface{}) error {
	return &A2Error{http.StatusNotFound, FmtMsgAndArgs(msgAndArgs...)}
}

func IsNotFoundError(err error) bool {
	e, ok := err.(*A2Error)
	return ok && e.StatusCode == http.StatusNotFound
}

func ForbiddenError(msgAndArgs ...interface{}) error {
	return &A2Error{http.StatusForbidden, FmtMsgAndArgs(msgAndArgs...)}
}

func IsForbiddenError(err error) bool {
	e, ok := err.(*A2Error)
	return ok && e.StatusCode == http.StatusForbidden
}

func InvalidArgumentError(msgAndArgs ...interface{}) error {
	return &A2Error{http.StatusBadRequest, FmtMsgAndArgs(msgAndArgs...)}
}

func IsInvalidArgumentError(err error) bool {
	e, ok := err.(*A2Error)
	return ok && e.StatusCode == http.StatusBadRequest
}

func ConflictError(msgAndArgs ...interface{}) error {
	return &A2Error{http.StatusConflict, FmtMsgAndArgs(msgAndArgs...)}
}

func IsConflictError(err error) bool {
	e, ok := err.(*A2Error)
	return ok && e.StatusCode == http.StatusConflict
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// getPrefix skips "skip" stack frames to get the file & line number
// of original caller.
func getPrefix(skip int, format string) string {
	if _, file, line, ok := runtime.Caller(skip); ok {
		return fmt.Sprintf(format, filepath.Base(file), line)
	}
	return ""
}

// Errorf is a passthrough to fmt.Errorf, with an additional prefix
// containing the filename and line number.
func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(getPrefix(defaultSkip, errorPrefixFormat)+format, a...)
}

// ErrorfSkipFrames allows the skip count for stack frames to be
// specified. This is useful when generating errors via helper
// methods. Skip should be specified as the number of additional stack
// frames between the location at which the error is caused and the
// location at which the error is generated.
func ErrorfSkipFrames(skip int, format string, a ...interface{}) error {
	return fmt.Errorf(getPrefix(defaultSkip+skip, errorPrefixFormat)+format, a...)
}

// Error is a passthrough to fmt.Error, with an additional prefix
// containing the filename and line number.
func Error(a ...interface{}) error {
	return ErrorSkipFrames(1, a...)
}

// ErrorSkipFrames allows the skip count for stack frames to be
// specified. See the comments for ErrorfSkip.
func ErrorSkipFrames(skip int, a ...interface{}) error {
	prefix := getPrefix(defaultSkip+skip, errorPrefixFormat)
	if prefix != "" {
		a = append([]interface{}{prefix}, a...)
	}
	return fmt.Errorf("%s", fmt.Sprint(a...))
}
