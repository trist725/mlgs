//go:generate gen-static-data-go
//go:generate protoc --enum-go_out=. enum.proto global.proto
package sd

import "mlgs/src/conf"

func init() {
	success := LoadAll(conf.Server.XlsxPath)
	if success != true {
		panic("sd LoadAll faild")
	}
	success = AfterLoadAll(conf.Server.XlsxPath)
	if success != true {
		panic("sd AfterLoadAll faild")
	}
}
