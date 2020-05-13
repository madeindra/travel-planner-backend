package helpers

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pwd string) string {
	var stringHash string

	bytePwd := []byte(pwd)
	hash, _ := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)

	stringHash = string(hash)
	return stringHash
}

func ComparePasswords(hashedPwd string, plainPwd string) error {
	var err error

	bytePwd := []byte(plainPwd)
	byteHash := []byte(hashedPwd)
	err = bcrypt.CompareHashAndPassword(byteHash, bytePwd)

	return err
}
