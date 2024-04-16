package cache

import (
	"context"
	"social-network/config"
	"social-network/middleware/logger"
	"social-network/models"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb      *redis.Client
	Rctx     context.Context
	RrwMutex sync.RWMutex
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	Rctx = context.Background()

	_, err := Rdb.Ping(Rctx).Result()
	if err != nil {
		logger.Error(map[string]interface{}{"Redis connect error": err.Error()})
	}
}

func DelLike(momentid, userid uint64) (int64, error) {
	userID := strconv.FormatUint(userid, 10)
	momentID := strconv.FormatUint(momentid, 10)
	SRemCmd := Rdb.SRem(Rctx, "likes:"+momentID, userID) // 删除点赞成员
	count, err := SRemCmd.Result()
	if err != nil || count == 0 {
		return -1, err
	}
	ZIncrByCmd := Rdb.ZIncrBy(Rctx, "likes:ranked", float64(0-count), momentID) // 减少排名点赞数
	likes_count, err := ZIncrByCmd.Result()
	if err != nil {
		return -1, err
	}
	return int64(likes_count), nil
}

func AddLike(momentid, userid uint64) (int64, error) {
	userID := strconv.FormatUint(userid, 10)
	momentID := strconv.FormatUint(momentid, 10)
	SAddCmd := Rdb.SAdd(Rctx, "likes:"+momentID, userID)

	count, err := SAddCmd.Result()
	if err != nil || count == 0 {
		return -1, err
	}
	ZIncrByCmd := Rdb.ZIncrBy(Rctx, "likes:ranked", float64(count), momentID)
	likes_count, err := ZIncrByCmd.Result()
	if err != nil {
		return -1, err
	}
	return int64(likes_count), nil
}
func AddLikes(likes []models.Like) error {

	pipeRedis := Rdb.TxPipeline()
	var SAddCmd []*redis.IntCmd
	var ZIncrByCmd []*redis.FloatCmd
	for _, like := range likes {
		userID := strconv.FormatUint(like.UserID, 10)
		momentID := strconv.FormatUint(like.MomentID, 10)
		SAddCmd = append(SAddCmd, pipeRedis.SAdd(Rctx, "likes:"+momentID, userID))
		ZIncrByCmd = append(ZIncrByCmd, pipeRedis.ZIncrBy(Rctx, "likes:ranked", 1, momentID))
	}
	_, err := pipeRedis.Exec(Rctx)
	if err != nil {
		return err
	}
	var likescount_all int64
	for _, cmd := range SAddCmd {
		likes_count, err := cmd.Result()
		if err != nil {
			return err
		}
		likescount_all += likes_count
	}
	if likescount_all != int64(len(likes)) {

		return redis.Nil
	}
	for _, cmd := range ZIncrByCmd {
		if cmd.Err() != nil {
			return cmd.Err()
		}
	}

	return nil

}
func DelMonent(momentid uint64) error {
	momentID := strconv.FormatUint(momentid, 10)
	skey := "likes:" + momentID
	pipeRedis := Rdb.Pipeline()
	DelKeyCmd := pipeRedis.Del(Rctx, skey)
	ZRemCmd := pipeRedis.ZRem(Rctx, "likes:ranked", momentID)
	_, err := pipeRedis.Exec(Rctx)
	if err != nil {
		return err
	}
	if DelKeyCmd.Err() != nil {
		return DelKeyCmd.Err()
	}
	if ZRemCmd.Err() != nil {
		return ZRemCmd.Err()
	}
	return nil
}
