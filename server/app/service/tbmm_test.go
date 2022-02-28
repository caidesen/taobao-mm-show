package service

import (
	"encoding/json"
	"testing"
)

func TestGetOneMMPicInCycle(t *testing.T) {
	data := GetOneMMPicInCycle()

	if len(data.Url) == 0 {
		t.Error("没有url")
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	t.Log("res:", string(marshal))
}
func TestGetTbMMPicList(t *testing.T) {
	list := GetTbMMPicList(10)
	for _, data := range list {
		marshal, err := json.Marshal(data)
		if err != nil {
			return
		}
		t.Log("res:", string(marshal))
	}
}
