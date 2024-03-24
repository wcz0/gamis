package renderers


/**
 * List 列表展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/card
 *

*/
type ListRenderer struct {
	*BaseRenderer
}

func NewListRenderer() *ListRenderer {
    a := &ListRenderer{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "list")
    return a
}
/**
 * 是否隐藏勾选框
 */
func (a *ListRenderer) HideCheckToggler(value string) *ListRenderer {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * 是否固底
 */
func (a *ListRenderer) AffixFooter(value string) *ListRenderer {
    a.Set("affixFooter", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ListRenderer) DisabledOn(value string) *ListRenderer {
    a.Set("disabledOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ListRenderer) EditorSetting(value string) *ListRenderer {
    a.Set("editorSetting", value)
    return a
}

/**
 * 底部区域类名
 */
func (a *ListRenderer) FooterClassName(value string) *ListRenderer {
    a.Set("footerClassName", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *ListRenderer) ShowFooter(value string) *ListRenderer {
    a.Set("showFooter", value)
    return a
}

/**
 * 大小
 * 可选值: sm | base
 */
func (a *ListRenderer) Size(value string) *ListRenderer {
    a.Set("size", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ListRenderer) ClassName(value string) *ListRenderer {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ListRenderer) HiddenOn(value string) *ListRenderer {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ListRenderer) UseMobileUI(value string) *ListRenderer {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否固顶
 */
func (a *ListRenderer) AffixHeader(value string) *ListRenderer {
    a.Set("affixHeader", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ListRenderer) Hidden(value string) *ListRenderer {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ListRenderer) Static(value string) *ListRenderer {
    a.Set("static", value)
    return a
}

/**
 * 数据源: 绑定当前环境变量
 */
func (a *ListRenderer) Source(value string) *ListRenderer {
    a.Set("source", value)
    return a
}

/**
 * 配置某项是否可以点选
 */
func (a *ListRenderer) ItemCheckableOn(value string) *ListRenderer {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * 点击列表单行时，是否选择
 */
func (a *ListRenderer) CheckOnItemClick(value string) *ListRenderer {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ListRenderer) Disabled(value string) *ListRenderer {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ListRenderer) StaticOn(value string) *ListRenderer {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ListRenderer) StaticInputClassName(value string) *ListRenderer {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *ListRenderer) StaticSchema(value string) *ListRenderer {
    a.Set("staticSchema", value)
    return a
}

/**
 * 底部区域
 */
func (a *ListRenderer) Footer(value string) *ListRenderer {
    a.Set("footer", value)
    return a
}

/**
 * 顶部区域类名
 */
func (a *ListRenderer) HeaderClassName(value string) *ListRenderer {
    a.Set("headerClassName", value)
    return a
}

/**
 * 是否显示头部
 */
func (a *ListRenderer) ShowHeader(value string) *ListRenderer {
    a.Set("showHeader", value)
    return a
}

/**
 * 是否显示
 */
func (a *ListRenderer) Visible(value string) *ListRenderer {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ListRenderer) Id(value string) *ListRenderer {
    a.Set("id", value)
    return a
}

/**
 * 无数据提示
 */
func (a *ListRenderer) Placeholder(value string) *ListRenderer {
    a.Set("placeholder", value)
    return a
}

/**
 * 点击列表项的行为
 */
func (a *ListRenderer) ItemAction(value string) *ListRenderer {
    a.Set("itemAction", value)
    return a
}

/**
 * 组件样式
 */
func (a *ListRenderer) Style(value string) *ListRenderer {
    a.Set("style", value)
    return a
}

/**
 * 标题
 */
func (a *ListRenderer) Title(value string) *ListRenderer {
    a.Set("title", value)
    return a
}

/**
 * 可以用来作为值的字段
 */
func (a *ListRenderer) ValueField(value string) *ListRenderer {
    a.Set("valueField", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ListRenderer) VisibleOn(value string) *ListRenderer {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ListRenderer) OnEvent(value string) *ListRenderer {
    a.Set("onEvent", value)
    return a
}

/**
 * 顶部区域
 */
func (a *ListRenderer) Header(value string) *ListRenderer {
    a.Set("header", value)
    return a
}

/**
 * 单条数据展示内容配置
 */
func (a *ListRenderer) ListItem(value string) *ListRenderer {
    a.Set("listItem", value)
    return a
}

/**
 * 配置某项是否可拖拽排序，前提是要开启拖拽功能
 */
func (a *ListRenderer) ItemDraggableOn(value string) *ListRenderer {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ListRenderer) StaticLabelClassName(value string) *ListRenderer {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 指定为 List 列表展示控件。
 * 可选值: list | static-list
 */
func (a *ListRenderer) Type(value string) *ListRenderer {
    a.Set("type", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ListRenderer) StaticPlaceholder(value string) *ListRenderer {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ListRenderer) StaticClassName(value string) *ListRenderer {
    a.Set("staticClassName", value)
    return a
}
