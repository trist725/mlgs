// 本文件由gen_static_data_go生成
// 请遵照提示添加修改！！！

package sd

import "encoding/json"
import "fmt"
import "log"
import "path/filepath"

import "github.com/tealeg/xlsx"
import "github.com/trist725/mgsu/util"

//////////////////////////////////////////////////////////////////////////////////////////////////
// TODO 添加扩展import代码
//import_extend_begin
//import_extend_end
//////////////////////////////////////////////////////////////////////////////////////////////////

type Room struct {
	ID int64 `excel_column:"0" excel_name:"id"` // 编号

	Des string `excel_column:"3" excel_name:"des"` // 描述

	Chairlimit int `excel_column:"4" excel_name:"chairlimit"` // 座位数限制（玩家个数）

	Playerlimit int `excel_column:"5" excel_name:"playerlimit"` // 房间内的总人数限制

	Sb int64 `excel_column:"6" excel_name:"sb"` // 小盲注

	Bb int64 `excel_column:"7" excel_name:"bb"` // 大盲注

	Limit int64 `excel_column:"8" excel_name:"limit"` // 押注上限（0为无上限）

	Chip int64 `excel_column:"9" excel_name:"chip"` // 带入场的筹码

	Inning int `excel_column:"10" excel_name:"inning"` // 局数（0代表无限局）

	Handcard int `excel_column:"11" excel_name:"handcard"` // 手牌数

	Pool int `excel_column:"12" excel_name:"pool"` // 牌池

	Startpublicpokes int `excel_column:"13" excel_name:"startpublicpokes"` // 起始公共牌

	Totalpublicpokes int `excel_column:"14" excel_name:"totalpublicpokes"` // 总共公共牌

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体扩展字段
	//struct_extend_begin
	//struct_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func NewRoom() *Room {
	sd := &Room{}
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体New代码
	//struct_new_begin
	//struct_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return sd
}

func (sd Room) String() string {
	ba, _ := json.Marshal(sd)
	return string(ba)
}

func (sd Room) Clone() *Room {
	n := NewRoom()
	*n = sd

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体Clone代码
	//struct_clone_begin
	//struct_clone_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return n
}

func (sd *Room) load(row *xlsx.Row) error {
	return util.DeserializeStructFromXlsxRow(sd, row)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
type RoomManager struct {
	dataArray []*Room
	dataMap   map[int64]*Room

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager扩展字段
	//manager_extend_begin
	//manager_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func newRoomManager() *RoomManager {
	mgr := &RoomManager{
		dataArray: []*Room{},
		dataMap:   make(map[int64]*Room),
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager的New代码
	//manager_new_begin
	//manager_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return mgr
}

func (mgr *RoomManager) Load(excelFilePath string) (success bool) {
	success = true

	absExcelFilePath, err := filepath.Abs(excelFilePath)
	if err != nil {
		log.Printf("获取 %s 的绝对路径失败, %s", excelFilePath, err)
		return false
	}

	xl, err := xlsx.OpenFile(absExcelFilePath)
	if err != nil {
		log.Printf("打开 %s 失败, %s\n", excelFilePath, err)
		return false
	}

	if len(xl.Sheets) == 0 {
		log.Printf("%s 没有分页可加载\n", excelFilePath)
		return false
	}

	dataSheet, ok := xl.Sheet["data"]
	if !ok {
		log.Printf("%s 没有data分页\n", excelFilePath)
		return false
	}

	if len(dataSheet.Rows) < 3 {
		log.Printf("%s 数据少于3行\n", excelFilePath)
		return false
	}

	for i := 3; i < len(dataSheet.Rows); i++ {
		row := dataSheet.Rows[i]
		if len(row.Cells) <= 0 {
			continue
		}

		firstColumn := row.Cells[0]
		firstComment := firstColumn.String()
		if firstComment != "" {
			if firstComment[0] == '#' {
				// 跳过被注释掉的行
				continue
			}
		}

		sd := NewRoom()
		err = sd.load(row)
		if err != nil {
			log.Printf("%s 加载第%d行失败, %s\n", excelFilePath, i+1, err)
			success = false
			continue
		}

		if sd.ID == 0 {
			continue
		}

		//////////////////////////////////////////////////////////////////////////////////////////////////
		// TODO 添加结构体加载代码
		//struct_load_begin
		//struct_load_end
		//////////////////////////////////////////////////////////////////////////////////////////////////

		if err := mgr.check(excelFilePath, i+1, sd); err != nil {
			log.Println(err)
			success = false
			continue
		}

		mgr.dataArray = append(mgr.dataArray, sd)
		mgr.dataMap[sd.ID] = sd

		//////////////////////////////////////////////////////////////////////////////////////////////////
		// TODO 添加manager加载代码
		//manager_load_begin
		//manager_load_end
		//////////////////////////////////////////////////////////////////////////////////////////////////
	}

	return
}

func (mgr RoomManager) Size() int {
	return len(mgr.dataArray)
}

func (mgr RoomManager) Get(id int64) *Room {
	sd, ok := mgr.dataMap[id]
	if !ok {
		return nil
	}
	return sd.Clone()
}

func (mgr RoomManager) Each(f func(sd *Room) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd.Clone()) {
			break
		}
	}
}

func (mgr *RoomManager) each(f func(sd *Room) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd) {
			break
		}
	}
}

func (mgr RoomManager) findIf(f func(sd *Room) bool) *Room {
	for _, sd := range mgr.dataArray {
		if f(sd) {
			return sd
		}
	}
	return nil
}

func (mgr RoomManager) FindIf(f func(sd *Room) bool) *Room {
	for _, sd := range mgr.dataArray {
		n := sd.Clone()
		if f(n) {
			return n
		}
	}
	return nil
}

func (mgr RoomManager) check(excelFilePath string, row int, sd *Room) error {
	if _, ok := mgr.dataMap[sd.ID]; ok {
		return fmt.Errorf("%s 第%d行的id重复", excelFilePath, row)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加检查代码
	//check_begin
	//check_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return nil
}

func (mgr *RoomManager) AfterLoadAll(excelFilePath string) (success bool) {
	success = true
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加加载后处理代码
	//after_load_all_begin
	//after_load_all_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// TODO 添加扩展代码
//extend_begin
//extend_end
//////////////////////////////////////////////////////////////////////////////////////////////////
