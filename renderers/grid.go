package renderers


/**
 * Grid 格子布局渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/grid

 * @author wcz0
 * @version 6.2.2
 */
type Grid struct {
	*BaseRenderer
}

func NewGrid() *Grid {
    a := &Grid{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *Grid) Set(name string, value interface{}) *Grid {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "grid")
    return a
}

/**
 * 指定为 Grid 格子布局渲染器。
 */
func (a *Grid) Type(value interface{}) *Grid {
    a.Set("type", value)
    return a
}

/**
 * 列集合
 */
func (a *Grid) Columns(value interface{}) *Grid {
    a.Set("columns", value)
    return a
}

/**
 * 垂直对齐方式
 * 可选值: top | middle | bottom | between
 */
func (a *Grid) Valign(value interface{}) *Grid {
    a.Set("valign", value)
    return a
}

/**
 */
func (a *Grid) Staticschema(value interface{}) *Grid {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Grid) Editorsetting(value interface{}) *Grid {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Grid) Usemobileui(value interface{}) *Grid {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 水平对齐方式
 * 可选值: left | right | between | center
 */
func (a *Grid) Align(value interface{}) *Grid {
    a.Set("align", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Grid) Hiddenon(value interface{}) *Grid {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Grid) Static(value interface{}) *Grid {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Grid) Staticlabelclassname(value interface{}) *Grid {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 水平间距
 * 可选值: xs | sm | base | none | md | lg
 */
func (a *Grid) Gap(value interface{}) *Grid {
    a.Set("gap", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Grid) Disabled(value interface{}) *Grid {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Grid) Onevent(value interface{}) *Grid {
    a.Set("onEvent", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Grid) Classname(value interface{}) *Grid {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Grid) Visible(value interface{}) *Grid {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Grid) Visibleon(value interface{}) *Grid {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Grid) Staticplaceholder(value interface{}) *Grid {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *Grid) Testid(value interface{}) *Grid {
    a.Set("testid", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Grid) Disabledon(value interface{}) *Grid {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *Grid) Testidbuilder(value interface{}) *Grid {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 组件样式
 */
func (a *Grid) Style(value interface{}) *Grid {
    a.Set("style", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Grid) Id(value interface{}) *Grid {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Grid) Staticon(value interface{}) *Grid {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Grid) Staticclassname(value interface{}) *Grid {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Grid) Staticinputclassname(value interface{}) *Grid {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Grid) Hidden(value interface{}) *Grid {
    a.Set("hidden", value)
    return a
}
