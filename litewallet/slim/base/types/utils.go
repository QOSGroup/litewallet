package types

import (
	"bytes"
	"encoding/binary"
	"log"
)

// 函数：int64 转化为 []byte
func Int2Byte(in int64) []byte {
	var ret = bytes.NewBuffer([]byte{})
	err := binary.Write(ret, binary.BigEndian, in)
	if err != nil {
		log.Printf("Int2Byte error:%s", err.Error())
		return nil
	}

	return ret.Bytes()
}

// 函数：bool 转化为 []byte
func Bool2Byte(in bool) []byte {
	if in {
		return []byte{1}
	}
	return []byte{0}
}
