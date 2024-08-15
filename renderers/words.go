package renderers


/**
 * Words

 * @author wcz0
 * @version 6.2.2
 */
type Words struct {
	*BaseRenderer
}

func NewWords() *Words {
    a := &Words{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *Words) Set(name string, value interface{}) *Words {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "words")
    return a
}

/**
 * 是否禁用
 */
func (a *Words) Disabled(value interface{}) *Words {
    a.Set("disabled", value)
    return a
}

/**
 * 展示文字
 */
func (a *Words) Expendbuttontext(value interface{}) *Words {
    a.Set("expendButtonText", value)
    return a
}

/**
 * 收起文字
 */
func (a *Words) Collapsebuttontext(value interface{}) *Words {
    a.Set("collapseButtonText", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Words) Classname(value interface{}) *Words {
    a.Set("className", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Words) Staticplaceholder(value interface{}) *Words {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *Words) Staticschema(value interface{}) *Words {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Words) Testidbuilder(value interface{}) *Words {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 分割符
 */
func (a *Words) Delimiter(value interface{}) *Words {
    a.Set("delimiter", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Words) Disabledon(value interface{}) *Words {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Words) Staticclassname(value interface{}) *Words {
    a.Set("staticClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Words) Staticinputclassname(value interface{}) *Words {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Words) Staticon(value interface{}) *Words {
    a.Set("staticOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Words) Editorsetting(value interface{}) *Words {
    a.Set("editorSetting", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Words) Staticlabelclassname(value interface{}) *Words {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Words) Usemobileui(value interface{}) *Words {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Words) Type(value interface{}) *Words {
    a.Set("type", value)
    return a
}

/**
 * 展示文字
 */
func (a *Words) Collapsebutton(value interface{}) *Words {
    a.Set("collapseButton", value)
    return a
}

/**
 * tags数据
 */
func (a *Words) Words(value interface{}) *Words {
    a.Set("words", value)
    return a
}

/**
 * useTag 当数据是数组时，是否使用tag的方式展示
 */
func (a *Words) Intag(value interface{}) *Words {
    a.Set("inTag", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Words) Static(value interface{}) *Words {
    a.Set("static", value)
    return a
}

/**
 * 组件样式
 */
func (a *Words) Style(value interface{}) *Words {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Words) Testid(value interface{}) *Words {
    a.Set("testid", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Words) Onevent(value interface{}) *Words {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Words) Hiddenon(value interface{}) *Words {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *Words) Visible(value interface{}) *Words {
    a.Set("visible", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Words) Id(value interface{}) *Words {
    a.Set("id", value)
    return a
}

/**
 * 展示文字
 */
func (a *Words) Expendbutton(value interface{}) *Words {
    a.Set("expendButton", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Words) Hidden(value interface{}) *Words {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Words) Visibleon(value interface{}) *Words {
    a.Set("visibleOn", value)
    return a
}

/**
 * 展示限制, 为0时也无限制
 */
func (a *Words) Limit(value interface{}) *Words {
    a.Set("limit", value)
    return a
}
