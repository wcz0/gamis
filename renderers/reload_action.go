package renderers


/**

* @author wcz0
* @version 6.2.2
*/
type ReloadAction struct {
	*BaseRenderer
}

func NewReloadAction() *ReloadAction {
    a := &ReloadAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "reload")
    a.Set("type", "button")
    return a
}
/**
 * 激活状态时的类名
 */
func (a *ReloadAction) ActiveClassName(value string) *ReloadAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *ReloadAction) MergeData(value string) *ReloadAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *ReloadAction) OnClick(value string) *ReloadAction {
    a.Set("onClick", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ReloadAction) HiddenOn(value string) *ReloadAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *ReloadAction) Level(value string) *ReloadAction {
    a.Set("level", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *ReloadAction) Type(value string) *ReloadAction {
    a.Set("type", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *ReloadAction) Size(value string) *ReloadAction {
    a.Set("size", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *ReloadAction) TooltipPlacement(value string) *ReloadAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *ReloadAction) CountDown(value string) *ReloadAction {
    a.Set("countDown", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ReloadAction) VisibleOn(value string) *ReloadAction {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *ReloadAction) StaticSchema(value string) *ReloadAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *ReloadAction) IconClassName(value string) *ReloadAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *ReloadAction) RightIconClassName(value string) *ReloadAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *ReloadAction) HotKey(value string) *ReloadAction {
    a.Set("hotKey", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *ReloadAction) LoadingOn(value string) *ReloadAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ReloadAction) OnEvent(value string) *ReloadAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ReloadAction) EditorSetting(value string) *ReloadAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *ReloadAction) RightIcon(value string) *ReloadAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *ReloadAction) LoadingClassName(value string) *ReloadAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 */
func (a *ReloadAction) Tooltip(value string) *ReloadAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *ReloadAction) CountDownTpl(value string) *ReloadAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *ReloadAction) Block(value string) *ReloadAction {
    a.Set("block", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *ReloadAction) Id(value string) *ReloadAction {
    a.Set("id", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ReloadAction) StaticPlaceholder(value string) *ReloadAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ReloadAction) StaticInputClassName(value string) *ReloadAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *ReloadAction) Style(value string) *ReloadAction {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ReloadAction) UseMobileUI(value string) *ReloadAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *ReloadAction) Testid(value string) *ReloadAction {
    a.Set("testid", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ReloadAction) Disabled(value string) *ReloadAction {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ReloadAction) DisabledOn(value string) *ReloadAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *ReloadAction) Close(value string) *ReloadAction {
    a.Set("close", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *ReloadAction) RequireSelected(value string) *ReloadAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 按钮文字
 */
func (a *ReloadAction) Label(value string) *ReloadAction {
    a.Set("label", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *ReloadAction) ConfirmText(value string) *ReloadAction {
    a.Set("confirmText", value)
    return a
}

/**
 * 是否显示
 */
func (a *ReloadAction) Visible(value string) *ReloadAction {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *ReloadAction) Primary(value string) *ReloadAction {
    a.Set("primary", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *ReloadAction) ActiveLevel(value string) *ReloadAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 子内容
 */
func (a *ReloadAction) Body(value string) *ReloadAction {
    a.Set("body", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ReloadAction) ClassName(value string) *ReloadAction {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ReloadAction) Hidden(value string) *ReloadAction {
    a.Set("hidden", value)
    return a
}

/**
 * 角标
 */
func (a *ReloadAction) Badge(value string) *ReloadAction {
    a.Set("badge", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ReloadAction) StaticLabelClassName(value string) *ReloadAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *ReloadAction) Icon(value string) *ReloadAction {
    a.Set("icon", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ReloadAction) StaticClassName(value string) *ReloadAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *ReloadAction) DisabledTip(value string) *ReloadAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *ReloadAction) Required(value string) *ReloadAction {
    a.Set("required", value)
    return a
}

/**
 * 指定目标组件。
 */
func (a *ReloadAction) Target(value string) *ReloadAction {
    a.Set("target", value)
    return a
}

/**
 * 指定为刷新目标组件。
 */
func (a *ReloadAction) ActionType(value string) *ReloadAction {
    a.Set("actionType", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ReloadAction) Static(value string) *ReloadAction {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ReloadAction) StaticOn(value string) *ReloadAction {
    a.Set("staticOn", value)
    return a
}
