package dao

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
	"log"
)

// GenerateSalt 生成一个随机的salt字符串
func GenerateSalt() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

// PasswordEncrypt 对密码进行加密
func PasswordEncrypt(password string, salt string) (string, error) {
	publishedPassword := []byte(salt + password)
	hashedPassword, err := bcrypt.GenerateFromPassword(publishedPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	passwordHash := string(hashedPassword)
	return passwordHash, nil
}

// PasswordMatch 比较明文密码和密码哈希是否匹配
func PasswordMatch(password string, passwordHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}
