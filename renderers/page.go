package renderers


/**
 * amis Page 渲染器。详情请见：https://aisuda.bce.baidu.com/amis/zh-CN/components/page

 * @author wcz0
 * @version 6.2.2
 */
type Page struct {
	*BaseRenderer
}

func NewPage() *Page {
    a := &Page{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "page")
    return a
}


func (a *Page) Set(name string, value interface{}) *Page {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 */
func (a *Page) Loadingconfig(value interface{}) *Page {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Page) Disabledon(value interface{}) *Page {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Page) Visibleon(value interface{}) *Page {
    a.Set("visibleOn", value)
    return a
}

/**
 * 内容区域
 */
func (a *Page) Body(value interface{}) *Page {
    a.Set("body", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Page) Style(value interface{}) *Page {
    a.Set("style", value)
    return a
}

/**
 * 下拉刷新配置
 */
func (a *Page) Pullrefresh(value interface{}) *Page {
    a.Set("pullRefresh", value)
    return a
}

/**
 */
func (a *Page) Name(value interface{}) *Page {
    a.Set("name", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Page) Staticon(value interface{}) *Page {
    a.Set("staticOn", value)
    return a
}

/**
 * 指定为 page 渲染器。
 */
func (a *Page) Type(value interface{}) *Page {
    a.Set("type", value)
    return a
}

/**
 * 是否显示错误信息，默认是显示的。
 */
func (a *Page) Showerrormsg(value interface{}) *Page {
    a.Set("showErrorMsg", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Page) Hiddenon(value interface{}) *Page {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 边栏内容是否粘住，即不跟随滚动。
 */
func (a *Page) Asidesticky(value interface{}) *Page {
    a.Set("asideSticky", value)
    return a
}

/**
 * 是否默认就拉取表达式
 */
func (a *Page) Initfetchon(value interface{}) *Page {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 页面顶部区域，当存在 title 时在右上角显示。
 */
func (a *Page) Toolbar(value interface{}) *Page {
    a.Set("toolbar", value)
    return a
}

/**
 * 配置 toolbar 容器 className
 */
func (a *Page) Toolbarclassname(value interface{}) *Page {
    a.Set("toolbarClassName", value)
    return a
}

/**
 */
func (a *Page) Testidbuilder(value interface{}) *Page {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 配置 header 容器 className
 */
func (a *Page) Headerclassname(value interface{}) *Page {
    a.Set("headerClassName", value)
    return a
}

/**
 * 是否要静默加载，也就是说不显示进度
 */
func (a *Page) Silentpolling(value interface{}) *Page {
    a.Set("silentPolling", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Page) Static(value interface{}) *Page {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Page) Staticinputclassname(value interface{}) *Page {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Page) Editorsetting(value interface{}) *Page {
    a.Set("editorSetting", value)
    return a
}

/**
 * 自定义页面级别样式表
 */
func (a *Page) Css(value interface{}) *Page {
    a.Set("css", value)
    return a
}

/**
 */
func (a *Page) Messages(value interface{}) *Page {
    a.Set("messages", value)
    return a
}

/**
 */
func (a *Page) Definitions(value interface{}) *Page {
    a.Set("definitions", value)
    return a
}

/**
 * 配置轮询间隔，配置后 initApi 将轮询加载。
 */
func (a *Page) Interval(value interface{}) *Page {
    a.Set("interval", value)
    return a
}

/**
 * 配置容器 className
 */
func (a *Page) Classname(value interface{}) *Page {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Page) Visible(value interface{}) *Page {
    a.Set("visible", value)
    return a
}

/**
 * 是否默认就拉取？
 */
func (a *Page) Initfetch(value interface{}) *Page {
    a.Set("initFetch", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Page) Disabled(value interface{}) *Page {
    a.Set("disabled", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Page) Id(value interface{}) *Page {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Page) Staticclassname(value interface{}) *Page {
    a.Set("staticClassName", value)
    return a
}

/**
 * 内容区 css 类名
 */
func (a *Page) Bodyclassname(value interface{}) *Page {
    a.Set("bodyClassName", value)
    return a
}

/**
 * css 变量
 */
func (a *Page) Cssvars(value interface{}) *Page {
    a.Set("cssVars", value)
    return a
}

/**
 * 页面副标题
 */
func (a *Page) Subtitle(value interface{}) *Page {
    a.Set("subTitle", value)
    return a
}

/**
 * 页面描述, 标题旁边会出现个小图标，放上去会显示这个属性配置的内容。
 */
func (a *Page) Remark(value interface{}) *Page {
    a.Set("remark", value)
    return a
}

/**
 * 页面初始化的时候，可以设置一个 API 让其取拉取，发送数据会携带当前 data 数据（包含地址栏参数），获取得数据会合并到 data 中，供组件内使用。
 */
func (a *Page) Initapi(value interface{}) *Page {
    a.Set("initApi", value)
    return a
}

/**
 * 默认不设置自动感觉内容来决定要不要展示这些区域 如果配置了，以配置为主。
 */
func (a *Page) Regions(value interface{}) *Page {
    a.Set("regions", value)
    return a
}

/**
 */
func (a *Page) Staticschema(value interface{}) *Page {
    a.Set("staticSchema", value)
    return a
}

/**
 * 边栏最小宽度
 */
func (a *Page) Asideminwidth(value interface{}) *Page {
    a.Set("asideMinWidth", value)
    return a
}

/**
 * 边栏最小宽度
 */
func (a *Page) Asidemaxwidth(value interface{}) *Page {
    a.Set("asideMaxWidth", value)
    return a
}

/**
 * 边栏区 css 类名
 */
func (a *Page) Asideclassname(value interface{}) *Page {
    a.Set("asideClassName", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Page) Onevent(value interface{}) *Page {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *Page) Testid(value interface{}) *Page {
    a.Set("testid", value)
    return a
}

/**
 * 页面标题
 */
func (a *Page) Title(value interface{}) *Page {
    a.Set("title", value)
    return a
}

/**
 * 边栏区域
 */
func (a *Page) Aside(value interface{}) *Page {
    a.Set("aside", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Page) Staticplaceholder(value interface{}) *Page {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 边栏是否允许拖动
 */
func (a *Page) Asideresizor(value interface{}) *Page {
    a.Set("asideResizor", value)
    return a
}

/**
 * 配置停止轮询的条件。
 */
func (a *Page) Stopautorefreshwhen(value interface{}) *Page {
    a.Set("stopAutoRefreshWhen", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Page) Hidden(value interface{}) *Page {
    a.Set("hidden", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Page) Usemobileui(value interface{}) *Page {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 页面级别的初始数据
 */
func (a *Page) Data(value interface{}) *Page {
    a.Set("data", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Page) Staticlabelclassname(value interface{}) *Page {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 移动端下的样式表
 */
func (a *Page) Mobilecss(value interface{}) *Page {
    a.Set("mobileCSS", value)
    return a
}
