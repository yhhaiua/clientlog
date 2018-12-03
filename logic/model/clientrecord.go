package model

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ClientRecord struct {
	Id 		int
	Account string
	Type 	int
	Time    time.Time `orm:"auto_now_add;type(datetime)"`
	Step    int
	Onlyid  int64
	Logintime time.Time `orm:"auto_now_add;type(datetime)"`
	Os      string
	Msg		string		`orm:"type(text)"`
	Username string
	Platformid string
}

func (u *ClientRecord) TableName() string {
	return "client_record"
}

//Insert 插入一条日志
func (record *ClientRecord)Insert(Type int,Time int64,Step int,Onlyid int64,Logintime int64,Os,Msg,UserName,Platformid,account string) error  {

	record.Type = Type
	record.Time = time.Unix(Time/1000, 0)
	record.Step = Step
	record.Onlyid = Onlyid
	record.Logintime = time.Unix(Logintime/1000, 0)
	record.Os = Os
	record.Msg = Msg
	record.Username = UserName
	record.Platformid = Platformid
	record.Account = account
	//orm.Debug = true
	o := orm.NewOrm()
	_, err := o.Insert(record)
	return err
}
