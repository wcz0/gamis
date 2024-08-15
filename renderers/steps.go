package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Steps struct {
	*BaseRenderer
}

func NewSteps() *Steps {
    a := &Steps{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *Steps) Set(name string, value interface{}) *Steps {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "steps")
    return a
}

/**
 * 展示模式
 * 可选值: horizontal | vertical
 */
func (a *Steps) Mode(value interface{}) *Steps {
    a.Set("mode", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Steps) Classname(value interface{}) *Steps {
    a.Set("className", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Steps) Onevent(value interface{}) *Steps {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *Steps) Staticschema(value interface{}) *Steps {
    a.Set("staticSchema", value)
    return a
}

/**
 * 指定当前步骤
 */
func (a *Steps) Value(value interface{}) *Steps {
    a.Set("value", value)
    return a
}

/**
 * 标签放置位置
 * 可选值: horizontal | vertical
 */
func (a *Steps) Labelplacement(value interface{}) *Steps {
    a.Set("labelPlacement", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Steps) Usemobileui(value interface{}) *Steps {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Steps) Testid(value interface{}) *Steps {
    a.Set("testid", value)
    return a
}

/**
 * API 或 数据映射
 */
func (a *Steps) Source(value interface{}) *Steps {
    a.Set("source", value)
    return a
}

/**
 * 变量映射
 */
func (a *Steps) Name(value interface{}) *Steps {
    a.Set("name", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Steps) Static(value interface{}) *Steps {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Steps) Staticon(value interface{}) *Steps {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Steps) Editorsetting(value interface{}) *Steps {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Steps) Status(value interface{}) *Steps {
    a.Set("status", value)
    return a
}

/**
 * 指定为 Steps 步骤条渲染器
 */
func (a *Steps) Type(value interface{}) *Steps {
    a.Set("type", value)
    return a
}

/**
 * 步骤
 */
func (a *Steps) Steps(value interface{}) *Steps {
    a.Set("steps", value)
    return a
}

/**
 * 点状步骤条
 */
func (a *Steps) Progressdot(value interface{}) *Steps {
    a.Set("progressDot", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Steps) Hiddenon(value interface{}) *Steps {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Steps) Id(value interface{}) *Steps {
    a.Set("id", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Steps) Staticplaceholder(value interface{}) *Steps {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Steps) Staticclassname(value interface{}) *Steps {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Steps) Visibleon(value interface{}) *Steps {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Steps) Staticinputclassname(value interface{}) *Steps {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Steps) Staticlabelclassname(value interface{}) *Steps {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Steps) Disabled(value interface{}) *Steps {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Steps) Hidden(value interface{}) *Steps {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Steps) Visible(value interface{}) *Steps {
    a.Set("visible", value)
    return a
}

/**
 * 组件样式
 */
func (a *Steps) Style(value interface{}) *Steps {
    a.Set("style", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Steps) Disabledon(value interface{}) *Steps {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *Steps) Testidbuilder(value interface{}) *Steps {
    a.Set("testIdBuilder", value)
    return a
}
