package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type VanillaAction struct {
	*BaseRenderer
}

func NewVanillaAction() *VanillaAction {
    a := &VanillaAction{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "button")
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *VanillaAction) StaticOn(value interface{}) *VanillaAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *VanillaAction) StaticLabelClassName(value interface{}) *VanillaAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *VanillaAction) Block(value interface{}) *VanillaAction {
    a.Set("block", value)
    return a
}

/**
 * 角标
 */
func (a *VanillaAction) Badge(value interface{}) *VanillaAction {
    a.Set("badge", value)
    return a
}

/**
 */
func (a *VanillaAction) DownloadFileName(value interface{}) *VanillaAction {
    a.Set("downloadFileName", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *VanillaAction) StaticPlaceholder(value interface{}) *VanillaAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *VanillaAction) StaticSchema(value interface{}) *VanillaAction {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *VanillaAction) ActionType(value interface{}) *VanillaAction {
    a.Set("actionType", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *VanillaAction) EditorSetting(value interface{}) *VanillaAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *VanillaAction) Size(value interface{}) *VanillaAction {
    a.Set("size", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *VanillaAction) Required(value interface{}) *VanillaAction {
    a.Set("required", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *VanillaAction) Id(value interface{}) *VanillaAction {
    a.Set("id", value)
    return a
}

/**
 */
func (a *VanillaAction) Testid(value interface{}) *VanillaAction {
    a.Set("testid", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *VanillaAction) Target(value interface{}) *VanillaAction {
    a.Set("target", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *VanillaAction) OnClick(value interface{}) *VanillaAction {
    a.Set("onClick", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *VanillaAction) ActiveClassName(value interface{}) *VanillaAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *VanillaAction) StaticClassName(value interface{}) *VanillaAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *VanillaAction) ConfirmText(value interface{}) *VanillaAction {
    a.Set("confirmText", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *VanillaAction) HiddenOn(value interface{}) *VanillaAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *VanillaAction) Style(value interface{}) *VanillaAction {
    a.Set("style", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *VanillaAction) RightIconClassName(value interface{}) *VanillaAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *VanillaAction) TooltipPlacement(value interface{}) *VanillaAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *VanillaAction) RequireSelected(value interface{}) *VanillaAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *VanillaAction) DisabledTip(value interface{}) *VanillaAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *VanillaAction) Icon(value interface{}) *VanillaAction {
    a.Set("icon", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *VanillaAction) IconClassName(value interface{}) *VanillaAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *VanillaAction) HotKey(value interface{}) *VanillaAction {
    a.Set("hotKey", value)
    return a
}

/**
 * 按钮文字
 */
func (a *VanillaAction) Label(value interface{}) *VanillaAction {
    a.Set("label", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *VanillaAction) Hidden(value interface{}) *VanillaAction {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *VanillaAction) Static(value interface{}) *VanillaAction {
    a.Set("static", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *VanillaAction) Type(value interface{}) *VanillaAction {
    a.Set("type", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *VanillaAction) LoadingClassName(value interface{}) *VanillaAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 */
func (a *VanillaAction) Primary(value interface{}) *VanillaAction {
    a.Set("primary", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *VanillaAction) CountDown(value interface{}) *VanillaAction {
    a.Set("countDown", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *VanillaAction) CountDownTpl(value interface{}) *VanillaAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *VanillaAction) OnEvent(value interface{}) *VanillaAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *VanillaAction) UseMobileUI(value interface{}) *VanillaAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *VanillaAction) MergeData(value interface{}) *VanillaAction {
    a.Set("mergeData", value)
    return a
}

/**
 */
func (a *VanillaAction) TestIdBuilder(value interface{}) *VanillaAction {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *VanillaAction) DisabledOn(value interface{}) *VanillaAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *VanillaAction) RightIcon(value interface{}) *VanillaAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *VanillaAction) Level(value interface{}) *VanillaAction {
    a.Set("level", value)
    return a
}

/**
 * 子内容
 */
func (a *VanillaAction) Body(value interface{}) *VanillaAction {
    a.Set("body", value)
    return a
}

/**
 * 是否禁用
 */
func (a *VanillaAction) Disabled(value interface{}) *VanillaAction {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *VanillaAction) StaticInputClassName(value interface{}) *VanillaAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *VanillaAction) Tooltip(value interface{}) *VanillaAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *VanillaAction) ActiveLevel(value interface{}) *VanillaAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *VanillaAction) Close(value interface{}) *VanillaAction {
    a.Set("close", value)
    return a
}

/**
 * 是否显示
 */
func (a *VanillaAction) Visible(value interface{}) *VanillaAction {
    a.Set("visible", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *VanillaAction) ClassName(value interface{}) *VanillaAction {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *VanillaAction) VisibleOn(value interface{}) *VanillaAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *VanillaAction) LoadingOn(value interface{}) *VanillaAction {
    a.Set("loadingOn", value)
    return a
}
