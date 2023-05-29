package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 10) //proses salting
	if err != nil {
		return "", err
	}
	return string(hashedByte), nil //kenapa harus di bungkus karena tipe dataannya string
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil //karena bool
}
