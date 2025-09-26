package renderers

type Words struct {
	*BaseRenderer
}

func NewWords() *Words {
	a := &Words{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "words")
	return a
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

func (a *Words) ClassName(value interface{}) *Words {
	a.Set("className", value)
	return a
}

func (a *Words) CollapseButton(value interface{}) *Words {
	a.Set("collapseButton", value)
	return a
}

func (a *Words) CollapseButtonText(value interface{}) *Words {
	a.Set("collapseButtonText", value)
	return a
}

func (a *Words) Delimiter(value interface{}) *Words {
	a.Set("delimiter", value)
	return a
}

func (a *Words) Disabled(value interface{}) *Words {
	a.Set("disabled", value)
	return a
}

func (a *Words) DisabledOn(value interface{}) *Words {
	a.Set("disabledOn", value)
	return a
}

func (a *Words) EditorSetting(value interface{}) *Words {
	a.Set("editorSetting", value)
	return a
}

func (a *Words) ExpendButton(value interface{}) *Words {
	a.Set("expendButton", value)
	return a
}

func (a *Words) ExpendButtonText(value interface{}) *Words {
	a.Set("expendButtonText", value)
	return a
}

func (a *Words) Hidden(value interface{}) *Words {
	a.Set("hidden", value)
	return a
}

func (a *Words) HiddenOn(value interface{}) *Words {
	a.Set("hiddenOn", value)
	return a
}

func (a *Words) Id(value interface{}) *Words {
	a.Set("id", value)
	return a
}

func (a *Words) InTag(value interface{}) *Words {
	a.Set("inTag", value)
	return a
}

func (a *Words) LabelTpl(value interface{}) *Words {
	a.Set("labelTpl", value)
	return a
}

func (a *Words) Limit(value interface{}) *Words {
	a.Set("limit", value)
	return a
}

func (a *Words) OnEvent(value interface{}) *Words {
	a.Set("onEvent", value)
	return a
}

func (a *Words) Static(value interface{}) *Words {
	a.Set("static", value)
	return a
}

func (a *Words) StaticClassName(value interface{}) *Words {
	a.Set("staticClassName", value)
	return a
}

func (a *Words) StaticInputClassName(value interface{}) *Words {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Words) StaticLabelClassName(value interface{}) *Words {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Words) StaticOn(value interface{}) *Words {
	a.Set("staticOn", value)
	return a
}

func (a *Words) StaticPlaceholder(value interface{}) *Words {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Words) StaticSchema(value interface{}) *Words {
	a.Set("staticSchema", value)
	return a
}

func (a *Words) Style(value interface{}) *Words {
	a.Set("style", value)
	return a
}

func (a *Words) TestIdBuilder(value interface{}) *Words {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Words) Testid(value interface{}) *Words {
	a.Set("testid", value)
	return a
}

func (a *Words) Type(value interface{}) *Words {
	a.Set("type", value)
	return a
}

func (a *Words) UseMobileUI(value interface{}) *Words {
	a.Set("useMobileUI", value)
	return a
}

func (a *Words) Visible(value interface{}) *Words {
	a.Set("visible", value)
	return a
}

func (a *Words) VisibleOn(value interface{}) *Words {
	a.Set("visibleOn", value)
	return a
}

func (a *Words) Words(value interface{}) *Words {
	a.Set("words", value)
	return a
}
