package main

import "encoding/binary"

func Int32ToBytes(i int) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}
