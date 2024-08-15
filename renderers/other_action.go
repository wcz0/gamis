package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type OtherAction struct {
	*BaseRenderer
}

func NewOtherAction() *OtherAction {
    a := &OtherAction{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("actionType", "prev")
    a.Set("type", "button")
    return a
}


func (a *OtherAction) Set(name string, value interface{}) *OtherAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否禁用
 */
func (a *OtherAction) Disabled(value interface{}) *OtherAction {
    a.Set("disabled", value)
    return a
}

/**
 */
func (a *OtherAction) Staticschema(value interface{}) *OtherAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *OtherAction) Style(value interface{}) *OtherAction {
    a.Set("style", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *OtherAction) Confirmtext(value interface{}) *OtherAction {
    a.Set("confirmText", value)
    return a
}

/**
 */
func (a *OtherAction) Testidbuilder(value interface{}) *OtherAction {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *OtherAction) Block(value interface{}) *OtherAction {
    a.Set("block", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *OtherAction) Iconclassname(value interface{}) *OtherAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *OtherAction) Loadingclassname(value interface{}) *OtherAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *OtherAction) Hiddenon(value interface{}) *OtherAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *OtherAction) Visibleon(value interface{}) *OtherAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *OtherAction) Icon(value interface{}) *OtherAction {
    a.Set("icon", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *OtherAction) Righticonclassname(value interface{}) *OtherAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *OtherAction) Countdowntpl(value interface{}) *OtherAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *OtherAction) Id(value interface{}) *OtherAction {
    a.Set("id", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *OtherAction) Mergedata(value interface{}) *OtherAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *OtherAction) Usemobileui(value interface{}) *OtherAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *OtherAction) Tooltipplacement(value interface{}) *OtherAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *OtherAction) Static(value interface{}) *OtherAction {
    a.Set("static", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *OtherAction) Required(value interface{}) *OtherAction {
    a.Set("required", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *OtherAction) Close(value interface{}) *OtherAction {
    a.Set("close", value)
    return a
}

/**
 * 按钮文字
 */
func (a *OtherAction) Label(value interface{}) *OtherAction {
    a.Set("label", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *OtherAction) Type(value interface{}) *OtherAction {
    a.Set("type", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *OtherAction) Level(value interface{}) *OtherAction {
    a.Set("level", value)
    return a
}

/**
 * 可选值: prev | next | cancel | close | submit | confirm | add | reset | reset-and-submit
 */
func (a *OtherAction) Actiontype(value interface{}) *OtherAction {
    a.Set("actionType", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *OtherAction) Classname(value interface{}) *OtherAction {
    a.Set("className", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *OtherAction) Activeclassname(value interface{}) *OtherAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *OtherAction) Onevent(value interface{}) *OtherAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *OtherAction) Disabledtip(value interface{}) *OtherAction {
    a.Set("disabledTip", value)
    return a
}

/**
 * 角标
 */
func (a *OtherAction) Badge(value interface{}) *OtherAction {
    a.Set("badge", value)
    return a
}

/**
 */
func (a *OtherAction) Tooltip(value interface{}) *OtherAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *OtherAction) Target(value interface{}) *OtherAction {
    a.Set("target", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *OtherAction) Onclick(value interface{}) *OtherAction {
    a.Set("onClick", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *OtherAction) Staticlabelclassname(value interface{}) *OtherAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *OtherAction) Righticon(value interface{}) *OtherAction {
    a.Set("rightIcon", value)
    return a
}

/**
 */
func (a *OtherAction) Primary(value interface{}) *OtherAction {
    a.Set("primary", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *OtherAction) Requireselected(value interface{}) *OtherAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *OtherAction) Countdown(value interface{}) *OtherAction {
    a.Set("countDown", value)
    return a
}

/**
 * 是否显示
 */
func (a *OtherAction) Visible(value interface{}) *OtherAction {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *OtherAction) Staticplaceholder(value interface{}) *OtherAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *OtherAction) Editorsetting(value interface{}) *OtherAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *OtherAction) Activelevel(value interface{}) *OtherAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *OtherAction) Hotkey(value interface{}) *OtherAction {
    a.Set("hotKey", value)
    return a
}

/**
 * 子内容
 */
func (a *OtherAction) Body(value interface{}) *OtherAction {
    a.Set("body", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *OtherAction) Staticclassname(value interface{}) *OtherAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *OtherAction) Size(value interface{}) *OtherAction {
    a.Set("size", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *OtherAction) Hidden(value interface{}) *OtherAction {
    a.Set("hidden", value)
    return a
}

/**
 */
func (a *OtherAction) Testid(value interface{}) *OtherAction {
    a.Set("testid", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *OtherAction) Staticon(value interface{}) *OtherAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *OtherAction) Staticinputclassname(value interface{}) *OtherAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *OtherAction) Loadingon(value interface{}) *OtherAction {
    a.Set("loadingOn", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *OtherAction) Disabledon(value interface{}) *OtherAction {
    a.Set("disabledOn", value)
    return a
}
