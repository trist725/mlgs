// 本文件由gen_static_data_go生成
// 请遵照提示添加修改！！！

package sd

import (
	"encoding/json"
	"strconv"
)
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

type Global struct {
	ID int64 `excel_column:"0" excel_name:"id"` // 公共配置表ID

	Value string `excel_column:"1" excel_name:"value"` // 配置值

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体扩展字段
	//struct_extend_begin
	//struct_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func NewGlobal() *Global {
	sd := &Global{}
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体New代码
	//struct_new_begin
	//struct_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return sd
}

func (sd Global) String() string {
	ba, _ := json.Marshal(sd)
	return string(ba)
}

func (sd Global) Clone() *Global {
	n := NewGlobal()
	*n = sd

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体Clone代码
	//struct_clone_begin
	//struct_clone_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return n
}

func (sd *Global) load(row *xlsx.Row) error {
	return util.DeserializeStructFromXlsxRow(sd, row)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
type GlobalManager struct {
	dataArray []*Global
	dataMap   map[int64]*Global

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager扩展字段
	//manager_extend_begin
	//manager_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func newGlobalManager() *GlobalManager {
	mgr := &GlobalManager{
		dataArray: []*Global{},
		dataMap:   make(map[int64]*Global),
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager的New代码
	//manager_new_begin
	//manager_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return mgr
}

func (mgr *GlobalManager) Load(excelFilePath string) (success bool) {
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

		sd := NewGlobal()
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

func (mgr GlobalManager) Size() int {
	return len(mgr.dataArray)
}

func (mgr GlobalManager) Get(id int64) *Global {
	sd, ok := mgr.dataMap[id]
	if !ok {
		return nil
	}
	return sd.Clone()
}

func (mgr GlobalManager) Each(f func(sd *Global) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd.Clone()) {
			break
		}
	}
}

func (mgr *GlobalManager) each(f func(sd *Global) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd) {
			break
		}
	}
}

func (mgr GlobalManager) findIf(f func(sd *Global) bool) *Global {
	for _, sd := range mgr.dataArray {
		if f(sd) {
			return sd
		}
	}
	return nil
}

func (mgr GlobalManager) FindIf(f func(sd *Global) bool) *Global {
	for _, sd := range mgr.dataArray {
		n := sd.Clone()
		if f(n) {
			return n
		}
	}
	return nil
}

func (mgr GlobalManager) check(excelFilePath string, row int, sd *Global) error {
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

func (mgr *GlobalManager) AfterLoadAll(excelFilePath string) (success bool) {
	success = true
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加加载后处理代码
	//after_load_all_begin
	{
		d, ok := mgr.dataMap[int64(E_Global_InitUserDataId)]
		if !ok {
			log.Fatal("获取用户初始数据在person表id失败")
			return false
		}
		tid, err := strconv.Atoi(d.Value)
		if err != nil {
			log.Fatal("获取用户初始数据在person表id失败, %v", err)
			return false
		}
		gInitUserDataId = int64(tid)
		if gInitUserDataId <= 0 {
			log.Fatal("获取用户初始数据在person表id值有误, [%d]", tid)
			return false
		}
	}
	//after_load_all_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// TODO 添加扩展代码
//extend_begin

//用户初始化数据在person表的id
var gInitUserDataId int64

func InitUserDataId() int64 {
	return gInitUserDataId
}

//extend_end
//////////////////////////////////////////////////////////////////////////////////////////////////
