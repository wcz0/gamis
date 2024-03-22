package gamis

import "testing"

func TestAmis(t *testing.T){
	a := Amis("action")
	json := a.ActionType("ajax").ActiveClassName("active").ActiveLevel("info").ToJson()
	if json == nil {
		t.Error("Expected non-nil result, but got nil")
	}
}