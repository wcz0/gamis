package renderers


/**
 * Password

 * @author wcz0
 * @version 6.2.2
 */
type Password struct {
	*BaseRenderer
}

func NewPassword() *Password {
    a := &Password{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "password")
    return a
}


func (a *Password) Set(name string, value interface{}) *Password {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Password) Id(value interface{}) *Password {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Password) Staticon(value interface{}) *Password {
    a.Set("staticOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Password) Usemobileui(value interface{}) *Password {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 打码模式的文本
 */
func (a *Password) Mosaictext(value interface{}) *Password {
    a.Set("mosaicText", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Password) Classname(value interface{}) *Password {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Password) Visible(value interface{}) *Password {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Password) Staticclassname(value interface{}) *Password {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Password) Style(value interface{}) *Password {
    a.Set("style", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Password) Disabled(value interface{}) *Password {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Password) Hidden(value interface{}) *Password {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Password) Hiddenon(value interface{}) *Password {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Password) Visibleon(value interface{}) *Password {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Password) Onevent(value interface{}) *Password {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Password) Staticplaceholder(value interface{}) *Password {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Password) Editorsetting(value interface{}) *Password {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Password) Testidbuilder(value interface{}) *Password {
    a.Set("testIdBuilder", value)
    return a
}

/**
 */
func (a *Password) Testid(value interface{}) *Password {
    a.Set("testid", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Password) Staticlabelclassname(value interface{}) *Password {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Password) Staticinputclassname(value interface{}) *Password {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Password) Staticschema(value interface{}) *Password {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Password) Type(value interface{}) *Password {
    a.Set("type", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Password) Disabledon(value interface{}) *Password {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Password) Static(value interface{}) *Password {
    a.Set("static", value)
    return a
}
