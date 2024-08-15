package renderers


/**
 * Form 表单渲染器。说明：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/index

 * @author wcz0
 * @version 6.2.2
 */
type Form struct {
	*BaseRenderer
}

func NewForm() *Form {
    a := &Form{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "form")
    return a
}


func (a *Form) Set(name string, value interface{}) *Form {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 容器 css 类名
 */
func (a *Form) Classname(value interface{}) *Form {
    a.Set("className", value)
    return a
}

/**
 */
func (a *Form) Testid(value interface{}) *Form {
    a.Set("testid", value)
    return a
}

/**
 * 表单标题
 */
func (a *Form) Title(value interface{}) *Form {
    a.Set("title", value)
    return a
}

/**
 * 按钮集合，会固定在底部显示。
 */
func (a *Form) Actions(value interface{}) *Form {
    a.Set("actions", value)
    return a
}

/**
 */
func (a *Form) Data(value interface{}) *Form {
    a.Set("data", value)
    return a
}

/**
 * 是否静默拉取
 */
func (a *Form) Silentpolling(value interface{}) *Form {
    a.Set("silentPolling", value)
    return a
}

/**
 * 提交后清空表单
 */
func (a *Form) Clearaftersubmit(value interface{}) *Form {
    a.Set("clearAfterSubmit", value)
    return a
}

/**
 */
func (a *Form) Staticschema(value interface{}) *Form {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否开启调试，开启后会在顶部实时显示表单项数据。
 */
func (a *Form) Debug(value interface{}) *Form {
    a.Set("debug", value)
    return a
}

/**
 * 是否开启本地缓存
 */
func (a *Form) Persistdata(value interface{}) *Form {
    a.Set("persistData", value)
    return a
}

/**
 * 开启本地缓存后限制保存哪些 key
 */
func (a *Form) Persistdatakeys(value interface{}) *Form {
    a.Set("persistDataKeys", value)
    return a
}

/**
 * 提交完后重置表单
 */
func (a *Form) Resetaftersubmit(value interface{}) *Form {
    a.Set("resetAfterSubmit", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Form) Id(value interface{}) *Form {
    a.Set("id", value)
    return a
}

/**
 */
func (a *Form) Staticclassname(value interface{}) *Form {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Form) Testidbuilder(value interface{}) *Form {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 用来初始化表单数据
 */
func (a *Form) Initapi(value interface{}) *Form {
    a.Set("initApi", value)
    return a
}

/**
 * 表单初始先提交一次，联动的时候有用
 */
func (a *Form) Submitoninit(value interface{}) *Form {
    a.Set("submitOnInit", value)
    return a
}

/**
 * 组件样式
 */
func (a *Form) Style(value interface{}) *Form {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Form) Tabs(value interface{}) *Form {
    a.Set("tabs", value)
    return a
}

/**
 * 设置此属性后，表单提交发送保存接口后，还会继续轮询请求该接口，直到返回 finished 属性为 true 才 结束。
 */
func (a *Form) Asyncapi(value interface{}) *Form {
    a.Set("asyncApi", value)
    return a
}

/**
 * 是否用 panel 包裹起来
 */
func (a *Form) Wrapwithpanel(value interface{}) *Form {
    a.Set("wrapWithPanel", value)
    return a
}

/**
 * Form 用来获取初始数据的 api,与initApi不同的是，会一直轮询请求该接口，直到返回 finished 属性为 true 才 结束。
 */
func (a *Form) Initasyncapi(value interface{}) *Form {
    a.Set("initAsyncApi", value)
    return a
}

/**
 * 修改的时候是否直接提交表单。
 */
func (a *Form) Submitonchange(value interface{}) *Form {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 页面离开提示，为了防止页面不小心跳转而导致表单没有保存。
 */
func (a *Form) Promptpageleave(value interface{}) *Form {
    a.Set("promptPageLeave", value)
    return a
}

/**
 * 组合校验规则，选填
 */
func (a *Form) Rules(value interface{}) *Form {
    a.Set("rules", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Form) Disabled(value interface{}) *Form {
    a.Set("disabled", value)
    return a
}

/**
 * Form 用来保存数据的 api。详情：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/index#%E8%A1%A8%E5%8D%95%E6%8F%90%E4%BA%A4
 */
func (a *Form) Api(value interface{}) *Form {
    a.Set("api", value)
    return a
}

/**
 * 默认的提交按钮名称，如果设置成空，则可以把默认按钮去掉。
 */
func (a *Form) Submittext(value interface{}) *Form {
    a.Set("submitText", value)
    return a
}

/**
 * 是否固定底下的按钮在底部。
 */
func (a *Form) Affixfooter(value interface{}) *Form {
    a.Set("affixFooter", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Form) Staticlabelclassname(value interface{}) *Form {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *Form) Horizontal(value interface{}) *Form {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Form) Hiddenon(value interface{}) *Form {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Form) Onevent(value interface{}) *Form {
    a.Set("onEvent", value)
    return a
}

/**
 * 展示态时的className
 */
func (a *Form) Static(value interface{}) *Form {
    a.Set("static", value)
    return a
}

/**
 */
func (a *Form) Staticon(value interface{}) *Form {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Form) Name(value interface{}) *Form {
    a.Set("name", value)
    return a
}

/**
 * 配置容器 panel className
 */
func (a *Form) Panelclassname(value interface{}) *Form {
    a.Set("panelClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Form) Hidden(value interface{}) *Form {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Form) Visible(value interface{}) *Form {
    a.Set("visible", value)
    return a
}

/**
 * Form 也可以配置 feedback。
 */
func (a *Form) Feedback(value interface{}) *Form {
    a.Set("feedback", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Form) Staticplaceholder(value interface{}) *Form {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *Form) Fieldset(value interface{}) *Form {
    a.Set("fieldSet", value)
    return a
}

/**
 * Debug面板配置
 */
func (a *Form) Debugconfig(value interface{}) *Form {
    a.Set("debugConfig", value)
    return a
}

/**
 * 是否初始加载
 */
func (a *Form) Initfetch(value interface{}) *Form {
    a.Set("initFetch", value)
    return a
}

/**
 * 建议改成 api 的 sendOn 属性。
 */
func (a *Form) Initfetchon(value interface{}) *Form {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 设置后将轮询调用 initApi
 */
func (a *Form) Interval(value interface{}) *Form {
    a.Set("interval", value)
    return a
}

/**
 * 配置表单项默认的展示方式。
 * 可选值: normal | inline | horizontal | flex
 */
func (a *Form) Mode(value interface{}) *Form {
    a.Set("mode", value)
    return a
}

/**
 * 默认表单提交自己会通过发送 api 保存数据，但是也可以设定另外一个 form 的 name 值，或者另外一个 `CRUD` 模型的 name 值。 如果 target 目标是一个 `Form` ，则目标 `Form` 会重新触发 `initApi` 和 `schemaApi`，api 可以拿到当前 form 数据。如果目标是一个 `CRUD` 模型，则目标模型会重新触发搜索，参数为当前 Form 数据。
 */
func (a *Form) Target(value interface{}) *Form {
    a.Set("target", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Form) Staticinputclassname(value interface{}) *Form {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Form) Editorsetting(value interface{}) *Form {
    a.Set("editorSetting", value)
    return a
}

/**
 * 设置了initAsyncApi后，默认会从返回数据的data.finished来判断是否完成，也可以设置成其他的xxx，就会从data.xxx中获取
 */
func (a *Form) Initfinishedfield(value interface{}) *Form {
    a.Set("initFinishedField", value)
    return a
}

/**
 * 设置了initAsyncApi以后，默认拉取的时间间隔
 */
func (a *Form) Initcheckinterval(value interface{}) *Form {
    a.Set("initCheckInterval", value)
    return a
}

/**
 * 提交成功后清空本地缓存
 */
func (a *Form) Clearpersistdataaftersubmit(value interface{}) *Form {
    a.Set("clearPersistDataAfterSubmit", value)
    return a
}

/**
 * 轮询请求的时间间隔，默认为 3秒。设置 asyncApi 才有效
 */
func (a *Form) Checkinterval(value interface{}) *Form {
    a.Set("checkInterval", value)
    return a
}

/**
 */
func (a *Form) Redirect(value interface{}) *Form {
    a.Set("redirect", value)
    return a
}

/**
 * 具体的提示信息，选填。
 */
func (a *Form) Promptpageleavemessage(value interface{}) *Form {
    a.Set("promptPageLeaveMessage", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Form) Disabledon(value interface{}) *Form {
    a.Set("disabledOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Form) Usemobileui(value interface{}) *Form {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Form) Reload(value interface{}) *Form {
    a.Set("reload", value)
    return a
}

/**
 * 表单项集合
 */
func (a *Form) Body(value interface{}) *Form {
    a.Set("body", value)
    return a
}

/**
 * 配置停止轮询的条件
 */
func (a *Form) Stopautorefreshwhen(value interface{}) *Form {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 如果决定结束的字段名不是 `finished` 请设置此属性，比如 `is_success`
 */
func (a *Form) Finishedfield(value interface{}) *Form {
    a.Set("finishedField", value)
    return a
}

/**
 * 是否自动将第一个表单元素聚焦。
 */
func (a *Form) Autofocus(value interface{}) *Form {
    a.Set("autoFocus", value)
    return a
}

/**
 * 设置主键 id, 当设置后，检测表单是否完成时（asyncApi），只会携带此数据。
 */
func (a *Form) Primaryfield(value interface{}) *Form {
    a.Set("primaryField", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Form) Visibleon(value interface{}) *Form {
    a.Set("visibleOn", value)
    return a
}

/**
 * 消息文案配置，记住这个优先级是最低的，如果你的接口返回了 msg，接口返回的优先。
 */
func (a *Form) Messages(value interface{}) *Form {
    a.Set("messages", value)
    return a
}

/**
 * 禁用回车提交
 */
func (a *Form) Prevententersubmit(value interface{}) *Form {
    a.Set("preventEnterSubmit", value)
    return a
}

/**
 * 表单项显示为几列
 */
func (a *Form) Columncount(value interface{}) *Form {
    a.Set("columnCount", value)
    return a
}

/**
 * 表单label的对齐方式
 */
func (a *Form) Labelalign(value interface{}) *Form {
    a.Set("labelAlign", value)
    return a
}

/**
 * 指定为表单渲染器。
 */
func (a *Form) Type(value interface{}) *Form {
    a.Set("type", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *Form) Labelwidth(value interface{}) *Form {
    a.Set("labelWidth", value)
    return a
}
