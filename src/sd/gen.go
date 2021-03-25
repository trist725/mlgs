//go:generate gen_static_data_go
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
