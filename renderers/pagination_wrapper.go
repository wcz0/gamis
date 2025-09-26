package renderers


/**
 * 分页容器功能性渲染器。详情请见：https://aisuda.bce.baidu.com/amis/zh-CN/components/pagination-wrapper

 * @author wcz0
 * @version 6.2.2
 */
type PaginationWrapper struct {
	*BaseRenderer
}

func NewPaginationWrapper() *PaginationWrapper {
    a := &PaginationWrapper{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "pagination-wrapper")
    return a
}


func (a *PaginationWrapper) Set(name string, value interface{}) *PaginationWrapper {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * static
 */
func (a *PaginationWrapper) Static(value interface{}) *PaginationWrapper {
    a.Set("static", value)
    return a
}

/**
 * staticClassName
 */
func (a *PaginationWrapper) StaticClassName(value interface{}) *PaginationWrapper {
    a.Set("staticClassName", value)
    return a
}

/**
 * outputName
 */
func (a *PaginationWrapper) OutputName(value interface{}) *PaginationWrapper {
    a.Set("outputName", value)
    return a
}

/**
 * disabled
 */
func (a *PaginationWrapper) Disabled(value interface{}) *PaginationWrapper {
    a.Set("disabled", value)
    return a
}

/**
 * disabledOn
 */
func (a *PaginationWrapper) DisabledOn(value interface{}) *PaginationWrapper {
    a.Set("disabledOn", value)
    return a
}

/**
 * hiddenOn
 */
func (a *PaginationWrapper) HiddenOn(value interface{}) *PaginationWrapper {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *PaginationWrapper) StaticInputClassName(value interface{}) *PaginationWrapper {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * staticSchema
 */
func (a *PaginationWrapper) StaticSchema(value interface{}) *PaginationWrapper {
    a.Set("staticSchema", value)
    return a
}

/**
 * className
 */
func (a *PaginationWrapper) ClassName(value interface{}) *PaginationWrapper {
    a.Set("className", value)
    return a
}

/**
 */
func (a *PaginationWrapper) Type(value interface{}) *PaginationWrapper {
    a.Set("type", value)
    return a
}

/**
 * testid
 */
func (a *PaginationWrapper) Testid(value interface{}) *PaginationWrapper {
    a.Set("testid", value)
    return a
}

/**
 * inputName
 */
func (a *PaginationWrapper) InputName(value interface{}) *PaginationWrapper {
    a.Set("inputName", value)
    return a
}

/**
 * maxButtons
 */
func (a *PaginationWrapper) MaxButtons(value interface{}) *PaginationWrapper {
    a.Set("maxButtons", value)
    return a
}

/**
 * position
 */
func (a *PaginationWrapper) Position(value interface{}) *PaginationWrapper {
    a.Set("position", value)
    return a
}

/**
 * useMobileUI
 */
func (a *PaginationWrapper) UseMobileUI(value interface{}) *PaginationWrapper {
    a.Set("useMobileUI", value)
    return a
}

/**
 * perPage
 */
func (a *PaginationWrapper) PerPage(value interface{}) *PaginationWrapper {
    a.Set("perPage", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *PaginationWrapper) StaticPlaceholder(value interface{}) *PaginationWrapper {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * visible
 */
func (a *PaginationWrapper) Visible(value interface{}) *PaginationWrapper {
    a.Set("visible", value)
    return a
}

/**
 * body
 */
func (a *PaginationWrapper) Body(value interface{}) *PaginationWrapper {
    a.Set("body", value)
    return a
}

/**
 * onEvent
 */
func (a *PaginationWrapper) OnEvent(value interface{}) *PaginationWrapper {
    a.Set("onEvent", value)
    return a
}

/**
 * staticOn
 */
func (a *PaginationWrapper) StaticOn(value interface{}) *PaginationWrapper {
    a.Set("staticOn", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *PaginationWrapper) StaticLabelClassName(value interface{}) *PaginationWrapper {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * style
 */
func (a *PaginationWrapper) Style(value interface{}) *PaginationWrapper {
    a.Set("style", value)
    return a
}

/**
 * id
 */
func (a *PaginationWrapper) Id(value interface{}) *PaginationWrapper {
    a.Set("id", value)
    return a
}

/**
 * hidden
 */
func (a *PaginationWrapper) Hidden(value interface{}) *PaginationWrapper {
    a.Set("hidden", value)
    return a
}

/**
 * visibleOn
 */
func (a *PaginationWrapper) VisibleOn(value interface{}) *PaginationWrapper {
    a.Set("visibleOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *PaginationWrapper) EditorSetting(value interface{}) *PaginationWrapper {
    a.Set("editorSetting", value)
    return a
}

/**
 * showPageInput
 */
func (a *PaginationWrapper) ShowPageInput(value interface{}) *PaginationWrapper {
    a.Set("showPageInput", value)
    return a
}
