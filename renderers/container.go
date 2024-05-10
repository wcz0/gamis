package renderers


/**
 * Container 容器渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/container

 * @author wcz0
 * @version 6.2.2
 */
type Container struct {
	*BaseRenderer
}

func NewContainer() *Container {
    a := &Container{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "container")
    return a
}

/**
 * 事件动作配置
 */
func (a *Container) OnEvent(value interface{}) *Container {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Container) StaticOn(value interface{}) *Container {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Container) StaticSchema(value interface{}) *Container {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Container) TestIdBuilder(value interface{}) *Container {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * body 类名
 */
func (a *Container) BodyClassName(value interface{}) *Container {
    a.Set("bodyClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Container) HiddenOn(value interface{}) *Container {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Container) VisibleOn(value interface{}) *Container {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Container) StaticInputClassName(value interface{}) *Container {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Container) UseMobileUI(value interface{}) *Container {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Container) DisabledOn(value interface{}) *Container {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Container) Id(value interface{}) *Container {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Container) StaticClassName(value interface{}) *Container {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Container) StaticPlaceholder(value interface{}) *Container {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 使用的标签
 */
func (a *Container) WrapperComponent(value interface{}) *Container {
    a.Set("wrapperComponent", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Container) ClassName(value interface{}) *Container {
    a.Set("className", value)
    return a
}

/**
 * 是否需要对body加一层div包裹，默认为 true
 */
func (a *Container) WrapperBody(value interface{}) *Container {
    a.Set("wrapperBody", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Container) Static(value interface{}) *Container {
    a.Set("static", value)
    return a
}

/**
 * 是否开启容器拖拽
 */
func (a *Container) Draggable(value interface{}) *Container {
    a.Set("draggable", value)
    return a
}

/**
 * 是否开启容器拖拽配置
 */
func (a *Container) DraggableConfig(value interface{}) *Container {
    a.Set("draggableConfig", value)
    return a
}

/**
 */
func (a *Container) Testid(value interface{}) *Container {
    a.Set("testid", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Container) Disabled(value interface{}) *Container {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Container) StaticLabelClassName(value interface{}) *Container {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Container) Style(value interface{}) *Container {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Container) EditorSetting(value interface{}) *Container {
    a.Set("editorSetting", value)
    return a
}

/**
 * 内容
 */
func (a *Container) Body(value interface{}) *Container {
    a.Set("body", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Container) Hidden(value interface{}) *Container {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Container) Visible(value interface{}) *Container {
    a.Set("visible", value)
    return a
}

/**
 * 指定为 container 类型
 */
func (a *Container) Type(value interface{}) *Container {
    a.Set("type", value)
    return a
}
