package renderers


/**
 * 选项卡控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/tabs

 * @author wcz0
 * @version 6.2.2
 */
type Tabs struct {
	*BaseRenderer
}

func NewTabs() *Tabs {
    a := &Tabs{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "tabs")
    return a
}
/**
 * 激活的选项卡，hash值或索引值，支持使用表达式
 */
func (a *Tabs) ActiveKey(value interface{}) *Tabs {
    a.Set("activeKey", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Tabs) ClassName(value interface{}) *Tabs {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Tabs) DisabledOn(value interface{}) *Tabs {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Tabs) StaticPlaceholder(value interface{}) *Tabs {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *Tabs) Type(value interface{}) *Tabs {
    a.Set("type", value)
    return a
}

/**
 * 卡片隐藏的时候是否销毁卡片内容
 */
func (a *Tabs) UnmountOnExit(value interface{}) *Tabs {
    a.Set("unmountOnExit", value)
    return a
}

/**
 * 配置子表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *Tabs) SubFormMode(value interface{}) *Tabs {
    a.Set("subFormMode", value)
    return a
}

/**
 * 编辑器模式，侧边的位置
 * 可选值: left | right
 */
func (a *Tabs) SidePosition(value interface{}) *Tabs {
    a.Set("sidePosition", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Tabs) VisibleOn(value interface{}) *Tabs {
    a.Set("visibleOn", value)
    return a
}

/**
 * 内容类名
 */
func (a *Tabs) ContentClassName(value interface{}) *Tabs {
    a.Set("contentClassName", value)
    return a
}

/**
 * tooltip 提示的类名
 */
func (a *Tabs) ShowTipClassName(value interface{}) *Tabs {
    a.Set("showTipClassName", value)
    return a
}

/**
 * 超过多少个时折叠按钮
 */
func (a *Tabs) CollapseOnExceed(value interface{}) *Tabs {
    a.Set("collapseOnExceed", value)
    return a
}

/**
 * 是否支持拖拽
 */
func (a *Tabs) Draggable(value interface{}) *Tabs {
    a.Set("draggable", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Tabs) Hidden(value interface{}) *Tabs {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Tabs) Visible(value interface{}) *Tabs {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Tabs) Id(value interface{}) *Tabs {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Tabs) StaticClassName(value interface{}) *Tabs {
    a.Set("staticClassName", value)
    return a
}

/**
 * 展示形式
 */
func (a *Tabs) TabsMode(value interface{}) *Tabs {
    a.Set("tabsMode", value)
    return a
}

/**
 * 链接外层类名
 */
func (a *Tabs) LinksClassName(value interface{}) *Tabs {
    a.Set("linksClassName", value)
    return a
}

/**
 * 是否支持新增
 */
func (a *Tabs) Addable(value interface{}) *Tabs {
    a.Set("addable", value)
    return a
}

/**
 * 是否显示提示
 */
func (a *Tabs) ShowTip(value interface{}) *Tabs {
    a.Set("showTip", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Tabs) HiddenOn(value interface{}) *Tabs {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Tabs) StaticInputClassName(value interface{}) *Tabs {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 初始化激活的选项卡，hash值或索引值，支持使用表达式
 */
func (a *Tabs) DefaultKey(value interface{}) *Tabs {
    a.Set("defaultKey", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Tabs) StaticOn(value interface{}) *Tabs {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Tabs) StaticLabelClassName(value interface{}) *Tabs {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Tabs) EditorSetting(value interface{}) *Tabs {
    a.Set("editorSetting", value)
    return a
}

/**
 * 卡片是否只有在点开的时候加载？
 */
func (a *Tabs) MountOnEnter(value interface{}) *Tabs {
    a.Set("mountOnEnter", value)
    return a
}

/**
 * 折叠按钮文字
 */
func (a *Tabs) CollapseBtnLabel(value interface{}) *Tabs {
    a.Set("collapseBtnLabel", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Tabs) Disabled(value interface{}) *Tabs {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Tabs) OnEvent(value interface{}) *Tabs {
    a.Set("onEvent", value)
    return a
}

/**
 * 关联已有数据，选项卡直接根据目标数据重复。
 */
func (a *Tabs) Source(value interface{}) *Tabs {
    a.Set("source", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *Tabs) SubFormHorizontal(value interface{}) *Tabs {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * 是否导航支持内容溢出滚动。属性废弃，为了兼容暂且保留
 */
func (a *Tabs) Scrollable(value interface{}) *Tabs {
    a.Set("scrollable", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Tabs) Static(value interface{}) *Tabs {
    a.Set("static", value)
    return a
}

/**
 */
func (a *Tabs) StaticSchema(value interface{}) *Tabs {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *Tabs) Style(value interface{}) *Tabs {
    a.Set("style", value)
    return a
}

/**
 * 选项卡成员。当配置了 source 时，选项卡成员，将会根据目标数据进行重复。
 */
func (a *Tabs) Tabs(value interface{}) *Tabs {
    a.Set("tabs", value)
    return a
}

/**
 * 是否可编辑标签名
 */
func (a *Tabs) Editable(value interface{}) *Tabs {
    a.Set("editable", value)
    return a
}

/**
 * 自定义增加按钮文案
 */
func (a *Tabs) AddBtnText(value interface{}) *Tabs {
    a.Set("addBtnText", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Tabs) UseMobileUI(value interface{}) *Tabs {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 可以在右侧配置点其他功能按钮。
 */
func (a *Tabs) Toolbar(value interface{}) *Tabs {
    a.Set("toolbar", value)
    return a
}

/**
 * 是否支持删除
 */
func (a *Tabs) Closable(value interface{}) *Tabs {
    a.Set("closable", value)
    return a
}

/**
 * 是否滑动切换只在移动端生效
 */
func (a *Tabs) Swipeable(value interface{}) *Tabs {
    a.Set("swipeable", value)
    return a
}
