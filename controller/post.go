package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/models"
	"web_app/service"
)

// CreatePostHandler
// @Tags 帖子相关接口
// @Summary 创建贴子
// @param authorID query string require "用户ID"
// @param communityID query string require "社区ID"
// @param title query string require "标题"
// @param content query string require "内容"
// @Success 200 {string} json{"code","message"}
// @Router /api/v1/post [post]
func CreatePostHandler(c *gin.Context) {
	// 1.获取参数以校验参数
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("error", err))
		zap.L().Error("create post with invalid param")
		RespError(c, CodeInvalidParam)
		return
	}

	// 从 c 中获取当前请求的用户ID 在token中有保存
	userID, err := GetCurrentUser(c)
	if err != nil {
		RespError(c, CodeNeedLogin)
		return
	}
	fmt.Println(userID)
	p.AuthorID = userID

	//	2.创建帖子
	if err := service.CreatePost(p); err != nil {
		zap.L().Error("service.CreatePost(p) failed", zap.Error(err))
		RespError(c, CodeServerBusy)
		return
	}

	//	3.返回响应
	RespSuccess(c, nil)
}

// PostDetailHandler
// @Tags 帖子相关接口
// @Summary 通过ID获取贴子详情
// @param postID query string require "贴子ID"
// @Success 200 {string} json{"code","message"}
// @Router /api/v1/post/:id [post]
func PostDetailHandler(c *gin.Context) {
	// 1.获取参数postID与校验参数
	idStr := c.Param("id")
	// 转成10机制 64为的整数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		RespError(c, CodeInvalidParam)
		return
	}

	// 2.获取详情
	data, err := service.GetPostDetail(id)
	if err != nil {
		zap.L().Error("service.PostDetailHandler() failed", zap.Error(err))
		RespError(c, CodeServerBusy)
		return
	}

	// 3.返回响应
	RespSuccess(c, data)
}

// GetPostListHandler
// @Tags 帖子相关接口
// @Summary 获取贴子列表
// @param postID query string require "贴子ID"
// @Success 200 {string} json{"code","message"}
// @Router /posts [post]
func GetPostListHandler(c *gin.Context) {
	// 1.获取数据
	pageNo, pageSize := getPageInfo(c)

	data, err := service.GetPostList(pageNo, pageSize)
	if err != nil {
		zap.L().Error("service.GetPostList() failed", zap.Error(err))
		return
	}

	// 2.返回响应
	RespSuccess(c, data)
}

// GetPostListHandler2
// @Summary 根据时间或分数或社区获取贴子列表
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} RespPostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// 1.获取参数 /api/v1/posts?pageNo=1&pageSize=10&o=time
	// 初始结构体指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	err := c.ShouldBindQuery(&p)
	if err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		RespError(c, CodeInvalidParam)
		return
	}

	// 2.业务逻辑
	var data []*models.ApiPostDetail
	if p.CommunityID == 0 {
		// 查所有
		data, err = service.GetPostList2(p)
	} else {
		// 根据社区查询post
		data, err = service.ParamCommunityPostList(p)
	}

	if err != nil {
		zap.L().Error("service.GetPostList2() failed", zap.Error(err))
		return
	}

	// 3.返回响应
	RespSuccess(c, data)
}

//func GetCommunityPostListHandler(c *gin.Context) {
//	// 1.获取参数 /api/v1/posts?pageNo=1&pageSize=10&o=time
//	// 初始结构体指定初始参数
//	p := &models.ParamPostList{
//		Page:  1,
//		Size:  10,
//		Order: models.OrderTime,
//	}
//	err := c.ShouldBindQuery(&p)
//	if err != nil {
//		zap.L().Error("ParamCommunityPostList with invalid params", zap.Error(err))
//		RespError(c, CodeInvalidParam)
//		return
//	}
//
//	// 2.业务逻辑
//	data, err := service.ParamCommunityPostList(p)
//	if err != nil {
//		zap.L().Error("service.ParamCommunityPostList() failed", zap.Error(err))
//		return
//	}
//
//	// 3.返回响应
//	RespSuccess(c, data)
//}
