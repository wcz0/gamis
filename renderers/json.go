package renderers


/**
 * JSON 数据展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/json

 * @author wcz0
 * @version 6.2.2
 */
type Json struct {
	*BaseRenderer
}

func NewJson() *Json {
    a := &Json{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "json")
    return a
}


func (a *Json) Set(name string, value interface{}) *Json {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 组件样式
 */
func (a *Json) Style(value interface{}) *Json {
    a.Set("style", value)
    return a
}

/**
 * 支持从数据链取值
 */
func (a *Json) Source(value interface{}) *Json {
    a.Set("source", value)
    return a
}

/**
 * 图标风格
 * 可选值: square | circle | triangle
 */
func (a *Json) IconStyle(value interface{}) *Json {
    a.Set("iconStyle", value)
    return a
}

/**
 * 是否显示键的引号
 */
func (a *Json) QuotesOnKeys(value interface{}) *Json {
    a.Set("quotesOnKeys", value)
    return a
}

/**
 */
func (a *Json) TestIdBuilder(value interface{}) *Json {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否显示数据类型
 */
func (a *Json) DisplayDataTypes(value interface{}) *Json {
    a.Set("displayDataTypes", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Json) DisabledOn(value interface{}) *Json {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *Json) StaticSchema(value interface{}) *Json {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Json) EditorSetting(value interface{}) *Json {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Json) VisibleOn(value interface{}) *Json {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Json) Id(value interface{}) *Json {
    a.Set("id", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Json) UseMobileUI(value interface{}) *Json {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 要展示的 JSON 数据
 */
func (a *Json) Value(value interface{}) *Json {
    a.Set("value", value)
    return a
}

/**
 * 是否为键排序
 */
func (a *Json) SortKeys(value interface{}) *Json {
    a.Set("sortKeys", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Json) Disabled(value interface{}) *Json {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示
 */
func (a *Json) Visible(value interface{}) *Json {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Json) StaticClassName(value interface{}) *Json {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Json) StaticLabelClassName(value interface{}) *Json {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否可修改
 */
func (a *Json) Mutable(value interface{}) *Json {
    a.Set("mutable", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Json) HiddenOn(value interface{}) *Json {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Json) Static(value interface{}) *Json {
    a.Set("static", value)
    return a
}

/**
 */
func (a *Json) Testid(value interface{}) *Json {
    a.Set("testid", value)
    return a
}

/**
 * 默认展开的级别
 */
func (a *Json) LevelExpand(value interface{}) *Json {
    a.Set("levelExpand", value)
    return a
}

/**
 * 设置字符串的最大展示长度，超出长度阈值的字符串将被截断，点击value可切换字符串展示方式，默认为false
 */
func (a *Json) EllipsisThreshold(value interface{}) *Json {
    a.Set("ellipsisThreshold", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Json) OnEvent(value interface{}) *Json {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Json) StaticInputClassName(value interface{}) *Json {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否可复制
 */
func (a *Json) EnableClipboard(value interface{}) *Json {
    a.Set("enableClipboard", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Json) ClassName(value interface{}) *Json {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Json) Hidden(value interface{}) *Json {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Json) StaticOn(value interface{}) *Json {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Json) StaticPlaceholder(value interface{}) *Json {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 指定为Json展示类型
 * 可选值: json | static-json
 */
func (a *Json) Type(value interface{}) *Json {
    a.Set("type", value)
    return a
}
