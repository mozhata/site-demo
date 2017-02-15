package authenticate

import (
	"crypto/md5"
	"encoding/hex"
)

// LocalAuth local authorize
type LocalAuth struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	UUID   string `json:"uuid"`
	Phone  int64  `json:"phone"`
	Email  string `joson:"email"`
	Secret string `josn:"secret"`
}

func (loc *LocalAuth) GenerateSecret(pwd string) {
	loc.Secret = generateSecret(pwd, loc.UUID, loc.UserID)
}

func (loc *LocalAuth) IsValidPWD(pwd string) bool {
	return generateSecret(pwd, loc.UUID, loc.UserID) == loc.Secret
}

// IsValid should check valid before insert into DB
func (loc *LocalAuth) IsValid() bool {
	return loc.Secret != "" &&
		(loc.Email != "" || loc.Phone > 0) &&
		(loc.UserID >= 0 && loc.UUID != "")
}

func generateSecret(pwd, uuid string, userID int64) string {
	h := md5.New()
	h.Write([]byte(uuid[:userID%7]))
	h.Write([]byte(pwd))
	h.Write([]byte(uuid[userID%9:]))
	return hex.EncodeToString(h.Sum(nil))[:9]
}
