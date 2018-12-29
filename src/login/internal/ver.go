package internal

import (
	"encoding/json"
	"github.com/trist725/myleaf/log"
	"gopkg.in/resty.v1"
	"strconv"
	"time"
)

var (
	gAllowCltVer   CltVer
	gLastUpVerTime int64
	gUpVerInterval int64 = 10
)

func AllowCltVer() *CltVer {
	return &gAllowCltVer
}

type CltVer struct {
	//大版本
	BigVer int32
	//小版本
	SmallVer int32
	//修复版本
	FixVer int32
}

func GetAllowCltVer(dt int32, ver *CltVer) error {
	if time.Now().Unix()-gLastUpVerTime < gUpVerInterval {
		return nil
	}
	uri := "http://192.168.101.230:8000/manage/game/client/version/" + strconv.Itoa(int(dt))
	resp, err := resty.R().Get(uri)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), &ver)
	if err != nil {
		return err
	}

	gLastUpVerTime = time.Now().Unix()
	return nil
}

func CheckCltVer(loginVer *CltVer, dt int32) bool {
	//暂时不用
	return true

	if err := GetAllowCltVer(dt, &gAllowCltVer); err != nil {
		log.Error("GetAllowCltVer failed: %v", err)
		return false
	}
	if loginVer.BigVer == gAllowCltVer.BigVer &&
		loginVer.SmallVer == gAllowCltVer.SmallVer &&
		loginVer.FixVer == gAllowCltVer.FixVer {
		return true
	}
	return false
}
