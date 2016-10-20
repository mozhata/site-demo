package sqldb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Env struct {
	*gorm.DB
}

func NewEnv() *Env {
	db, err := gorm.Open("mysql", "root:root@tcp(mysql:3306)/pad?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return &Env{db}
}
