package renderers


/**
 * 搜索框渲染器

 * @author wcz0
 * @version 6.2.2
 */
type SearchBox struct {
	*BaseRenderer
}

func NewSearchBox() *SearchBox {
    a := &SearchBox{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "search-box")
    return a
}


func (a *SearchBox) Set(name string, value interface{}) *SearchBox {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 编辑器配置，运行时可以忽略
 */
func (a *SearchBox) EditorSetting(value interface{}) *SearchBox {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *SearchBox) HiddenOn(value interface{}) *SearchBox {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *SearchBox) StaticPlaceholder(value interface{}) *SearchBox {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否为加强样式
 */
func (a *SearchBox) Enhance(value interface{}) *SearchBox {
    a.Set("enhance", value)
    return a
}

/**
 * 是否处于加载状态
 */
func (a *SearchBox) Loading(value interface{}) *SearchBox {
    a.Set("loading", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *SearchBox) Id(value interface{}) *SearchBox {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *SearchBox) StaticClassName(value interface{}) *SearchBox {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否为 Mini 样式。
 */
func (a *SearchBox) Mini(value interface{}) *SearchBox {
    a.Set("mini", value)
    return a
}

/**
 * 外层 css 类名
 */
func (a *SearchBox) ClassName(value interface{}) *SearchBox {
    a.Set("className", value)
    return a
}

/**
 * 关键字名字。
 */
func (a *SearchBox) Name(value interface{}) *SearchBox {
    a.Set("name", value)
    return a
}

/**
 * 占位符
 */
func (a *SearchBox) Placeholder(value interface{}) *SearchBox {
    a.Set("placeholder", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *SearchBox) OnEvent(value interface{}) *SearchBox {
    a.Set("onEvent", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *SearchBox) UseMobileUI(value interface{}) *SearchBox {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *SearchBox) Testid(value interface{}) *SearchBox {
    a.Set("testid", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *SearchBox) Hidden(value interface{}) *SearchBox {
    a.Set("hidden", value)
    return a
}

/**
 */
func (a *SearchBox) StaticSchema(value interface{}) *SearchBox {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *SearchBox) Style(value interface{}) *SearchBox {
    a.Set("style", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *SearchBox) StaticInputClassName(value interface{}) *SearchBox {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *SearchBox) TestIdBuilder(value interface{}) *SearchBox {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否立马搜索。
 */
func (a *SearchBox) SearchImediately(value interface{}) *SearchBox {
    a.Set("searchImediately", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *SearchBox) DisabledOn(value interface{}) *SearchBox {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *SearchBox) Static(value interface{}) *SearchBox {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *SearchBox) StaticLabelClassName(value interface{}) *SearchBox {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否可清除
 */
func (a *SearchBox) Clearable(value interface{}) *SearchBox {
    a.Set("clearable", value)
    return a
}

/**
 * 是否开启清空内容后立即重新搜索
 */
func (a *SearchBox) ClearAndSubmit(value interface{}) *SearchBox {
    a.Set("clearAndSubmit", value)
    return a
}

/**
 * 是否显示
 */
func (a *SearchBox) Visible(value interface{}) *SearchBox {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *SearchBox) StaticOn(value interface{}) *SearchBox {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否禁用
 */
func (a *SearchBox) Disabled(value interface{}) *SearchBox {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *SearchBox) VisibleOn(value interface{}) *SearchBox {
    a.Set("visibleOn", value)
    return a
}

/**
 * 指定为搜索框。文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/search-box
 */
func (a *SearchBox) Type(value interface{}) *SearchBox {
    a.Set("type", value)
    return a
}
