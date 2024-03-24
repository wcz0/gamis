package renderers


/**
 * Drawer 抽出式弹框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/drawer

 * @author wcz0
 * @version 6.2.2
 */
type Drawer struct {
	*BaseRenderer
}

func NewDrawer() *Drawer {
    a := &Drawer{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "drawer")
    return a
}
/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Drawer) Id(value interface{}) *Drawer {
    a.Set("id", value)
    return a
}

/**
 * 内容区域
 */
func (a *Drawer) Body(value interface{}) *Drawer {
    a.Set("body", value)
    return a
}

/**
 * 配置 头部 容器 className
 */
func (a *Drawer) FooterClassName(value interface{}) *Drawer {
    a.Set("footerClassName", value)
    return a
}

/**
 * 是否支持按 ESC 关闭 Dialog
 */
func (a *Drawer) CloseOnEsc(value interface{}) *Drawer {
    a.Set("closeOnEsc", value)
    return a
}

/**
 * 是否展示关闭按钮 当值为false时，默认开启closeOnOutside
 */
func (a *Drawer) ShowCloseButton(value interface{}) *Drawer {
    a.Set("showCloseButton", value)
    return a
}

/**
 * 是否显示错误信息
 */
func (a *Drawer) ShowErrorMsg(value interface{}) *Drawer {
    a.Set("showErrorMsg", value)
    return a
}

/**
 * 配置 外层 className
 */
func (a *Drawer) ClassName(value interface{}) *Drawer {
    a.Set("className", value)
    return a
}

/**
 * 抽屉的宽度 （当position为left | right时生效）
 */
func (a *Drawer) Width(value interface{}) *Drawer {
    a.Set("width", value)
    return a
}

/**
 * 头部
 */
func (a *Drawer) Header(value interface{}) *Drawer {
    a.Set("header", value)
    return a
}

/**
 * 影响自动生成的按钮，如果自己配置了按钮这个配置无效。
 */
func (a *Drawer) Confirm(value interface{}) *Drawer {
    a.Set("confirm", value)
    return a
}

/**
 * 点击外部的时候是否关闭弹框。
 */
func (a *Drawer) CloseOnOutside(value interface{}) *Drawer {
    a.Set("closeOnOutside", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Drawer) OnEvent(value interface{}) *Drawer {
    a.Set("onEvent", value)
    return a
}

/**
 * 请通过配置 title 设置标题
 */
func (a *Drawer) Title(value interface{}) *Drawer {
    a.Set("title", value)
    return a
}

/**
 * 从什么位置弹出
 * 可选值: left | right | top | bottom
 */
func (a *Drawer) Position(value interface{}) *Drawer {
    a.Set("position", value)
    return a
}

/**
 */
func (a *Drawer) Testid(value interface{}) *Drawer {
    a.Set("testid", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Drawer) HiddenOn(value interface{}) *Drawer {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 配置 头部 容器 className
 */
func (a *Drawer) HeaderClassName(value interface{}) *Drawer {
    a.Set("headerClassName", value)
    return a
}

/**
 * 是否显示蒙层
 */
func (a *Drawer) Overlay(value interface{}) *Drawer {
    a.Set("overlay", value)
    return a
}

/**
 * 配置 Body 容器 className
 */
func (a *Drawer) BodyClassName(value interface{}) *Drawer {
    a.Set("bodyClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Drawer) VisibleOn(value interface{}) *Drawer {
    a.Set("visibleOn", value)
    return a
}

/**
 * 默认不用填写，自动会创建确认和取消按钮。
 */
func (a *Drawer) Actions(value interface{}) *Drawer {
    a.Set("actions", value)
    return a
}

/**
 */
func (a *Drawer) Name(value interface{}) *Drawer {
    a.Set("name", value)
    return a
}

/**
 * 是否显示
 */
func (a *Drawer) Visible(value interface{}) *Drawer {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Drawer) StaticInputClassName(value interface{}) *Drawer {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Drawer) Style(value interface{}) *Drawer {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Drawer) UseMobileUI(value interface{}) *Drawer {
    a.Set("useMobileUI", value)
    return a
}

/**
 * Dialog 大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *Drawer) Size(value interface{}) *Drawer {
    a.Set("size", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Drawer) DisabledOn(value interface{}) *Drawer {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Drawer) StaticClassName(value interface{}) *Drawer {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Drawer) StaticLabelClassName(value interface{}) *Drawer {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Drawer) StaticSchema(value interface{}) *Drawer {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Drawer) EditorSetting(value interface{}) *Drawer {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Drawer) Type(value interface{}) *Drawer {
    a.Set("type", value)
    return a
}

/**
 * 抽屉的高度 （当position为top | bottom时生效）
 */
func (a *Drawer) Height(value interface{}) *Drawer {
    a.Set("height", value)
    return a
}

/**
 * 底部
 */
func (a *Drawer) Footer(value interface{}) *Drawer {
    a.Set("footer", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Drawer) Hidden(value interface{}) *Drawer {
    a.Set("hidden", value)
    return a
}

/**
 * 是否可以拖动弹窗大小
 */
func (a *Drawer) Resizable(value interface{}) *Drawer {
    a.Set("resizable", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Drawer) Disabled(value interface{}) *Drawer {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Drawer) Static(value interface{}) *Drawer {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Drawer) StaticOn(value interface{}) *Drawer {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Drawer) StaticPlaceholder(value interface{}) *Drawer {
    a.Set("staticPlaceholder", value)
    return a
}
