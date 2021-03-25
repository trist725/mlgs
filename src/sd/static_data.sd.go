// 本文件由gen_static_data_go生成
// 请勿修改！！！

package sd

import "log"
import "path/filepath"

var (
	GlobalMgr = newGlobalManager()
)

func LoadAll(excelDir string) (success bool) {
	absExcelDir, err := filepath.Abs(excelDir)
	if err != nil {
		log.Println(err)
		return false
	}

	success = true

	success = GlobalMgr.Load(filepath.Join(absExcelDir, "global.xlsx")) && success

	return
}

func AfterLoadAll(excelDir string) (success bool) {
	absExcelDir, err := filepath.Abs(excelDir)
	if err != nil {
		log.Println(err)
		return false
	}

	success = true

	success = GlobalMgr.AfterLoadAll(filepath.Join(absExcelDir, "global.xlsx")) && success

	return
}
