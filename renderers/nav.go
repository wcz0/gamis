package renderers


/**
 * Nav 导航渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/nav

* @author wcz0
* @version 6.2.2
*/
type Nav struct {
	*BaseRenderer
}

func NewNav() *Nav {
    a := &Nav{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "nav")
    return a
}
/**
 * 角标
 */
func (a *Nav) ItemBadge(value string) *Nav {
    a.Set("itemBadge", value)
    return a
}

/**
 * 垂直模式 非折叠状态下 控制菜单打开方式
 * 可选值: float | inline
 */
func (a *Nav) Mode(value string) *Nav {
    a.Set("mode", value)
    return a
}

/**
 * 自定义展开图标位置 默认在前面 before after
 */
func (a *Nav) ExpandPosition(value string) *Nav {
    a.Set("expandPosition", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Nav) Hidden(value string) *Nav {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Nav) Static(value string) *Nav {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Nav) UseMobileUI(value string) *Nav {
    a.Set("useMobileUI", value)
    return a
}

/**
 * true 为垂直排列，false 为水平排列类似如 tabs。
 * 可选值: true | false
 */
func (a *Nav) Stacked(value string) *Nav {
    a.Set("stacked", value)
    return a
}

/**
 * 更多操作菜单列表
 */
func (a *Nav) ItemActions(value string) *Nav {
    a.Set("itemActions", value)
    return a
}

/**
 * 控制仅展示指定key菜单下的子菜单项
 */
func (a *Nav) ShowKey(value string) *Nav {
    a.Set("showKey", value)
    return a
}

/**
 * 子菜单项展开浮层样式
 */
func (a *Nav) PopupClassName(value string) *Nav {
    a.Set("popupClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Nav) Disabled(value string) *Nav {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Nav) HiddenOn(value string) *Nav {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Nav) StaticOn(value string) *Nav {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Nav) StaticInputClassName(value string) *Nav {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 仅允许同层级拖拽
 */
func (a *Nav) DragOnSameLevel(value string) *Nav {
    a.Set("dragOnSameLevel", value)
    return a
}

/**
 * 是否开启搜索
 */
func (a *Nav) Searchable(value string) *Nav {
    a.Set("searchable", value)
    return a
}

/**
 */
func (a *Nav) Testid(value string) *Nav {
    a.Set("testid", value)
    return a
}

/**
 * 懒加载 api，如果不配置复用 source 接口。
 */
func (a *Nav) DeferApi(value string) *Nav {
    a.Set("deferApi", value)
    return a
}

/**
 * 角标
 */
func (a *Nav) Badge(value string) *Nav {
    a.Set("badge", value)
    return a
}

/**
 * 最多展示多少层级
 */
func (a *Nav) Level(value string) *Nav {
    a.Set("level", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Nav) DisabledOn(value string) *Nav {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Nav) VisibleOn(value string) *Nav {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Nav) StaticLabelClassName(value string) *Nav {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Nav) StaticPlaceholder(value string) *Nav {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 主题配色 默认light
 * 可选值: light | dark
 */
func (a *Nav) ThemeColor(value string) *Nav {
    a.Set("themeColor", value)
    return a
}

/**
 * 手风琴展开 仅垂直inline模式支持
 */
func (a *Nav) Accordion(value string) *Nav {
    a.Set("accordion", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Nav) ClassName(value string) *Nav {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Nav) Visible(value string) *Nav {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Nav) Id(value string) *Nav {
    a.Set("id", value)
    return a
}

/**
 * 默认展开层级 小于等于该层数的节点默认全部打开
 */
func (a *Nav) DefaultOpenLevel(value string) *Nav {
    a.Set("defaultOpenLevel", value)
    return a
}

/**
 * 控制菜单缩起
 */
func (a *Nav) Collapsed(value string) *Nav {
    a.Set("collapsed", value)
    return a
}

/**
 * 自定义展开图标
 */
func (a *Nav) ExpandIcon(value string) *Nav {
    a.Set("expandIcon", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Nav) StaticClassName(value string) *Nav {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Nav) StaticSchema(value string) *Nav {
    a.Set("staticSchema", value)
    return a
}

/**
 * 链接地址集合
 */
func (a *Nav) Links(value string) *Nav {
    a.Set("links", value)
    return a
}

/**
 * 指定为 Nav 导航渲染器
 */
func (a *Nav) Type(value string) *Nav {
    a.Set("type", value)
    return a
}

/**
 */
func (a *Nav) IndentSize(value string) *Nav {
    a.Set("indentSize", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Nav) OnEvent(value string) *Nav {
    a.Set("onEvent", value)
    return a
}

/**
 * 组件样式
 */
func (a *Nav) Style(value string) *Nav {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Nav) EditorSetting(value string) *Nav {
    a.Set("editorSetting", value)
    return a
}

/**
 * 横向导航时自动收纳配置
 */
func (a *Nav) Overflow(value string) *Nav {
    a.Set("overflow", value)
    return a
}

/**
 * 搜索框相关配置
 */
func (a *Nav) SearchConfig(value string) *Nav {
    a.Set("searchConfig", value)
    return a
}

/**
 * 可以通过 API 拉取。
 */
func (a *Nav) Source(value string) *Nav {
    a.Set("source", value)
    return a
}

/**
 * 可拖拽
 */
func (a *Nav) Draggable(value string) *Nav {
    a.Set("draggable", value)
    return a
}

/**
 * 保存排序的 api
 */
func (a *Nav) SaveOrderApi(value string) *Nav {
    a.Set("saveOrderApi", value)
    return a
}
