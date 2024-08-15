package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Spinner struct {
	*BaseRenderer
}

func NewSpinner() *Spinner {
    a := &Spinner{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *Spinner) Set(name string, value interface{}) *Spinner {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "spinner")
    return a
}

/**
 * 自定义icon
 */
func (a *Spinner) Icon(value interface{}) *Spinner {
    a.Set("icon", value)
    return a
}

/**
 */
func (a *Spinner) Loadingconfig(value interface{}) *Spinner {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Spinner) Visibleon(value interface{}) *Spinner {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Spinner) Staticon(value interface{}) *Spinner {
    a.Set("staticOn", value)
    return a
}

/**
 * 控制Spinner显示与隐藏
 */
func (a *Spinner) Show(value interface{}) *Spinner {
    a.Set("show", value)
    return a
}

/**
 * spin图标位置包裹元素的自定义class
 */
func (a *Spinner) Spinnerclassname(value interface{}) *Spinner {
    a.Set("spinnerClassName", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Spinner) Onevent(value interface{}) *Spinner {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Spinner) Static(value interface{}) *Spinner {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Spinner) Staticinputclassname(value interface{}) *Spinner {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Spinner) Testidbuilder(value interface{}) *Spinner {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * spinner文案位置
 * 可选值: top | right | bottom | left
 */
func (a *Spinner) Tipplacement(value interface{}) *Spinner {
    a.Set("tipPlacement", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Spinner) Disabled(value interface{}) *Spinner {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Spinner) Hidden(value interface{}) *Spinner {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Spinner) Visible(value interface{}) *Spinner {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Spinner) Id(value interface{}) *Spinner {
    a.Set("id", value)
    return a
}

/**
 * 是否显示遮罩层
 */
func (a *Spinner) Overlay(value interface{}) *Spinner {
    a.Set("overlay", value)
    return a
}

/**
 * 组件样式
 */
func (a *Spinner) Style(value interface{}) *Spinner {
    a.Set("style", value)
    return a
}

/**
 * 作为容器使用时内容
 */
func (a *Spinner) Body(value interface{}) *Spinner {
    a.Set("body", value)
    return a
}

/**
 * spinner文案
 */
func (a *Spinner) Tip(value interface{}) *Spinner {
    a.Set("tip", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Spinner) Disabledon(value interface{}) *Spinner {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Spinner) Staticlabelclassname(value interface{}) *Spinner {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件类型
 */
func (a *Spinner) Type(value interface{}) *Spinner {
    a.Set("type", value)
    return a
}

/**
 */
func (a *Spinner) Testid(value interface{}) *Spinner {
    a.Set("testid", value)
    return a
}

/**
 * 作为容器使用时最外层元素的class
 */
func (a *Spinner) Spinnerwrapclassname(value interface{}) *Spinner {
    a.Set("spinnerWrapClassName", value)
    return a
}

/**
 * 自定义spinner的class
 */
func (a *Spinner) Classname(value interface{}) *Spinner {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Spinner) Hiddenon(value interface{}) *Spinner {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Spinner) Editorsetting(value interface{}) *Spinner {
    a.Set("editorSetting", value)
    return a
}

/**
 * spinner Icon 大小
 * 可选值: sm | lg | 
 */
func (a *Spinner) Size(value interface{}) *Spinner {
    a.Set("size", value)
    return a
}

/**
 */
func (a *Spinner) Mode(value interface{}) *Spinner {
    a.Set("mode", value)
    return a
}

/**
 * 延迟显示
 */
func (a *Spinner) Delay(value interface{}) *Spinner {
    a.Set("delay", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Spinner) Staticplaceholder(value interface{}) *Spinner {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Spinner) Staticclassname(value interface{}) *Spinner {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Spinner) Staticschema(value interface{}) *Spinner {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Spinner) Usemobileui(value interface{}) *Spinner {
    a.Set("useMobileUI", value)
    return a
}
