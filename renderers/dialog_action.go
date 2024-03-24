package renderers


/**

* @author wcz0
* @version 6.2.2
*/
type DialogAction struct {
	*BaseRenderer
}

func NewDialogAction() *DialogAction {
    a := &DialogAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "button")
    a.Set("actionType", "dialog")
    return a
}
/**
 * 事件动作配置
 */
func (a *DialogAction) OnEvent(value string) *DialogAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 组件样式
 */
func (a *DialogAction) Style(value string) *DialogAction {
    a.Set("style", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *DialogAction) Icon(value string) *DialogAction {
    a.Set("icon", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *DialogAction) DisabledTip(value string) *DialogAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *DialogAction) RequireSelected(value string) *DialogAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *DialogAction) MergeData(value string) *DialogAction {
    a.Set("mergeData", value)
    return a
}

/**
 */
func (a *DialogAction) Redirect(value string) *DialogAction {
    a.Set("redirect", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *DialogAction) Close(value string) *DialogAction {
    a.Set("close", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *DialogAction) LoadingClassName(value string) *DialogAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * 是否有下一个的表达式，正常可以不用配置，如果想要刷掉某些数据可以配置这个。
 */
func (a *DialogAction) NextCondition(value string) *DialogAction {
    a.Set("nextCondition", value)
    return a
}

/**
 * 是否禁用
 */
func (a *DialogAction) Disabled(value string) *DialogAction {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示
 */
func (a *DialogAction) Visible(value string) *DialogAction {
    a.Set("visible", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *DialogAction) Target(value string) *DialogAction {
    a.Set("target", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *DialogAction) LoadingOn(value string) *DialogAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *DialogAction) StaticPlaceholder(value string) *DialogAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *DialogAction) StaticClassName(value string) *DialogAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 指定为打开弹窗
 */
func (a *DialogAction) ActionType(value string) *DialogAction {
    a.Set("actionType", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *DialogAction) VisibleOn(value string) *DialogAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *DialogAction) Static(value string) *DialogAction {
    a.Set("static", value)
    return a
}

/**
 */
func (a *DialogAction) Testid(value string) *DialogAction {
    a.Set("testid", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *DialogAction) IconClassName(value string) *DialogAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *DialogAction) DisabledOn(value string) *DialogAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *DialogAction) RightIconClassName(value string) *DialogAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *DialogAction) RightIcon(value string) *DialogAction {
    a.Set("rightIcon", value)
    return a
}

/**
 */
func (a *DialogAction) Tooltip(value string) *DialogAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *DialogAction) ActiveClassName(value string) *DialogAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *DialogAction) StaticOn(value string) *DialogAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *DialogAction) CountDownTpl(value string) *DialogAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 弹框详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/dialog
 */
func (a *DialogAction) Dialog(value string) *DialogAction {
    a.Set("dialog", value)
    return a
}

/**
 * 按钮文字
 */
func (a *DialogAction) Label(value string) *DialogAction {
    a.Set("label", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *DialogAction) HiddenOn(value string) *DialogAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *DialogAction) StaticSchema(value string) *DialogAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *DialogAction) OnClick(value string) *DialogAction {
    a.Set("onClick", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *DialogAction) StaticInputClassName(value string) *DialogAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *DialogAction) UseMobileUI(value string) *DialogAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *DialogAction) Required(value string) *DialogAction {
    a.Set("required", value)
    return a
}

/**
 * 角标
 */
func (a *DialogAction) Badge(value string) *DialogAction {
    a.Set("badge", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *DialogAction) HotKey(value string) *DialogAction {
    a.Set("hotKey", value)
    return a
}

/**
 */
func (a *DialogAction) Reload(value string) *DialogAction {
    a.Set("reload", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *DialogAction) ClassName(value string) *DialogAction {
    a.Set("className", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *DialogAction) EditorSetting(value string) *DialogAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *DialogAction) Level(value string) *DialogAction {
    a.Set("level", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *DialogAction) Size(value string) *DialogAction {
    a.Set("size", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *DialogAction) Hidden(value string) *DialogAction {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *DialogAction) StaticLabelClassName(value string) *DialogAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *DialogAction) TooltipPlacement(value string) *DialogAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *DialogAction) CountDown(value string) *DialogAction {
    a.Set("countDown", value)
    return a
}

/**
 */
func (a *DialogAction) Primary(value string) *DialogAction {
    a.Set("primary", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *DialogAction) ActiveLevel(value string) *DialogAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *DialogAction) ConfirmText(value string) *DialogAction {
    a.Set("confirmText", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *DialogAction) Id(value string) *DialogAction {
    a.Set("id", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *DialogAction) Type(value string) *DialogAction {
    a.Set("type", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *DialogAction) Block(value string) *DialogAction {
    a.Set("block", value)
    return a
}

/**
 * 子内容
 */
func (a *DialogAction) Body(value string) *DialogAction {
    a.Set("body", value)
    return a
}
