package service

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		// 数据库查询出错
		return err
	}

	// 2.生成UID
	userID := snowflake.GenID()

	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 密码加密

	// 3.保存进数据库
	return mysql.InsertUser(user)

	// redis.xxx ...
}

func Login(p *models.ParamLogin) (user models.User, err error) {
	// 1.获取实例
	user.Username = p.Username
	user.Password = p.Password

	// 传递的是指针，就能拿到user.UserID
	if err = mysql.Login(&user); err != nil {
		return
	}

	// 生成JWT
	user.Token, err = jwt.GenToken(user.UserID, user.Username)
	return
}
