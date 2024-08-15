package renderers

type Log struct {
	*BaseRenderer
}

func NewLog() *Log {
	l := &Log{
		BaseRenderer: NewBaseRenderer(),
	}
	l.Set("type", "log")
	return l
}

func (l *Log) Set(name string, value interface{}) *Log {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	l.AmisSchema[name] = value
	return l
}

/**
 * 是否自动滚动到底部
 */
func (l *Log) AutoScroll(value interface{}) *Log {
	l.Set("autoScroll", value)
	return l
}

/**
 * 外层 CSS 类名
 */
func (l *Log) ClassName(value interface{}) *Log {
	l.Set("className", value)
	return l
}

/**
 * 关闭 ANSI 颜色支持
 */
func (l *Log) DisableColor(value interface{}) *Log {
	l.Set("disableColor", value)
	return l
}

/**
 * 返回内容的字符编码
 */
func (l *Log) Encoding(value interface{}) *Log {
	l.Set("encoding", value)
	return l
}

/**
 * 展示区域高度
 */
func (l *Log) Height(value interface{}) *Log {
	l.Set("height", value)
	return l
}

/**
 * 最大显示行数
 */
func (l *Log) MaxLength(value interface{}) *Log {
	l.Set("maxLength", value)
	return l
}

/**
 * 可选日志操作：['stop', 'clear', 'showLineNumber', 'filter']
 */
func (l *Log) Operation(value interface{}) *Log {
	l.Set("operation", value)
	return l
}

/**
 * 加载中的文字
 */
func (l *Log) Placeholder(value interface{}) *Log {
	l.Set("placeholder", value)
	return l
}

/**
 * 设置每行高度，将会开启虚拟渲染
 */
func (l *Log) RowHeight(value interface{}) *Log {
	l.Set("rowHeight", value)
	return l
}

/**
 * 接口
 */
func (l *Log) Source(value interface{}) *Log {
	l.Set("source", value)
	return l
}

/**
 * 指定为 log 渲染器。
 */
func (l *Log) Type(value interface{}) *Log {
	l.Set("type", value)
	return l
}
