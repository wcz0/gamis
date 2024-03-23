package gamis

import (
	"testing"
)

func TestAmis(t *testing.T) {
	str := Amis().Page().Title("这是一个标题").ToJson()
	if str == "" {
		t.Error("error")
	}
	t.Log(str)
}

