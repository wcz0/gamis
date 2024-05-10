package renderers


/**
 * 分页容器功能性渲染器。详情请见：https://aisuda.bce.baidu.com/amis/zh-CN/components/pagination-wrapper

 * @author wcz0
 * @version 6.2.2
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
 * 是否禁用
 */
func (a *PaginationWrapper) Disabled(value interface{}) *PaginationWrapper {
    a.Set("disabled", value)
    return a
}

/**
 * 每页显示多条数据。
 */
func (a *PaginationWrapper) PerPage(value interface{}) *PaginationWrapper {
    a.Set("perPage", value)
    return a
}

/**
 * 分页显示位置，如果配置为 none 则需要自己在内容区域配置 pagination 组件，否则不显示。
 * 可选值: top | bottom | none
 */
func (a *PaginationWrapper) Position(value interface{}) *PaginationWrapper {
    a.Set("position", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *PaginationWrapper) Static(value interface{}) *PaginationWrapper {
    a.Set("static", value)
    return a
}

/**
 */
func (a *PaginationWrapper) StaticSchema(value interface{}) *PaginationWrapper {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *PaginationWrapper) EditorSetting(value interface{}) *PaginationWrapper {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *PaginationWrapper) UseMobileUI(value interface{}) *PaginationWrapper {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *PaginationWrapper) ClassName(value interface{}) *PaginationWrapper {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *PaginationWrapper) Hidden(value interface{}) *PaginationWrapper {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *PaginationWrapper) Visible(value interface{}) *PaginationWrapper {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *PaginationWrapper) VisibleOn(value interface{}) *PaginationWrapper {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *PaginationWrapper) StaticLabelClassName(value interface{}) *PaginationWrapper {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *PaginationWrapper) TestIdBuilder(value interface{}) *PaginationWrapper {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 指定为分页容器功能性渲染器
 */
func (a *PaginationWrapper) Type(value interface{}) *PaginationWrapper {
    a.Set("type", value)
    return a
}

/**
 * 是否显示快速跳转输入框
 */
func (a *PaginationWrapper) ShowPageInput(value interface{}) *PaginationWrapper {
    a.Set("showPageInput", value)
    return a
}

/**
 * 内容区域
 */
func (a *PaginationWrapper) Body(value interface{}) *PaginationWrapper {
    a.Set("body", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *PaginationWrapper) StaticOn(value interface{}) *PaginationWrapper {
    a.Set("staticOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *PaginationWrapper) Style(value interface{}) *PaginationWrapper {
    a.Set("style", value)
    return a
}

/**
 * 输入字段名
 */
func (a *PaginationWrapper) InputName(value interface{}) *PaginationWrapper {
    a.Set("inputName", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *PaginationWrapper) StaticClassName(value interface{}) *PaginationWrapper {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *PaginationWrapper) StaticInputClassName(value interface{}) *PaginationWrapper {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 输出字段名
 */
func (a *PaginationWrapper) OutputName(value interface{}) *PaginationWrapper {
    a.Set("outputName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *PaginationWrapper) DisabledOn(value interface{}) *PaginationWrapper {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *PaginationWrapper) HiddenOn(value interface{}) *PaginationWrapper {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *PaginationWrapper) Id(value interface{}) *PaginationWrapper {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *PaginationWrapper) OnEvent(value interface{}) *PaginationWrapper {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *PaginationWrapper) Testid(value interface{}) *PaginationWrapper {
    a.Set("testid", value)
    return a
}

/**
 * 最多显示多少个分页按钮。
 */
func (a *PaginationWrapper) MaxButtons(value interface{}) *PaginationWrapper {
    a.Set("maxButtons", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *PaginationWrapper) StaticPlaceholder(value interface{}) *PaginationWrapper {
    a.Set("staticPlaceholder", value)
    return a
}
