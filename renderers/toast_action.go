package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type ToastAction struct {
	*BaseRenderer
}

func NewToastAction() *ToastAction {
    a := &ToastAction{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "button")
    a.Set("actionType", "toast")
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *ToastAction) Type(value interface{}) *ToastAction {
    a.Set("type", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *ToastAction) RightIconClassName(value interface{}) *ToastAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ToastAction) OnEvent(value interface{}) *ToastAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ToastAction) StaticInputClassName(value interface{}) *ToastAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *ToastAction) Testid(value interface{}) *ToastAction {
    a.Set("testid", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *ToastAction) Icon(value interface{}) *ToastAction {
    a.Set("icon", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *ToastAction) RightIcon(value interface{}) *ToastAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *ToastAction) ActiveLevel(value interface{}) *ToastAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *ToastAction) HotKey(value interface{}) *ToastAction {
    a.Set("hotKey", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *ToastAction) LoadingOn(value interface{}) *ToastAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *ToastAction) Style(value interface{}) *ToastAction {
    a.Set("style", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *ToastAction) Block(value interface{}) *ToastAction {
    a.Set("block", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *ToastAction) IconClassName(value interface{}) *ToastAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 是否显示
 */
func (a *ToastAction) Visible(value interface{}) *ToastAction {
    a.Set("visible", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ToastAction) EditorSetting(value interface{}) *ToastAction {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *ToastAction) Tooltip(value interface{}) *ToastAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *ToastAction) ActiveClassName(value interface{}) *ToastAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *ToastAction) RequireSelected(value interface{}) *ToastAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *ToastAction) OnClick(value interface{}) *ToastAction {
    a.Set("onClick", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ToastAction) StaticClassName(value interface{}) *ToastAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *ToastAction) CountDown(value interface{}) *ToastAction {
    a.Set("countDown", value)
    return a
}

/**
 * 子内容
 */
func (a *ToastAction) Body(value interface{}) *ToastAction {
    a.Set("body", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ToastAction) HiddenOn(value interface{}) *ToastAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ToastAction) StaticOn(value interface{}) *ToastAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *ToastAction) CountDownTpl(value interface{}) *ToastAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ToastAction) Static(value interface{}) *ToastAction {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ToastAction) UseMobileUI(value interface{}) *ToastAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *ToastAction) Required(value interface{}) *ToastAction {
    a.Set("required", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ToastAction) StaticLabelClassName(value interface{}) *ToastAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *ToastAction) TestIdBuilder(value interface{}) *ToastAction {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 按钮文字
 */
func (a *ToastAction) Label(value interface{}) *ToastAction {
    a.Set("label", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *ToastAction) Id(value interface{}) *ToastAction {
    a.Set("id", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ToastAction) StaticPlaceholder(value interface{}) *ToastAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *ToastAction) Size(value interface{}) *ToastAction {
    a.Set("size", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *ToastAction) DisabledTip(value interface{}) *ToastAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *ToastAction) TooltipPlacement(value interface{}) *ToastAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 */
func (a *ToastAction) StaticSchema(value interface{}) *ToastAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *ToastAction) Level(value interface{}) *ToastAction {
    a.Set("level", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *ToastAction) Target(value interface{}) *ToastAction {
    a.Set("target", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ToastAction) VisibleOn(value interface{}) *ToastAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *ToastAction) Close(value interface{}) *ToastAction {
    a.Set("close", value)
    return a
}

/**
 * 指定为打开弹窗，抽出式弹窗
 */
func (a *ToastAction) ActionType(value interface{}) *ToastAction {
    a.Set("actionType", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ToastAction) ClassName(value interface{}) *ToastAction {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ToastAction) Disabled(value interface{}) *ToastAction {
    a.Set("disabled", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *ToastAction) ConfirmText(value interface{}) *ToastAction {
    a.Set("confirmText", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *ToastAction) MergeData(value interface{}) *ToastAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ToastAction) Hidden(value interface{}) *ToastAction {
    a.Set("hidden", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ToastAction) DisabledOn(value interface{}) *ToastAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 轻提示详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/toast
 */
func (a *ToastAction) Toast(value interface{}) *ToastAction {
    a.Set("toast", value)
    return a
}

/**
 * 角标
 */
func (a *ToastAction) Badge(value interface{}) *ToastAction {
    a.Set("badge", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *ToastAction) LoadingClassName(value interface{}) *ToastAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 */
func (a *ToastAction) Primary(value interface{}) *ToastAction {
    a.Set("primary", value)
    return a
}
