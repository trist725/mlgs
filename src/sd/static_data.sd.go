// 本文件由gen_static_data_go生成
// 请勿修改！！！

package sd

import "log"
import "path/filepath"

var (
	GlobalMgr = newGlobalManager()
	ItemMgr   = newItemManager()
	PersonMgr = newPersonManager()
	RoomMgr   = newRoomManager()
	ShopMgr   = newShopManager()
	TimeMgr   = newTimeManager()
)

func LoadAll(excelDir string) (success bool) {
	absExcelDir, err := filepath.Abs(excelDir)
	if err != nil {
		log.Println(err)
		return false
	}

	success = true

	success = GlobalMgr.Load(filepath.Join(absExcelDir, "global.xlsx")) && success

	success = ItemMgr.Load(filepath.Join(absExcelDir, "item.xlsx")) && success

	success = PersonMgr.Load(filepath.Join(absExcelDir, "person.xlsx")) && success

	success = RoomMgr.Load(filepath.Join(absExcelDir, "room.xlsx")) && success

	success = ShopMgr.Load(filepath.Join(absExcelDir, "shop.xlsx")) && success

	success = TimeMgr.Load(filepath.Join(absExcelDir, "time.xlsx")) && success

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

	success = ItemMgr.AfterLoadAll(filepath.Join(absExcelDir, "item.xlsx")) && success

	success = PersonMgr.AfterLoadAll(filepath.Join(absExcelDir, "person.xlsx")) && success

	success = RoomMgr.AfterLoadAll(filepath.Join(absExcelDir, "room.xlsx")) && success

	success = ShopMgr.AfterLoadAll(filepath.Join(absExcelDir, "shop.xlsx")) && success

	success = TimeMgr.AfterLoadAll(filepath.Join(absExcelDir, "time.xlsx")) && success

	return
}
