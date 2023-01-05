package mysql

import (
	"testing"
	"time"
	"web_app/config"
	"web_app/models"
)

func init() {
	dbCfg := config.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "admin123",
		Port:         3006,
		DbName:       "bluebell",
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	Init(&dbCfg).Error()
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
		Status:      0,
		Title:       "Test",
		Content:     "just a test",
		CreateTime:  time.Time{},
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err: %v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
