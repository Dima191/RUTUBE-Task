package hash

import (
	"golang.org/x/crypto/bcrypt"
	"unsafe"
)

func Password(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(([]byte)(unsafe.Slice(unsafe.StringData(password), len(password))), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return unsafe.String(unsafe.SliceData(hashed), len(hashed)), nil
}
