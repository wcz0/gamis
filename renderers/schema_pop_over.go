package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type SchemaPopOver struct {
	*BaseRenderer
}

func NewSchemaPopOver() *SchemaPopOver {
    a := &SchemaPopOver{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *SchemaPopOver) Set(name string, value interface{}) *SchemaPopOver {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    return a
}

/**
 * 配置当前行是否启动，要用表达式
 */
func (a *SchemaPopOver) Popoverenableon(value interface{}) *SchemaPopOver {
    a.Set("popOverEnableOn", value)
    return a
}

/**
 * 是弹窗形式的时候有用。
 * 可选值: sm | md | lg | xl
 */
func (a *SchemaPopOver) Size(value interface{}) *SchemaPopOver {
    a.Set("size", value)
    return a
}

/**
 * 弹出位置
 * 可选值: center | left-top | left-top-left-top | left-top-left-center | left-top-left-bottom | left-top-center-top | left-top-center-center | left-top-center-bottom | left-top-right-top | left-top-right-center | left-top-right-bottom | right-top | right-top-left-top | right-top-left-center | right-top-left-bottom | right-top-center-top | right-top-center-center | right-top-center-bottom | right-top-right-top | right-top-right-center | right-top-right-bottom | left-bottom | left-bottom-left-top | left-bottom-left-center | left-bottom-left-bottom | left-bottom-center-top | left-bottom-center-center | left-bottom-center-bottom | left-bottom-right-top | left-bottom-right-center | left-bottom-right-bottom | right-bottom | right-bottom-left-top | right-bottom-left-center | right-bottom-left-bottom | right-bottom-center-top | right-bottom-center-center | right-bottom-center-bottom | right-bottom-right-top | right-bottom-right-center | right-bottom-right-bottom | fixed-center | fixed-left-top | fixed-right-top | fixed-left-bottom | fixed-right-bottom
 */
func (a *SchemaPopOver) Position(value interface{}) *SchemaPopOver {
    a.Set("position", value)
    return a
}

/**
 */
func (a *SchemaPopOver) Body(value interface{}) *SchemaPopOver {
    a.Set("body", value)
    return a
}

/**
 * 是否显示查看更多的 icon，通常是放大图标。
 */
func (a *SchemaPopOver) Showicon(value interface{}) *SchemaPopOver {
    a.Set("showIcon", value)
    return a
}

/**
 * 偏移量
 */
func (a *SchemaPopOver) Offset(value interface{}) *SchemaPopOver {
    a.Set("offset", value)
    return a
}

/**
 * 标题
 */
func (a *SchemaPopOver) Title(value interface{}) *SchemaPopOver {
    a.Set("title", value)
    return a
}

/**
 * 类名
 */
func (a *SchemaPopOver) Classname(value interface{}) *SchemaPopOver {
    a.Set("className", value)
    return a
}

/**
 * 弹框外层类名
 */
func (a *SchemaPopOver) Popoverclassname(value interface{}) *SchemaPopOver {
    a.Set("popOverClassName", value)
    return a
}

/**
 * 弹出模式
 * 可选值: dialog | drawer | popOver
 */
func (a *SchemaPopOver) Mode(value interface{}) *SchemaPopOver {
    a.Set("mode", value)
    return a
}

/**
 * 触发条件，默认是 click
 * 可选值: click | hover
 */
func (a *SchemaPopOver) Trigger(value interface{}) *SchemaPopOver {
    a.Set("trigger", value)
    return a
}
