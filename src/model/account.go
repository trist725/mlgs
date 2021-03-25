package model

import (
	"mlgs/src/msg"
)

func CreateAccount(recv *msg.C2S_Login) (d *Account, err error) {

	return
}

//func CreateVisitorAccount() (d *Account, err error) {
//	nextSeq, err := NextSeq(TblAccount)
//	if err != nil {
//		return nil, err
//	}
//	nextVisitorNameSeq, err := NextSeq(visitorNameSeq)
//	if err != nil {
//		return nil, err
//	}
//	d = Get_Account()
//	d.ID = int64(nextSeq)
//	d.Name = fmt.Sprintf("游客%d", nextVisitorNameSeq)
//	d.VName = d.Name
//	d.Password = util.GenRandomString(8)
//	d.RegisterTime = time.Now().Unix()
//	return
//}
