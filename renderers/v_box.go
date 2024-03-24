package renderers


/**
 * 垂直布局控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/vbox

 * @author wcz0
 * @version 6.2.2
 */
type VBox struct {
	*BaseRenderer
}

func NewVBox() *VBox {
    a := &VBox{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "vbox")
    return a
}
/**
 * 是否禁用表达式
 */
func (a *VBox) DisabledOn(value interface{}) *VBox {
    a.Set("disabledOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *VBox) Id(value interface{}) *VBox {
    a.Set("id", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *VBox) EditorSetting(value interface{}) *VBox {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *VBox) UseMobileUI(value interface{}) *VBox {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 行集合
 */
func (a *VBox) Rows(value interface{}) *VBox {
    a.Set("rows", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *VBox) ClassName(value interface{}) *VBox {
    a.Set("className", value)
    return a
}

/**
 */
func (a *VBox) Type(value interface{}) *VBox {
    a.Set("type", value)
    return a
}

/**
 * 是否显示
 */
func (a *VBox) Visible(value interface{}) *VBox {
    a.Set("visible", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *VBox) OnEvent(value interface{}) *VBox {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *VBox) StaticClassName(value interface{}) *VBox {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *VBox) StaticLabelClassName(value interface{}) *VBox {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *VBox) StaticInputClassName(value interface{}) *VBox {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *VBox) Disabled(value interface{}) *VBox {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *VBox) Hidden(value interface{}) *VBox {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *VBox) HiddenOn(value interface{}) *VBox {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *VBox) VisibleOn(value interface{}) *VBox {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *VBox) Static(value interface{}) *VBox {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *VBox) StaticOn(value interface{}) *VBox {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *VBox) StaticPlaceholder(value interface{}) *VBox {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *VBox) StaticSchema(value interface{}) *VBox {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *VBox) Style(value interface{}) *VBox {
    a.Set("style", value)
    return a
}
