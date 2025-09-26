package renderers


/**
 * Image 图片上传控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/image

 * @author wcz0
 * @version 6.2.2
 */
type ImageControl struct {
	*BaseRenderer
}

func NewImageControl() *ImageControl {
    a := &ImageControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-image")
    return a
}


func (a *ImageControl) Set(name string, value interface{}) *ImageControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * disabled
 */
func (a *ImageControl) Disabled(value interface{}) *ImageControl {
    a.Set("disabled", value)
    return a
}

/**
 * btnUploadClassName
 */
func (a *ImageControl) BtnUploadClassName(value interface{}) *ImageControl {
    a.Set("btnUploadClassName", value)
    return a
}

/**
 * fixedSize
 */
func (a *ImageControl) FixedSize(value interface{}) *ImageControl {
    a.Set("fixedSize", value)
    return a
}

/**
 * name
 */
func (a *ImageControl) Name(value interface{}) *ImageControl {
    a.Set("name", value)
    return a
}

/**
 * mode
 */
func (a *ImageControl) Mode(value interface{}) *ImageControl {
    a.Set("mode", value)
    return a
}

/**
 * editorSetting
 */
func (a *ImageControl) EditorSetting(value interface{}) *ImageControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * thumbMode
 */
func (a *ImageControl) ThumbMode(value interface{}) *ImageControl {
    a.Set("thumbMode", value)
    return a
}

/**
 * thumbRatio
 */
func (a *ImageControl) ThumbRatio(value interface{}) *ImageControl {
    a.Set("thumbRatio", value)
    return a
}

/**
 * initCrop
 */
func (a *ImageControl) InitCrop(value interface{}) *ImageControl {
    a.Set("initCrop", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ImageControl) Width(value interface{}) *ImageControl {
    a.Set("width", value)
    return a
}

/**
 * inputClassName
 */
func (a *ImageControl) InputClassName(value interface{}) *ImageControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * maxSize
 */
func (a *ImageControl) MaxSize(value interface{}) *ImageControl {
    a.Set("maxSize", value)
    return a
}

/**
 * capture
 */
func (a *ImageControl) Capture(value interface{}) *ImageControl {
    a.Set("capture", value)
    return a
}

/**
 * labelRemark
 */
func (a *ImageControl) LabelRemark(value interface{}) *ImageControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * extraName
 */
func (a *ImageControl) ExtraName(value interface{}) *ImageControl {
    a.Set("extraName", value)
    return a
}

/**
 * validateApi
 */
func (a *ImageControl) ValidateApi(value interface{}) *ImageControl {
    a.Set("validateApi", value)
    return a
}

/**
 * id
 */
func (a *ImageControl) Id(value interface{}) *ImageControl {
    a.Set("id", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *ImageControl) StaticLabelClassName(value interface{}) *ImageControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * crop
 */
func (a *ImageControl) Crop(value interface{}) *ImageControl {
    a.Set("crop", value)
    return a
}

/**
 * maxLength
 */
func (a *ImageControl) MaxLength(value interface{}) *ImageControl {
    a.Set("maxLength", value)
    return a
}

/**
 * multiple
 */
func (a *ImageControl) Multiple(value interface{}) *ImageControl {
    a.Set("multiple", value)
    return a
}

/**
 * draggableTip
 */
func (a *ImageControl) DraggableTip(value interface{}) *ImageControl {
    a.Set("draggableTip", value)
    return a
}

/**
 * useMobileUI
 */
func (a *ImageControl) UseMobileUI(value interface{}) *ImageControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * labelOverflow
 */
func (a *ImageControl) LabelOverflow(value interface{}) *ImageControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * placeholder
 */
func (a *ImageControl) Placeholder(value interface{}) *ImageControl {
    a.Set("placeholder", value)
    return a
}

/**
 * autoFill
 */
func (a *ImageControl) AutoFill(value interface{}) *ImageControl {
    a.Set("autoFill", value)
    return a
}

/**
 * onEvent
 */
func (a *ImageControl) OnEvent(value interface{}) *ImageControl {
    a.Set("onEvent", value)
    return a
}

/**
 * src
 */
func (a *ImageControl) Src(value interface{}) *ImageControl {
    a.Set("src", value)
    return a
}

/**
 * cropFormat
 */
func (a *ImageControl) CropFormat(value interface{}) *ImageControl {
    a.Set("cropFormat", value)
    return a
}

/**
 * delimiter
 */
func (a *ImageControl) Delimiter(value interface{}) *ImageControl {
    a.Set("delimiter", value)
    return a
}

/**
 * staticOn
 */
func (a *ImageControl) StaticOn(value interface{}) *ImageControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *ImageControl) StaticClassName(value interface{}) *ImageControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * style
 */
func (a *ImageControl) Style(value interface{}) *ImageControl {
    a.Set("style", value)
    return a
}

/**
 * label
 */
func (a *ImageControl) Label(value interface{}) *ImageControl {
    a.Set("label", value)
    return a
}

/**
 * hint
 */
func (a *ImageControl) Hint(value interface{}) *ImageControl {
    a.Set("hint", value)
    return a
}

/**
 * accept
 */
func (a *ImageControl) Accept(value interface{}) *ImageControl {
    a.Set("accept", value)
    return a
}

/**
 * btnClassName
 */
func (a *ImageControl) BtnClassName(value interface{}) *ImageControl {
    a.Set("btnClassName", value)
    return a
}

/**
 * extractValue
 */
func (a *ImageControl) ExtractValue(value interface{}) *ImageControl {
    a.Set("extractValue", value)
    return a
}

/**
 * labelAlign
 */
func (a *ImageControl) LabelAlign(value interface{}) *ImageControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * submitOnChange
 */
func (a *ImageControl) SubmitOnChange(value interface{}) *ImageControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * value
 */
func (a *ImageControl) Value(value interface{}) *ImageControl {
    a.Set("value", value)
    return a
}

/**
 * disabledOn
 */
func (a *ImageControl) DisabledOn(value interface{}) *ImageControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * draggable
 */
func (a *ImageControl) Draggable(value interface{}) *ImageControl {
    a.Set("draggable", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *ImageControl) ClearValueOnHidden(value interface{}) *ImageControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * invalidSizeMessage
 */
func (a *ImageControl) InvalidSizeMessage(value interface{}) *ImageControl {
    a.Set("invalidSizeMessage", value)
    return a
}

/**
 * autoUpload
 */
func (a *ImageControl) AutoUpload(value interface{}) *ImageControl {
    a.Set("autoUpload", value)
    return a
}

/**
 * compress
 */
func (a *ImageControl) Compress(value interface{}) *ImageControl {
    a.Set("compress", value)
    return a
}

/**
 * fixedSizeClassName
 */
func (a *ImageControl) FixedSizeClassName(value interface{}) *ImageControl {
    a.Set("fixedSizeClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *ImageControl) ValidationErrors(value interface{}) *ImageControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * hidden
 */
func (a *ImageControl) Hidden(value interface{}) *ImageControl {
    a.Set("hidden", value)
    return a
}

/**
 * reCropable
 */
func (a *ImageControl) ReCropable(value interface{}) *ImageControl {
    a.Set("reCropable", value)
    return a
}

/**
 * resetValue
 */
func (a *ImageControl) ResetValue(value interface{}) *ImageControl {
    a.Set("resetValue", value)
    return a
}

/**
 * dropCrop
 */
func (a *ImageControl) DropCrop(value interface{}) *ImageControl {
    a.Set("dropCrop", value)
    return a
}

/**
 * hideUploadButton
 */
func (a *ImageControl) HideUploadButton(value interface{}) *ImageControl {
    a.Set("hideUploadButton", value)
    return a
}

/**
 * initAutoFill
 */
func (a *ImageControl) InitAutoFill(value interface{}) *ImageControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * visible
 */
func (a *ImageControl) Visible(value interface{}) *ImageControl {
    a.Set("visible", value)
    return a
}

/**
 * size
 */
func (a *ImageControl) Size(value interface{}) *ImageControl {
    a.Set("size", value)
    return a
}

/**
 * showErrorModal
 */
func (a *ImageControl) ShowErrorModal(value interface{}) *ImageControl {
    a.Set("showErrorModal", value)
    return a
}

/**
 * invalidTypeMessage
 */
func (a *ImageControl) InvalidTypeMessage(value interface{}) *ImageControl {
    a.Set("invalidTypeMessage", value)
    return a
}

/**
 * remark
 */
func (a *ImageControl) Remark(value interface{}) *ImageControl {
    a.Set("remark", value)
    return a
}

/**
 * description
 */
func (a *ImageControl) Description(value interface{}) *ImageControl {
    a.Set("description", value)
    return a
}

/**
 * desc
 */
func (a *ImageControl) Desc(value interface{}) *ImageControl {
    a.Set("desc", value)
    return a
}

/**
 * inline
 */
func (a *ImageControl) Inline(value interface{}) *ImageControl {
    a.Set("inline", value)
    return a
}

/**
 * hiddenOn
 */
func (a *ImageControl) HiddenOn(value interface{}) *ImageControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *ImageControl) StaticInputClassName(value interface{}) *ImageControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * allowInput
 */
func (a *ImageControl) AllowInput(value interface{}) *ImageControl {
    a.Set("allowInput", value)
    return a
}

/**
 * labelWidth
 */
func (a *ImageControl) LabelWidth(value interface{}) *ImageControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *ImageControl) StaticPlaceholder(value interface{}) *ImageControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * uploadBtnText
 */
func (a *ImageControl) UploadBtnText(value interface{}) *ImageControl {
    a.Set("uploadBtnText", value)
    return a
}

/**
 * compressOptions
 */
func (a *ImageControl) CompressOptions(value interface{}) *ImageControl {
    a.Set("compressOptions", value)
    return a
}

/**
 * receiver
 */
func (a *ImageControl) Receiver(value interface{}) *ImageControl {
    a.Set("receiver", value)
    return a
}

/**
 * joinValues
 */
func (a *ImageControl) JoinValues(value interface{}) *ImageControl {
    a.Set("joinValues", value)
    return a
}

/**
 * frameImage
 */
func (a *ImageControl) FrameImage(value interface{}) *ImageControl {
    a.Set("frameImage", value)
    return a
}

/**
 */
func (a *ImageControl) Type(value interface{}) *ImageControl {
    a.Set("type", value)
    return a
}

/**
 * horizontal
 */
func (a *ImageControl) Horizontal(value interface{}) *ImageControl {
    a.Set("horizontal", value)
    return a
}

/**
 * required
 */
func (a *ImageControl) Required(value interface{}) *ImageControl {
    a.Set("required", value)
    return a
}

/**
 * row
 */
func (a *ImageControl) Row(value interface{}) *ImageControl {
    a.Set("row", value)
    return a
}

/**
 * visibleOn
 */
func (a *ImageControl) VisibleOn(value interface{}) *ImageControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * static
 */
func (a *ImageControl) Static(value interface{}) *ImageControl {
    a.Set("static", value)
    return a
}

/**
 * cropQuality
 */
func (a *ImageControl) CropQuality(value interface{}) *ImageControl {
    a.Set("cropQuality", value)
    return a
}

/**
 * staticSchema
 */
func (a *ImageControl) StaticSchema(value interface{}) *ImageControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *ImageControl) ReadOnlyOn(value interface{}) *ImageControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *ImageControl) DescriptionClassName(value interface{}) *ImageControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * validations
 */
func (a *ImageControl) Validations(value interface{}) *ImageControl {
    a.Set("validations", value)
    return a
}

/**
 * className
 */
func (a *ImageControl) ClassName(value interface{}) *ImageControl {
    a.Set("className", value)
    return a
}

/**
 * imageClassName
 */
func (a *ImageControl) ImageClassName(value interface{}) *ImageControl {
    a.Set("imageClassName", value)
    return a
}

/**
 * limit
 */
func (a *ImageControl) Limit(value interface{}) *ImageControl {
    a.Set("limit", value)
    return a
}

/**
 * showCompressOptions
 */
func (a *ImageControl) ShowCompressOptions(value interface{}) *ImageControl {
    a.Set("showCompressOptions", value)
    return a
}

/**
 * labelClassName
 */
func (a *ImageControl) LabelClassName(value interface{}) *ImageControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * readOnly
 */
func (a *ImageControl) ReadOnly(value interface{}) *ImageControl {
    a.Set("readOnly", value)
    return a
}

/**
 * validateOnChange
 */
func (a *ImageControl) ValidateOnChange(value interface{}) *ImageControl {
    a.Set("validateOnChange", value)
    return a
}
