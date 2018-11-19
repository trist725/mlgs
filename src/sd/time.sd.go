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

type Time struct {
	ID int64 `excel_column:"0" excel_name:"id"` // 编号

	Value int `excel_column:"2" excel_name:"value"` // 数值

	Des string `excel_column:"3" excel_name:"des"` // 描述

	Res string `excel_column:"4" excel_name:"res"` // 结果

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体扩展字段
	//struct_extend_begin
	//struct_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func NewTime() *Time {
	sd := &Time{}
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体New代码
	//struct_new_begin
	//struct_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return sd
}

func (sd Time) String() string {
	ba, _ := json.Marshal(sd)
	return string(ba)
}

func (sd Time) Clone() *Time {
	n := NewTime()
	*n = sd

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体Clone代码
	//struct_clone_begin
	//struct_clone_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return n
}

func (sd *Time) load(row *xlsx.Row) error {
	return util.DeserializeStructFromXlsxRow(sd, row)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
type TimeManager struct {
	dataArray []*Time
	dataMap   map[int64]*Time

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager扩展字段
	//manager_extend_begin
	//manager_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func newTimeManager() *TimeManager {
	mgr := &TimeManager{
		dataArray: []*Time{},
		dataMap:   make(map[int64]*Time),
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager的New代码
	//manager_new_begin
	//manager_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return mgr
}

func (mgr *TimeManager) Load(excelFilePath string) (success bool) {
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

		sd := NewTime()
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

func (mgr TimeManager) Size() int {
	return len(mgr.dataArray)
}

func (mgr TimeManager) Get(id int64) *Time {
	sd, ok := mgr.dataMap[id]
	if !ok {
		return nil
	}
	return sd.Clone()
}

func (mgr TimeManager) Each(f func(sd *Time) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd.Clone()) {
			break
		}
	}
}

func (mgr *TimeManager) each(f func(sd *Time) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd) {
			break
		}
	}
}

func (mgr TimeManager) findIf(f func(sd *Time) bool) *Time {
	for _, sd := range mgr.dataArray {
		if f(sd) {
			return sd
		}
	}
	return nil
}

func (mgr TimeManager) FindIf(f func(sd *Time) bool) *Time {
	for _, sd := range mgr.dataArray {
		n := sd.Clone()
		if f(n) {
			return n
		}
	}
	return nil
}

func (mgr TimeManager) check(excelFilePath string, row int, sd *Time) error {
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

func (mgr *TimeManager) AfterLoadAll(excelFilePath string) (success bool) {
	success = true
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加加载后处理代码
	//after_load_all_begin

	//开局发牌时间
	//菜鸡策划真心教不会配表，先写死了
	{
		timeSd := TimeMgr.Get(5)
		gDealCardTime = timeSd.Value
	}
	//第1阶段行动时间
	//菜鸡策划真心教不会配表，先写死了
	{
		timeSd := TimeMgr.Get(6)
		gActionTime_S1 = timeSd.Value
	}
	//第2阶段行动时间
	//菜鸡策划真心教不会配表，先写死了
	{
		timeSd := TimeMgr.Get(7)
		gActionTime_S2 = timeSd.Value
	}
	//第3阶段行动时间
	//菜鸡策划真心教不会配表，先写死了
	{
		timeSd := TimeMgr.Get(8)
		gActionTime_S3 = timeSd.Value
	}
	//第4阶段行动时间
	//菜鸡策划真心教不会配表，先写死了
	{
		timeSd := TimeMgr.Get(9)
		gActionTime_S4 = timeSd.Value
	}
	//after_load_all_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// TODO 添加扩展代码
//extend_begin

//开局发牌时间
var gDealCardTime int

func InitDealCardTime() int {
	if gDealCardTime < 1 {
		log.Fatal("开局发牌时间 配表有误")
	}
	return gMinStartGamePlayer
}

//第1阶段行动时间
var gActionTime_S1 int

func InitActionTime_S1() int {
	if gActionTime_S1 < 1 {
		log.Fatal("第1阶段行动时间 配表有误")
	}
	return gActionTime_S1
}

//第2阶段行动时间
var gActionTime_S2 int

func InitActionTime_S2() int {
	if gActionTime_S2 < 1 {
		log.Fatal("第2阶段行动时间 配表有误")
	}
	return gActionTime_S2
}

//第3阶段行动时间
var gActionTime_S3 int

func InitActionTime_S3() int {
	if gActionTime_S3 < 1 {
		log.Fatal("第3阶段行动时间 配表有误")
	}
	return gActionTime_S3
}

//第4阶段行动时间
var gActionTime_S4 int

func InitActionTime_S4() int {
	if gActionTime_S4 < 1 {
		log.Fatal("第4阶段行动时间 配表有误")
	}
	return gActionTime_S4
}

//extend_end
//////////////////////////////////////////////////////////////////////////////////////////////////
