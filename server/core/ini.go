package core

import (
	"bistoury-sync/server/global"
	"gopkg.in/ini.v1"
	"os"
)

func ConfigLoad() {
	conf, err := ini.Load("app.ini")
	if err != nil {
		global.GL_LOG.Info("Failed to read the configuration file app.ini!")
		os.Exit(2)
	} else {
		global.GL_LOG.Info("The configuration file was read successfully. Procedure")
	}
	global.GL_SERVER_HOST = conf.Section("server").Key("host").String()
	global.GL_SERVER_PORT = conf.Section("server").Key("port").String()
	global.GL_MYSQL_HOST = conf.Section("mysql").Key("host").String()
	global.GL_MYSQL_DB = conf.Section("mysql").Key("db").String()
	global.GL_MYSQL_USER = conf.Section("mysql").Key("user").String()
	global.GL_MYSQL_PASS = conf.Section("mysql").Key("password").String()
}
