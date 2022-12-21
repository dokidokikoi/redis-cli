package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"redis-cli/internal/define"

	"github.com/go-redis/redis/v9"
)

func GetConnection(identity string) (*define.Connection, error) {
	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	data, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	for i, _ := range conf.Connections {
		if conf.Connections[i].Identity == identity {
			return conf.Connections[i], nil
		}
	}
	return nil, errors.New("连接数据不存在")
}

func GetRedisClient(identity string, db int) (*redis.Client, error) {
	conn, err := GetConnection(identity)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conn.Addr, conn.Port),
		Username: conn.Username,
		Password: conn.Password,
		DB:       db,
	})

	return rdb, nil
}
