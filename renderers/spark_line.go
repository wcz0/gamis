package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type SparkLine struct {
	*BaseRenderer
}

func NewSparkLine() *SparkLine {
    a := &SparkLine{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "sparkline")
    return a
}


func (a *SparkLine) Set(name string, value interface{}) *SparkLine {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否隐藏表达式
 */
func (a *SparkLine) HiddenOn(value interface{}) *SparkLine {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *SparkLine) Visible(value interface{}) *SparkLine {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *SparkLine) StaticOn(value interface{}) *SparkLine {
    a.Set("staticOn", value)
    return a
}

/**
 * 宽度
 */
func (a *SparkLine) Width(value interface{}) *SparkLine {
    a.Set("width", value)
    return a
}

/**
 * 空数据时显示的内容
 */
func (a *SparkLine) Placeholder(value interface{}) *SparkLine {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否禁用
 */
func (a *SparkLine) Disabled(value interface{}) *SparkLine {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *SparkLine) DisabledOn(value interface{}) *SparkLine {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *SparkLine) Id(value interface{}) *SparkLine {
    a.Set("id", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *SparkLine) UseMobileUI(value interface{}) *SparkLine {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *SparkLine) Type(value interface{}) *SparkLine {
    a.Set("type", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *SparkLine) EditorSetting(value interface{}) *SparkLine {
    a.Set("editorSetting", value)
    return a
}

/**
 * css 类名
 */
func (a *SparkLine) ClassName(value interface{}) *SparkLine {
    a.Set("className", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *SparkLine) OnEvent(value interface{}) *SparkLine {
    a.Set("onEvent", value)
    return a
}

/**
 * 组件样式
 */
func (a *SparkLine) Style(value interface{}) *SparkLine {
    a.Set("style", value)
    return a
}

/**
 * 高度
 */
func (a *SparkLine) Height(value interface{}) *SparkLine {
    a.Set("height", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *SparkLine) StaticInputClassName(value interface{}) *SparkLine {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 点击行为
 */
func (a *SparkLine) ClickAction(value interface{}) *SparkLine {
    a.Set("clickAction", value)
    return a
}

/**
 */
func (a *SparkLine) Value(value interface{}) *SparkLine {
    a.Set("value", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *SparkLine) Static(value interface{}) *SparkLine {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *SparkLine) StaticLabelClassName(value interface{}) *SparkLine {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *SparkLine) StaticSchema(value interface{}) *SparkLine {
    a.Set("staticSchema", value)
    return a
}

/**
 * 关联数据变量。
 */
func (a *SparkLine) Name(value interface{}) *SparkLine {
    a.Set("name", value)
    return a
}

/**
 */
func (a *SparkLine) Testid(value interface{}) *SparkLine {
    a.Set("testid", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *SparkLine) Hidden(value interface{}) *SparkLine {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *SparkLine) VisibleOn(value interface{}) *SparkLine {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *SparkLine) StaticPlaceholder(value interface{}) *SparkLine {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *SparkLine) StaticClassName(value interface{}) *SparkLine {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *SparkLine) TestIdBuilder(value interface{}) *SparkLine {
    a.Set("testIdBuilder", value)
    return a
}
