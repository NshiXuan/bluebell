package service

import (
	"strconv"
	"web_app/dao/redis"
	"web_app/models"
)

// 推荐阅读
// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm

// 本项目使用简化版的投票分数
// 投一票就加432分 86400/200 -> 200张赞成票可以给你的贴子需续一天 -> 《redis实战》

/* 投票的几种情况
derection=1时，有两种情况：
	1. 之前没有投过票，现在投赞成票		--> 更新分数和投票记录
	2. 之前头反对票，现在改投赞成票		--> 更新分数和投票记录
derection=0时，有两种情况：
	1. 之前投过赞成票，现在取消投票		--> 更新分数和投票记录
	2，之前投过反对票，现在取消投票		--> 更新分数和投票记录
derection=-1时，有两种情况：
	1. 之前没有投过票，现在投反对票		--> 更新分数和投票记录
	2. 之前投赞成票，现在投反对票		--> 更新分数和投票记录

投票的限制：
每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不许投票
	1. 到期之后将redis中保存的赞成票及反对票数存储到mysql中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

// 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteDate) error {
	//zap.L().Debug(
	//	"VoteForPost",
	//	zap.Int64("userID", userID),
	//	zap.String("postID", p.PostID),
	//	zap.Int8("direction", p.Direction),
	//)
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
