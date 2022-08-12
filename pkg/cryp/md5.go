package cryp

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5(str string) string {
	md5Ctx := md5.New()
	io.WriteString(md5Ctx, str)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
