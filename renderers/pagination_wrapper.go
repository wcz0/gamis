package renderers


/**
 * 分页容器功能性渲染器。详情请见：https://aisuda.bce.baidu.com/amis/zh-CN/components/pagination-wrapper
 *

*/
type PaginationWrapper struct {
	*BaseRenderer
}

func NewPaginationWrapper() *PaginationWrapper {
    a := &PaginationWrapper{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "pagination-wrapper")
    return a
}
/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *PaginationWrapper) UseMobileUI(value string) *PaginationWrapper {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 输入字段名
 */
func (a *PaginationWrapper) InputName(value string) *PaginationWrapper {
    a.Set("inputName", value)
    return a
}

/**
 * 分页显示位置，如果配置为 none 则需要自己在内容区域配置 pagination 组件，否则不显示。
 * 可选值: top | bottom | none
 */
func (a *PaginationWrapper) Position(value string) *PaginationWrapper {
    a.Set("position", value)
    return a
}

/**
 * 是否显示
 */
func (a *PaginationWrapper) Visible(value string) *PaginationWrapper {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *PaginationWrapper) StaticLabelClassName(value string) *PaginationWrapper {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *PaginationWrapper) EditorSetting(value string) *PaginationWrapper {
    a.Set("editorSetting", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *PaginationWrapper) OnEvent(value string) *PaginationWrapper {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *PaginationWrapper) StaticPlaceholder(value string) *PaginationWrapper {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *PaginationWrapper) StaticClassName(value string) *PaginationWrapper {
    a.Set("staticClassName", value)
    return a
}

/**
 * 输出字段名
 */
func (a *PaginationWrapper) OutputName(value string) *PaginationWrapper {
    a.Set("outputName", value)
    return a
}

/**
 * 每页显示多条数据。
 */
func (a *PaginationWrapper) PerPage(value string) *PaginationWrapper {
    a.Set("perPage", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *PaginationWrapper) DisabledOn(value string) *PaginationWrapper {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *PaginationWrapper) HiddenOn(value string) *PaginationWrapper {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *PaginationWrapper) StaticOn(value string) *PaginationWrapper {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *PaginationWrapper) StaticSchema(value string) *PaginationWrapper {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *PaginationWrapper) Style(value string) *PaginationWrapper {
    a.Set("style", value)
    return a
}

/**
 * 指定为分页容器功能性渲染器
 */
func (a *PaginationWrapper) Type(value string) *PaginationWrapper {
    a.Set("type", value)
    return a
}

/**
 * 是否禁用
 */
func (a *PaginationWrapper) Disabled(value string) *PaginationWrapper {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *PaginationWrapper) StaticInputClassName(value string) *PaginationWrapper {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 最多显示多少个分页按钮。
 */
func (a *PaginationWrapper) MaxButtons(value string) *PaginationWrapper {
    a.Set("maxButtons", value)
    return a
}

/**
 * 内容区域
 */
func (a *PaginationWrapper) Body(value string) *PaginationWrapper {
    a.Set("body", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *PaginationWrapper) ClassName(value string) *PaginationWrapper {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *PaginationWrapper) Hidden(value string) *PaginationWrapper {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *PaginationWrapper) Id(value string) *PaginationWrapper {
    a.Set("id", value)
    return a
}

/**
 * 是否显示快速跳转输入框
 */
func (a *PaginationWrapper) ShowPageInput(value string) *PaginationWrapper {
    a.Set("showPageInput", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *PaginationWrapper) VisibleOn(value string) *PaginationWrapper {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *PaginationWrapper) Static(value string) *PaginationWrapper {
    a.Set("static", value)
    return a
}
