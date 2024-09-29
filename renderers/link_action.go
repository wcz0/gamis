package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type LinkAction struct {
	*BaseRenderer
}

func NewLinkAction() *LinkAction {
    a := &LinkAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "button")
    a.Set("actionType", "link")
    return a
}


func (a *LinkAction) Set(name string, value interface{}) *LinkAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 激活状态时的样式
 */
func (a *LinkAction) ActiveLevel(value interface{}) *LinkAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *LinkAction) OnClick(value interface{}) *LinkAction {
    a.Set("onClick", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *LinkAction) VisibleOn(value interface{}) *LinkAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *LinkAction) EditorSetting(value interface{}) *LinkAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *LinkAction) Required(value interface{}) *LinkAction {
    a.Set("required", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *LinkAction) TooltipPlacement(value interface{}) *LinkAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *LinkAction) RequireSelected(value interface{}) *LinkAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *LinkAction) OnEvent(value interface{}) *LinkAction {
    a.Set("onEvent", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *LinkAction) LoadingClassName(value interface{}) *LinkAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 */
func (a *LinkAction) Tooltip(value interface{}) *LinkAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *LinkAction) Target(value interface{}) *LinkAction {
    a.Set("target", value)
    return a
}

/**
 */
func (a *LinkAction) Primary(value interface{}) *LinkAction {
    a.Set("primary", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *LinkAction) ActiveClassName(value interface{}) *LinkAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *LinkAction) CountDownTpl(value interface{}) *LinkAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *LinkAction) StaticClassName(value interface{}) *LinkAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *LinkAction) Level(value interface{}) *LinkAction {
    a.Set("level", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *LinkAction) StaticPlaceholder(value interface{}) *LinkAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *LinkAction) ClassName(value interface{}) *LinkAction {
    a.Set("className", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *LinkAction) Type(value interface{}) *LinkAction {
    a.Set("type", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *LinkAction) LoadingOn(value interface{}) *LinkAction {
    a.Set("loadingOn", value)
    return a
}

/**
 */
func (a *LinkAction) StaticSchema(value interface{}) *LinkAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *LinkAction) Block(value interface{}) *LinkAction {
    a.Set("block", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *LinkAction) IconClassName(value interface{}) *LinkAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 指定为打开链接行为，跟 url 不同的时这个行为为单页模式。
 */
func (a *LinkAction) ActionType(value interface{}) *LinkAction {
    a.Set("actionType", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *LinkAction) DisabledOn(value interface{}) *LinkAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *LinkAction) StaticOn(value interface{}) *LinkAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *LinkAction) StaticInputClassName(value interface{}) *LinkAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *LinkAction) ConfirmText(value interface{}) *LinkAction {
    a.Set("confirmText", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *LinkAction) CountDown(value interface{}) *LinkAction {
    a.Set("countDown", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *LinkAction) HotKey(value interface{}) *LinkAction {
    a.Set("hotKey", value)
    return a
}

/**
 * 跳转到哪？支持配置相对路径。
 */
func (a *LinkAction) Link(value interface{}) *LinkAction {
    a.Set("link", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *LinkAction) RightIcon(value interface{}) *LinkAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *LinkAction) RightIconClassName(value interface{}) *LinkAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 角标
 */
func (a *LinkAction) Badge(value interface{}) *LinkAction {
    a.Set("badge", value)
    return a
}

/**
 * 组件样式
 */
func (a *LinkAction) Style(value interface{}) *LinkAction {
    a.Set("style", value)
    return a
}

/**
 */
func (a *LinkAction) TestIdBuilder(value interface{}) *LinkAction {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 按钮文字
 */
func (a *LinkAction) Label(value interface{}) *LinkAction {
    a.Set("label", value)
    return a
}

/**
 * 是否禁用
 */
func (a *LinkAction) Disabled(value interface{}) *LinkAction {
    a.Set("disabled", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *LinkAction) Icon(value interface{}) *LinkAction {
    a.Set("icon", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *LinkAction) MergeData(value interface{}) *LinkAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *LinkAction) StaticLabelClassName(value interface{}) *LinkAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *LinkAction) Hidden(value interface{}) *LinkAction {
    a.Set("hidden", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *LinkAction) Id(value interface{}) *LinkAction {
    a.Set("id", value)
    return a
}

/**
 */
func (a *LinkAction) Testid(value interface{}) *LinkAction {
    a.Set("testid", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *LinkAction) Size(value interface{}) *LinkAction {
    a.Set("size", value)
    return a
}

/**
 * 子内容
 */
func (a *LinkAction) Body(value interface{}) *LinkAction {
    a.Set("body", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *LinkAction) Static(value interface{}) *LinkAction {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *LinkAction) UseMobileUI(value interface{}) *LinkAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *LinkAction) Close(value interface{}) *LinkAction {
    a.Set("close", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *LinkAction) HiddenOn(value interface{}) *LinkAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *LinkAction) Visible(value interface{}) *LinkAction {
    a.Set("visible", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *LinkAction) DisabledTip(value interface{}) *LinkAction {
    a.Set("disabledTip", value)
    return a
}
