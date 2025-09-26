package renderers

type Action struct {
	*BaseRenderer
}

func NewAction() *Action {
	a := &Action{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "action")
	return a
}

func (a *Action) Set(name string, value interface{}) *Action {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Action) ActionType(value interface{}) *Action {
	a.Set("actionType", value)
	return a
}

func (a *Action) ActiveClassName(value interface{}) *Action {
	a.Set("activeClassName", value)
	return a
}

func (a *Action) ActiveLevel(value interface{}) *Action {
	a.Set("activeLevel", value)
	return a
}

func (a *Action) Api(value interface{}) *Action {
	a.Set("api", value)
	return a
}

func (a *Action) ClassName(value interface{}) *Action {
	a.Set("className", value)
	return a
}

func (a *Action) Close(value interface{}) *Action {
	a.Set("close", value)
	return a
}

func (a *Action) ConfirmText(value interface{}) *Action {
	a.Set("confirmText", value)
	return a
}

func (a *Action) ConfirmTitle(value interface{}) *Action {
	a.Set("confirmTitle", value)
	return a
}

func (a *Action) DisabledTip(value interface{}) *Action {
	a.Set("disabledTip", value)
	return a
}

func (a *Action) Icon(value interface{}) *Action {
	a.Set("icon", value)
	return a
}

func (a *Action) IconClassName(value interface{}) *Action {
	a.Set("iconClassName", value)
	return a
}

func (a *Action) Label(value interface{}) *Action {
	a.Set("label", value)
	return a
}

func (a *Action) Level(value interface{}) *Action {
	a.Set("level", value)
	return a
}

func (a *Action) Link(value interface{}) *Action {
	a.Set("link", value)
	return a
}

func (a *Action) Reload(value interface{}) *Action {
	a.Set("reload", value)
	return a
}

func (a *Action) Required(value interface{}) *Action {
	a.Set("required", value)
	return a
}

func (a *Action) RightIcon(value interface{}) *Action {
	a.Set("rightIcon", value)
	return a
}

func (a *Action) RightIconClassName(value interface{}) *Action {
	a.Set("rightIconClassName", value)
	return a
}

func (a *Action) Size(value interface{}) *Action {
	a.Set("size", value)
	return a
}

func (a *Action) Tooltip(value interface{}) *Action {
	a.Set("tooltip", value)
	return a
}

func (a *Action) TooltipPlacement(value interface{}) *Action {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *Action) Type(value interface{}) *Action {
	a.Set("type", value)
	return a
}
