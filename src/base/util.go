package base

import (
	"mlgs/src/conf"

	"github.com/trist725/mgsu/util"
)

func Int32ToByteArr(src int32) []byte {
	return util.UInt32ToByteArr(uint32(src), conf.LittleEndian)
}

func Uint16ToByteArr(src uint16) []byte {
	return util.UInt16ToByteArr(src, conf.LittleEndian)
}
