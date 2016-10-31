package i18n

import "bitbucket.org/applysquare/applysquare-go/pkg/util/h"

type S string

func (s S) ICanBeToken() {}

func (s S) ToHTML(w h.StringWriter) error {
	return h.S(s).ToHTML(w)
}

func (s S) String() string {
	return string(s)
}
