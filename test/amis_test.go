package test

import (
	"testing"

	"github.com/wcz0/gamis"
)

// MockAdminLang 模拟 tools.AdminLang 函数
func MockAdminLang(ctx interface{}, key string) string {
	switch key {
	case "Login":
		return "登录"
	default:
		return key
	}
}

func TestVanillaActionCreation(t *testing.T) {
	// 模拟 ctx
	ctx := map[string]interface{}{}

	// 创建 VanillaAction 组件
	action := gamis.VanillaAction().
		ActionType("submit").
		Label(MockAdminLang(ctx, "Login")).
		Level("primary").
		ClassName("w-full")

	// 验证组件是否创建成功
	if action == nil {
		t.Fatal("VanillaAction 创建失败")
	}

	// 获取 JSON 输出
	json, err := action.ToJson()
	if err != nil {
		t.Fatalf("转换为 JSON 失败: %v", err)
	}

	// 验证 JSON 不为空
	if json == "" {
		t.Fatal("生成的 JSON 为空")
	}

	t.Logf("生成的 JSON: %s", json)

	// 验证 JSON 包含预期的字段
	expectedFields := []string{
		`"actionType":"submit"`,
		`"label":"登录"`,
		`"level":"primary"`,
		`"className":"w-full"`,
	}

	for _, field := range expectedFields {
		if !contains(json, field) {
			t.Errorf("JSON 中缺少预期字段: %s", field)
		}
	}
}

// contains 检查字符串是否包含子字符串
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[len(s)-len(substr):] == substr ||
		len(s) > len(substr) && findSubstring(s, substr)
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
