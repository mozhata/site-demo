package i18n

type LocalizedString struct {
	ZhCN string `json:"zh-cn"`
	EnUS string `json:"en-us"`
}

func (s LocalizedString) String(intl *Context) string {
	switch intl.lang {
	case ZhCN:
		if s.ZhCN == "" {
			return s.EnUS
		}
		return s.ZhCN
	default:
		return s.EnUS
	}
}

// L constructs LocalizedString for deferred translation.
func L(key string) LocalizedString {
	return LocalizedString{
		ZhCN: CnContext.S(key),
		EnUS: EnContext.S(key),
	}
}
