package renderers

type Calendar struct {
	*BaseRenderer
}

func NewCalendar() *Calendar {
	a := &Calendar{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "calendar")
	return a
}

func (a *Calendar) Set(name string, value interface{}) *Calendar {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Calendar) ClassName(value interface{}) *Calendar {
	a.Set("className", value)
	return a
}

func (a *Calendar) Disabled(value interface{}) *Calendar {
	a.Set("disabled", value)
	return a
}

func (a *Calendar) DisabledOn(value interface{}) *Calendar {
	a.Set("disabledOn", value)
	return a
}

func (a *Calendar) EditorSetting(value interface{}) *Calendar {
	a.Set("editorSetting", value)
	return a
}

func (a *Calendar) Hidden(value interface{}) *Calendar {
	a.Set("hidden", value)
	return a
}

func (a *Calendar) HiddenOn(value interface{}) *Calendar {
	a.Set("hiddenOn", value)
	return a
}

func (a *Calendar) Id(value interface{}) *Calendar {
	a.Set("id", value)
	return a
}

func (a *Calendar) LargeMode(value interface{}) *Calendar {
	a.Set("largeMode", value)
	return a
}

func (a *Calendar) OnEvent(value interface{}) *Calendar {
	a.Set("onEvent", value)
	return a
}

func (a *Calendar) ScheduleAction(value interface{}) *Calendar {
	a.Set("scheduleAction", value)
	return a
}

func (a *Calendar) ScheduleClassNames(value interface{}) *Calendar {
	a.Set("scheduleClassNames", value)
	return a
}

func (a *Calendar) Schedules(value interface{}) *Calendar {
	a.Set("schedules", value)
	return a
}

func (a *Calendar) Static(value interface{}) *Calendar {
	a.Set("static", value)
	return a
}

func (a *Calendar) StaticClassName(value interface{}) *Calendar {
	a.Set("staticClassName", value)
	return a
}

func (a *Calendar) StaticInputClassName(value interface{}) *Calendar {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Calendar) StaticLabelClassName(value interface{}) *Calendar {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Calendar) StaticOn(value interface{}) *Calendar {
	a.Set("staticOn", value)
	return a
}

func (a *Calendar) StaticPlaceholder(value interface{}) *Calendar {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Calendar) StaticSchema(value interface{}) *Calendar {
	a.Set("staticSchema", value)
	return a
}

func (a *Calendar) Style(value interface{}) *Calendar {
	a.Set("style", value)
	return a
}

func (a *Calendar) TestIdBuilder(value interface{}) *Calendar {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Calendar) Testid(value interface{}) *Calendar {
	a.Set("testid", value)
	return a
}

func (a *Calendar) TodayActiveStyle(value interface{}) *Calendar {
	a.Set("todayActiveStyle", value)
	return a
}

func (a *Calendar) Type(value interface{}) *Calendar {
	a.Set("type", value)
	return a
}

func (a *Calendar) UseMobileUI(value interface{}) *Calendar {
	a.Set("useMobileUI", value)
	return a
}

func (a *Calendar) Visible(value interface{}) *Calendar {
	a.Set("visible", value)
	return a
}

func (a *Calendar) VisibleOn(value interface{}) *Calendar {
	a.Set("visibleOn", value)
	return a
}
