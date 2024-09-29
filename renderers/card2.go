package renderers


/**
 * Card2 新卡片渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/card2

 * @author wcz0
 * @version 6.2.2
 */
type Card2 struct {
	*BaseRenderer
}

func NewCard2() *Card2 {
    a := &Card2{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "card2")
    return a
}


func (a *Card2) Set(name string, value interface{}) *Card2 {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否隐藏
 */
func (a *Card2) Hidden(value interface{}) *Card2 {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Card2) Id(value interface{}) *Card2 {
    a.Set("id", value)
    return a
}

/**
 */
func (a *Card2) TestIdBuilder(value interface{}) *Card2 {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Card2) HiddenOn(value interface{}) *Card2 {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Card2) Style(value interface{}) *Card2 {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Card2) UseMobileUI(value interface{}) *Card2 {
    a.Set("useMobileUI", value)
    return a
}

/**
 * body 类名
 */
func (a *Card2) BodyClassName(value interface{}) *Card2 {
    a.Set("bodyClassName", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Card2) Static(value interface{}) *Card2 {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Card2) StaticOn(value interface{}) *Card2 {
    a.Set("staticOn", value)
    return a
}

/**
 * 渲染标签
 */
func (a *Card2) WrapperComponent(value interface{}) *Card2 {
    a.Set("wrapperComponent", value)
    return a
}

/**
 * 指定为 card2 类型
 */
func (a *Card2) Type(value interface{}) *Card2 {
    a.Set("type", value)
    return a
}

/**
 * 内容
 */
func (a *Card2) Body(value interface{}) *Card2 {
    a.Set("body", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Card2) DisabledOn(value interface{}) *Card2 {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Card2) Visible(value interface{}) *Card2 {
    a.Set("visible", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Card2) OnEvent(value interface{}) *Card2 {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Card2) StaticLabelClassName(value interface{}) *Card2 {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Card2) Testid(value interface{}) *Card2 {
    a.Set("testid", value)
    return a
}

/**
 * 隐藏选框
 */
func (a *Card2) HideCheckToggler(value interface{}) *Card2 {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Card2) Disabled(value interface{}) *Card2 {
    a.Set("disabled", value)
    return a
}

/**
 */
func (a *Card2) StaticSchema(value interface{}) *Card2 {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Card2) EditorSetting(value interface{}) *Card2 {
    a.Set("editorSetting", value)
    return a
}

/**
 * 不配置href且cards容器下生效，点击整个卡片触发选中
 */
func (a *Card2) CheckOnItemClick(value interface{}) *Card2 {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Card2) ClassName(value interface{}) *Card2 {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Card2) VisibleOn(value interface{}) *Card2 {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Card2) StaticPlaceholder(value interface{}) *Card2 {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Card2) StaticClassName(value interface{}) *Card2 {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Card2) StaticInputClassName(value interface{}) *Card2 {
    a.Set("staticInputClassName", value)
    return a
}
