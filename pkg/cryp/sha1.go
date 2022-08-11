package cryp

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
