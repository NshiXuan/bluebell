package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/models"
	"web_app/service"
)

//type VoteDate struct {
//	// UserID 从请求中获取当前用户的ID
//	PostID    int64 `json:"post_id,string"`   // 帖子ID
//	Direction int   `json:"direction,string"` // 赞成票(1) 反对票(-1)
//}

// PostVoteController 投票
// @Tags 投票
// @Summary 投票
// @param postID query string require "贴子ID"
// @param direction query string require "1赞成 0取消 -1反对"
// @Success 200 {string} json{"code","message"}
// @Router /api/v1/posts2 [post]
func PostVoteController(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamVoteDate)

	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			RespError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		RespErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	// 2.业务逻辑
	// 获取当前用户ID
	userID, err := GetCurrentUser(c)
	if err != nil {
		RespError(c, CodeNeedLogin)
		return
	}

	if err := service.VoteForPost(userID, p); err != nil {
		zap.L().Error("service.VoteForPost() failed", zap.Error(err))
		RespError(c, CodeServerBusy)
		return
	}

	// 3.返回响应
	RespSuccess(c, nil)
}
