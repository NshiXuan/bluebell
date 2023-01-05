package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const CtxtUserIDKey = "userId"

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser 获取当前登录用户id
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxtUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageInfo(c *gin.Context) (int, int) {
	pageNoStr := c.Query("pageNo")
	PageSizeStr := c.Query("pageSize")

	var (
		pageNo   int
		pageSize int
		err      error
	)

	pageNo, err = strconv.Atoi(pageNoStr)
	if err != nil {
		pageNo = 0
	}
	pageSize, err = strconv.Atoi(PageSizeStr)
	if err != nil {
		pageSize = 10
	}

	return pageNo, pageSize
}
