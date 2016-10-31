package common

func MapStringSlice(s []string, f func(string) string) []string {
	if s == nil {
		return nil
	}
	o := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		o[i] = f(s[i])
	}
	return o
}

func FilterStringSlice(s []string, f func(string) bool) []string {
	if s == nil {
		return nil
	}
	o := make([]string, 0)
	for _, si := range s {
		if f(si) {
			o = append(o, si)
		}
	}
	return o
}
