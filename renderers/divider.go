package renderers


/**
 * Divider 分割线渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/divider

 * @author wcz0
 * @version 6.2.2
 */
type Divider struct {
	*BaseRenderer
}

func NewDivider() *Divider {
    a := &Divider{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "divider")
    return a
}


func (a *Divider) Set(name string, value interface{}) *Divider {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 静态展示表单项类名
 */
func (a *Divider) Staticclassname(value interface{}) *Divider {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Divider) Staticinputclassname(value interface{}) *Divider {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Divider) Disabled(value interface{}) *Divider {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Divider) Static(value interface{}) *Divider {
    a.Set("static", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Divider) Editorsetting(value interface{}) *Divider {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Divider) Testid(value interface{}) *Divider {
    a.Set("testid", value)
    return a
}

/**
 */
func (a *Divider) Titleclassname(value interface{}) *Divider {
    a.Set("titleClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Divider) Hidden(value interface{}) *Divider {
    a.Set("hidden", value)
    return a
}

/**
 */
func (a *Divider) Staticschema(value interface{}) *Divider {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Divider) Rotate(value interface{}) *Divider {
    a.Set("rotate", value)
    return a
}

/**
 * 可选值: left | center | right
 */
func (a *Divider) Titleposition(value interface{}) *Divider {
    a.Set("titlePosition", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Divider) Onevent(value interface{}) *Divider {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Divider) Staticon(value interface{}) *Divider {
    a.Set("staticOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *Divider) Style(value interface{}) *Divider {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Divider) Title(value interface{}) *Divider {
    a.Set("title", value)
    return a
}

/**
 * 可选值: horizontal | vertical
 */
func (a *Divider) Direction(value interface{}) *Divider {
    a.Set("direction", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Divider) Visibleon(value interface{}) *Divider {
    a.Set("visibleOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Divider) Usemobileui(value interface{}) *Divider {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Divider) Classname(value interface{}) *Divider {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Divider) Hiddenon(value interface{}) *Divider {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Divider) Visible(value interface{}) *Divider {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Divider) Id(value interface{}) *Divider {
    a.Set("id", value)
    return a
}

/**
 */
func (a *Divider) Type(value interface{}) *Divider {
    a.Set("type", value)
    return a
}

/**
 */
func (a *Divider) Color(value interface{}) *Divider {
    a.Set("color", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Divider) Disabledon(value interface{}) *Divider {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Divider) Staticplaceholder(value interface{}) *Divider {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Divider) Staticlabelclassname(value interface{}) *Divider {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Divider) Testidbuilder(value interface{}) *Divider {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 可选值: dashed | solid
 */
func (a *Divider) Linestyle(value interface{}) *Divider {
    a.Set("lineStyle", value)
    return a
}
