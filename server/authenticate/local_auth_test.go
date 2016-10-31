package authenticate

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"testing"

	"github.com/pborman/uuid"

	"golang.org/x/crypto/bcrypt"
)

const (
	pwd  = "this is a password"
	salt = "salt"
)

func TestGenerateSecret(t *testing.T) {
	uuid := uuid.New()
	fmt.Println(generateSecret(pwd, uuid, 99))
}

/*
func BenchmarkEncryptPWD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashed := encryptPWD(pwd)
		fmt.Sprintf("%s", hashed)
	}
}*/

func BenchmarkStr(b *testing.B) {
	str := ""
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%b", str == "")
	}
}

func BenchmarkLenStr(b *testing.B) {
	str := ""
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%b", len(str) == 0)
	}
}

/*func BenchmarkCheckPWD(b *testing.B) {
	hashed := encryptPWD(pwd)
	for i := 0; i < b.N; i++ {
		checkPWD(hashed, pwd)
	}
}

func BenchmarkEncryByMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		secrect := encryByMd5(pwd)
		fmt.Sprintf("%s", secrect)
	}
}

func BenchmarkCheckMD5(b *testing.B) {
	secret := encryptPWD(pwd)
	for i := 0; i < b.N; i++ {
		checkMD5(secret, pwd)
	}
}

func BenchmarkEncryBySha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		secrect := encryBySha(pwd)
		fmt.Sprintf("%s", secrect)
	}
}

func BenchmarkCheckSha(b *testing.B) {
	secret := encryBySha(pwd)
	for i := 0; i < b.N; i++ {
		checkSha(secret, pwd)
	}
}*/

func encryByMd5(pwd string) string {
	encryed := md5.Sum([]byte(fmt.Sprintf("%s%s", pwd, salt)))
	return fmt.Sprintf("%x", encryed)
}

func checkMD5(encryed, pwd string) bool {
	actual := md5.Sum([]byte(fmt.Sprintf("%s%s", pwd, salt)))
	return fmt.Sprintf("%x", actual) == encryed
}

func encryBySha(pwd string) string {
	encryed := sha1.Sum([]byte(fmt.Sprintf("%s%s", pwd, salt)))
	return fmt.Sprintf("%x", encryed)
}

func checkSha(encryed, pwd string) bool {
	actual := sha1.Sum([]byte(fmt.Sprintf("%s%s", pwd, salt)))
	return fmt.Sprintf("%x", actual) == encryed
}

func encryptPWD(pwd string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), 3)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

func checkPWD(hashed, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd))
	return err != nil
}
