package renderers


/**
 * tpl 渲染器

 * @author wcz0
 * @version 6.2.2
 */
type Tpl struct {
	*BaseRenderer
}

func NewTpl() *Tpl {
    a := &Tpl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "tpl")
    return a
}


func (a *Tpl) Set(name string, value interface{}) *Tpl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否禁用表达式
 */
func (a *Tpl) DisabledOn(value interface{}) *Tpl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Tpl) StaticLabelClassName(value interface{}) *Tpl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Tpl) UseMobileUI(value interface{}) *Tpl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Tpl) HiddenOn(value interface{}) *Tpl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Tpl) StaticPlaceholder(value interface{}) *Tpl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Tpl) StaticClassName(value interface{}) *Tpl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Tpl) StaticInputClassName(value interface{}) *Tpl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Tpl) TestIdBuilder(value interface{}) *Tpl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 */
func (a *Tpl) Tpl(value interface{}) *Tpl {
    a.Set("tpl", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Tpl) Id(value interface{}) *Tpl {
    a.Set("id", value)
    return a
}

/**
 * 是否内联显示？
 */
func (a *Tpl) Inline(value interface{}) *Tpl {
    a.Set("inline", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Tpl) Disabled(value interface{}) *Tpl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Tpl) StaticOn(value interface{}) *Tpl {
    a.Set("staticOn", value)
    return a
}

/**
 * 角标
 */
func (a *Tpl) Badge(value interface{}) *Tpl {
    a.Set("badge", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Tpl) Hidden(value interface{}) *Tpl {
    a.Set("hidden", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Tpl) EditorSetting(value interface{}) *Tpl {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *Tpl) Html(value interface{}) *Tpl {
    a.Set("html", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Tpl) ClassName(value interface{}) *Tpl {
    a.Set("className", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Tpl) OnEvent(value interface{}) *Tpl {
    a.Set("onEvent", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Tpl) Style(value interface{}) *Tpl {
    a.Set("style", value)
    return a
}

/**
 * 指定为模板渲染器。文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template
 * 可选值: tpl | html
 */
func (a *Tpl) Type(value interface{}) *Tpl {
    a.Set("type", value)
    return a
}

/**
 */
func (a *Tpl) Text(value interface{}) *Tpl {
    a.Set("text", value)
    return a
}

/**
 * 是否显示
 */
func (a *Tpl) Visible(value interface{}) *Tpl {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Tpl) VisibleOn(value interface{}) *Tpl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Tpl) Static(value interface{}) *Tpl {
    a.Set("static", value)
    return a
}

/**
 */
func (a *Tpl) StaticSchema(value interface{}) *Tpl {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Tpl) Testid(value interface{}) *Tpl {
    a.Set("testid", value)
    return a
}

/**
 */
func (a *Tpl) Raw(value interface{}) *Tpl {
    a.Set("raw", value)
    return a
}

/**
 * 标签类型
 */
func (a *Tpl) WrapperComponent(value interface{}) *Tpl {
    a.Set("wrapperComponent", value)
    return a
}

/**
 */
func (a *Tpl) TestidBuilder(value interface{}) *Tpl {
    a.Set("testidBuilder", value)
    return a
}
