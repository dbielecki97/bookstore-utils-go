package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dbielecki97/bookstore-utils-go/logger"
	"golang.org/x/crypto/bcrypt"
)

func GetMd5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		logger.Error("error when hashing password", err)
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

func Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
