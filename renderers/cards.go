package renderers


/**
 * Cards 卡片集合渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/card

* @author wcz0
* @version 6.2.2
*/
type Cards struct {
	*BaseRenderer
}

func NewCards() *Cards {
    a := &Cards{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "cards")
    return a
}
/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Cards) Id(value string) *Cards {
    a.Set("id", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Cards) StaticOn(value string) *Cards {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Cards) StaticPlaceholder(value string) *Cards {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Cards) UseMobileUI(value string) *Cards {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Cards) Card(value string) *Cards {
    a.Set("card", value)
    return a
}

/**
 * 头部 CSS 类名
 */
func (a *Cards) HeaderClassName(value string) *Cards {
    a.Set("headerClassName", value)
    return a
}

/**
 * 无数据提示
 */
func (a *Cards) Placeholder(value string) *Cards {
    a.Set("placeholder", value)
    return a
}

/**
 * 配置某项是否可以点选
 */
func (a *Cards) ItemCheckableOn(value string) *Cards {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * 配置某项是否可拖拽排序，前提是要开启拖拽功能
 */
func (a *Cards) ItemDraggableOn(value string) *Cards {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Cards) Hidden(value string) *Cards {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Cards) Visible(value string) *Cards {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Cards) StaticInputClassName(value string) *Cards {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 数据源: 绑定当前环境变量
 */
func (a *Cards) Source(value string) *Cards {
    a.Set("source", value)
    return a
}

/**
 * 点击卡片的时候是否勾选卡片。
 */
func (a *Cards) CheckOnItemClick(value string) *Cards {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 */
func (a *Cards) LoadingConfig(value string) *Cards {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Cards) HiddenOn(value string) *Cards {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *Cards) StaticSchema(value string) *Cards {
    a.Set("staticSchema", value)
    return a
}

/**
 * 底部 CSS 类名
 */
func (a *Cards) FooterClassName(value string) *Cards {
    a.Set("footerClassName", value)
    return a
}

/**
 * 卡片 CSS 类名
 */
func (a *Cards) ItemClassName(value string) *Cards {
    a.Set("itemClassName", value)
    return a
}

/**
 * 底部区域
 */
func (a *Cards) Footer(value string) *Cards {
    a.Set("footer", value)
    return a
}

/**
 * 是否为瀑布流布局？
 */
func (a *Cards) MasonryLayout(value string) *Cards {
    a.Set("masonryLayout", value)
    return a
}

/**
 * 可以用来作为值的字段
 */
func (a *Cards) ValueField(value string) *Cards {
    a.Set("valueField", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Cards) ClassName(value string) *Cards {
    a.Set("className", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Cards) Static(value string) *Cards {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Cards) StaticClassName(value string) *Cards {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否固底
 */
func (a *Cards) AffixFooter(value string) *Cards {
    a.Set("affixFooter", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Cards) Disabled(value string) *Cards {
    a.Set("disabled", value)
    return a
}

/**
 * 组件样式
 */
func (a *Cards) Style(value string) *Cards {
    a.Set("style", value)
    return a
}

/**
 * 是否隐藏勾选框
 */
func (a *Cards) HideCheckToggler(value string) *Cards {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * 是否固顶
 */
func (a *Cards) AffixHeader(value string) *Cards {
    a.Set("affixHeader", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Cards) StaticLabelClassName(value string) *Cards {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Cards) VisibleOn(value string) *Cards {
    a.Set("visibleOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Cards) EditorSetting(value string) *Cards {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Cards) DisabledOn(value string) *Cards {
    a.Set("disabledOn", value)
    return a
}

/**
 * 指定为 cards 类型
 */
func (a *Cards) Type(value string) *Cards {
    a.Set("type", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *Cards) ShowFooter(value string) *Cards {
    a.Set("showFooter", value)
    return a
}

/**
 * 是否显示头部
 */
func (a *Cards) ShowHeader(value string) *Cards {
    a.Set("showHeader", value)
    return a
}

/**
 * 标题
 */
func (a *Cards) Title(value string) *Cards {
    a.Set("title", value)
    return a
}

/**
 * 顶部区域
 */
func (a *Cards) Header(value string) *Cards {
    a.Set("header", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Cards) OnEvent(value string) *Cards {
    a.Set("onEvent", value)
    return a
}
