package service

import (
	"context"
	"errors"
	"fmt"
	"redis-cli/internal/define"
	"redis-cli/internal/helper"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v9"
)

func DBList(indentity string) ([]*define.DBItem, error) {
	conn, err := helper.GetConnection(indentity)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conn.Addr, conn.Port),
		Username: conn.Username,
		Password: conn.Password,
	})

	keySpace, err := rdb.Info(context.Background(), "keyspace").Result()
	if err != nil {
		return nil, err
	}

	m := make(map[string]int)
	v := strings.Split(keySpace, "\n")
	for i := 1; i < len(v)-1; i++ {
		database := strings.Split(v[i], ":")
		if len(database) < 2 {
			continue
		}
		vv := strings.Split(database[1], ",")
		if len(vv) < 1 {
			continue
		}
		keyNum := strings.Split(vv[0], "=")
		if len(keyNum) < 2 {
			continue
		}
		num, err := strconv.Atoi(keyNum[1])
		if err != nil {
			continue
		}
		m[database[0]] = num
	}
	databaseRes, err := rdb.ConfigGet(context.Background(), "databases").Result()
	if err != nil {
		return nil, err
	}
	numStr, ok := databaseRes["databases"]
	if !ok {
		return nil, errors.New("连接数据异常")
	}
	dbNum, err := strconv.Atoi(numStr)
	if err != nil {
		return nil, err
	}
	data := make([]*define.DBItem, dbNum)
	for i, _ := range data {
		data[i] = new(define.DBItem)
		data[i].Key = fmt.Sprintf("db%d", i)
		if n, ok := m[data[i].Key]; ok {
			data[i].Number = n
		}
	}

	return data, nil
}
