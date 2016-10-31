package common

func StringInSlice(value string, array []string) bool {
	for _, s := range array {
		if s == value {
			return true
		}
	}
	return false
}

func SliceContainDuplicate(array1 []string, array2 []string) bool {
	for _, s1 := range array1 {
		for _, s2 := range array2 {
			if s1 == s2 {
				return true
			}
		}
	}
	return false
}

func IsEmptySlice(array []string) bool {
	for _, s := range array {
		if s != "" {
			return false
		}
	}
	return true
}
