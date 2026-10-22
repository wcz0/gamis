package gamis

import "github.com/wcz0/gamis/renderers"

// Root 保留仅为兼容旧代码；它与 Page 等价。
//
// Deprecated: 使用 Page() 代替。
func Root() *renderers.Root {
	return renderers.NewRoot()
}
