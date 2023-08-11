package api

import (
	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, code int, message string) {
	data := gin.H{
		"code":    code,
		"message": message,
	}
	//global.GL_LOG.Infof("JsonResponse:%s", data)
	c.AbortWithStatusJSON(code, data)
}
