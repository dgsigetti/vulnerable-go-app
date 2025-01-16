package utils

import (
	"crypto/md5" // Vulnerable: Weak hashing algorithm
	"fmt"
)

func Hash(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

