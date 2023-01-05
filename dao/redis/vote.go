package redis

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"math"
	"time"

	"github.com/go-redis/redis"
)

/* 投票的几种情况 4-34
derection=1时，有两种情况：
	1. 之前没有投过票，现在投赞成票		--> 更新分数和投票记录  差值的绝对值：1  +432
	2. 之前头反对票，现在改投赞成票		--> 更新分数和投票记录  差值的绝对值：2  +432*2
derection=0时，有两种情况：
	1，之前投过反对票，现在取消投票		--> 更新分数和投票记录  差值的绝对值：1	+432
	2. 之前投过赞成票，现在取消投票		--> 更新分数和投票记录  差值的绝对值：1  -432
derection=-1时，有两种情况：
	1. 之前没有投过票，现在投反对票		--> 更新分数和投票记录  差值的绝对值：1  -432
	2. 之前投赞成票，现在投反对票		--> 更新分数和投票记录  差值的绝对值：2  -432*2

投票的限制：
每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不许投票
	1. 到期之后将redis中保存的赞成票及反对票数存储到mysql中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

const (
	oneWeekSeconds = 7 * 24 * 3600
	scorePerVote   = 432 // 每一票的分数
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepested   = errors.New("不允许重复投票")
)

func VoteForPost(userID, postID string, value float64) error {
	// 1.判断投票的限制
	// 取redis贴子发布的时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	fmt.Println("postTime", postTime)
	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		return ErrVoteTimeExpire
	}

	// 2和3需要放到事务中
	// 2.更新分数
	// 先查当前用户给当前贴子的投票记录 value oValue为 0 1 -1 ,ZScore差不多默认为0
	oValue := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	zap.L().Debug("redis.VoteForPost", zap.Any("value", value))
	zap.L().Debug("redis.VoteForPost", zap.Any("oValue", oValue))
	if value == oValue {
		return ErrVoteRepested
	}
	var op float64
	if value > oValue { // 加
		op = 1
	} else { // 减
		op = -1
	}

	diff := math.Abs(oValue - value) // 计算两次投票的差值
	zap.L().Debug("redis.VoteForPost", zap.Any("diff", diff))

	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)

	// 3.记录用户为该贴子投票
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
