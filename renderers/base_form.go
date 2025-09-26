package renderers


/**
 * Form 表单渲染器。说明：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/index

 * @author wcz0
 * @version 6.2.2
 */
type BaseForm struct {
	*BaseRenderer
}

func NewBaseForm() *BaseForm {
    a := &BaseForm{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *BaseForm) Set(name string, value interface{}) *BaseForm {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * body
 */
func (a *BaseForm) Body(value interface{}) *BaseForm {
    a.Set("body", value)
    return a
}

/**
 * initAsyncApi
 */
func (a *BaseForm) InitAsyncApi(value interface{}) *BaseForm {
    a.Set("initAsyncApi", value)
    return a
}

/**
 * mode
 */
func (a *BaseForm) Mode(value interface{}) *BaseForm {
    a.Set("mode", value)
    return a
}

/**
 * messages
 */
func (a *BaseForm) Messages(value interface{}) *BaseForm {
    a.Set("messages", value)
    return a
}

/**
 * disabledOn
 */
func (a *BaseForm) DisabledOn(value interface{}) *BaseForm {
    a.Set("disabledOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *BaseForm) StaticLabelClassName(value interface{}) *BaseForm {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * debug
 */
func (a *BaseForm) Debug(value interface{}) *BaseForm {
    a.Set("debug", value)
    return a
}

/**
 * debugConfig
 */
func (a *BaseForm) DebugConfig(value interface{}) *BaseForm {
    a.Set("debugConfig", value)
    return a
}

/**
 * hidden
 */
func (a *BaseForm) Hidden(value interface{}) *BaseForm {
    a.Set("hidden", value)
    return a
}

/**
 * interval
 */
func (a *BaseForm) Interval(value interface{}) *BaseForm {
    a.Set("interval", value)
    return a
}

/**
 * autoFocus
 */
func (a *BaseForm) AutoFocus(value interface{}) *BaseForm {
    a.Set("autoFocus", value)
    return a
}

/**
 * name
 */
func (a *BaseForm) Name(value interface{}) *BaseForm {
    a.Set("name", value)
    return a
}

/**
 * reload
 */
func (a *BaseForm) Reload(value interface{}) *BaseForm {
    a.Set("reload", value)
    return a
}

/**
 * onEvent
 */
func (a *BaseForm) OnEvent(value interface{}) *BaseForm {
    a.Set("onEvent", value)
    return a
}

/**
 * data
 */
func (a *BaseForm) Data(value interface{}) *BaseForm {
    a.Set("data", value)
    return a
}

/**
 * initFetchOn
 */
func (a *BaseForm) InitFetchOn(value interface{}) *BaseForm {
    a.Set("initFetchOn", value)
    return a
}

/**
 * silentPolling
 */
func (a *BaseForm) SilentPolling(value interface{}) *BaseForm {
    a.Set("silentPolling", value)
    return a
}

/**
 * feedback
 */
func (a *BaseForm) Feedback(value interface{}) *BaseForm {
    a.Set("feedback", value)
    return a
}

/**
 * preventEnterSubmit
 */
func (a *BaseForm) PreventEnterSubmit(value interface{}) *BaseForm {
    a.Set("preventEnterSubmit", value)
    return a
}

/**
 * horizontal
 */
func (a *BaseForm) Horizontal(value interface{}) *BaseForm {
    a.Set("horizontal", value)
    return a
}

/**
 * className
 */
func (a *BaseForm) ClassName(value interface{}) *BaseForm {
    a.Set("className", value)
    return a
}

/**
 * checkInterval
 */
func (a *BaseForm) CheckInterval(value interface{}) *BaseForm {
    a.Set("checkInterval", value)
    return a
}

/**
 * staticOn
 */
func (a *BaseForm) StaticOn(value interface{}) *BaseForm {
    a.Set("staticOn", value)
    return a
}

/**
 * fieldSet
 */
func (a *BaseForm) FieldSet(value interface{}) *BaseForm {
    a.Set("fieldSet", value)
    return a
}

/**
 * initFetch
 */
func (a *BaseForm) InitFetch(value interface{}) *BaseForm {
    a.Set("initFetch", value)
    return a
}

/**
 * visible
 */
func (a *BaseForm) Visible(value interface{}) *BaseForm {
    a.Set("visible", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *BaseForm) StaticPlaceholder(value interface{}) *BaseForm {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticClassName
 */
func (a *BaseForm) StaticClassName(value interface{}) *BaseForm {
    a.Set("staticClassName", value)
    return a
}

/**
 * useMobileUI
 */
func (a *BaseForm) UseMobileUI(value interface{}) *BaseForm {
    a.Set("useMobileUI", value)
    return a
}

/**
 * persistDataKeys
 */
func (a *BaseForm) PersistDataKeys(value interface{}) *BaseForm {
    a.Set("persistDataKeys", value)
    return a
}

/**
 * asyncApi
 */
func (a *BaseForm) AsyncApi(value interface{}) *BaseForm {
    a.Set("asyncApi", value)
    return a
}

/**
 * hiddenOn
 */
func (a *BaseForm) HiddenOn(value interface{}) *BaseForm {
    a.Set("hiddenOn", value)
    return a
}

/**
 * actions
 */
func (a *BaseForm) Actions(value interface{}) *BaseForm {
    a.Set("actions", value)
    return a
}

/**
 * api
 */
func (a *BaseForm) Api(value interface{}) *BaseForm {
    a.Set("api", value)
    return a
}

/**
 * panelClassName
 */
func (a *BaseForm) PanelClassName(value interface{}) *BaseForm {
    a.Set("panelClassName", value)
    return a
}

/**
 * target
 */
func (a *BaseForm) Target(value interface{}) *BaseForm {
    a.Set("target", value)
    return a
}

/**
 * rules
 */
func (a *BaseForm) Rules(value interface{}) *BaseForm {
    a.Set("rules", value)
    return a
}

/**
 * submitText
 */
func (a *BaseForm) SubmitText(value interface{}) *BaseForm {
    a.Set("submitText", value)
    return a
}

/**
 * editorSetting
 */
func (a *BaseForm) EditorSetting(value interface{}) *BaseForm {
    a.Set("editorSetting", value)
    return a
}

/**
 * clearPersistDataAfterSubmit
 */
func (a *BaseForm) ClearPersistDataAfterSubmit(value interface{}) *BaseForm {
    a.Set("clearPersistDataAfterSubmit", value)
    return a
}

/**
 * finishedField
 */
func (a *BaseForm) FinishedField(value interface{}) *BaseForm {
    a.Set("finishedField", value)
    return a
}

/**
 * primaryField
 */
func (a *BaseForm) PrimaryField(value interface{}) *BaseForm {
    a.Set("primaryField", value)
    return a
}

/**
 * redirect
 */
func (a *BaseForm) Redirect(value interface{}) *BaseForm {
    a.Set("redirect", value)
    return a
}

/**
 * labelAlign
 */
func (a *BaseForm) LabelAlign(value interface{}) *BaseForm {
    a.Set("labelAlign", value)
    return a
}

/**
 * title
 */
func (a *BaseForm) Title(value interface{}) *BaseForm {
    a.Set("title", value)
    return a
}

/**
 * tabs
 */
func (a *BaseForm) Tabs(value interface{}) *BaseForm {
    a.Set("tabs", value)
    return a
}

/**
 * submitOnInit
 */
func (a *BaseForm) SubmitOnInit(value interface{}) *BaseForm {
    a.Set("submitOnInit", value)
    return a
}

/**
 * promptPageLeave
 */
func (a *BaseForm) PromptPageLeave(value interface{}) *BaseForm {
    a.Set("promptPageLeave", value)
    return a
}

/**
 * disabled
 */
func (a *BaseForm) Disabled(value interface{}) *BaseForm {
    a.Set("disabled", value)
    return a
}

/**
 * staticSchema
 */
func (a *BaseForm) StaticSchema(value interface{}) *BaseForm {
    a.Set("staticSchema", value)
    return a
}

/**
 * style
 */
func (a *BaseForm) Style(value interface{}) *BaseForm {
    a.Set("style", value)
    return a
}

/**
 * initApi
 */
func (a *BaseForm) InitApi(value interface{}) *BaseForm {
    a.Set("initApi", value)
    return a
}

/**
 * columnCount
 */
func (a *BaseForm) ColumnCount(value interface{}) *BaseForm {
    a.Set("columnCount", value)
    return a
}

/**
 * wrapWithPanel
 */
func (a *BaseForm) WrapWithPanel(value interface{}) *BaseForm {
    a.Set("wrapWithPanel", value)
    return a
}

/**
 * labelWidth
 */
func (a *BaseForm) LabelWidth(value interface{}) *BaseForm {
    a.Set("labelWidth", value)
    return a
}

/**
 * static
 */
func (a *BaseForm) Static(value interface{}) *BaseForm {
    a.Set("static", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *BaseForm) StaticInputClassName(value interface{}) *BaseForm {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * initFinishedField
 */
func (a *BaseForm) InitFinishedField(value interface{}) *BaseForm {
    a.Set("initFinishedField", value)
    return a
}

/**
 * initCheckInterval
 */
func (a *BaseForm) InitCheckInterval(value interface{}) *BaseForm {
    a.Set("initCheckInterval", value)
    return a
}

/**
 * resetAfterSubmit
 */
func (a *BaseForm) ResetAfterSubmit(value interface{}) *BaseForm {
    a.Set("resetAfterSubmit", value)
    return a
}

/**
 * affixFooter
 */
func (a *BaseForm) AffixFooter(value interface{}) *BaseForm {
    a.Set("affixFooter", value)
    return a
}

/**
 * id
 */
func (a *BaseForm) Id(value interface{}) *BaseForm {
    a.Set("id", value)
    return a
}

/**
 * stopAutoRefreshWhen
 */
func (a *BaseForm) StopAutoRefreshWhen(value interface{}) *BaseForm {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * persistData
 */
func (a *BaseForm) PersistData(value interface{}) *BaseForm {
    a.Set("persistData", value)
    return a
}

/**
 * submitOnChange
 */
func (a *BaseForm) SubmitOnChange(value interface{}) *BaseForm {
    a.Set("submitOnChange", value)
    return a
}

/**
 * visibleOn
 */
func (a *BaseForm) VisibleOn(value interface{}) *BaseForm {
    a.Set("visibleOn", value)
    return a
}

/**
 * clearAfterSubmit
 */
func (a *BaseForm) ClearAfterSubmit(value interface{}) *BaseForm {
    a.Set("clearAfterSubmit", value)
    return a
}

/**
 * promptPageLeaveMessage
 */
func (a *BaseForm) PromptPageLeaveMessage(value interface{}) *BaseForm {
    a.Set("promptPageLeaveMessage", value)
    return a
}
