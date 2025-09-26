package renderers

type WangEditor struct {
	*BaseRenderer
}

func NewWangEditor() *WangEditor {
	a := &WangEditor{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "custom-wang-editor")
	return a
}

func (a *WangEditor) Set(name string, value interface{}) *WangEditor {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *WangEditor) AutoFill(value interface{}) *WangEditor {
	a.Set("autoFill", value)
	return a
}

func (a *WangEditor) AutoFocus(value interface{}) *WangEditor {
	a.Set("autoFocus", value)
	return a
}

func (a *WangEditor) ClassName(value interface{}) *WangEditor {
	a.Set("className", value)
	return a
}

func (a *WangEditor) ClearValueOnHidden(value interface{}) *WangEditor {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *WangEditor) Description(value interface{}) *WangEditor {
	a.Set("description", value)
	return a
}

func (a *WangEditor) Disabled(value interface{}) *WangEditor {
	a.Set("disabled", value)
	return a
}

func (a *WangEditor) DisabledOn(value interface{}) *WangEditor {
	a.Set("disabledOn", value)
	return a
}

func (a *WangEditor) ExcludeKeys(value interface{}) *WangEditor {
	a.Set("excludeKeys", value)
	return a
}

func (a *WangEditor) Height(value interface{}) *WangEditor {
	a.Set("height", value)
	return a
}

func (a *WangEditor) InsertKeys(value interface{}) *WangEditor {
	a.Set("insertKeys", value)
	return a
}

func (a *WangEditor) Label(value interface{}) *WangEditor {
	a.Set("label", value)
	return a
}

func (a *WangEditor) LabelAlign(value interface{}) *WangEditor {
	a.Set("labelAlign", value)
	return a
}

func (a *WangEditor) LabelRemark(value interface{}) *WangEditor {
	a.Set("labelRemark", value)
	return a
}

func (a *WangEditor) MaxLength(value interface{}) *WangEditor {
	a.Set("maxLength", value)
	return a
}

func (a *WangEditor) Name(value interface{}) *WangEditor {
	a.Set("name", value)
	return a
}

func (a *WangEditor) Placeholder(value interface{}) *WangEditor {
	a.Set("placeholder", value)
	return a
}

func (a *WangEditor) Remark(value interface{}) *WangEditor {
	a.Set("remark", value)
	return a
}

func (a *WangEditor) Required(value interface{}) *WangEditor {
	a.Set("required", value)
	return a
}

func (a *WangEditor) RequiredOn(value interface{}) *WangEditor {
	a.Set("requiredOn", value)
	return a
}

func (a *WangEditor) Static(value interface{}) *WangEditor {
	a.Set("static", value)
	return a
}

func (a *WangEditor) ToolbarKeys(value interface{}) *WangEditor {
	a.Set("toolbarKeys", value)
	return a
}

func (a *WangEditor) Type(value interface{}) *WangEditor {
	a.Set("type", value)
	return a
}

func (a *WangEditor) UploadImageMaxNumber(value interface{}) *WangEditor {
	a.Set("uploadImageMaxNumber", value)
	return a
}

func (a *WangEditor) UploadImageMaxSize(value interface{}) *WangEditor {
	a.Set("uploadImageMaxSize", value)
	return a
}

func (a *WangEditor) UploadImageServer(value interface{}) *WangEditor {
	a.Set("uploadImageServer", value)
	return a
}

func (a *WangEditor) UploadVideoMaxNumber(value interface{}) *WangEditor {
	a.Set("uploadVideoMaxNumber", value)
	return a
}

func (a *WangEditor) UploadVideoMaxSize(value interface{}) *WangEditor {
	a.Set("uploadVideoMaxSize", value)
	return a
}

func (a *WangEditor) UploadVideoServer(value interface{}) *WangEditor {
	a.Set("uploadVideoServer", value)
	return a
}

func (a *WangEditor) ValidateApi(value interface{}) *WangEditor {
	a.Set("validateApi", value)
	return a
}

func (a *WangEditor) ValidationErrors(value interface{}) *WangEditor {
	a.Set("validationErrors", value)
	return a
}

func (a *WangEditor) Validations(value interface{}) *WangEditor {
	a.Set("validations", value)
	return a
}

func (a *WangEditor) Value(value interface{}) *WangEditor {
	a.Set("value", value)
	return a
}

func (a *WangEditor) Visible(value interface{}) *WangEditor {
	a.Set("visible", value)
	return a
}

func (a *WangEditor) VisibleOn(value interface{}) *WangEditor {
	a.Set("visibleOn", value)
	return a
}
