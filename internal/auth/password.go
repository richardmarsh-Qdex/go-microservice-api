package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(plain string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
}

func ComparePassword(hashed []byte, plain string) error {
	return bcrypt.CompareHashAndPassword(hashed, []byte(plain))
}
