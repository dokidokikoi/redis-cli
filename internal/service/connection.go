package service

import (
	"encoding/json"
	"errors"
	"os"
	"redis-cli/internal/define"

	uuid "github.com/satori/go.uuid"
)

// 连接列表
func ConnectionList() ([]*define.Connection, error) {
	nowPath, _ := os.Getwd()
	data, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return nil, err
	}
	conf := new(define.Config)
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
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
	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	data, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(nowPath, 0666)
	} else if err != nil {
		return err
	} else {
		json.Unmarshal(data, conf)
	}
	conf.Connections = append(conf.Connections, conn)
	data, _ = json.Marshal(conf)
	os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
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
	nowPath, _ := os.Getwd()
	data, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
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
	os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
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
