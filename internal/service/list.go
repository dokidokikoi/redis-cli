package service

import (
	"context"
	"errors"
	"redis-cli/internal/define"
	"redis-cli/internal/helper"
	"time"

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

func ListUpdate(rdb *redis.Client, req *define.UpdateKeyValueRequest) error {
	values, ok := req.Value.([]interface{})
	if !ok {
		return errors.New("参数错误")
	}
	var vals []interface{}
	for _, v := range values {
		v, ok := v.(map[string]interface{})
		if !ok {
			return errors.New("参数错误")
		}
		vals = append(vals, v["value"])
	}
	rdb.Del(context.Background(), req.Key)
	err := rdb.RPush(context.Background(), req.Key, vals...).Err()
	if err != nil {
		return err
	}
	if req.TTL > 0 {
		rdb.Expire(context.Background(), req.Key, req.TTL*time.Second)
	}
	return nil
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
