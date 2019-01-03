package control

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/yhhaiua/clientlog/logic/model"
	"github.com/yhhaiua/engine/grouter"
	"github.com/yhhaiua/engine/log"
	"net/http"
	"strconv"
)

var gLog = log.GetLogger()

type LogControl struct {
	Clientkey string
}

func (control *LogControl)LogNote(w http.ResponseWriter, r *http.Request, _ grouter.Params)  {
	Type := r.FormValue("type")
	time := r.FormValue("time")
	account := r.FormValue("account")
	sign := r.FormValue("sign")
	md5str := Type + time +account + control.Clientkey
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(md5str))
	cipherStr := md5Ctx.Sum(nil)
	mysigon := hex.EncodeToString(cipherStr)
	if mysigon != sign{
		gLog.Error("md5 error : me:%s,client:%s",mysigon,sign)
		control.send(w,"md5 error")
		return;
	}
	onlyid := r.FormValue("onlyid")
	step := r.FormValue("step")
	logintime := r.FormValue("logintime")
	os := r.FormValue("os")
	msg := r.FormValue("msg")
	username := r.FormValue("username")
	platformid := r.FormValue("platformid")

	iType,_:= strconv.Atoi(Type)
	Time, _ := strconv.ParseInt(time, 10, 64)
	Step,_:= strconv.Atoi(step)
	Onlyid, _ := strconv.ParseInt(onlyid, 10, 64)
	Logintime, _ := strconv.ParseInt(logintime, 10, 64)

	record := new(model.ClientRecord)

	err := record.Insert(iType,Time,Step,Onlyid,Logintime,os,msg,username,platformid,account)
	if err != nil{
		gLog.Error("save error:%s",err)
		control.send(w,"save error")
	}else{
		control.send(w,"ok")
	}
}

func (control *LogControl)send(w http.ResponseWriter,msg string)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	fmt.Fprintf(w, "%s", msg)
}