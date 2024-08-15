package renderers


/**
 * SwitchContainer 状态容器渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/state-container

 * @author wcz0
 * @version 6.2.2
 */
type SwitchContainer struct {
	*BaseRenderer
}

func NewSwitchContainer() *SwitchContainer {
    a := &SwitchContainer{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "switch-container")
    return a
}


func (a *SwitchContainer) Set(name string, value interface{}) *SwitchContainer {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否显示
 */
func (a *SwitchContainer) Visible(value interface{}) *SwitchContainer {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *SwitchContainer) Visibleon(value interface{}) *SwitchContainer {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *SwitchContainer) Staticinputclassname(value interface{}) *SwitchContainer {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *SwitchContainer) Staticschema(value interface{}) *SwitchContainer {
    a.Set("staticSchema", value)
    return a
}

/**
 * 自定义样式
 */
func (a *SwitchContainer) Style(value interface{}) *SwitchContainer {
    a.Set("style", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *SwitchContainer) Classname(value interface{}) *SwitchContainer {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *SwitchContainer) Disabled(value interface{}) *SwitchContainer {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *SwitchContainer) Hidden(value interface{}) *SwitchContainer {
    a.Set("hidden", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *SwitchContainer) Usemobileui(value interface{}) *SwitchContainer {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 指定为 container 类型
 */
func (a *SwitchContainer) Type(value interface{}) *SwitchContainer {
    a.Set("type", value)
    return a
}

/**
 * 状态项列表
 */
func (a *SwitchContainer) Items(value interface{}) *SwitchContainer {
    a.Set("items", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *SwitchContainer) Id(value interface{}) *SwitchContainer {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *SwitchContainer) Staticlabelclassname(value interface{}) *SwitchContainer {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *SwitchContainer) Editorsetting(value interface{}) *SwitchContainer {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *SwitchContainer) Testidbuilder(value interface{}) *SwitchContainer {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *SwitchContainer) Disabledon(value interface{}) *SwitchContainer {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *SwitchContainer) Staticon(value interface{}) *SwitchContainer {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *SwitchContainer) Staticplaceholder(value interface{}) *SwitchContainer {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *SwitchContainer) Onevent(value interface{}) *SwitchContainer {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *SwitchContainer) Static(value interface{}) *SwitchContainer {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *SwitchContainer) Staticclassname(value interface{}) *SwitchContainer {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *SwitchContainer) Testid(value interface{}) *SwitchContainer {
    a.Set("testid", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *SwitchContainer) Hiddenon(value interface{}) *SwitchContainer {
    a.Set("hiddenOn", value)
    return a
}
