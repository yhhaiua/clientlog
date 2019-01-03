package main

import (
	"github.com/yhhaiua/clientlog/logic"
	"github.com/yhhaiua/engine/log"
	"time"
)

var gLog = log.GetLogger()

func main() {

	gLog.Config("config/log4j.xml")
	if logic.Instance().LogicInit(){
		gLog.Info("clientlog 启动成功")
	}else{
		gLog.Error("clientlog 启动失败")
	}
	time.Sleep(3*time.Second)
}
