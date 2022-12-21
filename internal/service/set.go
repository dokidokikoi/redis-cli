package service

import (
	"context"
	"redis-cli/internal/define"
	"redis-cli/internal/helper"

	"github.com/go-redis/redis/v9"
)

func GetSetData(rdb *redis.Client, key string) (interface{}, error) {
	sets, _, err := rdb.SScan(context.Background(), key, 0, "", define.MaxSetLen).Result()
	if err != nil {
		return nil, err
	}
	data := make([]*define.KeyValue, 0, len(sets))
	for i := 0; i < len(sets); i++ {
		data = append(data, &define.KeyValue{
			Value: sets[i],
		})
	}

	return data, nil
}

func SetValueDelete(req *define.SetValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return err
	}
	err = rdb.SRem(context.Background(), req.Key, req.Value).Err()
	return err
}

func SetValueCreate(req *define.SetValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return err
	}
	err = rdb.SAdd(context.Background(), req.Key, req.Value).Err()
	return err
}
