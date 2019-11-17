package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncDec(t *testing.T) {

	password := "123456"
	wrongPassword := "1234"
	data := []byte("你好,QOS钱包")

	encBytes, saltBytes := Encrypt(data, password)

	decData, err := Decrypt(encBytes, saltBytes, password)
	assert.Equal(t, data, decData)
	assert.Nil(t, err)

	decData, err = Decrypt(encBytes, saltBytes, wrongPassword)
	assert.Nil(t, decData)
	assert.Equal(t, err, NewErrWrongPassword())
}
