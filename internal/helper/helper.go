package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"redis-cli/internal/define"

	"github.com/go-redis/redis/v9"
)

func GetConfigPath() string {
	current, _ := user.Current()
	return current.HomeDir + string(os.PathSeparator) + ".redis-cli/conf" + string(os.PathSeparator)
}

func GetConfig() (*define.Config, error) {
	conf := new(define.Config)
	data, err := os.ReadFile(GetConfigPath() + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func SaveConfig(conn *define.Connection) error {
	conf := new(define.Config)
	data, err := os.ReadFile(GetConfigPath() + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(GetConfigPath(), 0777)
	} else if err != nil {
		return err
	} else {
		json.Unmarshal(data, conf)
	}
	conf.Connections = append(conf.Connections, conn)
	data, _ = json.Marshal(conf)
	os.WriteFile(GetConfigPath()+define.ConfigName, data, 0666)
	return nil
}

func GetConnection(identity string) (*define.Connection, error) {
	conf := new(define.Config)
	data, err := os.ReadFile(GetConfigPath() + define.ConfigName)
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
