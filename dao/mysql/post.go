package mysql

import (
	"fmt"
	"strings"
	"web_app/models"
)

func CreatePost(p *models.Post) (err error) {
	err = db.Create(&p).Error
	return
}

func GetPostDetail(postID int64) (models.Post, error) {
	// 1.定义模型
	var detail models.Post
	err := db.First(&detail, "post_id = ?", postID).Error
	if err != nil {
		err = ErrorInvalidID
	}
	return detail, err
}

func GetPostList(pageNo, pageSize int) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, 10)
	//Select("post_id", "title", "content", "author_id", "community_id", "create_time")
	db.Order("create_time desc").Limit(pageSize).Offset(pageNo).Find(&posts)
	return
}

// 根据id列表获取post数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	postList = make([]*models.Post, 0, 10)
	// 返回的数据按照给定的id顺序返回 需要使用FIND_IN_SET
	err = db.Raw("select post_id,title,content,author_id,community_id,create_time"+
		" from posts"+
		" WHERE post_id IN ?"+
		" order by FIND_IN_SET(post_id,?)", ids, strings.Join(ids, ",")).Scan(&postList).Error
	fmt.Println("postList", postList)

	return
}
