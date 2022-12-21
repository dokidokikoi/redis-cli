package service

import (
	"context"
	"redis-cli/internal/define"
	"redis-cli/internal/helper"

	"github.com/go-redis/redis/v9"
)

func GetZsetData(rdb *redis.Client, key string) (interface{}, error) {
	data, err := rdb.ZRevRangeWithScores(context.Background(), key, 0, define.MaxZSetLen-1).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ZSetValueDelete(req *define.ZSetValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return err
	}
	err = rdb.ZRem(context.Background(), req.Key, req.Member).Err()
	return err
}

func ZSetValueCreate(req *define.ZSetValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return err
	}
	err = rdb.ZAdd(context.Background(), req.Key, redis.Z{
		Score:  req.Score,
		Member: req.Member,
	}).Err()
	return err
}
