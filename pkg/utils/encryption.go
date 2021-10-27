package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// MD5 use md5 encrypt
func MD5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func MD5Byte(data []byte) [16]byte {
	return md5.Sum(data)
}

// Base64En encod []byte
func Base64En(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64De encod string
func Base64De(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func Sha256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}
