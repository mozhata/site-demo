package home

const (
	EMAIL = iota
	PHTONE
)

func RegisterByEmail(email, pwd string) string {

	return ""
}

// func IsRegistered(key interface{}, kind int) bool {
// 	var (result bool
// 		sql string
// 		count int)
// 	switch kind{
// 	case EMAIL:
// 		// ckeck email
// 		sql := fmt.Sprintf("selec count(email) from local_auth where email＝%q", key)

// 	case: PHTONE:
// 		// ckeck phone
// 	}

// 	return result
// }
