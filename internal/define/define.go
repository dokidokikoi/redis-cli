package define

import "time"

var ConfigName = "redis-client.conf"

const (
	MaxHashLen = 100
	MaxListLen = 100
	MaxSetLen  = 100
	MaxZSetLen = 100
)

type Connection struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Connections []*Connection `json:"connections"`
}

type DBItem struct {
	Key    string `json:"key"` // db0 db1 ...
	Number int    `json:"number"`
}

type KeyListRequest struct {
	ConnIdentity string `json:"conn_identity"`
	Db           int    `json:"db"`
	Keyword      string `json:"keyword"`
}

type KeyValueRequest struct {
	ConnIdentity string `json:"conn_identity"`
	Db           int    `json:"db"`
	Key          string `json:"key"`
}

type KeyValueReply struct {
	Value interface{}   `json:"value"`
	TTL   time.Duration `json:"ttl"`
	Type  string        `json:"type"`
}

type KeyValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type CreateKeyValueRequest struct {
	KeyValueRequest
	Type string `json:"type"`
}

type UpdateKeyValueRequest struct {
	CreateKeyValueRequest
	TTL   time.Duration `json:"ttl"`
	Value interface{}   `json:"value"`
}

type ListValueRequest struct {
	KeyValueRequest
	Value string `json:"value"`
}

type HashFieldDeleteRequest struct {
	KeyValueRequest
	Fields []string `json:"field"`
	Value  string   `json:"value"`
}

type HashAddOrUpdateFieldRequest struct {
	KeyValueRequest
	Field string `json:"field"`
	Value string `json:"value"`
}

type SetValueRequest struct {
	KeyValueRequest
	Value string `json:"value"`
}

type ZSetValueRequest struct {
	KeyValueRequest
	Score  float64     `json:"score"`
	Member interface{} `json:"member"`
}
