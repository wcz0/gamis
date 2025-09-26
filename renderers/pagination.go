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

    a.Set("type", "pagination")
    return a
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
/**
 * static
 */
func (a *Pagination) Static(value interface{}) *Pagination {
    a.Set("static", value)
    return a
}

/**
 * staticClassName
 */
func (a *Pagination) StaticClassName(value interface{}) *Pagination {
    a.Set("staticClassName", value)
    return a
}

/**
 * perPage
 */
func (a *Pagination) PerPage(value interface{}) *Pagination {
    a.Set("perPage", value)
    return a
}

/**
 * visibleOn
 */
func (a *Pagination) VisibleOn(value interface{}) *Pagination {
    a.Set("visibleOn", value)
    return a
}

/**
 * onEvent
 */
func (a *Pagination) OnEvent(value interface{}) *Pagination {
    a.Set("onEvent", value)
    return a
}

/**
 * useMobileUI
 */
func (a *Pagination) UseMobileUI(value interface{}) *Pagination {
    a.Set("useMobileUI", value)
    return a
}

/**
 * mode
 */
func (a *Pagination) Mode(value interface{}) *Pagination {
    a.Set("mode", value)
    return a
}

/**
 * showPageInput
 */
func (a *Pagination) ShowPageInput(value interface{}) *Pagination {
    a.Set("showPageInput", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *Pagination) StaticLabelClassName(value interface{}) *Pagination {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *Pagination) StaticInputClassName(value interface{}) *Pagination {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * maxButtons
 */
func (a *Pagination) MaxButtons(value interface{}) *Pagination {
    a.Set("maxButtons", value)
    return a
}

/**
 * activePage
 */
func (a *Pagination) ActivePage(value interface{}) *Pagination {
    a.Set("activePage", value)
    return a
}

/**
 * showPerPage
 */
func (a *Pagination) ShowPerPage(value interface{}) *Pagination {
    a.Set("showPerPage", value)
    return a
}

/**
 * popOverContainerSelector
 */
func (a *Pagination) PopOverContainerSelector(value interface{}) *Pagination {
    a.Set("popOverContainerSelector", value)
    return a
}

/**
 * className
 */
func (a *Pagination) ClassName(value interface{}) *Pagination {
    a.Set("className", value)
    return a
}

/**
 * disabled
 */
func (a *Pagination) Disabled(value interface{}) *Pagination {
    a.Set("disabled", value)
    return a
}

/**
 * id
 */
func (a *Pagination) Id(value interface{}) *Pagination {
    a.Set("id", value)
    return a
}

/**
 */
func (a *Pagination) Type(value interface{}) *Pagination {
    a.Set("type", value)
    return a
}

/**
 * hiddenOn
 */
func (a *Pagination) HiddenOn(value interface{}) *Pagination {
    a.Set("hiddenOn", value)
    return a
}

/**
 * editorSetting
 */
func (a *Pagination) EditorSetting(value interface{}) *Pagination {
    a.Set("editorSetting", value)
    return a
}

/**
 * layout
 */
func (a *Pagination) Layout(value interface{}) *Pagination {
    a.Set("layout", value)
    return a
}

/**
 * total
 */
func (a *Pagination) Total(value interface{}) *Pagination {
    a.Set("total", value)
    return a
}

/**
 * staticOn
 */
func (a *Pagination) StaticOn(value interface{}) *Pagination {
    a.Set("staticOn", value)
    return a
}

/**
 * perPageAvailable
 */
func (a *Pagination) PerPageAvailable(value interface{}) *Pagination {
    a.Set("perPageAvailable", value)
    return a
}

/**
 * disabledOn
 */
func (a *Pagination) DisabledOn(value interface{}) *Pagination {
    a.Set("disabledOn", value)
    return a
}

/**
 * hidden
 */
func (a *Pagination) Hidden(value interface{}) *Pagination {
    a.Set("hidden", value)
    return a
}

/**
 * staticSchema
 */
func (a *Pagination) StaticSchema(value interface{}) *Pagination {
    a.Set("staticSchema", value)
    return a
}

/**
 * testid
 */
func (a *Pagination) Testid(value interface{}) *Pagination {
    a.Set("testid", value)
    return a
}

/**
 * hasNext
 */
func (a *Pagination) HasNext(value interface{}) *Pagination {
    a.Set("hasNext", value)
    return a
}

/**
 * visible
 */
func (a *Pagination) Visible(value interface{}) *Pagination {
    a.Set("visible", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *Pagination) StaticPlaceholder(value interface{}) *Pagination {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * style
 */
func (a *Pagination) Style(value interface{}) *Pagination {
    a.Set("style", value)
    return a
}
