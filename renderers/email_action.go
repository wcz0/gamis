package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type EmailAction struct {
	*BaseRenderer
}

func NewEmailAction() *EmailAction {
    a := &EmailAction{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "button")
    a.Set("actionType", "email")
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *EmailAction) Target(value interface{}) *EmailAction {
    a.Set("target", value)
    return a
}

/**
 * 是否显示
 */
func (a *EmailAction) Visible(value interface{}) *EmailAction {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *EmailAction) StaticOn(value interface{}) *EmailAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *EmailAction) Type(value interface{}) *EmailAction {
    a.Set("type", value)
    return a
}

/**
 */
func (a *EmailAction) Tooltip(value interface{}) *EmailAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *EmailAction) Close(value interface{}) *EmailAction {
    a.Set("close", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *EmailAction) Static(value interface{}) *EmailAction {
    a.Set("static", value)
    return a
}

/**
 */
func (a *EmailAction) StaticSchema(value interface{}) *EmailAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *EmailAction) ActiveClassName(value interface{}) *EmailAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 抄送邮箱
 */
func (a *EmailAction) Cc(value interface{}) *EmailAction {
    a.Set("cc", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *EmailAction) ClassName(value interface{}) *EmailAction {
    a.Set("className", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *EmailAction) StaticPlaceholder(value interface{}) *EmailAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *EmailAction) IconClassName(value interface{}) *EmailAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *EmailAction) ConfirmText(value interface{}) *EmailAction {
    a.Set("confirmText", value)
    return a
}

/**
 * 邮件正文
 */
func (a *EmailAction) Body(value interface{}) *EmailAction {
    a.Set("body", value)
    return a
}

/**
 * 是否禁用
 */
func (a *EmailAction) Disabled(value interface{}) *EmailAction {
    a.Set("disabled", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *EmailAction) ActiveLevel(value interface{}) *EmailAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *EmailAction) RequireSelected(value interface{}) *EmailAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 邮件主题
 */
func (a *EmailAction) Subject(value interface{}) *EmailAction {
    a.Set("subject", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *EmailAction) StaticLabelClassName(value interface{}) *EmailAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *EmailAction) TestIdBuilder(value interface{}) *EmailAction {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 组件样式
 */
func (a *EmailAction) Style(value interface{}) *EmailAction {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *EmailAction) UseMobileUI(value interface{}) *EmailAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 按钮文字
 */
func (a *EmailAction) Label(value interface{}) *EmailAction {
    a.Set("label", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *EmailAction) MergeData(value interface{}) *EmailAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *EmailAction) LoadingOn(value interface{}) *EmailAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * 匿名抄送邮箱
 */
func (a *EmailAction) Bcc(value interface{}) *EmailAction {
    a.Set("bcc", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *EmailAction) LoadingClassName(value interface{}) *EmailAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *EmailAction) CountDown(value interface{}) *EmailAction {
    a.Set("countDown", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *EmailAction) OnClick(value interface{}) *EmailAction {
    a.Set("onClick", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *EmailAction) DisabledOn(value interface{}) *EmailAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *EmailAction) OnEvent(value interface{}) *EmailAction {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *EmailAction) Testid(value interface{}) *EmailAction {
    a.Set("testid", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *EmailAction) RightIconClassName(value interface{}) *EmailAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *EmailAction) Required(value interface{}) *EmailAction {
    a.Set("required", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *EmailAction) CountDownTpl(value interface{}) *EmailAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *EmailAction) HotKey(value interface{}) *EmailAction {
    a.Set("hotKey", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *EmailAction) StaticClassName(value interface{}) *EmailAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *EmailAction) Icon(value interface{}) *EmailAction {
    a.Set("icon", value)
    return a
}

/**
 * 指定为打开邮箱行为
 */
func (a *EmailAction) ActionType(value interface{}) *EmailAction {
    a.Set("actionType", value)
    return a
}

/**
 * 收件人邮箱
 */
func (a *EmailAction) To(value interface{}) *EmailAction {
    a.Set("to", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *EmailAction) StaticInputClassName(value interface{}) *EmailAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *EmailAction) Hidden(value interface{}) *EmailAction {
    a.Set("hidden", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *EmailAction) Level(value interface{}) *EmailAction {
    a.Set("level", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *EmailAction) HiddenOn(value interface{}) *EmailAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *EmailAction) EditorSetting(value interface{}) *EmailAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *EmailAction) Block(value interface{}) *EmailAction {
    a.Set("block", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *EmailAction) DisabledTip(value interface{}) *EmailAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *EmailAction) VisibleOn(value interface{}) *EmailAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *EmailAction) Id(value interface{}) *EmailAction {
    a.Set("id", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *EmailAction) RightIcon(value interface{}) *EmailAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *EmailAction) Size(value interface{}) *EmailAction {
    a.Set("size", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *EmailAction) TooltipPlacement(value interface{}) *EmailAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 */
func (a *EmailAction) Primary(value interface{}) *EmailAction {
    a.Set("primary", value)
    return a
}

/**
 * 角标
 */
func (a *EmailAction) Badge(value interface{}) *EmailAction {
    a.Set("badge", value)
    return a
}
