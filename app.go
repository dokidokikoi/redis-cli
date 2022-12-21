package main

import (
	"context"
	"fmt"

	"redis-cli/internal/define"
	"redis-cli/internal/service"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// 连接列表
func (a *App) ConnectionList(name string) H {
	conn, err := service.ConnectionList()
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}

	fmt.Println(conn)
	return M{
		"code": 200,
		"data": conn,
	}
}

func (a *App) ConnectionCreate(connection *define.Connection) H {
	err := service.ConnectionCreate(connection)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"msg":  "新建成功",
	}
}

func (a *App) ConnectionEdit(connection *define.Connection) H {
	err := service.ConnectionEdit(connection)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"msg":  "修改成功",
	}
}

func (a *App) ConnectionDelete(identity string) H {
	err := service.ConnectionDelete(identity)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"msg":  "删除成功",
	}
}

func (a *App) KeyList(req *define.KeyListRequest) H {
	if req.ConnIdentity == "" {
		return M{
			"code": -1,
			"msg":  "连接的唯一标识不能为空",
		}
	}
	data, err := service.KeyList(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"data": data,
	}
}

func (a *App) DBList(identity string) H {
	dbs, err := service.DBList(identity)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"data": dbs,
	}
}

func (a *App) GetKeyValue(req *define.KeyValueRequest) H {
	val, err := service.GetKeyValue(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"data": val,
	}
}

func (a *App) DeleteKeyValue(req *define.KeyValueRequest) H {
	if req.Key == "" || req.ConnIdentity == "" {
		return M{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.DeleteKeyValue(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"msg":  "删除成功",
	}
}

func (a *App) CreateKeyValue(req *define.CreateKeyValueRequest) H {
	if req.ConnIdentity == "" || req.Key == "" || req.Type == "" {
		return M{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.CreateKeyValue(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"msg":  "新建成功",
	}
}

func (a *App) UpdateKeyValue(req *define.UpdateKeyValueRequest) H {
	if req.ConnIdentity == "" || req.Key == "" || req.Type == "" || req.Value == "" {
		return M{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}

	err := service.UpdateKeyValue(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"msg":  "新建成功",
	}
}

func (a *App) ListValueDelete(req *define.ListValueRequest) H {
	if req.ConnIdentity == "" || req.Key == "" {
		return M{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.ListValueDelete(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"msg":  "删除成功",
	}
}

func (a *App) ListValueCreate(req *define.ListValueRequest) H {
	if req.ConnIdentity == "" || req.Key == "" || req.Value == "" {
		return M{
			"code": -1,
			"msg":  "必填参不能为空",
		}
	}
	err := service.ListValueCreate(req)
	if err != nil {
		return M{
			"code": -1,
			"msg":  fmt.Sprintf("ERROR: %s", err.Error()),
		}
	}
	return M{
		"code": 200,
		"msg":  "新增成功",
	}
}
