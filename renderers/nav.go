package renderers


/**
 * Nav 导航渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/nav

 * @author wcz0
 * @version 6.2.2
 */
type Nav struct {
	*BaseRenderer
}

func NewNav() *Nav {
    a := &Nav{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "nav")
    return a
}


func (a *Nav) Set(name string, value interface{}) *Nav {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 指定为 Nav 导航渲染器
 */
func (a *Nav) Type(value interface{}) *Nav {
    a.Set("type", value)
    return a
}

/**
 * 角标
 */
func (a *Nav) Itembadge(value interface{}) *Nav {
    a.Set("itemBadge", value)
    return a
}

/**
 * 垂直模式 非折叠状态下 控制菜单打开方式
 * 可选值: float | inline
 */
func (a *Nav) Mode(value interface{}) *Nav {
    a.Set("mode", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Nav) Hiddenon(value interface{}) *Nav {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Nav) Staticclassname(value interface{}) *Nav {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Nav) Staticinputclassname(value interface{}) *Nav {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Nav) Staticschema(value interface{}) *Nav {
    a.Set("staticSchema", value)
    return a
}

/**
 * 主题配色 默认light
 * 可选值: light | dark
 */
func (a *Nav) Themecolor(value interface{}) *Nav {
    a.Set("themeColor", value)
    return a
}

/**
 * 子菜单项展开浮层样式
 */
func (a *Nav) Popupclassname(value interface{}) *Nav {
    a.Set("popupClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Nav) Disabled(value interface{}) *Nav {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Nav) Static(value interface{}) *Nav {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Nav) Usemobileui(value interface{}) *Nav {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 横向导航时自动收纳配置
 */
func (a *Nav) Overflow(value interface{}) *Nav {
    a.Set("overflow", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Nav) Staticlabelclassname(value interface{}) *Nav {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以通过 API 拉取。
 */
func (a *Nav) Source(value interface{}) *Nav {
    a.Set("source", value)
    return a
}

/**
 * true 为垂直排列，false 为水平排列类似如 tabs。
 * 可选值: true | false
 */
func (a *Nav) Stacked(value interface{}) *Nav {
    a.Set("stacked", value)
    return a
}

/**
 * 是否开启搜索
 */
func (a *Nav) Searchable(value interface{}) *Nav {
    a.Set("searchable", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Nav) Disabledon(value interface{}) *Nav {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Nav) Hidden(value interface{}) *Nav {
    a.Set("hidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Nav) Onevent(value interface{}) *Nav {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Nav) Staticon(value interface{}) *Nav {
    a.Set("staticOn", value)
    return a
}

/**
 * 角标
 */
func (a *Nav) Badge(value interface{}) *Nav {
    a.Set("badge", value)
    return a
}

/**
 * 最多展示多少层级
 */
func (a *Nav) Level(value interface{}) *Nav {
    a.Set("level", value)
    return a
}

/**
 * 手风琴展开 仅垂直inline模式支持
 */
func (a *Nav) Accordion(value interface{}) *Nav {
    a.Set("accordion", value)
    return a
}

/**
 */
func (a *Nav) Testidbuilder(value interface{}) *Nav {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 懒加载 api，如果不配置复用 source 接口。
 */
func (a *Nav) Deferapi(value interface{}) *Nav {
    a.Set("deferApi", value)
    return a
}

/**
 * 可拖拽
 */
func (a *Nav) Draggable(value interface{}) *Nav {
    a.Set("draggable", value)
    return a
}

/**
 * 搜索框相关配置
 */
func (a *Nav) Searchconfig(value interface{}) *Nav {
    a.Set("searchConfig", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Nav) Id(value interface{}) *Nav {
    a.Set("id", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Nav) Staticplaceholder(value interface{}) *Nav {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 仅允许同层级拖拽
 */
func (a *Nav) Dragonsamelevel(value interface{}) *Nav {
    a.Set("dragOnSameLevel", value)
    return a
}

/**
 * 自定义展开图标位置 默认在前面 before after
 */
func (a *Nav) Expandposition(value interface{}) *Nav {
    a.Set("expandPosition", value)
    return a
}

/**
 * 链接地址集合
 */
func (a *Nav) Links(value interface{}) *Nav {
    a.Set("links", value)
    return a
}

/**
 */
func (a *Nav) Indentsize(value interface{}) *Nav {
    a.Set("indentSize", value)
    return a
}

/**
 * 保存排序的 api
 */
func (a *Nav) Saveorderapi(value interface{}) *Nav {
    a.Set("saveOrderApi", value)
    return a
}

/**
 * 控制菜单缩起
 */
func (a *Nav) Collapsed(value interface{}) *Nav {
    a.Set("collapsed", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Nav) Classname(value interface{}) *Nav {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Nav) Visible(value interface{}) *Nav {
    a.Set("visible", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Nav) Editorsetting(value interface{}) *Nav {
    a.Set("editorSetting", value)
    return a
}

/**
 * 更多操作菜单列表
 */
func (a *Nav) Itemactions(value interface{}) *Nav {
    a.Set("itemActions", value)
    return a
}

/**
 * 控制仅展示指定key菜单下的子菜单项
 */
func (a *Nav) Showkey(value interface{}) *Nav {
    a.Set("showKey", value)
    return a
}

/**
 * 自定义展开图标
 */
func (a *Nav) Expandicon(value interface{}) *Nav {
    a.Set("expandIcon", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Nav) Visibleon(value interface{}) *Nav {
    a.Set("visibleOn", value)
    return a
}

/**
 * 组件样式
 */
func (a *Nav) Style(value interface{}) *Nav {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Nav) Testid(value interface{}) *Nav {
    a.Set("testid", value)
    return a
}

/**
 * 默认展开层级 小于等于该层数的节点默认全部打开
 */
func (a *Nav) Defaultopenlevel(value interface{}) *Nav {
    a.Set("defaultOpenLevel", value)
    return a
}
