package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Pagination struct {
	*BaseRenderer
}

func NewPagination() *Pagination {
    a := &Pagination{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *Pagination) Set(name string, value interface{}) *Pagination {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "pagination")
    return a
}

/**
 * 是否隐藏
 */
func (a *Pagination) Hidden(value interface{}) *Pagination {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *Pagination) Visible(value interface{}) *Pagination {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Pagination) Static(value interface{}) *Pagination {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Pagination) Staticlabelclassname(value interface{}) *Pagination {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 最多显示多少个分页按钮。
 */
func (a *Pagination) Maxbuttons(value interface{}) *Pagination {
    a.Set("maxButtons", value)
    return a
}

/**
 * 是否展示分页切换，也同时受layout控制
 */
func (a *Pagination) Showperpage(value interface{}) *Pagination {
    a.Set("showPerPage", value)
    return a
}

/**
 * 总条数
 */
func (a *Pagination) Total(value interface{}) *Pagination {
    a.Set("total", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Pagination) Classname(value interface{}) *Pagination {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Pagination) Visibleon(value interface{}) *Pagination {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Pagination) Id(value interface{}) *Pagination {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Pagination) Staticclassname(value interface{}) *Pagination {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Pagination) Staticinputclassname(value interface{}) *Pagination {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Pagination) Type(value interface{}) *Pagination {
    a.Set("type", value)
    return a
}

/**
 * 当前页数
 */
func (a *Pagination) Activepage(value interface{}) *Pagination {
    a.Set("activePage", value)
    return a
}

/**
 * 弹层挂载节点
 */
func (a *Pagination) Popovercontainerselector(value interface{}) *Pagination {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 * 每页显示条数
 */
func (a *Pagination) Perpage(value interface{}) *Pagination {
    a.Set("perPage", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Pagination) Disabled(value interface{}) *Pagination {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Pagination) Disabledon(value interface{}) *Pagination {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Pagination) Staticplaceholder(value interface{}) *Pagination {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 组件样式
 */
func (a *Pagination) Style(value interface{}) *Pagination {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Pagination) Usemobileui(value interface{}) *Pagination {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Pagination) Testid(value interface{}) *Pagination {
    a.Set("testid", value)
    return a
}

/**
 * 是否显示快速跳转输入框
 */
func (a *Pagination) Showpageinput(value interface{}) *Pagination {
    a.Set("showPageInput", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Pagination) Onevent(value interface{}) *Pagination {
    a.Set("onEvent", value)
    return a
}

/**
 * 模式，默认normal，如果只想简单显示可以配置成 `simple`。
 */
func (a *Pagination) Mode(value interface{}) *Pagination {
    a.Set("mode", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Pagination) Hiddenon(value interface{}) *Pagination {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Pagination) Staticon(value interface{}) *Pagination {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *Pagination) Staticschema(value interface{}) *Pagination {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Pagination) Editorsetting(value interface{}) *Pagination {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Pagination) Hasnext(value interface{}) *Pagination {
    a.Set("hasNext", value)
    return a
}

/**
 */
func (a *Pagination) Testidbuilder(value interface{}) *Pagination {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 通过控制layout属性的顺序，调整分页结构 total,perPage,pager,go
 */
func (a *Pagination) Layout(value interface{}) *Pagination {
    a.Set("layout", value)
    return a
}

/**
 * 指定每页可以显示多少条
 */
func (a *Pagination) Perpageavailable(value interface{}) *Pagination {
    a.Set("perPageAvailable", value)
    return a
}
