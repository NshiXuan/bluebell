package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"web_app/models"
)

func GetCommunityList() (communityList []models.Community, err error) {
	var communities []models.Community
	err = db.Select("community_id", "community_name").Find(&communities).Error
	if err != nil {
		zap.L().Warn("find community failed")
		err = nil
	}
	fmt.Println(communities)
	return communities, err
}

func GetCommunityDetail(id int64) (models.CommunityDetail, error) {
	var detail models.CommunityDetail
	err := db.Table(models.CommunityTableName()).First(&detail, "community_id = ?", id).Error
	if err != nil {
		err = ErrorInvalidID
	}
	fmt.Println(detail)
	return detail, err
}
