package core

import "bistoury-sync/server/global"

func CoreInit() {
	global.GL_LOG = Zap()
	ConfigLoad()
	global.GL_DB = Gorm()
	SigInit()
}
