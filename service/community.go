package service

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() ([]models.Community, error) {
	// 查询数据库 查找所有的community并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (models.CommunityDetail, error) {
	return mysql.GetCommunityDetail(id)
}
