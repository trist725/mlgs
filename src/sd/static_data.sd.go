// 本文件由gen_static_data_go生成
// 请勿修改！！！

package sd

import "log"
import "path/filepath"

var (
	ItemMgr = newItemManager()
	ShopMgr = newShopManager()
)

func LoadAll(excelDir string) (success bool) {
	absExcelDir, err := filepath.Abs(excelDir)
	if err != nil {
		log.Println(err)
		return false
	}

	success = true

	success = ItemMgr.Load(filepath.Join(absExcelDir, "item.xlsx")) && success

	success = ShopMgr.Load(filepath.Join(absExcelDir, "shop.xlsx")) && success

	return
}

func AfterLoadAll(excelDir string) (success bool) {
	absExcelDir, err := filepath.Abs(excelDir)
	if err != nil {
		log.Println(err)
		return false
	}

	success = true

	success = ItemMgr.AfterLoadAll(filepath.Join(absExcelDir, "item.xlsx")) && success

	success = ShopMgr.AfterLoadAll(filepath.Join(absExcelDir, "shop.xlsx")) && success

	return
}
