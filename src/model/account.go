package model

import (
	"fmt"
	"time"

	"github.com/trist725/mgsu/util"
)

func CreateAccount(accountName string, password string) (d *Account, err error) {
	nextSeq, err := NextSeq(TblAccount)
	if err != nil {
		return nil, err
	}
	d = Get_Account()
	d.ID = int64(nextSeq)
	d.Name = accountName
	d.Password = password
	d.RegisterTime = time.Now().Unix()
	return
}

func CreateVisitorAccount() (d *Account, err error) {
	nextSeq, err := NextSeq(TblAccount)
	if err != nil {
		return nil, err
	}
	nextVisitorNameSeq, err := NextSeq(visitorNameSeq)
	if err != nil {
		return nil, err
	}
	d = Get_Account()
	d.ID = int64(nextSeq)
	d.Name = fmt.Sprintf("游客%d", nextVisitorNameSeq)
	d.VName = d.Name
	d.Password = util.GenRandomString(8)
	d.RegisterTime = time.Now().Unix()
	return
}
