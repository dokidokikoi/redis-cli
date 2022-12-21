package service

import "testing"

func TestDBList(t *testing.T) {
	data, err := DBList("cf610dd9-c6e8-4305-94c1-cefc28079428")
	if err != nil {
		t.Errorf("err: %s", err)
	}
	t.Logf("data: %#v", data)
}
