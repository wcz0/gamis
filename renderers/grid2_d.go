package renderers


/**
 * 二维布局渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/grid-2d

 * @author wcz0
 * @version 6.2.2
 */
type Grid2D struct {
	*BaseRenderer
}

func NewGrid2D() *Grid2D {
    a := &Grid2D{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "grid-2d")
    return a
}
/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Grid2D) Id(value interface{}) *Grid2D {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Grid2D) StaticLabelClassName(value interface{}) *Grid2D {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Grid2D) StaticInputClassName(value interface{}) *Grid2D {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Grid2D) Style(value interface{}) *Grid2D {
    a.Set("style", value)
    return a
}

/**
 * 单位行高度，默认 50 px
 */
func (a *Grid2D) RowHeight(value interface{}) *Grid2D {
    a.Set("rowHeight", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Grid2D) DisabledOn(value interface{}) *Grid2D {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Grid2D) HiddenOn(value interface{}) *Grid2D {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Grid2D) StaticPlaceholder(value interface{}) *Grid2D {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 指定为 grid-2d 展示类型
 */
func (a *Grid2D) Type(value interface{}) *Grid2D {
    a.Set("type", value)
    return a
}

/**
 * grid 2d 容器宽度，默认是 auto
 */
func (a *Grid2D) Width(value interface{}) *Grid2D {
    a.Set("width", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Grid2D) Disabled(value interface{}) *Grid2D {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Grid2D) VisibleOn(value interface{}) *Grid2D {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *Grid2D) StaticSchema(value interface{}) *Grid2D {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Grid2D) EditorSetting(value interface{}) *Grid2D {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Grid2D) Static(value interface{}) *Grid2D {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Grid2D) StaticOn(value interface{}) *Grid2D {
    a.Set("staticOn", value)
    return a
}

/**
 * 列数量，默认是 12
 */
func (a *Grid2D) Cols(value interface{}) *Grid2D {
    a.Set("cols", value)
    return a
}

/**
 * 格子间距，默认 0，包含行和列
 */
func (a *Grid2D) Gap(value interface{}) *Grid2D {
    a.Set("gap", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Grid2D) ClassName(value interface{}) *Grid2D {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Grid2D) Visible(value interface{}) *Grid2D {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Grid2D) StaticClassName(value interface{}) *Grid2D {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Grid2D) Hidden(value interface{}) *Grid2D {
    a.Set("hidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Grid2D) OnEvent(value interface{}) *Grid2D {
    a.Set("onEvent", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Grid2D) UseMobileUI(value interface{}) *Grid2D {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 格子行级别的间距，如果不设置就和 gap 一样
 */
func (a *Grid2D) GapRow(value interface{}) *Grid2D {
    a.Set("gapRow", value)
    return a
}

/**
 * 每个格子的配置
 */
func (a *Grid2D) Grids(value interface{}) *Grid2D {
    a.Set("grids", value)
    return a
}
