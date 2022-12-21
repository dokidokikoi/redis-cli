package service

import (
	"context"
	"redis-cli/internal/define"
	"redis-cli/internal/helper"

	"github.com/go-redis/redis/v9"
)

func GetListData(rdb *redis.Client, key string) (interface{}, error) {
	list, err := rdb.LRange(context.Background(), key, 0, define.MaxListLen-1).Result()
	if err != nil {
		return nil, err
	}
	data := make([]*define.KeyValue, 0, len(list))
	for i := 0; i < len(list); i++ {
		data = append(data, &define.KeyValue{
			Value: list[i],
		})
	}

	return data, nil
}

func ListValueDelete(req *define.ListValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return err
	}
	err = rdb.LRem(context.Background(), req.Key, 1, req.Value).Err()
	return err
}

func ListValueCreate(req *define.ListValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return err
	}
	err = rdb.RPush(context.Background(), req.Key, req.Value).Err()
	return err
}
