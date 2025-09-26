package renderers

type FileControl struct {
	*BaseRenderer
}

func NewFileControl() *FileControl {
	a := &FileControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-file")
	return a
}

func (a *FileControl) Set(name string, value interface{}) *FileControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *FileControl) Accept(value interface{}) *FileControl {
	a.Set("accept", value)
	return a
}

func (a *FileControl) AsBase64(value interface{}) *FileControl {
	a.Set("asBase64", value)
	return a
}

func (a *FileControl) AsBlob(value interface{}) *FileControl {
	a.Set("asBlob", value)
	return a
}

func (a *FileControl) AutoFill(value interface{}) *FileControl {
	a.Set("autoFill", value)
	return a
}

func (a *FileControl) AutoUpload(value interface{}) *FileControl {
	a.Set("autoUpload", value)
	return a
}

func (a *FileControl) BtnClassName(value interface{}) *FileControl {
	a.Set("btnClassName", value)
	return a
}

func (a *FileControl) BtnLabel(value interface{}) *FileControl {
	a.Set("btnLabel", value)
	return a
}

func (a *FileControl) BtnUploadClassName(value interface{}) *FileControl {
	a.Set("btnUploadClassName", value)
	return a
}

func (a *FileControl) Capture(value interface{}) *FileControl {
	a.Set("capture", value)
	return a
}

func (a *FileControl) ChunkApi(value interface{}) *FileControl {
	a.Set("chunkApi", value)
	return a
}

func (a *FileControl) ChunkSize(value interface{}) *FileControl {
	a.Set("chunkSize", value)
	return a
}

func (a *FileControl) ClassName(value interface{}) *FileControl {
	a.Set("className", value)
	return a
}

func (a *FileControl) ClearValueOnHidden(value interface{}) *FileControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *FileControl) Concurrency(value interface{}) *FileControl {
	a.Set("concurrency", value)
	return a
}

func (a *FileControl) Delimiter(value interface{}) *FileControl {
	a.Set("delimiter", value)
	return a
}

func (a *FileControl) Desc(value interface{}) *FileControl {
	a.Set("desc", value)
	return a
}

func (a *FileControl) Description(value interface{}) *FileControl {
	a.Set("description", value)
	return a
}

func (a *FileControl) DescriptionClassName(value interface{}) *FileControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *FileControl) Disabled(value interface{}) *FileControl {
	a.Set("disabled", value)
	return a
}

func (a *FileControl) DisabledOn(value interface{}) *FileControl {
	a.Set("disabledOn", value)
	return a
}

func (a *FileControl) DocumentLink(value interface{}) *FileControl {
	a.Set("documentLink", value)
	return a
}

func (a *FileControl) Documentation(value interface{}) *FileControl {
	a.Set("documentation", value)
	return a
}

func (a *FileControl) DownloadUrl(value interface{}) *FileControl {
	a.Set("downloadUrl", value)
	return a
}

func (a *FileControl) Drag(value interface{}) *FileControl {
	a.Set("drag", value)
	return a
}

func (a *FileControl) EditorSetting(value interface{}) *FileControl {
	a.Set("editorSetting", value)
	return a
}

func (a *FileControl) ExtraName(value interface{}) *FileControl {
	a.Set("extraName", value)
	return a
}

func (a *FileControl) ExtractValue(value interface{}) *FileControl {
	a.Set("extractValue", value)
	return a
}

func (a *FileControl) FileField(value interface{}) *FileControl {
	a.Set("fileField", value)
	return a
}

func (a *FileControl) FinishChunkApi(value interface{}) *FileControl {
	a.Set("finishChunkApi", value)
	return a
}

func (a *FileControl) Hidden(value interface{}) *FileControl {
	a.Set("hidden", value)
	return a
}

func (a *FileControl) HiddenOn(value interface{}) *FileControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *FileControl) HideUploadButton(value interface{}) *FileControl {
	a.Set("hideUploadButton", value)
	return a
}

func (a *FileControl) Hint(value interface{}) *FileControl {
	a.Set("hint", value)
	return a
}

func (a *FileControl) Horizontal(value interface{}) *FileControl {
	a.Set("horizontal", value)
	return a
}

func (a *FileControl) Id(value interface{}) *FileControl {
	a.Set("id", value)
	return a
}

func (a *FileControl) InitAutoFill(value interface{}) *FileControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *FileControl) Inline(value interface{}) *FileControl {
	a.Set("inline", value)
	return a
}

func (a *FileControl) InputClassName(value interface{}) *FileControl {
	a.Set("inputClassName", value)
	return a
}

func (a *FileControl) InvalidSizeMessage(value interface{}) *FileControl {
	a.Set("invalidSizeMessage", value)
	return a
}

func (a *FileControl) InvalidTypeMessage(value interface{}) *FileControl {
	a.Set("invalidTypeMessage", value)
	return a
}

func (a *FileControl) JoinValues(value interface{}) *FileControl {
	a.Set("joinValues", value)
	return a
}

func (a *FileControl) Label(value interface{}) *FileControl {
	a.Set("label", value)
	return a
}

func (a *FileControl) LabelAlign(value interface{}) *FileControl {
	a.Set("labelAlign", value)
	return a
}

func (a *FileControl) LabelClassName(value interface{}) *FileControl {
	a.Set("labelClassName", value)
	return a
}

func (a *FileControl) LabelRemark(value interface{}) *FileControl {
	a.Set("labelRemark", value)
	return a
}

func (a *FileControl) LabelWidth(value interface{}) *FileControl {
	a.Set("labelWidth", value)
	return a
}

func (a *FileControl) MaxLength(value interface{}) *FileControl {
	a.Set("maxLength", value)
	return a
}

func (a *FileControl) MaxSize(value interface{}) *FileControl {
	a.Set("maxSize", value)
	return a
}

func (a *FileControl) Mode(value interface{}) *FileControl {
	a.Set("mode", value)
	return a
}

func (a *FileControl) Multiple(value interface{}) *FileControl {
	a.Set("multiple", value)
	return a
}

func (a *FileControl) Name(value interface{}) *FileControl {
	a.Set("name", value)
	return a
}

func (a *FileControl) NameField(value interface{}) *FileControl {
	a.Set("nameField", value)
	return a
}

func (a *FileControl) OnEvent(value interface{}) *FileControl {
	a.Set("onEvent", value)
	return a
}

func (a *FileControl) Placeholder(value interface{}) *FileControl {
	a.Set("placeholder", value)
	return a
}

func (a *FileControl) ReadOnly(value interface{}) *FileControl {
	a.Set("readOnly", value)
	return a
}

func (a *FileControl) ReadOnlyOn(value interface{}) *FileControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *FileControl) Receiver(value interface{}) *FileControl {
	a.Set("receiver", value)
	return a
}

func (a *FileControl) Remark(value interface{}) *FileControl {
	a.Set("remark", value)
	return a
}

func (a *FileControl) Required(value interface{}) *FileControl {
	a.Set("required", value)
	return a
}

func (a *FileControl) ResetValue(value interface{}) *FileControl {
	a.Set("resetValue", value)
	return a
}

func (a *FileControl) Row(value interface{}) *FileControl {
	a.Set("row", value)
	return a
}

func (a *FileControl) SaveImmediately(value interface{}) *FileControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *FileControl) Size(value interface{}) *FileControl {
	a.Set("size", value)
	return a
}

func (a *FileControl) StartChunkApi(value interface{}) *FileControl {
	a.Set("startChunkApi", value)
	return a
}

func (a *FileControl) StateTextMap(value interface{}) *FileControl {
	a.Set("stateTextMap", value)
	return a
}

func (a *FileControl) Static(value interface{}) *FileControl {
	a.Set("static", value)
	return a
}

func (a *FileControl) StaticClassName(value interface{}) *FileControl {
	a.Set("staticClassName", value)
	return a
}

func (a *FileControl) StaticInputClassName(value interface{}) *FileControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *FileControl) StaticLabelClassName(value interface{}) *FileControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *FileControl) StaticOn(value interface{}) *FileControl {
	a.Set("staticOn", value)
	return a
}

func (a *FileControl) StaticPlaceholder(value interface{}) *FileControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *FileControl) StaticSchema(value interface{}) *FileControl {
	a.Set("staticSchema", value)
	return a
}

func (a *FileControl) Style(value interface{}) *FileControl {
	a.Set("style", value)
	return a
}

func (a *FileControl) SubmitOnChange(value interface{}) *FileControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *FileControl) TemplateUrl(value interface{}) *FileControl {
	a.Set("templateUrl", value)
	return a
}

func (a *FileControl) TestIdBuilder(value interface{}) *FileControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *FileControl) Type(value interface{}) *FileControl {
	a.Set("type", value)
	return a
}

func (a *FileControl) UrlField(value interface{}) *FileControl {
	a.Set("urlField", value)
	return a
}

func (a *FileControl) UseChunk(value interface{}) *FileControl {
	a.Set("useChunk", value)
	return a
}

func (a *FileControl) UseMobileUI(value interface{}) *FileControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *FileControl) ValidateApi(value interface{}) *FileControl {
	a.Set("validateApi", value)
	return a
}

func (a *FileControl) ValidateOnChange(value interface{}) *FileControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *FileControl) ValidationErrors(value interface{}) *FileControl {
	a.Set("validationErrors", value)
	return a
}

func (a *FileControl) Validations(value interface{}) *FileControl {
	a.Set("validations", value)
	return a
}

func (a *FileControl) Value(value interface{}) *FileControl {
	a.Set("value", value)
	return a
}

func (a *FileControl) ValueField(value interface{}) *FileControl {
	a.Set("valueField", value)
	return a
}

func (a *FileControl) Visible(value interface{}) *FileControl {
	a.Set("visible", value)
	return a
}

func (a *FileControl) VisibleOn(value interface{}) *FileControl {
	a.Set("visibleOn", value)
	return a
}

func (a *FileControl) Width(value interface{}) *FileControl {
	a.Set("width", value)
	return a
}
