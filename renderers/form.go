package renderers


/**

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
 * submitOnChange
 */
func (a *Form) SubmitOnChange(value interface{}) *Form {
    a.Set("submitOnChange", value)
    return a
}

/**
 * affixFooter
 */
func (a *Form) AffixFooter(value interface{}) *Form {
    a.Set("affixFooter", value)
    return a
}

/**
 * labelWidth
 */
func (a *Form) LabelWidth(value interface{}) *Form {
    a.Set("labelWidth", value)
    return a
}

/**
 * staticClassName
 */
func (a *Form) StaticClassName(value interface{}) *Form {
    a.Set("staticClassName", value)
    return a
}

/**
 * stopAutoRefreshWhen
 */
func (a *Form) StopAutoRefreshWhen(value interface{}) *Form {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * feedback
 */
func (a *Form) Feedback(value interface{}) *Form {
    a.Set("feedback", value)
    return a
}

/**
 */
func (a *Form) Type(value interface{}) *Form {
    a.Set("type", value)
    return a
}

/**
 * initApi
 */
func (a *Form) InitApi(value interface{}) *Form {
    a.Set("initApi", value)
    return a
}

/**
 * initFinishedField
 */
func (a *Form) InitFinishedField(value interface{}) *Form {
    a.Set("initFinishedField", value)
    return a
}

/**
 * visibleOn
 */
func (a *Form) VisibleOn(value interface{}) *Form {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Form) StaticPlaceholder(value interface{}) *Form {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * initFetchOn
 */
func (a *Form) InitFetchOn(value interface{}) *Form {
    a.Set("initFetchOn", value)
    return a
}

/**
 * panelClassName
 */
func (a *Form) PanelClassName(value interface{}) *Form {
    a.Set("panelClassName", value)
    return a
}

/**
 * redirect
 */
func (a *Form) Redirect(value interface{}) *Form {
    a.Set("redirect", value)
    return a
}

/**
 * wrapWithPanel
 */
func (a *Form) WrapWithPanel(value interface{}) *Form {
    a.Set("wrapWithPanel", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Form) HiddenOn(value interface{}) *Form {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *Form) Id(value interface{}) *Form {
    a.Set("id", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Form) StaticLabelClassName(value interface{}) *Form {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * rules
 */
func (a *Form) Rules(value interface{}) *Form {
    a.Set("rules", value)
    return a
}

/**
 * hidden
 */
func (a *Form) Hidden(value interface{}) *Form {
    a.Set("hidden", value)
    return a
}

/**
 * target
 */
func (a *Form) Target(value interface{}) *Form {
    a.Set("target", value)
    return a
}

/**
 * staticSchema
 */
func (a *Form) StaticSchema(value interface{}) *Form {
    a.Set("staticSchema", value)
    return a
}

/**
 * interval
 */
func (a *Form) Interval(value interface{}) *Form {
    a.Set("interval", value)
    return a
}

/**
 * preventEnterSubmit
 */
func (a *Form) PreventEnterSubmit(value interface{}) *Form {
    a.Set("preventEnterSubmit", value)
    return a
}

/**
 * horizontal
 */
func (a *Form) Horizontal(value interface{}) *Form {
    a.Set("horizontal", value)
    return a
}

/**
 * primaryField
 */
func (a *Form) PrimaryField(value interface{}) *Form {
    a.Set("primaryField", value)
    return a
}

/**
 * disabled
 */
func (a *Form) Disabled(value interface{}) *Form {
    a.Set("disabled", value)
    return a
}

/**
 * editorSetting
 */
func (a *Form) EditorSetting(value interface{}) *Form {
    a.Set("editorSetting", value)
    return a
}

/**
 * fieldSet
 */
func (a *Form) FieldSet(value interface{}) *Form {
    a.Set("fieldSet", value)
    return a
}

/**
 * mode
 */
func (a *Form) Mode(value interface{}) *Form {
    a.Set("mode", value)
    return a
}

/**
 * columnCount
 */
func (a *Form) ColumnCount(value interface{}) *Form {
    a.Set("columnCount", value)
    return a
}

/**
 * messages
 */
func (a *Form) Messages(value interface{}) *Form {
    a.Set("messages", value)
    return a
}

/**
 * submitText
 */
func (a *Form) SubmitText(value interface{}) *Form {
    a.Set("submitText", value)
    return a
}

/**
 * labelAlign
 */
func (a *Form) LabelAlign(value interface{}) *Form {
    a.Set("labelAlign", value)
    return a
}

/**
 * onEvent
 */
func (a *Form) OnEvent(value interface{}) *Form {
    a.Set("onEvent", value)
    return a
}

/**
 * autoFocus
 */
func (a *Form) AutoFocus(value interface{}) *Form {
    a.Set("autoFocus", value)
    return a
}

/**
 * submitOnInit
 */
func (a *Form) SubmitOnInit(value interface{}) *Form {
    a.Set("submitOnInit", value)
    return a
}

/**
 * disabledOn
 */
func (a *Form) DisabledOn(value interface{}) *Form {
    a.Set("disabledOn", value)
    return a
}

/**
 * title
 */
func (a *Form) Title(value interface{}) *Form {
    a.Set("title", value)
    return a
}

/**
 * data
 */
func (a *Form) Data(value interface{}) *Form {
    a.Set("data", value)
    return a
}

/**
 * initCheckInterval
 */
func (a *Form) InitCheckInterval(value interface{}) *Form {
    a.Set("initCheckInterval", value)
    return a
}

/**
 * promptPageLeaveMessage
 */
func (a *Form) PromptPageLeaveMessage(value interface{}) *Form {
    a.Set("promptPageLeaveMessage", value)
    return a
}

/**
 * visible
 */
func (a *Form) Visible(value interface{}) *Form {
    a.Set("visible", value)
    return a
}

/**
 * debugConfig
 */
func (a *Form) DebugConfig(value interface{}) *Form {
    a.Set("debugConfig", value)
    return a
}

/**
 * initAsyncApi
 */
func (a *Form) InitAsyncApi(value interface{}) *Form {
    a.Set("initAsyncApi", value)
    return a
}

/**
 * persistDataKeys
 */
func (a *Form) PersistDataKeys(value interface{}) *Form {
    a.Set("persistDataKeys", value)
    return a
}

/**
 * reload
 */
func (a *Form) Reload(value interface{}) *Form {
    a.Set("reload", value)
    return a
}

/**
 * className
 */
func (a *Form) ClassName(value interface{}) *Form {
    a.Set("className", value)
    return a
}

/**
 * style
 */
func (a *Form) Style(value interface{}) *Form {
    a.Set("style", value)
    return a
}

/**
 * actions
 */
func (a *Form) Actions(value interface{}) *Form {
    a.Set("actions", value)
    return a
}

/**
 * persistData
 */
func (a *Form) PersistData(value interface{}) *Form {
    a.Set("persistData", value)
    return a
}

/**
 * asyncApi
 */
func (a *Form) AsyncApi(value interface{}) *Form {
    a.Set("asyncApi", value)
    return a
}

/**
 * name
 */
func (a *Form) Name(value interface{}) *Form {
    a.Set("name", value)
    return a
}

/**
 * promptPageLeave
 */
func (a *Form) PromptPageLeave(value interface{}) *Form {
    a.Set("promptPageLeave", value)
    return a
}

/**
 * staticOn
 */
func (a *Form) StaticOn(value interface{}) *Form {
    a.Set("staticOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Form) UseMobileUI(value interface{}) *Form {
    a.Set("useMobileUI", value)
    return a
}

/**
 * tabs
 */
func (a *Form) Tabs(value interface{}) *Form {
    a.Set("tabs", value)
    return a
}

/**
 * checkInterval
 */
func (a *Form) CheckInterval(value interface{}) *Form {
    a.Set("checkInterval", value)
    return a
}

/**
 * static
 */
func (a *Form) Static(value interface{}) *Form {
    a.Set("static", value)
    return a
}

/**
 * initFetch
 */
func (a *Form) InitFetch(value interface{}) *Form {
    a.Set("initFetch", value)
    return a
}

/**
 * silentPolling
 */
func (a *Form) SilentPolling(value interface{}) *Form {
    a.Set("silentPolling", value)
    return a
}

/**
 * api
 */
func (a *Form) Api(value interface{}) *Form {
    a.Set("api", value)
    return a
}

/**
 * clearAfterSubmit
 */
func (a *Form) ClearAfterSubmit(value interface{}) *Form {
    a.Set("clearAfterSubmit", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Form) StaticInputClassName(value interface{}) *Form {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * body
 */
func (a *Form) Body(value interface{}) *Form {
    a.Set("body", value)
    return a
}

/**
 * debug
 */
func (a *Form) Debug(value interface{}) *Form {
    a.Set("debug", value)
    return a
}

/**
 * clearPersistDataAfterSubmit
 */
func (a *Form) ClearPersistDataAfterSubmit(value interface{}) *Form {
    a.Set("clearPersistDataAfterSubmit", value)
    return a
}

/**
 * finishedField
 */
func (a *Form) FinishedField(value interface{}) *Form {
    a.Set("finishedField", value)
    return a
}

/**
 * resetAfterSubmit
 */
func (a *Form) ResetAfterSubmit(value interface{}) *Form {
    a.Set("resetAfterSubmit", value)
    return a
}
