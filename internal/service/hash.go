package service

import (
	"context"
	"redis-cli/internal/define"
	"redis-cli/internal/helper"

	"github.com/go-redis/redis/v9"
)

func GetHashData(rdb *redis.Client, key string) (interface{}, error) {
	keys, _, err := rdb.HScan(context.Background(), key, 0, "", define.MaxHashLen).Result()
	if err != nil {
		return nil, err
	}
	data := make([]*define.KeyValue, 0, len(keys)/2)
	for i := 0; i < len(keys); i += 2 {
		data = append(data, &define.KeyValue{
			Key:   keys[i],
			Value: keys[i+1],
		})
	}

	return data, nil
}

// HashFieldDelete hash 字段删除
func HashFieldDelete(req *define.HashFieldDeleteRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return err
	}
	err = rdb.HDel(context.Background(), req.Key, req.Fields...).Err()
	return err
}

// HashAddOrUpdateField hash 字段新增、修改
func HashAddOrUpdateField(req *define.HashAddOrUpdateFieldRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return err
	}
	err = rdb.HSet(context.Background(), req.Key, map[string]interface{}{req.Field: req.Value}).Err()
	return err
}
