package renderers


/**
 * Mapping 映射展示控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/mapping

 * @author wcz0
 * @version 6.2.2
 */
type Mapping struct {
	*BaseRenderer
}

func NewMapping() *Mapping {
    a := &Mapping{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "map")
    return a
}


func (a *Mapping) Set(name string, value interface{}) *Mapping {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticLabelClassName
 */
func (a *Mapping) StaticLabelClassName(value interface{}) *Mapping {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Mapping) StaticInputClassName(value interface{}) *Mapping {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * testid
 */
func (a *Mapping) Testid(value interface{}) *Mapping {
    a.Set("testid", value)
    return a
}

/**
 * name
 */
func (a *Mapping) Name(value interface{}) *Mapping {
    a.Set("name", value)
    return a
}

/**
 * map
 */
func (a *Mapping) Map(value interface{}) *Mapping {
    a.Set("map", value)
    return a
}

/**
 * staticOn
 */
func (a *Mapping) StaticOn(value interface{}) *Mapping {
    a.Set("staticOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *Mapping) EditorSetting(value interface{}) *Mapping {
    a.Set("editorSetting", value)
    return a
}

/**
 * 可选值: map | mapping
 */
func (a *Mapping) Type(value interface{}) *Mapping {
    a.Set("type", value)
    return a
}

/**
 * onEvent
 */
func (a *Mapping) OnEvent(value interface{}) *Mapping {
    a.Set("onEvent", value)
    return a
}

/**
 * valueField
 */
func (a *Mapping) ValueField(value interface{}) *Mapping {
    a.Set("valueField", value)
    return a
}

/**
 * labelField
 */
func (a *Mapping) LabelField(value interface{}) *Mapping {
    a.Set("labelField", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Mapping) StaticPlaceholder(value interface{}) *Mapping {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * disabled
 */
func (a *Mapping) Disabled(value interface{}) *Mapping {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *Mapping) DisabledOn(value interface{}) *Mapping {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Mapping) HiddenOn(value interface{}) *Mapping {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *Mapping) Id(value interface{}) *Mapping {
    a.Set("id", value)
    return a
}

/**
 * staticSchema
 */
func (a *Mapping) StaticSchema(value interface{}) *Mapping {
    a.Set("staticSchema", value)
    return a
}

/**
 * className
 */
func (a *Mapping) ClassName(value interface{}) *Mapping {
    a.Set("className", value)
    return a
}

/**
 * static
 */
func (a *Mapping) Static(value interface{}) *Mapping {
    a.Set("static", value)
    return a
}

/**
 * style
 */
func (a *Mapping) Style(value interface{}) *Mapping {
    a.Set("style", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Mapping) UseMobileUI(value interface{}) *Mapping {
    a.Set("useMobileUI", value)
    return a
}

/**
 * source
 */
func (a *Mapping) Source(value interface{}) *Mapping {
    a.Set("source", value)
    return a
}

/**
 * placeholder
 */
func (a *Mapping) Placeholder(value interface{}) *Mapping {
    a.Set("placeholder", value)
    return a
}

/**
 * hidden
 */
func (a *Mapping) Hidden(value interface{}) *Mapping {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *Mapping) Visible(value interface{}) *Mapping {
    a.Set("visible", value)
    return a
}

/**
 * staticClassName
 */
func (a *Mapping) StaticClassName(value interface{}) *Mapping {
    a.Set("staticClassName", value)
    return a
}

/**
 * itemSchema
 */
func (a *Mapping) ItemSchema(value interface{}) *Mapping {
    a.Set("itemSchema", value)
    return a
}

/**
 * visibleOn
 */
func (a *Mapping) VisibleOn(value interface{}) *Mapping {
    a.Set("visibleOn", value)
    return a
}
