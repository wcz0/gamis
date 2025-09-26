package renderers


/**

 * @author wcz0
 * @version 6.2.2
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


func (a *WizardStep) Set(name string, value interface{}) *WizardStep {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * visible
 */
func (a *WizardStep) Visible(value interface{}) *WizardStep {
    a.Set("visible", value)
    return a
}

/**
 * style
 */
func (a *WizardStep) Style(value interface{}) *WizardStep {
    a.Set("style", value)
    return a
}

/**
 * data
 */
func (a *WizardStep) Data(value interface{}) *WizardStep {
    a.Set("data", value)
    return a
}

/**
 * initFetch
 */
func (a *WizardStep) InitFetch(value interface{}) *WizardStep {
    a.Set("initFetch", value)
    return a
}

/**
 * clearAfterSubmit
 */
func (a *WizardStep) ClearAfterSubmit(value interface{}) *WizardStep {
    a.Set("clearAfterSubmit", value)
    return a
}

/**
 * panelClassName
 */
func (a *WizardStep) PanelClassName(value interface{}) *WizardStep {
    a.Set("panelClassName", value)
    return a
}

/**
 * tabs
 */
func (a *WizardStep) Tabs(value interface{}) *WizardStep {
    a.Set("tabs", value)
    return a
}

/**
 * mode
 */
func (a *WizardStep) Mode(value interface{}) *WizardStep {
    a.Set("mode", value)
    return a
}

/**
 * autoFocus
 */
func (a *WizardStep) AutoFocus(value interface{}) *WizardStep {
    a.Set("autoFocus", value)
    return a
}

/**
 * title
 */
func (a *WizardStep) Title(value interface{}) *WizardStep {
    a.Set("title", value)
    return a
}

/**
 * debugConfig
 */
func (a *WizardStep) DebugConfig(value interface{}) *WizardStep {
    a.Set("debugConfig", value)
    return a
}

/**
 * initAsyncApi
 */
func (a *WizardStep) InitAsyncApi(value interface{}) *WizardStep {
    a.Set("initAsyncApi", value)
    return a
}

/**
 * interval
 */
func (a *WizardStep) Interval(value interface{}) *WizardStep {
    a.Set("interval", value)
    return a
}

/**
 * checkInterval
 */
func (a *WizardStep) CheckInterval(value interface{}) *WizardStep {
    a.Set("checkInterval", value)
    return a
}

/**
 * messages
 */
func (a *WizardStep) Messages(value interface{}) *WizardStep {
    a.Set("messages", value)
    return a
}

/**
 * labelWidth
 */
func (a *WizardStep) LabelWidth(value interface{}) *WizardStep {
    a.Set("labelWidth", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *WizardStep) StaticPlaceholder(value interface{}) *WizardStep {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * initCheckInterval
 */
func (a *WizardStep) InitCheckInterval(value interface{}) *WizardStep {
    a.Set("initCheckInterval", value)
    return a
}

/**
 * silentPolling
 */
func (a *WizardStep) SilentPolling(value interface{}) *WizardStep {
    a.Set("silentPolling", value)
    return a
}

/**
 * stopAutoRefreshWhen
 */
func (a *WizardStep) StopAutoRefreshWhen(value interface{}) *WizardStep {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * clearPersistDataAfterSubmit
 */
func (a *WizardStep) ClearPersistDataAfterSubmit(value interface{}) *WizardStep {
    a.Set("clearPersistDataAfterSubmit", value)
    return a
}

/**
 * submitOnChange
 */
func (a *WizardStep) SubmitOnChange(value interface{}) *WizardStep {
    a.Set("submitOnChange", value)
    return a
}

/**
 * jumpableOn
 */
func (a *WizardStep) JumpableOn(value interface{}) *WizardStep {
    a.Set("jumpableOn", value)
    return a
}

/**
 * id
 */
func (a *WizardStep) Id(value interface{}) *WizardStep {
    a.Set("id", value)
    return a
}

/**
 * finishedField
 */
func (a *WizardStep) FinishedField(value interface{}) *WizardStep {
    a.Set("finishedField", value)
    return a
}

/**
 * rules
 */
func (a *WizardStep) Rules(value interface{}) *WizardStep {
    a.Set("rules", value)
    return a
}

/**
 * preventEnterSubmit
 */
func (a *WizardStep) PreventEnterSubmit(value interface{}) *WizardStep {
    a.Set("preventEnterSubmit", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *WizardStep) StaticLabelClassName(value interface{}) *WizardStep {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *WizardStep) StaticSchema(value interface{}) *WizardStep {
    a.Set("staticSchema", value)
    return a
}

/**
 * editorSetting
 */
func (a *WizardStep) EditorSetting(value interface{}) *WizardStep {
    a.Set("editorSetting", value)
    return a
}

/**
 * description
 */
func (a *WizardStep) Description(value interface{}) *WizardStep {
    a.Set("description", value)
    return a
}

/**
 * actions
 */
func (a *WizardStep) Actions(value interface{}) *WizardStep {
    a.Set("actions", value)
    return a
}

/**
 * asyncApi
 */
func (a *WizardStep) AsyncApi(value interface{}) *WizardStep {
    a.Set("asyncApi", value)
    return a
}

/**
 * redirect
 */
func (a *WizardStep) Redirect(value interface{}) *WizardStep {
    a.Set("redirect", value)
    return a
}

/**
 * affixFooter
 */
func (a *WizardStep) AffixFooter(value interface{}) *WizardStep {
    a.Set("affixFooter", value)
    return a
}

/**
 * staticOn
 */
func (a *WizardStep) StaticOn(value interface{}) *WizardStep {
    a.Set("staticOn", value)
    return a
}

/**
 * fieldSet
 */
func (a *WizardStep) FieldSet(value interface{}) *WizardStep {
    a.Set("fieldSet", value)
    return a
}

/**
 * debug
 */
func (a *WizardStep) Debug(value interface{}) *WizardStep {
    a.Set("debug", value)
    return a
}

/**
 * initFinishedField
 */
func (a *WizardStep) InitFinishedField(value interface{}) *WizardStep {
    a.Set("initFinishedField", value)
    return a
}

/**
 * submitText
 */
func (a *WizardStep) SubmitText(value interface{}) *WizardStep {
    a.Set("submitText", value)
    return a
}

/**
 * promptPageLeave
 */
func (a *WizardStep) PromptPageLeave(value interface{}) *WizardStep {
    a.Set("promptPageLeave", value)
    return a
}

/**
 * className
 */
func (a *WizardStep) ClassName(value interface{}) *WizardStep {
    a.Set("className", value)
    return a
}

/**
 * hiddenOn
 */
func (a *WizardStep) HiddenOn(value interface{}) *WizardStep {
    a.Set("hiddenOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *WizardStep) VisibleOn(value interface{}) *WizardStep {
    a.Set("visibleOn", value)
    return a
}

/**
 * useMobileUI
 */
func (a *WizardStep) UseMobileUI(value interface{}) *WizardStep {
    a.Set("useMobileUI", value)
    return a
}

/**
 * icon
 */
func (a *WizardStep) Icon(value interface{}) *WizardStep {
    a.Set("icon", value)
    return a
}

/**
 * initFetchOn
 */
func (a *WizardStep) InitFetchOn(value interface{}) *WizardStep {
    a.Set("initFetchOn", value)
    return a
}

/**
 * persistDataKeys
 */
func (a *WizardStep) PersistDataKeys(value interface{}) *WizardStep {
    a.Set("persistDataKeys", value)
    return a
}

/**
 * reload
 */
func (a *WizardStep) Reload(value interface{}) *WizardStep {
    a.Set("reload", value)
    return a
}

/**
 * disabledOn
 */
func (a *WizardStep) DisabledOn(value interface{}) *WizardStep {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *WizardStep) Hidden(value interface{}) *WizardStep {
    a.Set("hidden", value)
    return a
}

/**
 * resetAfterSubmit
 */
func (a *WizardStep) ResetAfterSubmit(value interface{}) *WizardStep {
    a.Set("resetAfterSubmit", value)
    return a
}

/**
 * static
 */
func (a *WizardStep) Static(value interface{}) *WizardStep {
    a.Set("static", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *WizardStep) StaticInputClassName(value interface{}) *WizardStep {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * subTitle
 */
func (a *WizardStep) SubTitle(value interface{}) *WizardStep {
    a.Set("subTitle", value)
    return a
}

/**
 * feedback
 */
func (a *WizardStep) Feedback(value interface{}) *WizardStep {
    a.Set("feedback", value)
    return a
}

/**
 * target
 */
func (a *WizardStep) Target(value interface{}) *WizardStep {
    a.Set("target", value)
    return a
}

/**
 * onEvent
 */
func (a *WizardStep) OnEvent(value interface{}) *WizardStep {
    a.Set("onEvent", value)
    return a
}

/**
 * columnCount
 */
func (a *WizardStep) ColumnCount(value interface{}) *WizardStep {
    a.Set("columnCount", value)
    return a
}

/**
 * horizontal
 */
func (a *WizardStep) Horizontal(value interface{}) *WizardStep {
    a.Set("horizontal", value)
    return a
}

/**
 * staticClassName
 */
func (a *WizardStep) StaticClassName(value interface{}) *WizardStep {
    a.Set("staticClassName", value)
    return a
}

/**
 * value
 */
func (a *WizardStep) Value(value interface{}) *WizardStep {
    a.Set("value", value)
    return a
}

/**
 * initApi
 */
func (a *WizardStep) InitApi(value interface{}) *WizardStep {
    a.Set("initApi", value)
    return a
}

/**
 * persistData
 */
func (a *WizardStep) PersistData(value interface{}) *WizardStep {
    a.Set("persistData", value)
    return a
}

/**
 * name
 */
func (a *WizardStep) Name(value interface{}) *WizardStep {
    a.Set("name", value)
    return a
}

/**
 * primaryField
 */
func (a *WizardStep) PrimaryField(value interface{}) *WizardStep {
    a.Set("primaryField", value)
    return a
}

/**
 * wrapWithPanel
 */
func (a *WizardStep) WrapWithPanel(value interface{}) *WizardStep {
    a.Set("wrapWithPanel", value)
    return a
}

/**
 * labelAlign
 */
func (a *WizardStep) LabelAlign(value interface{}) *WizardStep {
    a.Set("labelAlign", value)
    return a
}

/**
 * body
 */
func (a *WizardStep) Body(value interface{}) *WizardStep {
    a.Set("body", value)
    return a
}

/**
 * api
 */
func (a *WizardStep) Api(value interface{}) *WizardStep {
    a.Set("api", value)
    return a
}

/**
 * submitOnInit
 */
func (a *WizardStep) SubmitOnInit(value interface{}) *WizardStep {
    a.Set("submitOnInit", value)
    return a
}

/**
 * promptPageLeaveMessage
 */
func (a *WizardStep) PromptPageLeaveMessage(value interface{}) *WizardStep {
    a.Set("promptPageLeaveMessage", value)
    return a
}

/**
 * label
 */
func (a *WizardStep) Label(value interface{}) *WizardStep {
    a.Set("label", value)
    return a
}

/**
 * jumpable
 */
func (a *WizardStep) Jumpable(value interface{}) *WizardStep {
    a.Set("jumpable", value)
    return a
}

/**
 * disabled
 */
func (a *WizardStep) Disabled(value interface{}) *WizardStep {
    a.Set("disabled", value)
    return a
}
