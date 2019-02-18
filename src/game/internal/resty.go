package internal

import (
	"encoding/json"
	"github.com/trist725/myleaf/log"
	"gopkg.in/resty.v1"
	"mlgs/src/conf"
	"mlgs/src/model"
	"mlgs/src/msg"
	"time"
)

var (
	gLastUpTime int64
	gUpInterval int64 = 3

	gPubMails []MailInfo
	gNotices  []NoticeInfo
)

func PublicMails() *[]MailInfo {
	return &gPubMails
}

func Notices() *[]NoticeInfo {
	return &gNotices
}

type MailInfo struct {
	///邮件编号
	Id int64
	///奖励类型,对应item表id
	RewardType int64
	///奖励数量
	RewardNum int64
	///邮件内容
	Content string
}

type NoticeInfo struct {
	///类型
	Type int32
	//名称
	Name string
	///正文标题
	Title string
	///正文内容
	Content string
}

func (mi *MailInfo) ToModel() *model.Mail {
	m := model.Get_Mail()
	m.Id = mi.Id
	m.RewardType = mi.RewardType
	m.Content = mi.Content
	m.RewardNum = mi.RewardNum
	return m
}

func (mi *NoticeInfo) ToMsg() *msg.Notice {
	m := msg.Get_Notice()
	m.Content = mi.Content
	m.Type = mi.Type
	m.Name = mi.Name
	m.Title = mi.Title
	return m
}

func ConvertMails() (modelMails []*model.Mail) {
	for _, mi := range gPubMails {
		modelMails = append(modelMails, mi.ToModel())
	}
	return modelMails
}

func ConvertNotices() (notices []*msg.Notice) {
	for _, n := range gNotices {
		notices = append(notices, n.ToMsg())
	}
	return notices
}

func NeedUpdate() bool {
	if time.Now().Unix()-gLastUpTime < gUpInterval {
		return false
	}
	return true
}

func UpdateMails(ud *model.User) {
	if !NeedUpdate() {
		return
	}

	if err := GetNewMails(); err != nil {
		log.Error("[%s], GetNewMails: %v", err)
		return
	}
	ud.AddMails(ConvertMails())

	ids := []int64{}
	if err := GetOverdueMails(ids); err != nil {
		log.Error("[%s], GetOverdueMails: %v", err)
		return
	}
	ud.DelMails(ids)

	log.Debug("update mails: [%v]", time.Now())
}

func GetNewMails() error {
	uri := conf.Server.WebUrl + "/manage/game/mail/rewardMails"
	resp, err := resty.R().Get(uri)
	if err != nil {
		return err
	}

	gPubMails = nil
	err = json.Unmarshal(resp.Body(), &gPubMails)
	if err != nil {
		return err
	}

	gLastUpTime = time.Now().Unix()
	//log.Debug("update mails: [%v]", time.Now())
	return nil
}

func GetOverdueMails(ids []int64) error {
	uri := conf.Server.WebUrl + "/manage/game/mail/dated"
	resp, err := resty.R().Get(uri)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), &ids)
	if err != nil {
		return err
	}

	gLastUpTime = time.Now().Unix()
	//log.Debug("update mails: [%v]", time.Now())
	return nil
}

//func MailReceived(ids []int64) error{
//	uri := gRestProtocol + gRestAddr + "/manage/game/mail/rewardMails"
//	resp, err := resty.R().SetBody(ids).Put(uri)
//	if err != nil {
//		return err
//	}
//
//	var r int
//	if err = json.Unmarshal(resp.Body(), &r); err != nil {
//		return err
//	}
//	if r != len(ids){
//		return errors.New("invalid received mail count")
//	}
//	return nil
//}

func GetNotices() error {
	if !NeedUpdate() {
		return nil
	}

	uri := conf.Server.WebUrl + "/manage/game/activity/info"
	resp, err := resty.R().Get(uri)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), &gNotices)
	if err != nil {
		return err
	}

	gLastUpTime = time.Now().Unix()
	log.Debug("update notification: [%v]", time.Now())
	return nil
}
