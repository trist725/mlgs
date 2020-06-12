// 本文件由gen_static_data_go生成
// 请勿修改！！！

package sd

import (
	"log"
	"path/filepath"
)

var (
	AchieveMgr     = newAchieveManager()
	CompetitionMgr = newCompetitionManager()
	EmailMgr       = newEmailManager()
	GlobalMgr      = newGlobalManager()
	ItemMgr        = newItemManager()
	PersonMgr      = newPersonManager()
	ShopMgr        = newShopManager()
	TaskMgr        = newTaskManager()
)

func LoadAll(excelDir string) (success bool) {
	absExcelDir, err := filepath.Abs(excelDir)
	if err != nil {
		log.Println(err)
		return false
	}

	success = true

	success = AchieveMgr.Load(filepath.Join(absExcelDir, "achieve.xlsx")) && success

	success = CompetitionMgr.Load(filepath.Join(absExcelDir, "competition.xlsx")) && success

	success = EmailMgr.Load(filepath.Join(absExcelDir, "email.xlsx")) && success

	success = GlobalMgr.Load(filepath.Join(absExcelDir, "global.xlsx")) && success

	success = ItemMgr.Load(filepath.Join(absExcelDir, "item.xlsx")) && success

	success = PersonMgr.Load(filepath.Join(absExcelDir, "person.xlsx")) && success

	success = ShopMgr.Load(filepath.Join(absExcelDir, "shop.xlsx")) && success

	success = TaskMgr.Load(filepath.Join(absExcelDir, "task.xlsx")) && success

	return
}

func AfterLoadAll(excelDir string) (success bool) {
	absExcelDir, err := filepath.Abs(excelDir)
	if err != nil {
		log.Println(err)
		return false
	}

	success = true

	success = AchieveMgr.AfterLoadAll(filepath.Join(absExcelDir, "achieve.xlsx")) && success

	success = CompetitionMgr.AfterLoadAll(filepath.Join(absExcelDir, "competition.xlsx")) && success

	success = EmailMgr.AfterLoadAll(filepath.Join(absExcelDir, "email.xlsx")) && success

	success = GlobalMgr.AfterLoadAll(filepath.Join(absExcelDir, "global.xlsx")) && success

	success = ItemMgr.AfterLoadAll(filepath.Join(absExcelDir, "item.xlsx")) && success

	success = PersonMgr.AfterLoadAll(filepath.Join(absExcelDir, "person.xlsx")) && success

	success = ShopMgr.AfterLoadAll(filepath.Join(absExcelDir, "shop.xlsx")) && success

	success = TaskMgr.AfterLoadAll(filepath.Join(absExcelDir, "task.xlsx")) && success

	return
}
