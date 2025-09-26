package gamis

import (
	"fmt"
)

// Example 展示如何使用便捷函数创建组件
func Example() {
	// 使用便捷函数创建组件
	page := Page().
		Set("title", "示例页面").
		Set("body", []interface{}{
			Alert().
				Set("level", "info").
				Set("body", "这是一个示例页面"),
			Button().
				Set("label", "点击我").
				Set("actionType", "ajax"),
		})

	// 输出 JSON
	json, err := page.ToJson()
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	
	fmt.Println(json)
}