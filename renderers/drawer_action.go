package renderers


/**

*/
type DrawerAction struct {
	*BaseRenderer
}

func NewDrawerAction() *DrawerAction {
    a := &DrawerAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "drawer")
    a.Set("type", "button")
    return a
}
/**
 */
func (a *DrawerAction) Primary(value string) *DrawerAction {
    a.Set("primary", value)
    return a
}

/**
 */
func (a *DrawerAction) Tooltip(value string) *DrawerAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 是否显示
 */
func (a *DrawerAction) Visible(value string) *DrawerAction {
    a.Set("visible", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *DrawerAction) Id(value string) *DrawerAction {
    a.Set("id", value)
    return a
}

/**
 */
func (a *DrawerAction) StaticSchema(value string) *DrawerAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *DrawerAction) DisabledTip(value string) *DrawerAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *DrawerAction) RightIconClassName(value string) *DrawerAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *DrawerAction) IconClassName(value string) *DrawerAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *DrawerAction) ConfirmText(value string) *DrawerAction {
    a.Set("confirmText", value)
    return a
}

/**
 */
func (a *DrawerAction) Reload(value string) *DrawerAction {
    a.Set("reload", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *DrawerAction) Size(value string) *DrawerAction {
    a.Set("size", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *DrawerAction) LoadingClassName(value string) *DrawerAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *DrawerAction) ActiveLevel(value string) *DrawerAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *DrawerAction) DisabledOn(value string) *DrawerAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *DrawerAction) StaticClassName(value string) *DrawerAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *DrawerAction) ActiveClassName(value string) *DrawerAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *DrawerAction) Target(value string) *DrawerAction {
    a.Set("target", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *DrawerAction) CountDownTpl(value string) *DrawerAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *DrawerAction) HotKey(value string) *DrawerAction {
    a.Set("hotKey", value)
    return a
}

/**
 * 指定为打开弹窗，抽出式弹窗
 */
func (a *DrawerAction) ActionType(value string) *DrawerAction {
    a.Set("actionType", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *DrawerAction) HiddenOn(value string) *DrawerAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *DrawerAction) StaticLabelClassName(value string) *DrawerAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *DrawerAction) Style(value string) *DrawerAction {
    a.Set("style", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *DrawerAction) StaticPlaceholder(value string) *DrawerAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 按钮文字
 */
func (a *DrawerAction) Label(value string) *DrawerAction {
    a.Set("label", value)
    return a
}

/**
 * 是否有下一个的表达式，正常可以不用配置，如果想要刷掉某些数据可以配置这个。
 */
func (a *DrawerAction) NextCondition(value string) *DrawerAction {
    a.Set("nextCondition", value)
    return a
}

/**
 * 设置对齐方式
 */
func (a *DrawerAction) Align(value string) *DrawerAction {
    a.Set("align", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *DrawerAction) Hidden(value string) *DrawerAction {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *DrawerAction) StaticOn(value string) *DrawerAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *DrawerAction) StaticInputClassName(value string) *DrawerAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *DrawerAction) Icon(value string) *DrawerAction {
    a.Set("icon", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *DrawerAction) TooltipPlacement(value string) *DrawerAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *DrawerAction) Required(value string) *DrawerAction {
    a.Set("required", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *DrawerAction) EditorSetting(value string) *DrawerAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *DrawerAction) RequireSelected(value string) *DrawerAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *DrawerAction) Type(value string) *DrawerAction {
    a.Set("type", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *DrawerAction) Static(value string) *DrawerAction {
    a.Set("static", value)
    return a
}

/**
 * 角标
 */
func (a *DrawerAction) Badge(value string) *DrawerAction {
    a.Set("badge", value)
    return a
}

/**
 * 子内容
 */
func (a *DrawerAction) Body(value string) *DrawerAction {
    a.Set("body", value)
    return a
}

/**
 */
func (a *DrawerAction) Redirect(value string) *DrawerAction {
    a.Set("redirect", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *DrawerAction) UseMobileUI(value string) *DrawerAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *DrawerAction) RightIcon(value string) *DrawerAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *DrawerAction) Close(value string) *DrawerAction {
    a.Set("close", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *DrawerAction) ClassName(value string) *DrawerAction {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *DrawerAction) Disabled(value string) *DrawerAction {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *DrawerAction) VisibleOn(value string) *DrawerAction {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *DrawerAction) Testid(value string) *DrawerAction {
    a.Set("testid", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *DrawerAction) LoadingOn(value string) *DrawerAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *DrawerAction) OnClick(value string) *DrawerAction {
    a.Set("onClick", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *DrawerAction) OnEvent(value string) *DrawerAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *DrawerAction) Level(value string) *DrawerAction {
    a.Set("level", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *DrawerAction) CountDown(value string) *DrawerAction {
    a.Set("countDown", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *DrawerAction) Block(value string) *DrawerAction {
    a.Set("block", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *DrawerAction) MergeData(value string) *DrawerAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 抽出式弹框详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/drawer
 */
func (a *DrawerAction) Drawer(value string) *DrawerAction {
    a.Set("drawer", value)
    return a
}