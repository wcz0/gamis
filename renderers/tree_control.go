package renderers


/**
 * Tree 下拉选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tree

 * @author wcz0
 * @version 6.2.2
 */
type TreeControl struct {
	*BaseRenderer
}

func NewTreeControl() *TreeControl {
    a := &TreeControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-tree")
    return a
}


func (a *TreeControl) Set(name string, value interface{}) *TreeControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * staticSchema
 */
func (a *TreeControl) StaticSchema(value interface{}) *TreeControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TreeControl) Width(value interface{}) *TreeControl {
    a.Set("width", value)
    return a
}

/**
 * description
 */
func (a *TreeControl) Description(value interface{}) *TreeControl {
    a.Set("description", value)
    return a
}

/**
 * desc
 */
func (a *TreeControl) Desc(value interface{}) *TreeControl {
    a.Set("desc", value)
    return a
}

/**
 * inputClassName
 */
func (a *TreeControl) InputClassName(value interface{}) *TreeControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * value
 */
func (a *TreeControl) Value(value interface{}) *TreeControl {
    a.Set("value", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TreeControl) StaticLabelClassName(value interface{}) *TreeControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *TreeControl) ClearValueOnSourceChange(value interface{}) *TreeControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * source
 */
func (a *TreeControl) Source(value interface{}) *TreeControl {
    a.Set("source", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *TreeControl) ValuesNoWrap(value interface{}) *TreeControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *TreeControl) ClearValueOnHidden(value interface{}) *TreeControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * validateApi
 */
func (a *TreeControl) ValidateApi(value interface{}) *TreeControl {
    a.Set("validateApi", value)
    return a
}

/**
 * staticOn
 */
func (a *TreeControl) StaticOn(value interface{}) *TreeControl {
    a.Set("staticOn", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TreeControl) StaticPlaceholder(value interface{}) *TreeControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TreeControl) StaticInputClassName(value interface{}) *TreeControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * creatable
 */
func (a *TreeControl) Creatable(value interface{}) *TreeControl {
    a.Set("creatable", value)
    return a
}

/**
 * labelRemark
 */
func (a *TreeControl) LabelRemark(value interface{}) *TreeControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * inline
 */
func (a *TreeControl) Inline(value interface{}) *TreeControl {
    a.Set("inline", value)
    return a
}

/**
 * withChildren
 */
func (a *TreeControl) WithChildren(value interface{}) *TreeControl {
    a.Set("withChildren", value)
    return a
}

/**
 * rootCreatable
 */
func (a *TreeControl) RootCreatable(value interface{}) *TreeControl {
    a.Set("rootCreatable", value)
    return a
}

/**
 * searchApi
 */
func (a *TreeControl) SearchApi(value interface{}) *TreeControl {
    a.Set("searchApi", value)
    return a
}

/**
 * heightAuto
 */
func (a *TreeControl) HeightAuto(value interface{}) *TreeControl {
    a.Set("heightAuto", value)
    return a
}

/**
 * iconField
 */
func (a *TreeControl) IconField(value interface{}) *TreeControl {
    a.Set("iconField", value)
    return a
}

/**
 * removable
 */
func (a *TreeControl) Removable(value interface{}) *TreeControl {
    a.Set("removable", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *TreeControl) DescriptionClassName(value interface{}) *TreeControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * rootLabel
 */
func (a *TreeControl) RootLabel(value interface{}) *TreeControl {
    a.Set("rootLabel", value)
    return a
}

/**
 * onlyChildren
 */
func (a *TreeControl) OnlyChildren(value interface{}) *TreeControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * editable
 */
func (a *TreeControl) Editable(value interface{}) *TreeControl {
    a.Set("editable", value)
    return a
}

/**
 * editDialog
 */
func (a *TreeControl) EditDialog(value interface{}) *TreeControl {
    a.Set("editDialog", value)
    return a
}

/**
 * disabledOn
 */
func (a *TreeControl) DisabledOn(value interface{}) *TreeControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * style
 */
func (a *TreeControl) Style(value interface{}) *TreeControl {
    a.Set("style", value)
    return a
}

/**
 * itemActions
 */
func (a *TreeControl) ItemActions(value interface{}) *TreeControl {
    a.Set("itemActions", value)
    return a
}

/**
 * extraName
 */
func (a *TreeControl) ExtraName(value interface{}) *TreeControl {
    a.Set("extraName", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *TreeControl) ReadOnlyOn(value interface{}) *TreeControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * deleteApi
 */
func (a *TreeControl) DeleteApi(value interface{}) *TreeControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * resetValue
 */
func (a *TreeControl) ResetValue(value interface{}) *TreeControl {
    a.Set("resetValue", value)
    return a
}

/**
 * labelWidth
 */
func (a *TreeControl) LabelWidth(value interface{}) *TreeControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * name
 */
func (a *TreeControl) Name(value interface{}) *TreeControl {
    a.Set("name", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TreeControl) UseMobileUI(value interface{}) *TreeControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * hideRoot
 */
func (a *TreeControl) HideRoot(value interface{}) *TreeControl {
    a.Set("hideRoot", value)
    return a
}

/**
 * initFetch
 */
func (a *TreeControl) InitFetch(value interface{}) *TreeControl {
    a.Set("initFetch", value)
    return a
}

/**
 * clearable
 */
func (a *TreeControl) Clearable(value interface{}) *TreeControl {
    a.Set("clearable", value)
    return a
}

/**
 * addApi
 */
func (a *TreeControl) AddApi(value interface{}) *TreeControl {
    a.Set("addApi", value)
    return a
}

/**
 * addControls
 */
func (a *TreeControl) AddControls(value interface{}) *TreeControl {
    a.Set("addControls", value)
    return a
}

/**
 * labelAlign
 */
func (a *TreeControl) LabelAlign(value interface{}) *TreeControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * validations
 */
func (a *TreeControl) Validations(value interface{}) *TreeControl {
    a.Set("validations", value)
    return a
}

/**
 * enableNodePath
 */
func (a *TreeControl) EnableNodePath(value interface{}) *TreeControl {
    a.Set("enableNodePath", value)
    return a
}

/**
 * pathSeparator
 */
func (a *TreeControl) PathSeparator(value interface{}) *TreeControl {
    a.Set("pathSeparator", value)
    return a
}

/**
 * addDialog
 */
func (a *TreeControl) AddDialog(value interface{}) *TreeControl {
    a.Set("addDialog", value)
    return a
}

/**
 * onEvent
 */
func (a *TreeControl) OnEvent(value interface{}) *TreeControl {
    a.Set("onEvent", value)
    return a
}

/**
 * deferField
 */
func (a *TreeControl) DeferField(value interface{}) *TreeControl {
    a.Set("deferField", value)
    return a
}

/**
 * submitOnChange
 */
func (a *TreeControl) SubmitOnChange(value interface{}) *TreeControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * horizontal
 */
func (a *TreeControl) Horizontal(value interface{}) *TreeControl {
    a.Set("horizontal", value)
    return a
}

/**
 * placeholder
 */
func (a *TreeControl) Placeholder(value interface{}) *TreeControl {
    a.Set("placeholder", value)
    return a
}

/**
 * disabled
 */
func (a *TreeControl) Disabled(value interface{}) *TreeControl {
    a.Set("disabled", value)
    return a
}

/**
 * hidden
 */
func (a *TreeControl) Hidden(value interface{}) *TreeControl {
    a.Set("hidden", value)
    return a
}

/**
 * visible
 */
func (a *TreeControl) Visible(value interface{}) *TreeControl {
    a.Set("visible", value)
    return a
}

/**
 * rootValue
 */
func (a *TreeControl) RootValue(value interface{}) *TreeControl {
    a.Set("rootValue", value)
    return a
}

/**
 * options
 */
func (a *TreeControl) Options(value interface{}) *TreeControl {
    a.Set("options", value)
    return a
}

/**
 * joinValues
 */
func (a *TreeControl) JoinValues(value interface{}) *TreeControl {
    a.Set("joinValues", value)
    return a
}

/**
 * extractValue
 */
func (a *TreeControl) ExtractValue(value interface{}) *TreeControl {
    a.Set("extractValue", value)
    return a
}

/**
 * readOnly
 */
func (a *TreeControl) ReadOnly(value interface{}) *TreeControl {
    a.Set("readOnly", value)
    return a
}

/**
 * visibleOn
 */
func (a *TreeControl) VisibleOn(value interface{}) *TreeControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * onlyLeaf
 */
func (a *TreeControl) OnlyLeaf(value interface{}) *TreeControl {
    a.Set("onlyLeaf", value)
    return a
}

/**
 * showOutline
 */
func (a *TreeControl) ShowOutline(value interface{}) *TreeControl {
    a.Set("showOutline", value)
    return a
}

/**
 * enableDefaultIcon
 */
func (a *TreeControl) EnableDefaultIcon(value interface{}) *TreeControl {
    a.Set("enableDefaultIcon", value)
    return a
}

/**
 * initFetchOn
 */
func (a *TreeControl) InitFetchOn(value interface{}) *TreeControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * multiple
 */
func (a *TreeControl) Multiple(value interface{}) *TreeControl {
    a.Set("multiple", value)
    return a
}

/**
 * deferApi
 */
func (a *TreeControl) DeferApi(value interface{}) *TreeControl {
    a.Set("deferApi", value)
    return a
}

/**
 * editApi
 */
func (a *TreeControl) EditApi(value interface{}) *TreeControl {
    a.Set("editApi", value)
    return a
}

/**
 * row
 */
func (a *TreeControl) Row(value interface{}) *TreeControl {
    a.Set("row", value)
    return a
}

/**
 * className
 */
func (a *TreeControl) ClassName(value interface{}) *TreeControl {
    a.Set("className", value)
    return a
}

/**
 * autoCheckChildren
 */
func (a *TreeControl) AutoCheckChildren(value interface{}) *TreeControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * cascade
 */
func (a *TreeControl) Cascade(value interface{}) *TreeControl {
    a.Set("cascade", value)
    return a
}

/**
 * validateOnChange
 */
func (a *TreeControl) ValidateOnChange(value interface{}) *TreeControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * validationErrors
 */
func (a *TreeControl) ValidationErrors(value interface{}) *TreeControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TreeControl) HiddenOn(value interface{}) *TreeControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * id
 */
func (a *TreeControl) Id(value interface{}) *TreeControl {
    a.Set("id", value)
    return a
}

/**
 * staticClassName
 */
func (a *TreeControl) StaticClassName(value interface{}) *TreeControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * editorSetting
 */
func (a *TreeControl) EditorSetting(value interface{}) *TreeControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * treeClassName
 */
func (a *TreeControl) TreeClassName(value interface{}) *TreeControl {
    a.Set("treeClassName", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *TreeControl) DeleteConfirmText(value interface{}) *TreeControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * size
 */
func (a *TreeControl) Size(value interface{}) *TreeControl {
    a.Set("size", value)
    return a
}

/**
 */
func (a *TreeControl) Type(value interface{}) *TreeControl {
    a.Set("type", value)
    return a
}

/**
 * showIcon
 */
func (a *TreeControl) ShowIcon(value interface{}) *TreeControl {
    a.Set("showIcon", value)
    return a
}

/**
 * autoCancelParent
 */
func (a *TreeControl) AutoCancelParent(value interface{}) *TreeControl {
    a.Set("autoCancelParent", value)
    return a
}

/**
 * highlightTxt
 */
func (a *TreeControl) HighlightTxt(value interface{}) *TreeControl {
    a.Set("highlightTxt", value)
    return a
}

/**
 * searchable
 */
func (a *TreeControl) Searchable(value interface{}) *TreeControl {
    a.Set("searchable", value)
    return a
}

/**
 * editControls
 */
func (a *TreeControl) EditControls(value interface{}) *TreeControl {
    a.Set("editControls", value)
    return a
}

/**
 * selectFirst
 */
func (a *TreeControl) SelectFirst(value interface{}) *TreeControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * checkAll
 */
func (a *TreeControl) CheckAll(value interface{}) *TreeControl {
    a.Set("checkAll", value)
    return a
}

/**
 * label
 */
func (a *TreeControl) Label(value interface{}) *TreeControl {
    a.Set("label", value)
    return a
}

/**
 * searchConfig
 */
func (a *TreeControl) SearchConfig(value interface{}) *TreeControl {
    a.Set("searchConfig", value)
    return a
}

/**
 * labelOverflow
 */
func (a *TreeControl) LabelOverflow(value interface{}) *TreeControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * mode
 */
func (a *TreeControl) Mode(value interface{}) *TreeControl {
    a.Set("mode", value)
    return a
}

/**
 * required
 */
func (a *TreeControl) Required(value interface{}) *TreeControl {
    a.Set("required", value)
    return a
}

/**
 * autoFill
 */
func (a *TreeControl) AutoFill(value interface{}) *TreeControl {
    a.Set("autoFill", value)
    return a
}

/**
 * static
 */
func (a *TreeControl) Static(value interface{}) *TreeControl {
    a.Set("static", value)
    return a
}

/**
 * nodeBehavior
 */
func (a *TreeControl) NodeBehavior(value interface{}) *TreeControl {
    a.Set("nodeBehavior", value)
    return a
}

/**
 * delimiter
 */
func (a *TreeControl) Delimiter(value interface{}) *TreeControl {
    a.Set("delimiter", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *TreeControl) CreateBtnLabel(value interface{}) *TreeControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * labelClassName
 */
func (a *TreeControl) LabelClassName(value interface{}) *TreeControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * remark
 */
func (a *TreeControl) Remark(value interface{}) *TreeControl {
    a.Set("remark", value)
    return a
}

/**
 * hint
 */
func (a *TreeControl) Hint(value interface{}) *TreeControl {
    a.Set("hint", value)
    return a
}

/**
 * initAutoFill
 */
func (a *TreeControl) InitAutoFill(value interface{}) *TreeControl {
    a.Set("initAutoFill", value)
    return a
}
