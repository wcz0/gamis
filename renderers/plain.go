package renderers


/**
 * Plain 纯文本渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/plain

 * @author wcz0
 * @version 6.2.2
 */
type Plain struct {
	*BaseRenderer
}

func NewPlain() *Plain {
    a := &Plain{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "plain")
    return a
}


func (a *Plain) Set(name string, value interface{}) *Plain {
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
func (a *Plain) TestIdBuilder(value interface{}) *Plain {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Plain) StaticOn(value interface{}) *Plain {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Plain) StaticClassName(value interface{}) *Plain {
    a.Set("staticClassName", value)
    return a
}

/**
 * 是否内联显示？
 */
func (a *Plain) Inline(value interface{}) *Plain {
    a.Set("inline", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Plain) StaticLabelClassName(value interface{}) *Plain {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Plain) StaticInputClassName(value interface{}) *Plain {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Plain) Disabled(value interface{}) *Plain {
    a.Set("disabled", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Plain) DisabledOn(value interface{}) *Plain {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Plain) Static(value interface{}) *Plain {
    a.Set("static", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Plain) EditorSetting(value interface{}) *Plain {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Plain) Hidden(value interface{}) *Plain {
    a.Set("hidden", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Plain) OnEvent(value interface{}) *Plain {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Plain) StaticPlaceholder(value interface{}) *Plain {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Plain) UseMobileUI(value interface{}) *Plain {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Plain) Text(value interface{}) *Plain {
    a.Set("text", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Plain) HiddenOn(value interface{}) *Plain {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Plain) VisibleOn(value interface{}) *Plain {
    a.Set("visibleOn", value)
    return a
}

/**
 */
func (a *Plain) StaticSchema(value interface{}) *Plain {
    a.Set("staticSchema", value)
    return a
}

/**
 * 组件样式
 */
func (a *Plain) Style(value interface{}) *Plain {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Plain) Testid(value interface{}) *Plain {
    a.Set("testid", value)
    return a
}

/**
 */
func (a *Plain) Tpl(value interface{}) *Plain {
    a.Set("tpl", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Plain) Id(value interface{}) *Plain {
    a.Set("id", value)
    return a
}

/**
 * 指定为模板渲染器。文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template
 * 可选值: plain | text
 */
func (a *Plain) Type(value interface{}) *Plain {
    a.Set("type", value)
    return a
}

/**
 * 占位符
 */
func (a *Plain) Placeholder(value interface{}) *Plain {
    a.Set("placeholder", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Plain) ClassName(value interface{}) *Plain {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Plain) Visible(value interface{}) *Plain {
    a.Set("visible", value)
    return a
}
