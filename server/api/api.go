package api

import (
	"bistoury-sync/server/core/model"
	"bistoury-sync/server/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

func AppCreate(c *gin.Context) {
	var err error
	var body createbody
	if err := c.ShouldBindJSON(&body); err != nil {
		global.GL_LOG.Info(err)
		JsonResponse(c, 200, fmt.Sprintf("请求参数错误:%s", err))
		return
	}
	app := model.BistouryApp{
		Code:       body.Code,
		Name:       body.Name,
		Status:     1,
		Creator:    "admin",
		Group_code: body.GroupCode,
	}
	err = global.GL_DB.Where("code = ?", body.Code).FirstOrCreate(&app).Error
	if err != nil {
		global.GL_LOG.Error(err)
		JsonResponse(c, 500, "数据库错误,联系管理员")
		return

	} else {
		global.GL_LOG.Infof("App:%s created successfully", app.Code)
	}
	uid := strings.Replace(uuid.New().String(), "-", "", -1)
	server := model.BistouryServer{
		Server_id: uid,
		Ip:        body.IP,
		Port:      body.Port,
		Host:      body.Hostname,
		Log_dir:   body.Logdir,
		Room:      "k8s",
		App_code:  body.Code,
	}
	err = global.GL_DB.Where("Host = ?", body.Hostname).FirstOrCreate(&server).Error
	if err != nil {
		global.GL_LOG.Error(err)
		JsonResponse(c, 500, "数据库错误,联系管理员")
		return
	} else {
		global.GL_LOG.Infof("Server:%s created successfully", server.Host)
	}

	userapp := model.BistouryUserApp{
		App_code:  body.Code,
		User_code: body.Code,
	}
	err = global.GL_DB.Where("app_code = ?", body.Code).Where("user_code = ?", body.Code).FirstOrCreate(&userapp).Error
	if err != nil {
		global.GL_LOG.Error(err)
		JsonResponse(c, 500, "数据库错误,联系管理员")
		return
	} else {
		global.GL_LOG.Infof("UserApp-code:%s created successfully", app.Code)
	}
	userapp = model.BistouryUserApp{
		App_code:  body.Code,
		User_code: "admin",
	}
	err = global.GL_DB.Where("app_code = ?", body.Code).Where("user_code = ?", "admin").FirstOrCreate(&userapp).Error
	if err != nil {
		global.GL_LOG.Error(err)
		JsonResponse(c, 500, "数据库错误,联系管理员")
		return
	} else {
		global.GL_LOG.Infof("UserApp-admin:%s created successfully", app.Code)
	}

	JsonResponse(c, 200, fmt.Sprintf("host:%s 注册成功", body.Hostname))
	return

}

func AppDelete(c *gin.Context) {
	var body deletebody
	if err := c.ShouldBindJSON(&body); err != nil {
		global.GL_LOG.Info(err)
		JsonResponse(c, 500, "请求参数错误")
		return
	}
	server := model.BistouryServer{
		Host: body.Hostname,
		Ip:   body.IP,
	}
	result := global.GL_DB.Where("Host = ?", body.Hostname).Delete(&server)
	if result.Error == nil {
		global.GL_LOG.Infof("删除成功 - hostname:%s", body.Hostname)
	} else {
		global.GL_LOG.Infof("删除失败 - hostname:%s", body.Hostname)
	}
	JsonResponse(c, 200, fmt.Sprintf("主机:%s删除成功", body.Hostname))
	return
}
