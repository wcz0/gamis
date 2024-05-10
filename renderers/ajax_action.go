package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type AjaxAction struct {
	*BaseRenderer
}

func NewAjaxAction() *AjaxAction {
    a := &AjaxAction{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("actionType", "ajax")
    a.Set("type", "button")
    return a
}
/**
 * icon 上的css 类名
 */
func (a *AjaxAction) IconClassName(value interface{}) *AjaxAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 角标
 */
func (a *AjaxAction) Badge(value interface{}) *AjaxAction {
    a.Set("badge", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *AjaxAction) Static(value interface{}) *AjaxAction {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *AjaxAction) StaticLabelClassName(value interface{}) *AjaxAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *AjaxAction) Primary(value interface{}) *AjaxAction {
    a.Set("primary", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *AjaxAction) VisibleOn(value interface{}) *AjaxAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *AjaxAction) Block(value interface{}) *AjaxAction {
    a.Set("block", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *AjaxAction) Icon(value interface{}) *AjaxAction {
    a.Set("icon", value)
    return a
}

/**
 */
func (a *AjaxAction) Feedback(value interface{}) *AjaxAction {
    a.Set("feedback", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *AjaxAction) OnEvent(value interface{}) *AjaxAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *AjaxAction) RightIcon(value interface{}) *AjaxAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *AjaxAction) CountDownTpl(value interface{}) *AjaxAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 子内容
 */
func (a *AjaxAction) Body(value interface{}) *AjaxAction {
    a.Set("body", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *AjaxAction) ClassName(value interface{}) *AjaxAction {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *AjaxAction) Hidden(value interface{}) *AjaxAction {
    a.Set("hidden", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *AjaxAction) Target(value interface{}) *AjaxAction {
    a.Set("target", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *AjaxAction) OnClick(value interface{}) *AjaxAction {
    a.Set("onClick", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *AjaxAction) DisabledOn(value interface{}) *AjaxAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *AjaxAction) StaticOn(value interface{}) *AjaxAction {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *AjaxAction) Reload(value interface{}) *AjaxAction {
    a.Set("reload", value)
    return a
}

/**
 */
func (a *AjaxAction) IgnoreConfirm(value interface{}) *AjaxAction {
    a.Set("ignoreConfirm", value)
    return a
}

/**
 * 是否显示
 */
func (a *AjaxAction) Visible(value interface{}) *AjaxAction {
    a.Set("visible", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *AjaxAction) ActiveClassName(value interface{}) *AjaxAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 配置 ajax 发送地址
 */
func (a *AjaxAction) Api(value interface{}) *AjaxAction {
    a.Set("api", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *AjaxAction) DisabledTip(value interface{}) *AjaxAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *AjaxAction) Close(value interface{}) *AjaxAction {
    a.Set("close", value)
    return a
}

/**
 * 是否禁用
 */
func (a *AjaxAction) Disabled(value interface{}) *AjaxAction {
    a.Set("disabled", value)
    return a
}

/**
 * 组件样式
 */
func (a *AjaxAction) Style(value interface{}) *AjaxAction {
    a.Set("style", value)
    return a
}

/**
 * 指定为发送 ajax 的行为。
 */
func (a *AjaxAction) ActionType(value interface{}) *AjaxAction {
    a.Set("actionType", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *AjaxAction) Size(value interface{}) *AjaxAction {
    a.Set("size", value)
    return a
}

/**
 */
func (a *AjaxAction) Tooltip(value interface{}) *AjaxAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 按钮文字
 */
func (a *AjaxAction) Label(value interface{}) *AjaxAction {
    a.Set("label", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *AjaxAction) Level(value interface{}) *AjaxAction {
    a.Set("level", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *AjaxAction) LoadingClassName(value interface{}) *AjaxAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *AjaxAction) TooltipPlacement(value interface{}) *AjaxAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *AjaxAction) CountDown(value interface{}) *AjaxAction {
    a.Set("countDown", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *AjaxAction) HiddenOn(value interface{}) *AjaxAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *AjaxAction) Id(value interface{}) *AjaxAction {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *AjaxAction) StaticInputClassName(value interface{}) *AjaxAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *AjaxAction) StaticSchema(value interface{}) *AjaxAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *AjaxAction) Type(value interface{}) *AjaxAction {
    a.Set("type", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *AjaxAction) StaticClassName(value interface{}) *AjaxAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *AjaxAction) Required(value interface{}) *AjaxAction {
    a.Set("required", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *AjaxAction) UseMobileUI(value interface{}) *AjaxAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *AjaxAction) ConfirmText(value interface{}) *AjaxAction {
    a.Set("confirmText", value)
    return a
}

/**
 */
func (a *AjaxAction) Redirect(value interface{}) *AjaxAction {
    a.Set("redirect", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *AjaxAction) StaticPlaceholder(value interface{}) *AjaxAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *AjaxAction) RightIconClassName(value interface{}) *AjaxAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *AjaxAction) RequireSelected(value interface{}) *AjaxAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *AjaxAction) MergeData(value interface{}) *AjaxAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *AjaxAction) HotKey(value interface{}) *AjaxAction {
    a.Set("hotKey", value)
    return a
}

/**
 */
func (a *AjaxAction) Testid(value interface{}) *AjaxAction {
    a.Set("testid", value)
    return a
}

/**
 * 是否开启请求隔离, 主要用于隔离联动CRUD, Service的请求
 */
func (a *AjaxAction) IsolateScope(value interface{}) *AjaxAction {
    a.Set("isolateScope", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *AjaxAction) EditorSetting(value interface{}) *AjaxAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *AjaxAction) ActiveLevel(value interface{}) *AjaxAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *AjaxAction) LoadingOn(value interface{}) *AjaxAction {
    a.Set("loadingOn", value)
    return a
}
