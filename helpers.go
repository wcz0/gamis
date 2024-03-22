package gamis

import "gamis/renderers"

func Amis(typeStr string) interface{} {
	if typeStr != "" {
		return renderers.NewComponent().SetType(typeStr)
	}
	return renderers.NewAmis()
}