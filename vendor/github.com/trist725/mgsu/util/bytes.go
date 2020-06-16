package util

import (
	"encoding/binary"
)

func UInt16ToByteArr(src uint16, littleEndian bool) []byte {
	dst := make([]byte, 2)
	if littleEndian {
		binary.LittleEndian.PutUint16(dst, src)
	} else {
		binary.BigEndian.PutUint16(dst, src)
	}
	return dst
}

func UInt32ToByteArr(src uint32, littleEndian bool) []byte {
	dst := make([]byte, 4)
	if littleEndian {
		binary.LittleEndian.PutUint32(dst, src)
	} else {
		binary.BigEndian.PutUint32(dst, src)
	}
	return dst
}

func UInt64ToByteArr(src uint64, littleEndian bool) []byte {
	dst := make([]byte, 8)
	if littleEndian {
		binary.LittleEndian.PutUint64(dst, src)
	} else {
		binary.BigEndian.PutUint64(dst, src)
	}
	return dst
}
