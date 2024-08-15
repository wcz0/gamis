package renderers


/**
 * AnchorNav 锚点导航渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav

 * @author wcz0
 * @version 6.2.2
 */
type AnchorNav struct {
	*BaseRenderer
}

func NewAnchorNav() *AnchorNav {
    a := &AnchorNav{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *AnchorNav) Set(name string, value interface{}) *AnchorNav {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "anchor-nav")
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *AnchorNav) Staticinputclassname(value interface{}) *AnchorNav {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *AnchorNav) Style(value interface{}) *AnchorNav {
    a.Set("style", value)
    return a
}

/**
 */
func (a *AnchorNav) Testidbuilder(value interface{}) *AnchorNav {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *AnchorNav) Onevent(value interface{}) *AnchorNav {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *AnchorNav) Staticclassname(value interface{}) *AnchorNav {
    a.Set("staticClassName", value)
    return a
}

/**
 * 指定为 AnchorNav 锚点导航渲染器
 */
func (a *AnchorNav) Type(value interface{}) *AnchorNav {
    a.Set("type", value)
    return a
}

/**
 * 样式名
 */
func (a *AnchorNav) Classname(value interface{}) *AnchorNav {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *AnchorNav) Disabled(value interface{}) *AnchorNav {
    a.Set("disabled", value)
    return a
}

/**
 */
func (a *AnchorNav) Staticschema(value interface{}) *AnchorNav {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *AnchorNav) Usemobileui(value interface{}) *AnchorNav {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 可选值: vertical | horizontal
 */
func (a *AnchorNav) Direction(value interface{}) *AnchorNav {
    a.Set("direction", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *AnchorNav) Hidden(value interface{}) *AnchorNav {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *AnchorNav) Staticon(value interface{}) *AnchorNav {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *AnchorNav) Testid(value interface{}) *AnchorNav {
    a.Set("testid", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *AnchorNav) Disabledon(value interface{}) *AnchorNav {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *AnchorNav) Id(value interface{}) *AnchorNav {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *AnchorNav) Static(value interface{}) *AnchorNav {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *AnchorNav) Staticplaceholder(value interface{}) *AnchorNav {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *AnchorNav) Staticlabelclassname(value interface{}) *AnchorNav {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *AnchorNav) Editorsetting(value interface{}) *AnchorNav {
    a.Set("editorSetting", value)
    return a
}

/**
 * 被激活（定位）的楼层
 */
func (a *AnchorNav) Active(value interface{}) *AnchorNav {
    a.Set("active", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *AnchorNav) Visibleon(value interface{}) *AnchorNav {
    a.Set("visibleOn", value)
    return a
}

/**
 * 楼层集合
 */
func (a *AnchorNav) Links(value interface{}) *AnchorNav {
    a.Set("links", value)
    return a
}

/**
 * 导航样式名
 */
func (a *AnchorNav) Linkclassname(value interface{}) *AnchorNav {
    a.Set("linkClassName", value)
    return a
}

/**
 * 楼层样式名
 */
func (a *AnchorNav) Sectionclassname(value interface{}) *AnchorNav {
    a.Set("sectionClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *AnchorNav) Hiddenon(value interface{}) *AnchorNav {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *AnchorNav) Visible(value interface{}) *AnchorNav {
    a.Set("visible", value)
    return a
}
