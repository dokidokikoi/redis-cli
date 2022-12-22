package service

import (
	"context"
	"errors"
	"redis-cli/internal/define"
	"redis-cli/internal/helper"
	"time"

	"github.com/go-redis/redis/v9"
)

func GetZsetData(rdb *redis.Client, key string) (interface{}, error) {
	data, err := rdb.ZRevRangeWithScores(context.Background(), key, 0, define.MaxZSetLen-1).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ZsetUpdate(rdb *redis.Client, req *define.UpdateKeyValueRequest) error {
	values, ok := req.Value.([]interface{})
	if !ok {
		return errors.New("参数错误")
	}
	var vals []redis.Z
	for _, v := range values {
		v, ok := v.(map[string]interface{})
		if !ok {
			return errors.New("参数错误")
		}
		vals = append(vals, redis.Z{
			Member: v["Member"],
			Score:  v["Score"].(float64),
		})
	}
	err := rdb.ZAdd(context.Background(), req.Key, vals...).Err()
	if err != nil {
		return err
	}
	if req.TTL > 0 {
		rdb.Expire(context.Background(), req.Key, req.TTL*time.Second)
	}
	return nil
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
