package logic

import (
	"github.com/yhhaiua/clientlog/logic/config"
	"github.com/yhhaiua/clientlog/logic/control"
	"github.com/yhhaiua/engine/grouter"
	"github.com/yhhaiua/log4go"
	"net/http"
	"sync"
	"time"
)

var (
	instance *LogicSvr
	mu       sync.Mutex
)

//LogicSvr 服务器数据
type LogicSvr struct {
	Myconfig 	config.Config
	logContol	control.LogControl
}

//Instance 实例化LogicSvr
func Instance() *LogicSvr {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = new(LogicSvr)
		}
	}
	return instance
}
//LogicInit 初始化
func (logic *LogicSvr) LogicInit() bool {
	if logic.Myconfig.ConfigInit(){
		logic.logContol.Clientkey = logic.Myconfig.Clientkey
		return logic.routerInit()
	}
	return false
}

func (logic *LogicSvr) routerInit() bool{

	router := grouter.New()

	router.GET("/clientlog", logic.logContol.LogNote)

	log4go.Info("http监听开启%s", logic.Myconfig.Sport)
	log4go.Info("当前版本:v1.0.0")

	srv := &http.Server{
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:logic.Myconfig.Sport,
		Handler : router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log4go.Error("http监听失败 %s", err)
		return false
	}
	return true
}