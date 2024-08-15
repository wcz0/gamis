package renderers


/**
 * Alert 提示渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/alert

 * @author wcz0
 * @version 6.2.2
 */
type Alert struct {
	*BaseRenderer
}

func NewAlert() *Alert {
    a := &Alert{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *Alert) Set(name string, value interface{}) *Alert {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "alert")
    return a
}

/**
 * 内容区域
 */
func (a *Alert) Body(value interface{}) *Alert {
    a.Set("body", value)
    return a
}

/**
 * 左侧图标
 */
func (a *Alert) Icon(value interface{}) *Alert {
    a.Set("icon", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Alert) Hiddenon(value interface{}) *Alert {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Alert) Onevent(value interface{}) *Alert {
    a.Set("onEvent", value)
    return a
}

/**
 * 指定为提示框类型
 */
func (a *Alert) Type(value interface{}) *Alert {
    a.Set("type", value)
    return a
}

/**
 * 提示框标题
 */
func (a *Alert) Title(value interface{}) *Alert {
    a.Set("title", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Alert) Staticplaceholder(value interface{}) *Alert {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 关闭按钮CSS类名
 */
func (a *Alert) Closebuttonclassname(value interface{}) *Alert {
    a.Set("closeButtonClassName", value)
    return a
}

/**
 * 图标CSS类名
 */
func (a *Alert) Iconclassname(value interface{}) *Alert {
    a.Set("iconClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Alert) Disabled(value interface{}) *Alert {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Alert) Staticclassname(value interface{}) *Alert {
    a.Set("staticClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Alert) Editorsetting(value interface{}) *Alert {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Alert) Usemobileui(value interface{}) *Alert {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Alert) Classname(value interface{}) *Alert {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Alert) Visible(value interface{}) *Alert {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *Alert) Testidbuilder(value interface{}) *Alert {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 提示类型
 * 可选值: info | warning | success | danger
 */
func (a *Alert) Level(value interface{}) *Alert {
    a.Set("level", value)
    return a
}

/**
 * 操作区域
 */
func (a *Alert) Actions(value interface{}) *Alert {
    a.Set("actions", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Alert) Id(value interface{}) *Alert {
    a.Set("id", value)
    return a
}

/**
 * 组件样式
 */
func (a *Alert) Style(value interface{}) *Alert {
    a.Set("style", value)
    return a
}

/**
 * 是否显示关闭按钮
 */
func (a *Alert) Showclosebutton(value interface{}) *Alert {
    a.Set("showCloseButton", value)
    return a
}

/**
 * 是否显示ICON
 */
func (a *Alert) Showicon(value interface{}) *Alert {
    a.Set("showIcon", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Alert) Static(value interface{}) *Alert {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Alert) Staticlabelclassname(value interface{}) *Alert {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Alert) Disabledon(value interface{}) *Alert {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Alert) Visibleon(value interface{}) *Alert {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Alert) Staticon(value interface{}) *Alert {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Alert) Staticinputclassname(value interface{}) *Alert {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Alert) Staticschema(value interface{}) *Alert {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Alert) Hidden(value interface{}) *Alert {
    a.Set("hidden", value)
    return a
}

/**
 */
func (a *Alert) Testid(value interface{}) *Alert {
    a.Set("testid", value)
    return a
}
