package service

import (
	"context"
	"errors"
	"redis-cli/internal/define"
	"redis-cli/internal/helper"
	"time"

	"github.com/go-redis/redis/v9"
)

func KeyList(req *define.KeyListRequest) ([]string, error) {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return nil, err
	}

	res, _, err := rdb.Scan(context.Background(), 0, "*"+req.Keyword+"*", 100).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetKeyValue(req *define.KeyValueRequest) (*define.KeyValueReply, error) {
	if req.Key == "" || req.ConnIdentity == "" {
		return nil, errors.New("传入参数错误")
	}

	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return nil, err
	}

	reply := new(define.KeyValueReply)

	_type, err := rdb.Type(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	reply.Type = _type

	switch _type {
	case "string":
		v, err := rdb.Get(context.Background(), req.Key).Result()
		if err != nil {
			return nil, err
		}
		reply.Value = v
	case "hash":
		data, err := GetHashData(rdb, req.Key)
		if err != nil {
			return nil, err
		}
		reply.Value = data
	case "list":
		data, err := GetListData(rdb, req.Key)
		if err != nil {
			return nil, err
		}
		reply.Value = data
	case "set":
		data, err := GetSetData(rdb, req.Key)
		if err != nil {
			return nil, err
		}
		reply.Value = data
	case "zset":
		data, err := GetZsetData(rdb, req.Key)
		if err != nil {
			return nil, err
		}
		reply.Value = data
	}

	ttl, err := rdb.TTL(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	reply.TTL = time.Duration(ttl.Seconds())

	return reply, nil
}

func DeleteKeyValue(req *define.KeyValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return err
	}

	_, err = rdb.Del(context.Background(), req.Key).Result()
	return err
}

func CreateKeyValue(req *define.CreateKeyValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return nil
	}

	switch req.Type {
	case "string":
		err = rdb.Set(context.Background(), req.Key, "New Key", -1).Err()
	case "hash":
		err = rdb.HSet(context.Background(), req.Key, map[string]string{"key": "value"}).Err()
	case "list":
		err = rdb.RPush(context.Background(), req.Key, "value").Err()
	case "set":
		err = rdb.SAdd(context.Background(), req.Key, "value").Err()
	case "zset":
		err = rdb.ZAdd(context.Background(), req.Key, redis.Z{
			Score:  0,
			Member: "value",
		}).Err()
	}

	return err
}

func UpdateKeyValue(req *define.UpdateKeyValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	if err != nil {
		return nil
	}

	switch req.Type {
	case "string":
		err = rdb.Set(context.Background(), req.Key, req.Value, req.TTL*time.Second).Err()
	case "hash":
		err = rdb.HSet(context.Background(), req.Key, req.Value, req.TTL*time.Second).Err()
	case "list":
		err = rdb.RPush(context.Background(), req.Key, req.Value).Err()
		rdb.Expire(context.Background(), req.Key, req.TTL*time.Second)
	case "set":
		err = rdb.SAdd(context.Background(), req.Key, req.Value).Err()
		rdb.Expire(context.Background(), req.Key, req.TTL*time.Second)
	case "zset":
		members, ok := req.Value.([]redis.Z)
		if !ok {
			return errors.New("参数错误")
		}
		err = rdb.ZAdd(context.Background(), req.Key, members...).Err()
		rdb.Expire(context.Background(), req.Key, req.TTL*time.Second)
	}

	return err
}
