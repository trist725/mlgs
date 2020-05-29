//go:generate gen_static_data_go
//go:generate protoc --enum-go_out=. enum.proto global.proto
package sd

import "mlgs/src/conf"

func init() {
	success := LoadAll(conf.Server.XlsxPath)
	if !success {
		panic("sd LoadAll failed")
	}
	success = AfterLoadAll(conf.Server.XlsxPath)
	if !success {
		panic("sd AfterLoadAll failed")
	}
}
