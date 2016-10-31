package common

func CopyStrings(from []string) []string {
	if from == nil {
		return nil
	}
	ret := make([]string, len(from))
	copy(ret, from)
	return ret
}
