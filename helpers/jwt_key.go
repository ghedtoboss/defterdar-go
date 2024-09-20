package helpers

import "os"

func GetJwtKey() []byte {
	return []byte(os.Getenv("JWT_KEY"))
}
