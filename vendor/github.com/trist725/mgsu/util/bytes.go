package util

import "encoding/binary"

func Int32ToByteArr(src int32, littleEndian bool) []byte {
	dst := make([]byte, 4)
	if littleEndian {
		binary.LittleEndian.PutUint32(dst, uint32(src))
	} else {
		binary.BigEndian.PutUint32(dst, uint32(src))
	}
	return dst
}
