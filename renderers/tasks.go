package renderers


/**
 * Tasks 渲染器，格式说明 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/tasks

 * @author wcz0
 * @version 6.2.2
 */
type Tasks struct {
	*BaseRenderer
}

func NewTasks() *Tasks {
    a := &Tasks{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "tasks")
    return a
}

/**
 * 备注列说明
 */
func (a *Tasks) RemarkLabel(value interface{}) *Tasks {
    a.Set("remarkLabel", value)
    return a
}

/**
 */
func (a *Tasks) Name(value interface{}) *Tasks {
    a.Set("name", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Tasks) ClassName(value interface{}) *Tasks {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Tasks) Disabled(value interface{}) *Tasks {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Tasks) HiddenOn(value interface{}) *Tasks {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Tasks) VisibleOn(value interface{}) *Tasks {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Tasks) Id(value interface{}) *Tasks {
    a.Set("id", value)
    return a
}

/**
 * 操作按钮文字
 */
func (a *Tasks) BtnText(value interface{}) *Tasks {
    a.Set("btnText", value)
    return a
}

/**
 * 用来获取任务状态的 API，当没有进行时任务时不会发送。
 */
func (a *Tasks) CheckApi(value interface{}) *Tasks {
    a.Set("checkApi", value)
    return a
}

/**
 * 重试操作按钮文字
 */
func (a *Tasks) RetryBtnText(value interface{}) *Tasks {
    a.Set("retryBtnText", value)
    return a
}

/**
 * 提交任务使用的 API
 */
func (a *Tasks) SubmitApi(value interface{}) *Tasks {
    a.Set("submitApi", value)
    return a
}

/**
 */
func (a *Tasks) InitialStatusCode(value interface{}) *Tasks {
    a.Set("initialStatusCode", value)
    return a
}

/**
 */
func (a *Tasks) Items(value interface{}) *Tasks {
    a.Set("items", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Tasks) StaticPlaceholder(value interface{}) *Tasks {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Tasks) StaticClassName(value interface{}) *Tasks {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Tasks) Style(value interface{}) *Tasks {
    a.Set("style", value)
    return a
}

/**
 * 如果任务失败，且可以重试，提交的时候会使用此 API
 */
func (a *Tasks) ReSubmitApi(value interface{}) *Tasks {
    a.Set("reSubmitApi", value)
    return a
}

/**
 * 配置 table className
 */
func (a *Tasks) TableClassName(value interface{}) *Tasks {
    a.Set("tableClassName", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Tasks) OnEvent(value interface{}) *Tasks {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否显示
 */
func (a *Tasks) Visible(value interface{}) *Tasks {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *Tasks) StaticSchema(value interface{}) *Tasks {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Tasks) TestIdBuilder(value interface{}) *Tasks {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 当有任务进行中，会每隔一段时间再次检测，而时间间隔就是通过此项配置，默认 3s。
 */
func (a *Tasks) Interval(value interface{}) *Tasks {
    a.Set("interval", value)
    return a
}

/**
 * 配置容器重试按钮 className
 */
func (a *Tasks) RetryBtnClassName(value interface{}) *Tasks {
    a.Set("retryBtnClassName", value)
    return a
}

/**
 * 状态列说明
 */
func (a *Tasks) StatusLabel(value interface{}) *Tasks {
    a.Set("statusLabel", value)
    return a
}

/**
 */
func (a *Tasks) ReadyStatusCode(value interface{}) *Tasks {
    a.Set("readyStatusCode", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Tasks) Hidden(value interface{}) *Tasks {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Tasks) StaticInputClassName(value interface{}) *Tasks {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Tasks) UseMobileUI(value interface{}) *Tasks {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Tasks) Testid(value interface{}) *Tasks {
    a.Set("testid", value)
    return a
}

/**
 */
func (a *Tasks) CanRetryStatusCode(value interface{}) *Tasks {
    a.Set("canRetryStatusCode", value)
    return a
}

/**
 */
func (a *Tasks) FinishStatusCode(value interface{}) *Tasks {
    a.Set("finishStatusCode", value)
    return a
}

/**
 */
func (a *Tasks) LoadingConfig(value interface{}) *Tasks {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Tasks) StaticOn(value interface{}) *Tasks {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Tasks) LoadingStatusCode(value interface{}) *Tasks {
    a.Set("loadingStatusCode", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Tasks) DisabledOn(value interface{}) *Tasks {
    a.Set("disabledOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Tasks) EditorSetting(value interface{}) *Tasks {
    a.Set("editorSetting", value)
    return a
}

/**
 * 操作列说明
 */
func (a *Tasks) OperationLabel(value interface{}) *Tasks {
    a.Set("operationLabel", value)
    return a
}

/**
 * 状态显示对应的类名配置。
 */
func (a *Tasks) StatusLabelMap(value interface{}) *Tasks {
    a.Set("statusLabelMap", value)
    return a
}

/**
 * 状态显示对应的文字显示配置。
 */
func (a *Tasks) StatusTextMap(value interface{}) *Tasks {
    a.Set("statusTextMap", value)
    return a
}

/**
 * 任务名称列说明
 */
func (a *Tasks) TaskNameLabel(value interface{}) *Tasks {
    a.Set("taskNameLabel", value)
    return a
}

/**
 */
func (a *Tasks) ErrorStatusCode(value interface{}) *Tasks {
    a.Set("errorStatusCode", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Tasks) Static(value interface{}) *Tasks {
    a.Set("static", value)
    return a
}

/**
 * 指定为任务类型
 */
func (a *Tasks) Type(value interface{}) *Tasks {
    a.Set("type", value)
    return a
}

/**
 */
func (a *Tasks) BtnClassName(value interface{}) *Tasks {
    a.Set("btnClassName", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Tasks) StaticLabelClassName(value interface{}) *Tasks {
    a.Set("staticLabelClassName", value)
    return a
}
