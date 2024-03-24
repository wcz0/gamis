package test

import (
	"testing"

	"github.com/wcz0/gamis"
)

func TestAmis(t *testing.T) {
	str := gamis.Page().Title("这是一个标题").ToJson()
	if str == "" {
		t.Error("error")
	}
	t.Log(str)
}

