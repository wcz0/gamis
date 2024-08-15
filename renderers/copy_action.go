package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type CopyAction struct {
	*BaseRenderer
}

func NewCopyAction() *CopyAction {
    a := &CopyAction{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("actionType", "copy")
    a.Set("type", "button")
    return a
}

/**
 * 是否隐藏
 */
func (a *CopyAction) Hidden(value interface{}) *CopyAction {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *CopyAction) StaticPlaceholder(value interface{}) *CopyAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CopyAction) StaticLabelClassName(value interface{}) *CopyAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *CopyAction) Testid(value interface{}) *CopyAction {
    a.Set("testid", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *CopyAction) CountDownTpl(value interface{}) *CopyAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 复制啥内容由此配置，支持模板语法。
 */
func (a *CopyAction) Copy(value interface{}) *CopyAction {
    a.Set("copy", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CopyAction) HiddenOn(value interface{}) *CopyAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *CopyAction) DisabledTip(value interface{}) *CopyAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *CopyAction) Close(value interface{}) *CopyAction {
    a.Set("close", value)
    return a
}

/**
 */
func (a *CopyAction) TestIdBuilder(value interface{}) *CopyAction {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 指定为复制内容行为
 */
func (a *CopyAction) ActionType(value interface{}) *CopyAction {
    a.Set("actionType", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CopyAction) VisibleOn(value interface{}) *CopyAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *CopyAction) Id(value interface{}) *CopyAction {
    a.Set("id", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *CopyAction) Type(value interface{}) *CopyAction {
    a.Set("type", value)
    return a
}

/**
 */
func (a *CopyAction) Tooltip(value interface{}) *CopyAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 是否显示
 */
func (a *CopyAction) Visible(value interface{}) *CopyAction {
    a.Set("visible", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *CopyAction) Block(value interface{}) *CopyAction {
    a.Set("block", value)
    return a
}

/**
 * 按钮文字
 */
func (a *CopyAction) Label(value interface{}) *CopyAction {
    a.Set("label", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *CopyAction) Size(value interface{}) *CopyAction {
    a.Set("size", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *CopyAction) CountDown(value interface{}) *CopyAction {
    a.Set("countDown", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CopyAction) Static(value interface{}) *CopyAction {
    a.Set("static", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *CopyAction) MergeData(value interface{}) *CopyAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *CopyAction) Level(value interface{}) *CopyAction {
    a.Set("level", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CopyAction) StaticInputClassName(value interface{}) *CopyAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *CopyAction) IconClassName(value interface{}) *CopyAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *CopyAction) LoadingClassName(value interface{}) *CopyAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 */
func (a *CopyAction) Primary(value interface{}) *CopyAction {
    a.Set("primary", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *CopyAction) Target(value interface{}) *CopyAction {
    a.Set("target", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CopyAction) StaticClassName(value interface{}) *CopyAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *CopyAction) Icon(value interface{}) *CopyAction {
    a.Set("icon", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *CopyAction) ConfirmText(value interface{}) *CopyAction {
    a.Set("confirmText", value)
    return a
}

/**
 */
func (a *CopyAction) StaticSchema(value interface{}) *CopyAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *CopyAction) UseMobileUI(value interface{}) *CopyAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *CopyAction) ActiveClassName(value interface{}) *CopyAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *CopyAction) HotKey(value interface{}) *CopyAction {
    a.Set("hotKey", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *CopyAction) OnClick(value interface{}) *CopyAction {
    a.Set("onClick", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CopyAction) OnEvent(value interface{}) *CopyAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 角标
 */
func (a *CopyAction) Badge(value interface{}) *CopyAction {
    a.Set("badge", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *CopyAction) StaticOn(value interface{}) *CopyAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *CopyAction) TooltipPlacement(value interface{}) *CopyAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *CopyAction) LoadingOn(value interface{}) *CopyAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CopyAction) ClassName(value interface{}) *CopyAction {
    a.Set("className", value)
    return a
}

/**
 * 组件样式
 */
func (a *CopyAction) Style(value interface{}) *CopyAction {
    a.Set("style", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CopyAction) EditorSetting(value interface{}) *CopyAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *CopyAction) ActiveLevel(value interface{}) *CopyAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *CopyAction) RequireSelected(value interface{}) *CopyAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CopyAction) DisabledOn(value interface{}) *CopyAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *CopyAction) Required(value interface{}) *CopyAction {
    a.Set("required", value)
    return a
}

/**
 * 子内容
 */
func (a *CopyAction) Body(value interface{}) *CopyAction {
    a.Set("body", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CopyAction) Disabled(value interface{}) *CopyAction {
    a.Set("disabled", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *CopyAction) RightIcon(value interface{}) *CopyAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *CopyAction) RightIconClassName(value interface{}) *CopyAction {
    a.Set("rightIconClassName", value)
    return a
}
