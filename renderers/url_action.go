package renderers


/**

* @author wcz0
* @version 6.2.2
*/
type UrlAction struct {
	*BaseRenderer
}

func NewUrlAction() *UrlAction {
    a := &UrlAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "url")
    a.Set("type", "button")
    return a
}
/**
 * 是否禁用
 */
func (a *UrlAction) Disabled(value string) *UrlAction {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *UrlAction) VisibleOn(value string) *UrlAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *UrlAction) StaticClassName(value string) *UrlAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *UrlAction) UseMobileUI(value string) *UrlAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *UrlAction) ConfirmText(value string) *UrlAction {
    a.Set("confirmText", value)
    return a
}

/**
 * 指定为打开链接
 */
func (a *UrlAction) ActionType(value string) *UrlAction {
    a.Set("actionType", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *UrlAction) DisabledOn(value string) *UrlAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *UrlAction) StaticInputClassName(value string) *UrlAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *UrlAction) LoadingOn(value string) *UrlAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * 打开的目标地址
 */
func (a *UrlAction) Url(value string) *UrlAction {
    a.Set("url", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *UrlAction) Static(value string) *UrlAction {
    a.Set("static", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *UrlAction) Target(value string) *UrlAction {
    a.Set("target", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *UrlAction) CountDownTpl(value string) *UrlAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 子内容
 */
func (a *UrlAction) Body(value string) *UrlAction {
    a.Set("body", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *UrlAction) Size(value string) *UrlAction {
    a.Set("size", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *UrlAction) ActiveLevel(value string) *UrlAction {
    a.Set("activeLevel", value)
    return a
}

/**
 */
func (a *UrlAction) StaticSchema(value string) *UrlAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *UrlAction) Level(value string) *UrlAction {
    a.Set("level", value)
    return a
}

/**
 */
func (a *UrlAction) Primary(value string) *UrlAction {
    a.Set("primary", value)
    return a
}

/**
 * 角标
 */
func (a *UrlAction) Badge(value string) *UrlAction {
    a.Set("badge", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *UrlAction) ClassName(value string) *UrlAction {
    a.Set("className", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *UrlAction) Type(value string) *UrlAction {
    a.Set("type", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *UrlAction) LoadingClassName(value string) *UrlAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *UrlAction) ActiveClassName(value string) *UrlAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *UrlAction) MergeData(value string) *UrlAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *UrlAction) Required(value string) *UrlAction {
    a.Set("required", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *UrlAction) HotKey(value string) *UrlAction {
    a.Set("hotKey", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *UrlAction) OnEvent(value string) *UrlAction {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *UrlAction) Testid(value string) *UrlAction {
    a.Set("testid", value)
    return a
}

/**
 */
func (a *UrlAction) Tooltip(value string) *UrlAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *UrlAction) RequireSelected(value string) *UrlAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *UrlAction) Hidden(value string) *UrlAction {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *UrlAction) HiddenOn(value string) *UrlAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *UrlAction) Visible(value string) *UrlAction {
    a.Set("visible", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *UrlAction) CountDown(value string) *UrlAction {
    a.Set("countDown", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *UrlAction) StaticOn(value string) *UrlAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *UrlAction) Style(value string) *UrlAction {
    a.Set("style", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *UrlAction) RightIcon(value string) *UrlAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *UrlAction) OnClick(value string) *UrlAction {
    a.Set("onClick", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *UrlAction) EditorSetting(value string) *UrlAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *UrlAction) Id(value string) *UrlAction {
    a.Set("id", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *UrlAction) IconClassName(value string) *UrlAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 按钮文字
 */
func (a *UrlAction) Label(value string) *UrlAction {
    a.Set("label", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *UrlAction) TooltipPlacement(value string) *UrlAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *UrlAction) Close(value string) *UrlAction {
    a.Set("close", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *UrlAction) Block(value string) *UrlAction {
    a.Set("block", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *UrlAction) Icon(value string) *UrlAction {
    a.Set("icon", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *UrlAction) RightIconClassName(value string) *UrlAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 是否新窗口打开
 */
func (a *UrlAction) Blank(value string) *UrlAction {
    a.Set("blank", value)
    return a
}

/**
 * 设置链接
 */
func (a *UrlAction) Link(value string) *UrlAction {
    a.Set("link", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *UrlAction) StaticPlaceholder(value string) *UrlAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *UrlAction) DisabledTip(value string) *UrlAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *UrlAction) StaticLabelClassName(value string) *UrlAction {
    a.Set("staticLabelClassName", value)
    return a
}
