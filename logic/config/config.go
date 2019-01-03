package config

import (
	"github.com/yhhaiua/clientlog/logic/model"
	"github.com/yhhaiua/engine/gjson"
	"github.com/yhhaiua/engine/log"
	"io/ioutil"
)
var gLog = log.GetLogger()

type Config struct {
	Sport string //http端口
	Clientkey string
	sqlconfig model.SqlConfig
}

func (config *Config) ConfigInit() bool {
	path := "./config/config.json"
	key := "clientlog"


	configdata, err := ioutil.ReadFile(path)
	if err != nil {
		gLog.Error("Failed to open config file '%s': %s\n", path, err)
		return false
	}

	jsondata, err := gjson.NewJSONByte(configdata)
	if err != nil {
		gLog.Error("Failed to NewJsonByte config file '%s': %s\n", path, err)
		return false
	}
	keydata := gjson.NewGet(jsondata, key)

	if !keydata.IsValid() {
		gLog.Error("Failed1 to config file '%s'", path)
		return false
	}

	data := gjson.NewGetindex(keydata, 0)

	if !data.IsValid(){
		gLog.Error("Failed2 to config file '%s'", path)
		return false
	}

	config.Sport = data.Getstring("port")
	config.Clientkey = data.Getstring("clientkey")

	mysqldata := gjson.NewGet(data, "mysql")
	if !mysqldata.IsValid() {
		gLog.Error("Failed to mysql config file '%s'", path)
		return false
	}
	config.sqlconfig.Shost = mysqldata.Getstring("host")
	config.sqlconfig.Sdbname = mysqldata.Getstring("dbname")
	config.sqlconfig.Suser = mysqldata.Getstring("user")
	config.sqlconfig.Spassword = mysqldata.Getstring("password")
	config.sqlconfig.Maxopen = mysqldata.Getint("open")
	config.sqlconfig.Maxidle = mysqldata.Getint("idle")

	err = config.sqlconfig.InitDB()
	if(err != nil){
		gLog.Error("Failed to mysql InitDB file '%s',err:%s", path,err)
		return false
	}
	return true
}

