package renderers


/**
 * Link 链接展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/link

 * @author wcz0
 * @version 6.2.2
 */
type Link struct {
	*BaseRenderer
}

func NewLink() *Link {
    a := &Link{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "link")
    return a
}


func (a *Link) Set(name string, value interface{}) *Link {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 链接地址
 */
func (a *Link) Href(value interface{}) *Link {
    a.Set("href", value)
    return a
}

/**
 * 角标
 */
func (a *Link) Badge(value interface{}) *Link {
    a.Set("badge", value)
    return a
}

/**
 * a标签原生target属性
 */
func (a *Link) HtmlTarget(value interface{}) *Link {
    a.Set("htmlTarget", value)
    return a
}

/**
 * 组件样式
 */
func (a *Link) Style(value interface{}) *Link {
    a.Set("style", value)
    return a
}

/**
 * 是否新窗口打开。
 */
func (a *Link) Blank(value interface{}) *Link {
    a.Set("blank", value)
    return a
}

/**
 * 链接内容，如果不配置将显示链接地址。
 */
func (a *Link) Body(value interface{}) *Link {
    a.Set("body", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Link) VisibleOn(value interface{}) *Link {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Link) StaticClassName(value interface{}) *Link {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Link) StaticSchema(value interface{}) *Link {
    a.Set("staticSchema", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Link) EditorSetting(value interface{}) *Link {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Link) Testid(value interface{}) *Link {
    a.Set("testid", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Link) Id(value interface{}) *Link {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Link) OnEvent(value interface{}) *Link {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Link) Static(value interface{}) *Link {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Link) StaticOn(value interface{}) *Link {
    a.Set("staticOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Link) UseMobileUI(value interface{}) *Link {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 指定为 link 链接展示控件
 */
func (a *Link) Type(value interface{}) *Link {
    a.Set("type", value)
    return a
}

/**
 * 图标
 */
func (a *Link) Icon(value interface{}) *Link {
    a.Set("icon", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Link) Hidden(value interface{}) *Link {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Link) HiddenOn(value interface{}) *Link {
    a.Set("hiddenOn", value)
    return a
}

/**
 */
func (a *Link) TestIdBuilder(value interface{}) *Link {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Link) StaticLabelClassName(value interface{}) *Link {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Link) StaticInputClassName(value interface{}) *Link {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 右侧图标
 */
func (a *Link) RightIcon(value interface{}) *Link {
    a.Set("rightIcon", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Link) DisabledOn(value interface{}) *Link {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Link) StaticPlaceholder(value interface{}) *Link {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 是否显示
 */
func (a *Link) Visible(value interface{}) *Link {
    a.Set("visible", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Link) ClassName(value interface{}) *Link {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Link) Disabled(value interface{}) *Link {
    a.Set("disabled", value)
    return a
}
