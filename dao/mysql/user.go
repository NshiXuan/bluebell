package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"web_app/models"
)

const secret = "codersx"

// CheckUserExist 检查用户名是否存在
func CheckUserExist(username string) (err error) {
	var count int64
	md := models.User{}
	err = db.Model(&md).Where("username=?", username).Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 插入一条用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)

	// 执行sql
	return db.Create(&user).Error
}

// Login 用户登录
func Login(user *models.User) (err error) {
	// 用户登录的密码
	oPassword := user.Password
	err = db.Select("user_id", "username", "password").Find(user).Error
	if err != nil {
		// 查询数据库失败
		return err
	}
	if user.Username == "" {
		return ErrorUserNotExist
	}

	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

// GetUserById 根据id获取用户信息
func GetUserById(id int64) (models.User, error) {
	var user models.User
	err := db.Select("user_id", "username").First(&user, "user_id = ?", id).Error
	if err != nil {
		err = ErrorInvalidID
	}
	return user, err
}

// 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	// 把加密的加盐字符串写入
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
