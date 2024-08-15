package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Step struct {
	*BaseRenderer
}

func NewStep() *Step {
    a := &Step{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *Step) Set(name string, value interface{}) *Step {
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
 * 子标题
 */
func (a *Step) Subtitle(value interface{}) *Step {
    a.Set("subTitle", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Step) Visibleon(value interface{}) *Step {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Step) Staticclassname(value interface{}) *Step {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Step) Staticschema(value interface{}) *Step {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Step) Usemobileui(value interface{}) *Step {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 标题
 */
func (a *Step) Title(value interface{}) *Step {
    a.Set("title", value)
    return a
}

/**
 */
func (a *Step) Testid(value interface{}) *Step {
    a.Set("testid", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Step) Disabledon(value interface{}) *Step {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Step) Id(value interface{}) *Step {
    a.Set("id", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Step) Disabled(value interface{}) *Step {
    a.Set("disabled", value)
    return a
}

/**
 */
func (a *Step) Value(value interface{}) *Step {
    a.Set("value", value)
    return a
}

/**
 * 是否显示
 */
func (a *Step) Visible(value interface{}) *Step {
    a.Set("visible", value)
    return a
}

/**
 * 组件样式
 */
func (a *Step) Style(value interface{}) *Step {
    a.Set("style", value)
    return a
}

/**
 * 图标
 */
func (a *Step) Icon(value interface{}) *Step {
    a.Set("icon", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Step) Staticplaceholder(value interface{}) *Step {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *Step) Testidbuilder(value interface{}) *Step {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Step) Static(value interface{}) *Step {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Step) Staticinputclassname(value interface{}) *Step {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 描述
 */
func (a *Step) Description(value interface{}) *Step {
    a.Set("description", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Step) Classname(value interface{}) *Step {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Step) Hidden(value interface{}) *Step {
    a.Set("hidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Step) Onevent(value interface{}) *Step {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Step) Staticon(value interface{}) *Step {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Step) Editorsetting(value interface{}) *Step {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Step) Hiddenon(value interface{}) *Step {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Step) Staticlabelclassname(value interface{}) *Step {
    a.Set("staticLabelClassName", value)
    return a
}
