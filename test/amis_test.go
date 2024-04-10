package test

import (
	"encoding/json"
	"testing"

	"github.com/wcz0/gamis"
)

func TestAmis(t *testing.T) {
	str, err := gamis.Page().Title("这是标题2").Body([]interface{}{
		gamis.Button().ClassName("test"),
	}).ToJson()
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(gamis.Page().Title("这是标题2").Body([]interface{}{
		gamis.Button().ClassName("test"),
	}))
	str2 := string(bytes)
	if str == "" || str2 == "" {
		t.Error("error")
	}
	t.Log(str)
	t.Log(str2)
}
