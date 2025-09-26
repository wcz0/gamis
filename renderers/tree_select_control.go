package renderers


/**
 * Tree 下拉选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tree

 * @author wcz0
 * @version 6.2.2
 */
type TreeSelectControl struct {
	*BaseRenderer
}

func NewTreeSelectControl() *TreeSelectControl {
    a := &TreeSelectControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "tree-select")
    return a
}


func (a *TreeSelectControl) Set(name string, value interface{}) *TreeSelectControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * enableDefaultIcon
 */
func (a *TreeSelectControl) EnableDefaultIcon(value interface{}) *TreeSelectControl {
    a.Set("enableDefaultIcon", value)
    return a
}

/**
 * createBtnLabel
 */
func (a *TreeSelectControl) CreateBtnLabel(value interface{}) *TreeSelectControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * row
 */
func (a *TreeSelectControl) Row(value interface{}) *TreeSelectControl {
    a.Set("row", value)
    return a
}

/**
 * useMobileUI
 */
func (a *TreeSelectControl) UseMobileUI(value interface{}) *TreeSelectControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * showIcon
 */
func (a *TreeSelectControl) ShowIcon(value interface{}) *TreeSelectControl {
    a.Set("showIcon", value)
    return a
}

/**
 * maxTagCount
 */
func (a *TreeSelectControl) MaxTagCount(value interface{}) *TreeSelectControl {
    a.Set("maxTagCount", value)
    return a
}

/**
 * options
 */
func (a *TreeSelectControl) Options(value interface{}) *TreeSelectControl {
    a.Set("options", value)
    return a
}

/**
 * selectFirst
 */
func (a *TreeSelectControl) SelectFirst(value interface{}) *TreeSelectControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * initFetchOn
 */
func (a *TreeSelectControl) InitFetchOn(value interface{}) *TreeSelectControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * deleteConfirmText
 */
func (a *TreeSelectControl) DeleteConfirmText(value interface{}) *TreeSelectControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * size
 */
func (a *TreeSelectControl) Size(value interface{}) *TreeSelectControl {
    a.Set("size", value)
    return a
}

/**
 * rootLabel
 */
func (a *TreeSelectControl) RootLabel(value interface{}) *TreeSelectControl {
    a.Set("rootLabel", value)
    return a
}

/**
 * deferField
 */
func (a *TreeSelectControl) DeferField(value interface{}) *TreeSelectControl {
    a.Set("deferField", value)
    return a
}

/**
 * editControls
 */
func (a *TreeSelectControl) EditControls(value interface{}) *TreeSelectControl {
    a.Set("editControls", value)
    return a
}

/**
 * horizontal
 */
func (a *TreeSelectControl) Horizontal(value interface{}) *TreeSelectControl {
    a.Set("horizontal", value)
    return a
}

/**
 * clearValueOnHidden
 */
func (a *TreeSelectControl) ClearValueOnHidden(value interface{}) *TreeSelectControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否自动选中子节点
 */
func (a *TreeSelectControl) AutoCheckChildren(value interface{}) *TreeSelectControl {
    a.Set("autoCheckChildren", value)
    return a
}

/**
 * joinValues
 */
func (a *TreeSelectControl) JoinValues(value interface{}) *TreeSelectControl {
    a.Set("joinValues", value)
    return a
}

/**
 * extractValue
 */
func (a *TreeSelectControl) ExtractValue(value interface{}) *TreeSelectControl {
    a.Set("extractValue", value)
    return a
}

/**
 * labelAlign
 */
func (a *TreeSelectControl) LabelAlign(value interface{}) *TreeSelectControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * extraName
 */
func (a *TreeSelectControl) ExtraName(value interface{}) *TreeSelectControl {
    a.Set("extraName", value)
    return a
}

/**
 * mode
 */
func (a *TreeSelectControl) Mode(value interface{}) *TreeSelectControl {
    a.Set("mode", value)
    return a
}

/**
 * inline
 */
func (a *TreeSelectControl) Inline(value interface{}) *TreeSelectControl {
    a.Set("inline", value)
    return a
}

/**
 * validations
 */
func (a *TreeSelectControl) Validations(value interface{}) *TreeSelectControl {
    a.Set("validations", value)
    return a
}

/**
 * staticOn
 */
func (a *TreeSelectControl) StaticOn(value interface{}) *TreeSelectControl {
    a.Set("staticOn", value)
    return a
}

/**
 * desc
 */
func (a *TreeSelectControl) Desc(value interface{}) *TreeSelectControl {
    a.Set("desc", value)
    return a
}

/**
 * required
 */
func (a *TreeSelectControl) Required(value interface{}) *TreeSelectControl {
    a.Set("required", value)
    return a
}

/**
 * hiddenOn
 */
func (a *TreeSelectControl) HiddenOn(value interface{}) *TreeSelectControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * staticSchema
 */
func (a *TreeSelectControl) StaticSchema(value interface{}) *TreeSelectControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * rootValue
 */
func (a *TreeSelectControl) RootValue(value interface{}) *TreeSelectControl {
    a.Set("rootValue", value)
    return a
}

/**
 * withChildren
 */
func (a *TreeSelectControl) WithChildren(value interface{}) *TreeSelectControl {
    a.Set("withChildren", value)
    return a
}

/**
 * showOutline
 */
func (a *TreeSelectControl) ShowOutline(value interface{}) *TreeSelectControl {
    a.Set("showOutline", value)
    return a
}

/**
 * deferApi
 */
func (a *TreeSelectControl) DeferApi(value interface{}) *TreeSelectControl {
    a.Set("deferApi", value)
    return a
}

/**
 * editable
 */
func (a *TreeSelectControl) Editable(value interface{}) *TreeSelectControl {
    a.Set("editable", value)
    return a
}

/**
 * labelClassName
 */
func (a *TreeSelectControl) LabelClassName(value interface{}) *TreeSelectControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * validationErrors
 */
func (a *TreeSelectControl) ValidationErrors(value interface{}) *TreeSelectControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * id
 */
func (a *TreeSelectControl) Id(value interface{}) *TreeSelectControl {
    a.Set("id", value)
    return a
}

/**
 * enableNodePath
 */
func (a *TreeSelectControl) EnableNodePath(value interface{}) *TreeSelectControl {
    a.Set("enableNodePath", value)
    return a
}

/**
 * overflowTagPopover
 */
func (a *TreeSelectControl) OverflowTagPopover(value interface{}) *TreeSelectControl {
    a.Set("overflowTagPopover", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TreeSelectControl) Width(value interface{}) *TreeSelectControl {
    a.Set("width", value)
    return a
}

/**
 * name
 */
func (a *TreeSelectControl) Name(value interface{}) *TreeSelectControl {
    a.Set("name", value)
    return a
}

/**
 * hint
 */
func (a *TreeSelectControl) Hint(value interface{}) *TreeSelectControl {
    a.Set("hint", value)
    return a
}

/**
 * staticInputClassName
 */
func (a *TreeSelectControl) StaticInputClassName(value interface{}) *TreeSelectControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * style
 */
func (a *TreeSelectControl) Style(value interface{}) *TreeSelectControl {
    a.Set("style", value)
    return a
}

/**
 */
func (a *TreeSelectControl) Type(value interface{}) *TreeSelectControl {
    a.Set("type", value)
    return a
}

/**
 * 设置label字段
 */
func (a *TreeSelectControl) LabelField(value interface{}) *TreeSelectControl {
    a.Set("labelField", value)
    return a
}

/**
 * 设置value字段
 */
func (a *TreeSelectControl) ValueField(value interface{}) *TreeSelectControl {
    a.Set("valueField", value)
    return a
}

/**
 * valuesNoWrap
 */
func (a *TreeSelectControl) ValuesNoWrap(value interface{}) *TreeSelectControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * resetValue
 */
func (a *TreeSelectControl) ResetValue(value interface{}) *TreeSelectControl {
    a.Set("resetValue", value)
    return a
}

/**
 * label
 */
func (a *TreeSelectControl) Label(value interface{}) *TreeSelectControl {
    a.Set("label", value)
    return a
}

/**
 * labelOverflow
 */
func (a *TreeSelectControl) LabelOverflow(value interface{}) *TreeSelectControl {
    a.Set("labelOverflow", value)
    return a
}

/**
 * submitOnChange
 */
func (a *TreeSelectControl) SubmitOnChange(value interface{}) *TreeSelectControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * descriptionClassName
 */
func (a *TreeSelectControl) DescriptionClassName(value interface{}) *TreeSelectControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * value
 */
func (a *TreeSelectControl) Value(value interface{}) *TreeSelectControl {
    a.Set("value", value)
    return a
}

/**
 * onEvent
 */
func (a *TreeSelectControl) OnEvent(value interface{}) *TreeSelectControl {
    a.Set("onEvent", value)
    return a
}

/**
 * checkAll
 */
func (a *TreeSelectControl) CheckAll(value interface{}) *TreeSelectControl {
    a.Set("checkAll", value)
    return a
}

/**
 * addControls
 */
func (a *TreeSelectControl) AddControls(value interface{}) *TreeSelectControl {
    a.Set("addControls", value)
    return a
}

/**
 * creatable
 */
func (a *TreeSelectControl) Creatable(value interface{}) *TreeSelectControl {
    a.Set("creatable", value)
    return a
}

/**
 * remark
 */
func (a *TreeSelectControl) Remark(value interface{}) *TreeSelectControl {
    a.Set("remark", value)
    return a
}

/**
 * placeholder
 */
func (a *TreeSelectControl) Placeholder(value interface{}) *TreeSelectControl {
    a.Set("placeholder", value)
    return a
}

/**
 * initAutoFill
 */
func (a *TreeSelectControl) InitAutoFill(value interface{}) *TreeSelectControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * static
 */
func (a *TreeSelectControl) Static(value interface{}) *TreeSelectControl {
    a.Set("static", value)
    return a
}

/**
 * className
 */
func (a *TreeSelectControl) ClassName(value interface{}) *TreeSelectControl {
    a.Set("className", value)
    return a
}

/**
 * editorSetting
 */
func (a *TreeSelectControl) EditorSetting(value interface{}) *TreeSelectControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * hideNodePathLabel
 */
func (a *TreeSelectControl) HideNodePathLabel(value interface{}) *TreeSelectControl {
    a.Set("hideNodePathLabel", value)
    return a
}

/**
 * pathSeparator
 */
func (a *TreeSelectControl) PathSeparator(value interface{}) *TreeSelectControl {
    a.Set("pathSeparator", value)
    return a
}

/**
 * autoCancelParent
 */
func (a *TreeSelectControl) AutoCancelParent(value interface{}) *TreeSelectControl {
    a.Set("autoCancelParent", value)
    return a
}

/**
 * testIdBuilder
 */
func (a *TreeSelectControl) TestIdBuilder(value interface{}) *TreeSelectControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * clearable
 */
func (a *TreeSelectControl) Clearable(value interface{}) *TreeSelectControl {
    a.Set("clearable", value)
    return a
}

/**
 * inputClassName
 */
func (a *TreeSelectControl) InputClassName(value interface{}) *TreeSelectControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * disabledOn
 */
func (a *TreeSelectControl) DisabledOn(value interface{}) *TreeSelectControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * visibleOn
 */
func (a *TreeSelectControl) VisibleOn(value interface{}) *TreeSelectControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * staticClassName
 */
func (a *TreeSelectControl) StaticClassName(value interface{}) *TreeSelectControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * hideRoot
 */
func (a *TreeSelectControl) HideRoot(value interface{}) *TreeSelectControl {
    a.Set("hideRoot", value)
    return a
}

/**
 * onlyChildren
 */
func (a *TreeSelectControl) OnlyChildren(value interface{}) *TreeSelectControl {
    a.Set("onlyChildren", value)
    return a
}

/**
 * menuTpl
 */
func (a *TreeSelectControl) MenuTpl(value interface{}) *TreeSelectControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * initFetch
 */
func (a *TreeSelectControl) InitFetch(value interface{}) *TreeSelectControl {
    a.Set("initFetch", value)
    return a
}

/**
 * addApi
 */
func (a *TreeSelectControl) AddApi(value interface{}) *TreeSelectControl {
    a.Set("addApi", value)
    return a
}

/**
 * removable
 */
func (a *TreeSelectControl) Removable(value interface{}) *TreeSelectControl {
    a.Set("removable", value)
    return a
}

/**
 * labelWidth
 */
func (a *TreeSelectControl) LabelWidth(value interface{}) *TreeSelectControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * readOnlyOn
 */
func (a *TreeSelectControl) ReadOnlyOn(value interface{}) *TreeSelectControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * description
 */
func (a *TreeSelectControl) Description(value interface{}) *TreeSelectControl {
    a.Set("description", value)
    return a
}

/**
 * staticLabelClassName
 */
func (a *TreeSelectControl) StaticLabelClassName(value interface{}) *TreeSelectControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * delimiter
 */
func (a *TreeSelectControl) Delimiter(value interface{}) *TreeSelectControl {
    a.Set("delimiter", value)
    return a
}

/**
 * editDialog
 */
func (a *TreeSelectControl) EditDialog(value interface{}) *TreeSelectControl {
    a.Set("editDialog", value)
    return a
}

/**
 * validateOnChange
 */
func (a *TreeSelectControl) ValidateOnChange(value interface{}) *TreeSelectControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * visible
 */
func (a *TreeSelectControl) Visible(value interface{}) *TreeSelectControl {
    a.Set("visible", value)
    return a
}

/**
 * nodeBehavior
 */
func (a *TreeSelectControl) NodeBehavior(value interface{}) *TreeSelectControl {
    a.Set("nodeBehavior", value)
    return a
}

/**
 * itemActions
 */
func (a *TreeSelectControl) ItemActions(value interface{}) *TreeSelectControl {
    a.Set("itemActions", value)
    return a
}

/**
 * 是否可搜索
 */
func (a *TreeSelectControl) Searchable(value interface{}) *TreeSelectControl {
    a.Set("searchable", value)
    return a
}

/**
 * multiple
 */
func (a *TreeSelectControl) Multiple(value interface{}) *TreeSelectControl {
    a.Set("multiple", value)
    return a
}

/**
 * clearValueOnSourceChange
 */
func (a *TreeSelectControl) ClearValueOnSourceChange(value interface{}) *TreeSelectControl {
    a.Set("clearValueOnSourceChange", value)
    return a
}

/**
 * validateApi
 */
func (a *TreeSelectControl) ValidateApi(value interface{}) *TreeSelectControl {
    a.Set("validateApi", value)
    return a
}

/**
 * autoFill
 */
func (a *TreeSelectControl) AutoFill(value interface{}) *TreeSelectControl {
    a.Set("autoFill", value)
    return a
}

/**
 * cascade
 */
func (a *TreeSelectControl) Cascade(value interface{}) *TreeSelectControl {
    a.Set("cascade", value)
    return a
}

/**
 * rootCreatable
 */
func (a *TreeSelectControl) RootCreatable(value interface{}) *TreeSelectControl {
    a.Set("rootCreatable", value)
    return a
}

/**
 * addDialog
 */
func (a *TreeSelectControl) AddDialog(value interface{}) *TreeSelectControl {
    a.Set("addDialog", value)
    return a
}

/**
 * deleteApi
 */
func (a *TreeSelectControl) DeleteApi(value interface{}) *TreeSelectControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * disabled
 */
func (a *TreeSelectControl) Disabled(value interface{}) *TreeSelectControl {
    a.Set("disabled", value)
    return a
}

/**
 * staticPlaceholder
 */
func (a *TreeSelectControl) StaticPlaceholder(value interface{}) *TreeSelectControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * source
 */
func (a *TreeSelectControl) Source(value interface{}) *TreeSelectControl {
    a.Set("source", value)
    return a
}

/**
 * editApi
 */
func (a *TreeSelectControl) EditApi(value interface{}) *TreeSelectControl {
    a.Set("editApi", value)
    return a
}

/**
 * labelRemark
 */
func (a *TreeSelectControl) LabelRemark(value interface{}) *TreeSelectControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * readOnly
 */
func (a *TreeSelectControl) ReadOnly(value interface{}) *TreeSelectControl {
    a.Set("readOnly", value)
    return a
}

/**
 * hidden
 */
func (a *TreeSelectControl) Hidden(value interface{}) *TreeSelectControl {
    a.Set("hidden", value)
    return a
}

/**
 * onlyLeaf
 */
func (a *TreeSelectControl) OnlyLeaf(value interface{}) *TreeSelectControl {
    a.Set("onlyLeaf", value)
    return a
}
