package main

import (
	"encoding/binary"
)

func Int32ToBytes(i int) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func StringToArray[V [userNameSize]byte | [emailSize]byte](s string) V {
	var res V
	for k, v := range s {
		res[k] = byte(v)
	}
	return res
}

func ArrayToString[V [userNameSize]byte | [emailSize]byte](arr V) string {
	s := ""
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			s += string(arr[i])
		}
	}
	return s
}
