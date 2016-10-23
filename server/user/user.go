package user

import "github.com/pborman/uuid"

type User struct {
	ID       int64  `json:"id"`
	UUID     string `json:"uuid"`
	NickName string `json:"nick_name"`
}

func NewUser(nickName string) *User {
	return &User{
		UUID:     uuid.New(),
		NickName: nickName,
	}
}
