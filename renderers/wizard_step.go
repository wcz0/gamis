package renderers


/**

*/
type WizardStep struct {
	*BaseRenderer
}

func NewWizardStep() *WizardStep {
    a := &WizardStep{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 */
func (a *WizardStep) Redirect(value string) *WizardStep {
    a.Set("redirect", value)
    return a
}

/**
 * 设置后将轮询调用 initApi
 */
func (a *WizardStep) Interval(value string) *WizardStep {
    a.Set("interval", value)
    return a
}

/**
 */
func (a *WizardStep) Name(value string) *WizardStep {
    a.Set("name", value)
    return a
}

/**
 * 表单label的对齐方式
 */
func (a *WizardStep) LabelAlign(value string) *WizardStep {
    a.Set("labelAlign", value)
    return a
}

/**
 * 组件样式
 */
func (a *WizardStep) Style(value string) *WizardStep {
    a.Set("style", value)
    return a
}

/**
 */
func (a *WizardStep) Data(value string) *WizardStep {
    a.Set("data", value)
    return a
}

/**
 * 如果决定结束的字段名不是 `finished` 请设置此属性，比如 `is_success`
 */
func (a *WizardStep) FinishedField(value string) *WizardStep {
    a.Set("finishedField", value)
    return a
}

/**
 * Form 用来获取初始数据的 api,与initApi不同的是，会一直轮询请求该接口，直到返回 finished 属性为 true 才 结束。
 */
func (a *WizardStep) InitAsyncApi(value string) *WizardStep {
    a.Set("initAsyncApi", value)
    return a
}

/**
 * 配置容器 panel className
 */
func (a *WizardStep) PanelClassName(value string) *WizardStep {
    a.Set("panelClassName", value)
    return a
}

/**
 * 页面离开提示，为了防止页面不小心跳转而导致表单没有保存。
 */
func (a *WizardStep) PromptPageLeave(value string) *WizardStep {
    a.Set("promptPageLeave", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *WizardStep) LabelWidth(value string) *WizardStep {
    a.Set("labelWidth", value)
    return a
}

/**
 * 用来初始化表单数据
 */
func (a *WizardStep) InitApi(value string) *WizardStep {
    a.Set("initApi", value)
    return a
}

/**
 * 通过 JS 表达式来配置当前步骤可否被直接跳转到。
 */
func (a *WizardStep) JumpableOn(value string) *WizardStep {
    a.Set("jumpableOn", value)
    return a
}

/**
 * 是否开启调试，开启后会在顶部实时显示表单项数据。
 */
func (a *WizardStep) Debug(value string) *WizardStep {
    a.Set("debug", value)
    return a
}

/**
 * Debug面板配置
 */
func (a *WizardStep) DebugConfig(value string) *WizardStep {
    a.Set("debugConfig", value)
    return a
}

/**
 */
func (a *WizardStep) Reload(value string) *WizardStep {
    a.Set("reload", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *WizardStep) DisabledOn(value string) *WizardStep {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *WizardStep) StaticOn(value string) *WizardStep {
    a.Set("staticOn", value)
    return a
}

/**
 * 配置表单项默认的展示方式。
 * 可选值: normal | inline | horizontal
 */
func (a *WizardStep) Mode(value string) *WizardStep {
    a.Set("mode", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *WizardStep) StaticLabelClassName(value string) *WizardStep {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否初始加载
 */
func (a *WizardStep) InitFetch(value string) *WizardStep {
    a.Set("initFetch", value)
    return a
}

/**
 * 开启本地缓存后限制保存哪些 key
 */
func (a *WizardStep) PersistDataKeys(value string) *WizardStep {
    a.Set("persistDataKeys", value)
    return a
}

/**
 * 具体的提示信息，选填。
 */
func (a *WizardStep) PromptPageLeaveMessage(value string) *WizardStep {
    a.Set("promptPageLeaveMessage", value)
    return a
}

/**
 * 提交完后重置表单
 */
func (a *WizardStep) ResetAfterSubmit(value string) *WizardStep {
    a.Set("resetAfterSubmit", value)
    return a
}

/**
 * 修改的时候是否直接提交表单。
 */
func (a *WizardStep) SubmitOnChange(value string) *WizardStep {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 提交后清空表单
 */
func (a *WizardStep) ClearAfterSubmit(value string) *WizardStep {
    a.Set("clearAfterSubmit", value)
    return a
}

/**
 * 是否自动将第一个表单元素聚焦。
 */
func (a *WizardStep) AutoFocus(value string) *WizardStep {
    a.Set("autoFocus", value)
    return a
}

/**
 * 消息文案配置，记住这个优先级是最低的，如果你的接口返回了 msg，接口返回的优先。
 */
func (a *WizardStep) Messages(value string) *WizardStep {
    a.Set("messages", value)
    return a
}

/**
 * 是否可直接跳转到该步骤，一般编辑模式需要可直接跳转查看。
 */
func (a *WizardStep) Jumpable(value string) *WizardStep {
    a.Set("jumpable", value)
    return a
}

/**
 */
func (a *WizardStep) StaticSchema(value string) *WizardStep {
    a.Set("staticSchema", value)
    return a
}

/**
 * 描述
 */
func (a *WizardStep) Description(value string) *WizardStep {
    a.Set("description", value)
    return a
}

/**
 * 禁用回车提交
 */
func (a *WizardStep) PreventEnterSubmit(value string) *WizardStep {
    a.Set("preventEnterSubmit", value)
    return a
}

/**
 * 设置此属性后，表单提交发送保存接口后，还会继续轮询请求该接口，直到返回 finished 属性为 true 才 结束。
 */
func (a *WizardStep) AsyncApi(value string) *WizardStep {
    a.Set("asyncApi", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *WizardStep) OnEvent(value string) *WizardStep {
    a.Set("onEvent", value)
    return a
}

/**
 * 图标
 */
func (a *WizardStep) Icon(value string) *WizardStep {
    a.Set("icon", value)
    return a
}

/**
 * 子标题
 */
func (a *WizardStep) SubTitle(value string) *WizardStep {
    a.Set("subTitle", value)
    return a
}

/**
 */
func (a *WizardStep) FieldSet(value string) *WizardStep {
    a.Set("fieldSet", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *WizardStep) Hidden(value string) *WizardStep {
    a.Set("hidden", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *WizardStep) Id(value string) *WizardStep {
    a.Set("id", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *WizardStep) EditorSetting(value string) *WizardStep {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *WizardStep) Value(value string) *WizardStep {
    a.Set("value", value)
    return a
}

/**
 * 提交成功后清空本地缓存
 */
func (a *WizardStep) ClearPersistDataAfterSubmit(value string) *WizardStep {
    a.Set("clearPersistDataAfterSubmit", value)
    return a
}

/**
 * 是否固定底下的按钮在底部。
 */
func (a *WizardStep) AffixFooter(value string) *WizardStep {
    a.Set("affixFooter", value)
    return a
}

/**
 */
func (a *WizardStep) Tabs(value string) *WizardStep {
    a.Set("tabs", value)
    return a
}

/**
 * Form 用来保存数据的 api。详情：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/index#%E8%A1%A8%E5%8D%95%E6%8F%90%E4%BA%A4
 */
func (a *WizardStep) Api(value string) *WizardStep {
    a.Set("api", value)
    return a
}

/**
 * 按钮集合，会固定在底部显示。
 */
func (a *WizardStep) Actions(value string) *WizardStep {
    a.Set("actions", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *WizardStep) HiddenOn(value string) *WizardStep {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 设置主键 id, 当设置后，检测表单是否完成时（asyncApi），只会携带此数据。
 */
func (a *WizardStep) PrimaryField(value string) *WizardStep {
    a.Set("primaryField", value)
    return a
}

/**
 * 默认的提交按钮名称，如果设置成空，则可以把默认按钮去掉。
 */
func (a *WizardStep) SubmitText(value string) *WizardStep {
    a.Set("submitText", value)
    return a
}

/**
 */
func (a *WizardStep) Label(value string) *WizardStep {
    a.Set("label", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *WizardStep) VisibleOn(value string) *WizardStep {
    a.Set("visibleOn", value)
    return a
}

/**
 * 设置了initAsyncApi后，默认会从返回数据的data.finished来判断是否完成，也可以设置成其他的xxx，就会从data.xxx中获取
 */
func (a *WizardStep) InitFinishedField(value string) *WizardStep {
    a.Set("initFinishedField", value)
    return a
}

/**
 * 配置停止轮询的条件
 */
func (a *WizardStep) StopAutoRefreshWhen(value string) *WizardStep {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 表单项显示为几列
 */
func (a *WizardStep) ColumnCount(value string) *WizardStep {
    a.Set("columnCount", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *WizardStep) StaticPlaceholder(value string) *WizardStep {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *WizardStep) StaticClassName(value string) *WizardStep {
    a.Set("staticClassName", value)
    return a
}

/**
 * 建议改成 api 的 sendOn 属性。
 */
func (a *WizardStep) InitFetchOn(value string) *WizardStep {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 表单初始先提交一次，联动的时候有用
 */
func (a *WizardStep) SubmitOnInit(value string) *WizardStep {
    a.Set("submitOnInit", value)
    return a
}

/**
 * 是否用 panel 包裹起来
 */
func (a *WizardStep) WrapWithPanel(value string) *WizardStep {
    a.Set("wrapWithPanel", value)
    return a
}

/**
 * 是否禁用
 */
func (a *WizardStep) Disabled(value string) *WizardStep {
    a.Set("disabled", value)
    return a
}

/**
 * 表单项集合
 */
func (a *WizardStep) Body(value string) *WizardStep {
    a.Set("body", value)
    return a
}

/**
 * 是否开启本地缓存
 */
func (a *WizardStep) PersistData(value string) *WizardStep {
    a.Set("persistData", value)
    return a
}

/**
 * 设置了initAsyncApi以后，默认拉取的时间间隔
 */
func (a *WizardStep) InitCheckInterval(value string) *WizardStep {
    a.Set("initCheckInterval", value)
    return a
}

/**
 * 组合校验规则，选填
 */
func (a *WizardStep) Rules(value string) *WizardStep {
    a.Set("rules", value)
    return a
}

/**
 * 表单标题
 */
func (a *WizardStep) Title(value string) *WizardStep {
    a.Set("title", value)
    return a
}

/**
 * 默认表单提交自己会通过发送 api 保存数据，但是也可以设定另外一个 form 的 name 值，或者另外一个 `CRUD` 模型的 name 值。 如果 target 目标是一个 `Form` ，则目标 `Form` 会重新触发 `initApi` 和 `schemaApi`，api 可以拿到当前 form 数据。如果目标是一个 `CRUD` 模型，则目标模型会重新触发搜索，参数为当前 Form 数据。
 */
func (a *WizardStep) Target(value string) *WizardStep {
    a.Set("target", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *WizardStep) ClassName(value string) *WizardStep {
    a.Set("className", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *WizardStep) UseMobileUI(value string) *WizardStep {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否静默拉取
 */
func (a *WizardStep) SilentPolling(value string) *WizardStep {
    a.Set("silentPolling", value)
    return a
}

/**
 * Form 也可以配置 feedback。
 */
func (a *WizardStep) Feedback(value string) *WizardStep {
    a.Set("feedback", value)
    return a
}

/**
 * 轮询请求的时间间隔，默认为 3秒。设置 asyncApi 才有效
 */
func (a *WizardStep) CheckInterval(value string) *WizardStep {
    a.Set("checkInterval", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *WizardStep) Horizontal(value string) *WizardStep {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否显示
 */
func (a *WizardStep) Visible(value string) *WizardStep {
    a.Set("visible", value)
    return a
}

/**
 * 展示态时的className
 */
func (a *WizardStep) Static(value string) *WizardStep {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *WizardStep) StaticInputClassName(value string) *WizardStep {
    a.Set("staticInputClassName", value)
    return a
}
