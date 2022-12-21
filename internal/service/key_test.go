package service

import (
	"context"
	"redis-cli/internal/define"
	"testing"
	"time"

	"github.com/go-redis/redis/v9"
)

func TestKeyList(t *testing.T) {
	data, err := KeyList(&define.KeyListRequest{
		ConnIdentity: "cf610dd9-c6e8-4305-94c1-cefc28079428",
		Db:           2,
	})
	if err != nil {
		t.Errorf("err: %s", err)
	}
	t.Logf("data: %#v", data)
}

func TestRedisCreate(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	err := rdb.SetEx(context.Background(), "test", "test", time.Duration(100)*time.Second).Err()
	t.Logf("err: %s", err)
}

func TestCreateKeyValue(t *testing.T) {
	err := CreateKeyValue(&define.CreateKeyValueRequest{
		KeyValueRequest: define.KeyValueRequest{
			ConnIdentity: "87aa9047-9995-49fa-8df2-1f22016b5aa1",
			Db:           0,
			Key:          "test",
		},
		Type: "string",
	})
	if err != nil {
		t.Errorf("err: %s", err)
	}
}

func TestGetKeyValue(t *testing.T) {
	data, err := GetKeyValue(&define.KeyValueRequest{
		ConnIdentity: "87aa9047-9995-49fa-8df2-1f22016b5aa1",
		Db:           0,
		Key:          "test",
	})
	if err != nil {
		t.Errorf("err: %s", err)
	}
	t.Logf("data: %#v", data)
}

func TestUpdateKeyValue(t *testing.T) {
	err := UpdateKeyValue(&define.UpdateKeyValueRequest{
		CreateKeyValueRequest: define.CreateKeyValueRequest{
			KeyValueRequest: define.KeyValueRequest{
				ConnIdentity: "87aa9047-9995-49fa-8df2-1f22016b5aa1",
				Db:           0,
				Key:          "test",
			},
			Type: "string",
		},
		TTL:   time.Duration(1000),
		Value: "test",
	})
	if err != nil {
		t.Errorf("err: %s", err)
	}
}
