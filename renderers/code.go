package renderers

type Code struct {
	*BaseRenderer
}

func NewCode() *Code {
	c := &Code{
		BaseRenderer: NewBaseRenderer(),
	}
	c.Set("type", "code")
	return c
}

func (c *Code) Set(name string, value interface{}) *Code {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	c.AmisSchema[name] = value
	return c
}

/**
 * 外层 CSS 类名
 */
func (c *Code) ClassName(value interface{}) *Code {
	c.Set("className", value)
	return c
}

/**
 * 主题，还有 'vs-dark'
 */
func (c *Code) EditorTheme(value interface{}) *Code {
	c.Set("editorTheme", value)
	return c
}

/**
 * 所使用的高亮语言，默认是 plaintext
 */
func (c *Code) Language(value interface{}) *Code {
	c.Set("language", value)
	return c
}

/**
 * 在其他组件中，时，用作变量映射
 */
func (c *Code) Name(value interface{}) *Code {
	c.Set("name", value)
	return c
}

/**
 * 默认 tab 大小
 */
func (c *Code) TabSize(value interface{}) *Code {
	c.Set("tabSize", value)
	return c
}

/**
 * 指定为 code 渲染器。
 */
func (c *Code) Type(value interface{}) *Code {
	c.Set("type", value)
	return c
}

/**
 * 显示的颜色值
 */
func (c *Code) Value(value interface{}) *Code {
	c.Set("value", value)
	return c
}

/**
 * 是否折行
 */
func (c *Code) WordWrap(value bool) *Code {
	c.Set("wordWrap", value)
	return c
}
