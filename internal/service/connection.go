package service

import (
	"encoding/json"
	"errors"
	"os"
	"redis-cli/internal/define"
	"redis-cli/internal/helper"

	uuid "github.com/satori/go.uuid"
)

// 连接列表
func ConnectionList() ([]*define.Connection, error) {
	conf, err := helper.GetConfig()
	if err != nil {
		return nil, err
	}
	if conf == nil {
		return nil, nil
	}

	return conf.Connections, nil
}

func ConnectionCreate(conn *define.Connection) error {
	if conn.Addr == "" {
		return errors.New("连接地址不能为空")
	}
	if conn.Name == "" {
		conn.Name = conn.Addr
	}
	if conn.Port == "" {
		conn.Port = "6379"
	}

	conn.Identity = uuid.NewV4().String()
	return helper.SaveConfig(conn)
}

func ConnectionEdit(conn *define.Connection) error {
	if conn.Identity == "" {
		return errors.New("连接唯一标识不能为空")
	}
	if conn.Addr == "" {
		return errors.New("连接地址不能为空")
	}
	if conn.Name == "" {
		conn.Name = conn.Addr
	}
	if conn.Port == "" {
		conn.Port = "6379"
	}

	conf := new(define.Config)
	data, err := os.ReadFile(helper.GetConfigPath() + define.ConfigName)
	if err != nil {
		return err
	}
	json.Unmarshal(data, conf)
	for i, con := range conf.Connections {
		if con.Identity == conn.Identity {
			conf.Connections[i] = conn
			break
		}
	}
	data, _ = json.Marshal(conf)
	os.WriteFile(helper.GetConfigPath()+define.ConfigName, data, 0666)
	return nil
}

func ConnectionDelete(identity string) error {
	if identity == "" {
		return errors.New("连接唯一标识不能为空")
	}

	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	data, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return err
	}
	json.Unmarshal(data, conf)
	for i, con := range conf.Connections {
		if con.Identity == identity {
			conf.Connections = append(conf.Connections[:i], conf.Connections[i+1:]...)
			break
		}
	}
	data, _ = json.Marshal(conf)
	os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
}
