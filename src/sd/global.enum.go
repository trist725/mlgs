// Code generated by protoc-gen-enum-go. DO NOT EDIT IT!!!
// source: global.proto

/*
It has these top-level messages:
*/

package sd

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [E_Global] begin

type E_Global int32

const (
	E_Global_ E_Global = 0
	///初始用户数据在person表的id
	E_Global_InitUserDataId E_Global = 1

	E_Global_InitMatchRoomId E_Global = 2

	E_Global_MinStartGamePlayer E_Global = 3
)

var E_Global_name = map[int32]string{
	0: "E_Global_",
	1: "E_Global_InitUserDataId",
	2: "E_Global_InitMatchRoomId",
	3: "E_Global_MinStartGamePlayer",
}

var E_Global_value = map[string]int32{
	"E_Global_":                   0,
	"E_Global_InitUserDataId":     1,
	"E_Global_InitMatchRoomId":    2,
	"E_Global_MinStartGamePlayer": 3,
}

var E_Global_Slice = []int32{
	0,
	1,
	2,
	3,
}

func (x E_Global) String() string {
	if name, ok := E_Global_name[int32(x)]; ok {
		return name
	}
	return ""
}

func E_Global_Len() int {
	return len(E_Global_Slice)
}

func Check_E_Global_I(value int32) bool {
	if _, ok := E_Global_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_E_Global(value E_Global) bool {
	return Check_E_Global_I(int32(value))
}

func Each_E_Global(f func(E_Global) bool) {
	for _, value := range E_Global_Slice {
		if !f(E_Global(value)) {
			break
		}
	}
}

func Each_E_Global_I(f func(int32) bool) {
	for _, value := range E_Global_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [E_Global] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////