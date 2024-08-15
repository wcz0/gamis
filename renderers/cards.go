package renderers


/**
 * Cards 卡片集合渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/card

 * @author wcz0
 * @version 6.2.2
 */
type Cards struct {
	*BaseRenderer
}

func NewCards() *Cards {
    a := &Cards{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "cards")
    return a
}


func (a *Cards) Set(name string, value interface{}) *Cards {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 可以用来作为值的字段
 */
func (a *Cards) Valuefield(value interface{}) *Cards {
    a.Set("valueField", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Cards) Classname(value interface{}) *Cards {
    a.Set("className", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Cards) Staticclassname(value interface{}) *Cards {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Cards) Staticon(value interface{}) *Cards {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否固顶
 */
func (a *Cards) Affixheader(value interface{}) *Cards {
    a.Set("affixHeader", value)
    return a
}

/**
 * 是否固底
 */
func (a *Cards) Affixfooter(value interface{}) *Cards {
    a.Set("affixFooter", value)
    return a
}

/**
 * 是否为瀑布流布局？
 */
func (a *Cards) Masonrylayout(value interface{}) *Cards {
    a.Set("masonryLayout", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Cards) Visibleon(value interface{}) *Cards {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Cards) Onevent(value interface{}) *Cards {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *Cards) Staticschema(value interface{}) *Cards {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Cards) Editorsetting(value interface{}) *Cards {
    a.Set("editorSetting", value)
    return a
}

/**
 * 头部 CSS 类名
 */
func (a *Cards) Headerclassname(value interface{}) *Cards {
    a.Set("headerClassName", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *Cards) Showfooter(value interface{}) *Cards {
    a.Set("showFooter", value)
    return a
}

/**
 * 是否隐藏勾选框
 */
func (a *Cards) Hidechecktoggler(value interface{}) *Cards {
    a.Set("hideCheckToggler", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Cards) Disabled(value interface{}) *Cards {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Cards) Hidden(value interface{}) *Cards {
    a.Set("hidden", value)
    return a
}

/**
 * 组件样式
 */
func (a *Cards) Style(value interface{}) *Cards {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Cards) Testid(value interface{}) *Cards {
    a.Set("testid", value)
    return a
}

/**
 * 顶部区域
 */
func (a *Cards) Header(value interface{}) *Cards {
    a.Set("header", value)
    return a
}

/**
 * 配置某项是否可拖拽排序，前提是要开启拖拽功能
 */
func (a *Cards) Itemdraggableon(value interface{}) *Cards {
    a.Set("itemDraggableOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Cards) Id(value interface{}) *Cards {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Cards) Staticlabelclassname(value interface{}) *Cards {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Cards) Card(value interface{}) *Cards {
    a.Set("card", value)
    return a
}

/**
 * 无数据提示
 */
func (a *Cards) Placeholder(value interface{}) *Cards {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否显示头部
 */
func (a *Cards) Showheader(value interface{}) *Cards {
    a.Set("showHeader", value)
    return a
}

/**
 * 点击卡片的时候是否勾选卡片。
 */
func (a *Cards) Checkonitemclick(value interface{}) *Cards {
    a.Set("checkOnItemClick", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Cards) Hiddenon(value interface{}) *Cards {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 指定为 cards 类型
 */
func (a *Cards) Type(value interface{}) *Cards {
    a.Set("type", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Cards) Disabledon(value interface{}) *Cards {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Cards) Staticinputclassname(value interface{}) *Cards {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 底部 CSS 类名
 */
func (a *Cards) Footerclassname(value interface{}) *Cards {
    a.Set("footerClassName", value)
    return a
}

/**
 */
func (a *Cards) Loadingconfig(value interface{}) *Cards {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Cards) Staticplaceholder(value interface{}) *Cards {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Cards) Usemobileui(value interface{}) *Cards {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 卡片 CSS 类名
 */
func (a *Cards) Itemclassname(value interface{}) *Cards {
    a.Set("itemClassName", value)
    return a
}

/**
 * 配置某项是否可以点选
 */
func (a *Cards) Itemcheckableon(value interface{}) *Cards {
    a.Set("itemCheckableOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Cards) Visible(value interface{}) *Cards {
    a.Set("visible", value)
    return a
}

/**
 * 数据源: 绑定当前环境变量
 */
func (a *Cards) Source(value interface{}) *Cards {
    a.Set("source", value)
    return a
}

/**
 * 标题
 */
func (a *Cards) Title(value interface{}) *Cards {
    a.Set("title", value)
    return a
}

/**
 * 底部区域
 */
func (a *Cards) Footer(value interface{}) *Cards {
    a.Set("footer", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Cards) Static(value interface{}) *Cards {
    a.Set("static", value)
    return a
}

/**
 */
func (a *Cards) Testidbuilder(value interface{}) *Cards {
    a.Set("testIdBuilder", value)
    return a
}
