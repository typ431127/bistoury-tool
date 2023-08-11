package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GL_LOG         *zap.SugaredLogger
	GL_DB          *gorm.DB
	GL_SERVER_HOST string
	GL_SERVER_PORT string
	GL_MYSQL_HOST  string
	GL_MYSQL_DB    string
	GL_MYSQL_USER  string
	GL_MYSQL_PASS  string
)
