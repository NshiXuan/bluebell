package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/service"
)

/* 社区相关的 */

// CommunityHandler
// @Tags 社区
// @Summary 查询所有的社区
// @Success 200 {string} json{"code","message"}
// @Router /api/v1/community [get]
func CommunityHandler(c *gin.Context) {
	data, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList() failed", zap.Error(err))
		RespError(c, CodeServerBusy)
		return
	}
	RespSuccess(c, data)
	return
}

// CommunityDetailHandler
// @Tags 社区
// @Summary 根据id查询社区详情
// @param id query string require "社区ID"
// @Success 200 {string} json{"code","message"}
// @Router /community/:id [get]
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id")
	// 转成10机制 64为的整数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		RespError(c, CodeInvalidParam)
		return
	}

	data, err := service.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("service.CommunityDetailHandler() failed", zap.Error(err))
		RespError(c, CodeServerBusy)
		return
	}
	RespSuccess(c, data)
}
