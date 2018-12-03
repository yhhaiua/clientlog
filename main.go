package main

import (
	"github.com/yhhaiua/clientlog/logic"
	"github.com/yhhaiua/log4go"
	"time"
)

func main() {

	log4go.LoadConfiguration("config/log4j.xml")
	if logic.Instance().LogicInit(){
		log4go.Info("clientlog 启动成功")
	}else{
		log4go.Error("clientlog 启动失败")
	}
	time.Sleep(3*time.Second)
}
