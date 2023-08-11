package core

import (
	"bistoury-sync/server/api"
	"bistoury-sync/server/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ServerRun() {
	router := gin.Default()
	v := router.Group("/api")
	{
		v.POST("/app", api.AppCreate)
		v.DELETE("/app", api.AppDelete)
	}
	err := router.Run(fmt.Sprintf("%s:%s", global.GL_SERVER_HOST, global.GL_SERVER_PORT))
	if err != nil {
		global.GL_LOG.Error(err)
	}

}
