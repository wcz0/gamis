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

    a.Set("actionType", "dialog")
    a.Set("type", "button")
    return a
}


func (a *DialogAction) Set(name string, value interface{}) *DialogAction {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否禁用表达式
 */
func (a *DialogAction) Disabledon(value interface{}) *DialogAction {
    a.Set("disabledOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *DialogAction) Onevent(value interface{}) *DialogAction {
    a.Set("onEvent", value)
    return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *DialogAction) Target(value interface{}) *DialogAction {
    a.Set("target", value)
    return a
}

/**
 * 是否有下一个的表达式，正常可以不用配置，如果想要刷掉某些数据可以配置这个。
 */
func (a *DialogAction) Nextcondition(value interface{}) *DialogAction {
    a.Set("nextCondition", value)
    return a
}

/**
 * 是否禁用
 */
func (a *DialogAction) Disabled(value interface{}) *DialogAction {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *DialogAction) Staticinputclassname(value interface{}) *DialogAction {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 激活状态时的样式
 */
func (a *DialogAction) Activelevel(value interface{}) *DialogAction {
    a.Set("activeLevel", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *DialogAction) Hiddenon(value interface{}) *DialogAction {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *DialogAction) Staticlabelclassname(value interface{}) *DialogAction {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *DialogAction) Testidbuilder(value interface{}) *DialogAction {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *DialogAction) Righticon(value interface{}) *DialogAction {
    a.Set("rightIcon", value)
    return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *DialogAction) Countdown(value interface{}) *DialogAction {
    a.Set("countDown", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *DialogAction) Staticon(value interface{}) *DialogAction {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *DialogAction) Staticclassname(value interface{}) *DialogAction {
    a.Set("staticClassName", value)
    return a
}

/**
 * 禁用时的文案提示。
 */
func (a *DialogAction) Disabledtip(value interface{}) *DialogAction {
    a.Set("disabledTip", value)
    return a
}

/**
 */
func (a *DialogAction) Tooltip(value interface{}) *DialogAction {
    a.Set("tooltip", value)
    return a
}

/**
 * 键盘快捷键
 */
func (a *DialogAction) Hotkey(value interface{}) *DialogAction {
    a.Set("hotKey", value)
    return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *DialogAction) Icon(value interface{}) *DialogAction {
    a.Set("icon", value)
    return a
}

/**
 * loading 上的css 类名
 */
func (a *DialogAction) Loadingclassname(value interface{}) *DialogAction {
    a.Set("loadingClassName", value)
    return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *DialogAction) Righticonclassname(value interface{}) *DialogAction {
    a.Set("rightIconClassName", value)
    return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *DialogAction) Requireselected(value interface{}) *DialogAction {
    a.Set("requireSelected", value)
    return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *DialogAction) Mergedata(value interface{}) *DialogAction {
    a.Set("mergeData", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *DialogAction) Visibleon(value interface{}) *DialogAction {
    a.Set("visibleOn", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *DialogAction) Classname(value interface{}) *DialogAction {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *DialogAction) Visible(value interface{}) *DialogAction {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *DialogAction) Static(value interface{}) *DialogAction {
    a.Set("static", value)
    return a
}

/**
 */
func (a *DialogAction) Primary(value interface{}) *DialogAction {
    a.Set("primary", value)
    return a
}

/**
 * 按钮大小
 * 可选值: xs | sm | md | lg
 */
func (a *DialogAction) Size(value interface{}) *DialogAction {
    a.Set("size", value)
    return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *DialogAction) Confirmtext(value interface{}) *DialogAction {
    a.Set("confirmText", value)
    return a
}

/**
 * 自定义事件处理函数
 */
func (a *DialogAction) Onclick(value interface{}) *DialogAction {
    a.Set("onClick", value)
    return a
}

/**
 */
func (a *DialogAction) Reload(value interface{}) *DialogAction {
    a.Set("reload", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *DialogAction) Staticplaceholder(value interface{}) *DialogAction {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *DialogAction) Block(value interface{}) *DialogAction {
    a.Set("block", value)
    return a
}

/**
 * 按钮样式
 * 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *DialogAction) Level(value interface{}) *DialogAction {
    a.Set("level", value)
    return a
}

/**
 * 激活状态时的类名
 */
func (a *DialogAction) Activeclassname(value interface{}) *DialogAction {
    a.Set("activeClassName", value)
    return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *DialogAction) Close(value interface{}) *DialogAction {
    a.Set("close", value)
    return a
}

/**
 * icon 上的css 类名
 */
func (a *DialogAction) Iconclassname(value interface{}) *DialogAction {
    a.Set("iconClassName", value)
    return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *DialogAction) Required(value interface{}) *DialogAction {
    a.Set("required", value)
    return a
}

/**
 * 倒计时文字自定义
 */
func (a *DialogAction) Countdowntpl(value interface{}) *DialogAction {
    a.Set("countDownTpl", value)
    return a
}

/**
 * 子内容
 */
func (a *DialogAction) Body(value interface{}) *DialogAction {
    a.Set("body", value)
    return a
}

/**
 * 组件样式
 */
func (a *DialogAction) Style(value interface{}) *DialogAction {
    a.Set("style", value)
    return a
}

/**
 * 按钮文字
 */
func (a *DialogAction) Label(value interface{}) *DialogAction {
    a.Set("label", value)
    return a
}

/**
 * 弹框详情 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/dialog
 */
func (a *DialogAction) Dialog(value interface{}) *DialogAction {
    a.Set("dialog", value)
    return a
}

/**
 */
func (a *DialogAction) Redirect(value interface{}) *DialogAction {
    a.Set("redirect", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *DialogAction) Hidden(value interface{}) *DialogAction {
    a.Set("hidden", value)
    return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *DialogAction) Id(value interface{}) *DialogAction {
    a.Set("id", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *DialogAction) Usemobileui(value interface{}) *DialogAction {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 指定为打开弹窗
 */
func (a *DialogAction) Actiontype(value interface{}) *DialogAction {
    a.Set("actionType", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *DialogAction) Editorsetting(value interface{}) *DialogAction {
    a.Set("editorSetting", value)
    return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。
 * 可选值: button | submit | reset
 */
func (a *DialogAction) Type(value interface{}) *DialogAction {
    a.Set("type", value)
    return a
}

/**
 */
func (a *DialogAction) Testid(value interface{}) *DialogAction {
    a.Set("testid", value)
    return a
}

/**
 * 可选值: top | right | bottom | left
 */
func (a *DialogAction) Tooltipplacement(value interface{}) *DialogAction {
    a.Set("tooltipPlacement", value)
    return a
}

/**
 * 数据映射
 */
func (a *DialogAction) Data(value interface{}) *DialogAction {
    a.Set("data", value)
    return a
}

/**
 */
func (a *DialogAction) Staticschema(value interface{}) *DialogAction {
    a.Set("staticSchema", value)
    return a
}

/**
 * 角标
 */
func (a *DialogAction) Badge(value interface{}) *DialogAction {
    a.Set("badge", value)
    return a
}

/**
 * 是否显示loading效果
 */
func (a *DialogAction) Loadingon(value interface{}) *DialogAction {
    a.Set("loadingOn", value)
    return a
}
