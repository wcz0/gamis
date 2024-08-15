package renderers


/**
 * Tag

 * @author wcz0
 * @version 6.2.2
 */
type Tag struct {
	*BaseRenderer
}

func NewTag() *Tag {
    a := &Tag{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "tag")
    return a
}


func (a *Tag) Set(name string, value interface{}) *Tag {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否选中
 */
func (a *Tag) Checked(value interface{}) *Tag {
    a.Set("checked", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Tag) Static(value interface{}) *Tag {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Tag) Usemobileui(value interface{}) *Tag {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Tag) Staticinputclassname(value interface{}) *Tag {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Tag) Testid(value interface{}) *Tag {
    a.Set("testid", value)
    return a
}

/**
 * 标签颜色
 */
func (a *Tag) Color(value interface{}) *Tag {
    a.Set("color", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Tag) Disabled(value interface{}) *Tag {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Tag) Hidden(value interface{}) *Tag {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Tag) Staticplaceholder(value interface{}) *Tag {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Tag) Staticclassname(value interface{}) *Tag {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Tag) Staticschema(value interface{}) *Tag {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Tag) Testidbuilder(value interface{}) *Tag {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 标签文本内容
 */
func (a *Tag) Label(value interface{}) *Tag {
    a.Set("label", value)
    return a
}

/**
 * 是否展示关闭按钮
 */
func (a *Tag) Closable(value interface{}) *Tag {
    a.Set("closable", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Tag) Id(value interface{}) *Tag {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Tag) Onevent(value interface{}) *Tag {
    a.Set("onEvent", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Tag) Style(value interface{}) *Tag {
    a.Set("style", value)
    return a
}

/**
 * 关闭图标
 */
func (a *Tag) Closeicon(value interface{}) *Tag {
    a.Set("closeIcon", value)
    return a
}

/**
 * 是否是可选的标签
 */
func (a *Tag) Checkable(value interface{}) *Tag {
    a.Set("checkable", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Tag) Staticon(value interface{}) *Tag {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Tag) Visible(value interface{}) *Tag {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Tag) Staticlabelclassname(value interface{}) *Tag {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * status模式时候设置的前置图标
 */
func (a *Tag) Icon(value interface{}) *Tag {
    a.Set("icon", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Tag) Hiddenon(value interface{}) *Tag {
    a.Set("hiddenOn", value)
    return a
}

/**
 * normal: 面性标签，对应color的背景色 rounded: 线性标签， 对应color的边框 status: 带图标的标签， 图标可以自定义
 * 可选值: normal | rounded | status
 */
func (a *Tag) Displaymode(value interface{}) *Tag {
    a.Set("displayMode", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Tag) Editorsetting(value interface{}) *Tag {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Tag) Type(value interface{}) *Tag {
    a.Set("type", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Tag) Visibleon(value interface{}) *Tag {
    a.Set("visibleOn", value)
    return a
}

/**
 * 类名
 */
func (a *Tag) Classname(value interface{}) *Tag {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Tag) Disabledon(value interface{}) *Tag {
    a.Set("disabledOn", value)
    return a
}
