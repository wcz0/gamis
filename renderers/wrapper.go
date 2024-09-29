package renderers


/**
 * Wrapper 容器渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/wrapper

 * @author wcz0
 * @version 6.2.2
 */
type Wrapper struct {
	*BaseRenderer
}

func NewWrapper() *Wrapper {
    a := &Wrapper{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "wrapper")
    return a
}


func (a *Wrapper) Set(name string, value interface{}) *Wrapper {
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
func (a *Wrapper) StaticClassName(value interface{}) *Wrapper {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Wrapper) StaticInputClassName(value interface{}) *Wrapper {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 内容
 */
func (a *Wrapper) Body(value interface{}) *Wrapper {
    a.Set("body", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Wrapper) Static(value interface{}) *Wrapper {
    a.Set("static", value)
    return a
}

/**
 * 可选值: xs | sm | md | lg | none
 */
func (a *Wrapper) Size(value interface{}) *Wrapper {
    a.Set("size", value)
    return a
}

/**
 * 是否显示
 */
func (a *Wrapper) Visible(value interface{}) *Wrapper {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Wrapper) StaticLabelClassName(value interface{}) *Wrapper {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Wrapper) Testid(value interface{}) *Wrapper {
    a.Set("testid", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Wrapper) HiddenOn(value interface{}) *Wrapper {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Wrapper) Id(value interface{}) *Wrapper {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Wrapper) OnEvent(value interface{}) *Wrapper {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Wrapper) StaticPlaceholder(value interface{}) *Wrapper {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Wrapper) VisibleOn(value interface{}) *Wrapper {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *Wrapper) StaticSchema(value interface{}) *Wrapper {
    a.Set("staticSchema", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Wrapper) ClassName(value interface{}) *Wrapper {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Wrapper) Hidden(value interface{}) *Wrapper {
    a.Set("hidden", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Wrapper) EditorSetting(value interface{}) *Wrapper {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Wrapper) UseMobileUI(value interface{}) *Wrapper {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Wrapper) DisabledOn(value interface{}) *Wrapper {
    a.Set("disabledOn", value)
    return a
}

/**
 * 指定为 container 类型
 */
func (a *Wrapper) Type(value interface{}) *Wrapper {
    a.Set("type", value)
    return a
}

/**
 */
func (a *Wrapper) Wrap(value interface{}) *Wrapper {
    a.Set("wrap", value)
    return a
}

/**
 */
func (a *Wrapper) TestIdBuilder(value interface{}) *Wrapper {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Wrapper) Disabled(value interface{}) *Wrapper {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Wrapper) StaticOn(value interface{}) *Wrapper {
    a.Set("staticOn", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Wrapper) Style(value interface{}) *Wrapper {
    a.Set("style", value)
    return a
}
