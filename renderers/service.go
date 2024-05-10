package renderers


/**
 * Service 服务类控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/service

 * @author wcz0
 * @version 6.2.2
 */
type Service struct {
	*BaseRenderer
}

func NewService() *Service {
    a := &Service{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "service")
    return a
}
/**
 * 用表达式来配置。
 */
func (a *Service) InitFetchSchemaOn(value interface{}) *Service {
    a.Set("initFetchSchemaOn", value)
    return a
}

/**
 * 是否静默拉取
 */
func (a *Service) SilentPolling(value interface{}) *Service {
    a.Set("silentPolling", value)
    return a
}

/**
 */
func (a *Service) Messages(value interface{}) *Service {
    a.Set("messages", value)
    return a
}

/**
 * 是否以Alert的形式显示api接口响应的错误信息，默认展示
 */
func (a *Service) ShowErrorMsg(value interface{}) *Service {
    a.Set("showErrorMsg", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Service) UseMobileUI(value interface{}) *Service {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 页面初始化的时候，可以设置一个 API 让其取拉取，发送数据会携带当前 data 数据（包含地址栏参数），获取得数据会合并到 data 中，供组件内使用。
 */
func (a *Service) Api(value interface{}) *Service {
    a.Set("api", value)
    return a
}

/**
 * 是否默认就拉取？通过表达式来决定.
 */
func (a *Service) InitFetchOn(value interface{}) *Service {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Service) Id(value interface{}) *Service {
    a.Set("id", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Service) EditorSetting(value interface{}) *Service {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Service) StaticOn(value interface{}) *Service {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Service) StaticLabelClassName(value interface{}) *Service {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Service) StaticInputClassName(value interface{}) *Service {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *Service) Style(value interface{}) *Service {
    a.Set("style", value)
    return a
}

/**
 * 通过调用外部函数来获取数据
 */
func (a *Service) DataProvider(value interface{}) *Service {
    a.Set("dataProvider", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Service) OnEvent(value interface{}) *Service {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Service) Static(value interface{}) *Service {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Service) StaticClassName(value interface{}) *Service {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Service) StaticSchema(value interface{}) *Service {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否默认就拉取？
 */
func (a *Service) InitFetch(value interface{}) *Service {
    a.Set("initFetch", value)
    return a
}

/**
 */
func (a *Service) LoadingConfig(value interface{}) *Service {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Service) Hidden(value interface{}) *Service {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Service) HiddenOn(value interface{}) *Service {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Service) VisibleOn(value interface{}) *Service {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Service) Visible(value interface{}) *Service {
    a.Set("visible", value)
    return a
}

/**
 * 指定为 Service 数据拉取控件。
 */
func (a *Service) Type(value interface{}) *Service {
    a.Set("type", value)
    return a
}

/**
 * 用来获取远程 Schema 的 api
 */
func (a *Service) SchemaApi(value interface{}) *Service {
    a.Set("schemaApi", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Service) DisabledOn(value interface{}) *Service {
    a.Set("disabledOn", value)
    return a
}

/**
 * 内容区域
 */
func (a *Service) Body(value interface{}) *Service {
    a.Set("body", value)
    return a
}

/**
 */
func (a *Service) FetchOn(value interface{}) *Service {
    a.Set("fetchOn", value)
    return a
}

/**
 */
func (a *Service) Name(value interface{}) *Service {
    a.Set("name", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Service) Disabled(value interface{}) *Service {
    a.Set("disabled", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Service) StaticPlaceholder(value interface{}) *Service {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否默认加载 schemaApi
 */
func (a *Service) InitFetchSchema(value interface{}) *Service {
    a.Set("initFetchSchema", value)
    return a
}

/**
 * 是否轮询拉取
 */
func (a *Service) Interval(value interface{}) *Service {
    a.Set("interval", value)
    return a
}

/**
 * 关闭轮询的条件。
 */
func (a *Service) StopAutoRefreshWhen(value interface{}) *Service {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Service) ClassName(value interface{}) *Service {
    a.Set("className", value)
    return a
}

/**
 * WebScocket 地址，用于实时获取数据
 */
func (a *Service) Ws(value interface{}) *Service {
    a.Set("ws", value)
    return a
}
