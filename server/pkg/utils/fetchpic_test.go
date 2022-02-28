package utils

import (
	"encoding/json"
	"testing"
)

func TestFindOneByUomg(t *testing.T) {
	data, err := FindOneByUomg()
	if err != nil {
		t.Error(err)
	}
	if len(data.Url) == 0 {
		t.Error("没有url")
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	t.Log("xiaoAPI:", string(marshal))

}
func TestFindOneByXiaoApi(t *testing.T) {
	data, err := FindOneByXiaoApi()
	if err != nil {
		t.Error(err)
	}
	if len(data.Url) == 0 {
		t.Error("没有url")
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	t.Log("xiaoAPI:", string(marshal))
}
