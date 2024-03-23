package renderers


/**

*/
type AjaxAction struct {
	*BaseRenderer
}

func NewAjaxAction() *AjaxAction {
    a := &AjaxAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "button")
    a.Set("actionType", "ajax")
    return a
}
/**
 * 静态展示表单项类名
 */
func (a *AjaxAction) StaticClassName(value string) *AjaxAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *AjaxAction) ClassName(value string) *AjaxAction {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *AjaxAction) DisabledOn(value string) *AjaxAction {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *AjaxAction) StaticSchema(value string) *AjaxAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *AjaxAction) HotKey(value string) *AjaxAction {
    a.Set("hotKey", value)
    return a
}

/**
 */
func (a *AjaxAction) Redirect(value string) *AjaxAction {
    a.Set("redirect", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *AjaxAction) Id(value string) *AjaxAction {
    a.Set("id", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *AjaxAction) Icon(value string) *AjaxAction {
    a.Set("icon", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *AjaxAction) Static(value string) *AjaxAction {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *AjaxAction) StaticOn(value string) *AjaxAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *AjaxAction) DisabledTip(value string) *AjaxAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *AjaxAction) ActiveLevel(value string) *AjaxAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *AjaxAction) MergeData(value string) *AjaxAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 是否禁用
 */
func (a *AjaxAction) Disabled(value string) *AjaxAction {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *AjaxAction) Hidden(value string) *AjaxAction {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *AjaxAction) StaticPlaceholder(value string) *AjaxAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *AjaxAction) ConfirmText(value string) *AjaxAction {
    a.Set("confirmText", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *AjaxAction) CountDown(value string) *AjaxAction {
    a.Set("countDown", value)
    return a
}

/**
 */
func (a *AjaxAction) Reload(value string) *AjaxAction {
    a.Set("reload", value)
    return a
}

/**
 * 是否开启请求隔离, 主要用于隔离联动CRUD, Service的请求
 */
func (a *AjaxAction) IsolateScope(value string) *AjaxAction {
    a.Set("isolateScope", value)
    return a
}

/**
 * 组件样式
 */
func (a *AjaxAction) Style(value string) *AjaxAction {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *AjaxAction) UseMobileUI(value string) *AjaxAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *AjaxAction) Block(value string) *AjaxAction {
    a.Set("block", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *AjaxAction) CountDownTpl(value string) *AjaxAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *AjaxAction) LoadingOn(value string) *AjaxAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *AjaxAction) RightIconClassName(value string) *AjaxAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *AjaxAction) Size(value string) *AjaxAction {
    a.Set("size", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *AjaxAction) Required(value string) *AjaxAction {
    a.Set("required", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *AjaxAction) IconClassName(value string) *AjaxAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *AjaxAction) ActiveClassName(value string) *AjaxAction {
    a.Set("activeClassName", value)
    return a
}

/**
 */
func (a *AjaxAction) Tooltip(value string) *AjaxAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *AjaxAction) Close(value string) *AjaxAction {
    a.Set("close", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *AjaxAction) OnClick(value string) *AjaxAction {
    a.Set("onClick", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *AjaxAction) OnEvent(value string) *AjaxAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *AjaxAction) StaticLabelClassName(value string) *AjaxAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *AjaxAction) Type(value string) *AjaxAction {
    a.Set("type", value)
    return a
}

/**
 */
func (a *AjaxAction) Feedback(value string) *AjaxAction {
    a.Set("feedback", value)
    return a
}

/**
 * 按钮文字
 */
func (a *AjaxAction) Label(value string) *AjaxAction {
    a.Set("label", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *AjaxAction) Level(value string) *AjaxAction {
    a.Set("level", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *AjaxAction) TooltipPlacement(value string) *AjaxAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 角标
 */
func (a *AjaxAction) Badge(value string) *AjaxAction {
    a.Set("badge", value)
    return a
}

/**
 * 子内容
 */
func (a *AjaxAction) Body(value string) *AjaxAction {
    a.Set("body", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *AjaxAction) HiddenOn(value string) *AjaxAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *AjaxAction) Visible(value string) *AjaxAction {
    a.Set("visible", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *AjaxAction) RequireSelected(value string) *AjaxAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 配置 ajax 发送地址
 */
func (a *AjaxAction) Api(value string) *AjaxAction {
    a.Set("api", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *AjaxAction) VisibleOn(value string) *AjaxAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *AjaxAction) StaticInputClassName(value string) *AjaxAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *AjaxAction) EditorSetting(value string) *AjaxAction {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *AjaxAction) Testid(value string) *AjaxAction {
    a.Set("testid", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *AjaxAction) RightIcon(value string) *AjaxAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *AjaxAction) LoadingClassName(value string) *AjaxAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 */
func (a *AjaxAction) Primary(value string) *AjaxAction {
    a.Set("primary", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *AjaxAction) Target(value string) *AjaxAction {
    a.Set("target", value)
    return a
}

/**
 * 指定为发送 ajax 的行为。
 */
func (a *AjaxAction) ActionType(value string) *AjaxAction {
    a.Set("actionType", value)
    return a
}

/**
 */
func (a *AjaxAction) IgnoreConfirm(value string) *AjaxAction {
    a.Set("ignoreConfirm", value)
    return a
}
