package security

import "golang.org/x/crypto/bcrypt"

func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func CheckPassword(senhaHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
