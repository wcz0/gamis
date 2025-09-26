package renderers

type TableControl struct {
	*BaseRenderer
}

func NewTableControl() *TableControl {
	a := &TableControl{
		BaseRenderer: NewBaseRenderer(),
	}

	a.Set("type", "input-table")
	return a
}

func (a *TableControl) Set(name string, value interface{}) *TableControl {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *TableControl) AddApi(value interface{}) *TableControl {
	a.Set("addApi", value)
	return a
}

func (a *TableControl) AddBtnIcon(value interface{}) *TableControl {
	a.Set("addBtnIcon", value)
	return a
}

func (a *TableControl) AddBtnLabel(value interface{}) *TableControl {
	a.Set("addBtnLabel", value)
	return a
}

func (a *TableControl) Addable(value interface{}) *TableControl {
	a.Set("addable", value)
	return a
}

func (a *TableControl) AffixFooter(value interface{}) *TableControl {
	a.Set("affixFooter", value)
	return a
}

func (a *TableControl) AffixHeader(value interface{}) *TableControl {
	a.Set("affixHeader", value)
	return a
}

func (a *TableControl) AffixRow(value interface{}) *TableControl {
	a.Set("affixRow", value)
	return a
}

func (a *TableControl) AutoFill(value interface{}) *TableControl {
	a.Set("autoFill", value)
	return a
}

func (a *TableControl) AutoFillHeight(value interface{}) *TableControl {
	a.Set("autoFillHeight", value)
	return a
}

func (a *TableControl) AutoGenerateFilter(value interface{}) *TableControl {
	a.Set("autoGenerateFilter", value)
	return a
}

func (a *TableControl) CanAccessSuperData(value interface{}) *TableControl {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *TableControl) CancelBtnIcon(value interface{}) *TableControl {
	a.Set("cancelBtnIcon", value)
	return a
}

func (a *TableControl) CancelBtnLabel(value interface{}) *TableControl {
	a.Set("cancelBtnLabel", value)
	return a
}

func (a *TableControl) ChildrenAddable(value interface{}) *TableControl {
	a.Set("childrenAddable", value)
	return a
}

func (a *TableControl) ClassName(value interface{}) *TableControl {
	a.Set("className", value)
	return a
}

func (a *TableControl) ClearValueOnHidden(value interface{}) *TableControl {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *TableControl) Columns(value interface{}) *TableControl {
	a.Set("columns", value)
	return a
}

func (a *TableControl) ColumnsTogglable(value interface{}) *TableControl {
	a.Set("columnsTogglable", value)
	return a
}

func (a *TableControl) CombineFromIndex(value interface{}) *TableControl {
	a.Set("combineFromIndex", value)
	return a
}

func (a *TableControl) CombineNum(value interface{}) *TableControl {
	a.Set("combineNum", value)
	return a
}

func (a *TableControl) ConfirmBtnIcon(value interface{}) *TableControl {
	a.Set("confirmBtnIcon", value)
	return a
}

func (a *TableControl) ConfirmBtnLabel(value interface{}) *TableControl {
	a.Set("confirmBtnLabel", value)
	return a
}

func (a *TableControl) CopyAddBtn(value interface{}) *TableControl {
	a.Set("copyAddBtn", value)
	return a
}

func (a *TableControl) CopyBtnIcon(value interface{}) *TableControl {
	a.Set("copyBtnIcon", value)
	return a
}

func (a *TableControl) CopyBtnLabel(value interface{}) *TableControl {
	a.Set("copyBtnLabel", value)
	return a
}

func (a *TableControl) CopyData(value interface{}) *TableControl {
	a.Set("copyData", value)
	return a
}

func (a *TableControl) Copyable(value interface{}) *TableControl {
	a.Set("copyable", value)
	return a
}

func (a *TableControl) DeferApi(value interface{}) *TableControl {
	a.Set("deferApi", value)
	return a
}

func (a *TableControl) DeleteApi(value interface{}) *TableControl {
	a.Set("deleteApi", value)
	return a
}

func (a *TableControl) DeleteBtnIcon(value interface{}) *TableControl {
	a.Set("deleteBtnIcon", value)
	return a
}

func (a *TableControl) DeleteBtnLabel(value interface{}) *TableControl {
	a.Set("deleteBtnLabel", value)
	return a
}

func (a *TableControl) DeleteConfirmText(value interface{}) *TableControl {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *TableControl) Desc(value interface{}) *TableControl {
	a.Set("desc", value)
	return a
}

func (a *TableControl) Description(value interface{}) *TableControl {
	a.Set("description", value)
	return a
}

func (a *TableControl) DescriptionClassName(value interface{}) *TableControl {
	a.Set("descriptionClassName", value)
	return a
}

func (a *TableControl) Disabled(value interface{}) *TableControl {
	a.Set("disabled", value)
	return a
}

func (a *TableControl) DisabledOn(value interface{}) *TableControl {
	a.Set("disabledOn", value)
	return a
}

func (a *TableControl) Draggable(value interface{}) *TableControl {
	a.Set("draggable", value)
	return a
}

func (a *TableControl) EditBtnIcon(value interface{}) *TableControl {
	a.Set("editBtnIcon", value)
	return a
}

func (a *TableControl) EditBtnLabel(value interface{}) *TableControl {
	a.Set("editBtnLabel", value)
	return a
}

func (a *TableControl) Editable(value interface{}) *TableControl {
	a.Set("editable", value)
	return a
}

func (a *TableControl) EditorSetting(value interface{}) *TableControl {
	a.Set("editorSetting", value)
	return a
}

func (a *TableControl) EnableStaticTransform(value interface{}) *TableControl {
	a.Set("enableStaticTransform", value)
	return a
}

func (a *TableControl) ExtraName(value interface{}) *TableControl {
	a.Set("extraName", value)
	return a
}

func (a *TableControl) Footable(value interface{}) *TableControl {
	a.Set("footable", value)
	return a
}

func (a *TableControl) FooterAddBtn(value interface{}) *TableControl {
	a.Set("footerAddBtn", value)
	return a
}

func (a *TableControl) FooterClassName(value interface{}) *TableControl {
	a.Set("footerClassName", value)
	return a
}

func (a *TableControl) HeaderClassName(value interface{}) *TableControl {
	a.Set("headerClassName", value)
	return a
}

func (a *TableControl) Hidden(value interface{}) *TableControl {
	a.Set("hidden", value)
	return a
}

func (a *TableControl) HiddenOn(value interface{}) *TableControl {
	a.Set("hiddenOn", value)
	return a
}

func (a *TableControl) Hint(value interface{}) *TableControl {
	a.Set("hint", value)
	return a
}

func (a *TableControl) Horizontal(value interface{}) *TableControl {
	a.Set("horizontal", value)
	return a
}

func (a *TableControl) Id(value interface{}) *TableControl {
	a.Set("id", value)
	return a
}

func (a *TableControl) InitAutoFill(value interface{}) *TableControl {
	a.Set("initAutoFill", value)
	return a
}

func (a *TableControl) Inline(value interface{}) *TableControl {
	a.Set("inline", value)
	return a
}

func (a *TableControl) InputClassName(value interface{}) *TableControl {
	a.Set("inputClassName", value)
	return a
}

func (a *TableControl) ItemBadge(value interface{}) *TableControl {
	a.Set("itemBadge", value)
	return a
}

func (a *TableControl) Label(value interface{}) *TableControl {
	a.Set("label", value)
	return a
}

func (a *TableControl) LabelAlign(value interface{}) *TableControl {
	a.Set("labelAlign", value)
	return a
}

func (a *TableControl) LabelClassName(value interface{}) *TableControl {
	a.Set("labelClassName", value)
	return a
}

func (a *TableControl) LabelRemark(value interface{}) *TableControl {
	a.Set("labelRemark", value)
	return a
}

func (a *TableControl) LabelWidth(value interface{}) *TableControl {
	a.Set("labelWidth", value)
	return a
}

func (a *TableControl) MatchFunc(value interface{}) *TableControl {
	a.Set("matchFunc", value)
	return a
}

func (a *TableControl) MaxLength(value interface{}) *TableControl {
	a.Set("maxLength", value)
	return a
}

func (a *TableControl) MinLength(value interface{}) *TableControl {
	a.Set("minLength", value)
	return a
}

func (a *TableControl) Mode(value interface{}) *TableControl {
	a.Set("mode", value)
	return a
}

func (a *TableControl) Name(value interface{}) *TableControl {
	a.Set("name", value)
	return a
}

func (a *TableControl) NeedConfirm(value interface{}) *TableControl {
	a.Set("needConfirm", value)
	return a
}

func (a *TableControl) OnEvent(value interface{}) *TableControl {
	a.Set("onEvent", value)
	return a
}

func (a *TableControl) PerPage(value interface{}) *TableControl {
	a.Set("perPage", value)
	return a
}

func (a *TableControl) Placeholder(value interface{}) *TableControl {
	a.Set("placeholder", value)
	return a
}

func (a *TableControl) PrefixRow(value interface{}) *TableControl {
	a.Set("prefixRow", value)
	return a
}

func (a *TableControl) ReadOnly(value interface{}) *TableControl {
	a.Set("readOnly", value)
	return a
}

func (a *TableControl) ReadOnlyOn(value interface{}) *TableControl {
	a.Set("readOnlyOn", value)
	return a
}

func (a *TableControl) Remark(value interface{}) *TableControl {
	a.Set("remark", value)
	return a
}

func (a *TableControl) Removable(value interface{}) *TableControl {
	a.Set("removable", value)
	return a
}

func (a *TableControl) Required(value interface{}) *TableControl {
	a.Set("required", value)
	return a
}

func (a *TableControl) Resizable(value interface{}) *TableControl {
	a.Set("resizable", value)
	return a
}

func (a *TableControl) Row(value interface{}) *TableControl {
	a.Set("row", value)
	return a
}

func (a *TableControl) RowClassNameExpr(value interface{}) *TableControl {
	a.Set("rowClassNameExpr", value)
	return a
}

func (a *TableControl) SaveImmediately(value interface{}) *TableControl {
	a.Set("saveImmediately", value)
	return a
}

func (a *TableControl) Scaffold(value interface{}) *TableControl {
	a.Set("scaffold", value)
	return a
}

func (a *TableControl) ShowFooter(value interface{}) *TableControl {
	a.Set("showFooter", value)
	return a
}

func (a *TableControl) ShowFooterAddBtn(value interface{}) *TableControl {
	a.Set("showFooterAddBtn", value)
	return a
}

func (a *TableControl) ShowHeader(value interface{}) *TableControl {
	a.Set("showHeader", value)
	return a
}

func (a *TableControl) ShowIndex(value interface{}) *TableControl {
	a.Set("showIndex", value)
	return a
}

func (a *TableControl) ShowTableAddBtn(value interface{}) *TableControl {
	a.Set("showTableAddBtn", value)
	return a
}

func (a *TableControl) Size(value interface{}) *TableControl {
	a.Set("size", value)
	return a
}

func (a *TableControl) Source(value interface{}) *TableControl {
	a.Set("source", value)
	return a
}

func (a *TableControl) Static(value interface{}) *TableControl {
	a.Set("static", value)
	return a
}

func (a *TableControl) StaticClassName(value interface{}) *TableControl {
	a.Set("staticClassName", value)
	return a
}

func (a *TableControl) StaticInputClassName(value interface{}) *TableControl {
	a.Set("staticInputClassName", value)
	return a
}

func (a *TableControl) StaticLabelClassName(value interface{}) *TableControl {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *TableControl) StaticOn(value interface{}) *TableControl {
	a.Set("staticOn", value)
	return a
}

func (a *TableControl) StaticPlaceholder(value interface{}) *TableControl {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *TableControl) StaticSchema(value interface{}) *TableControl {
	a.Set("staticSchema", value)
	return a
}

func (a *TableControl) Style(value interface{}) *TableControl {
	a.Set("style", value)
	return a
}

func (a *TableControl) SubAddBtnIcon(value interface{}) *TableControl {
	a.Set("subAddBtnIcon", value)
	return a
}

func (a *TableControl) SubAddBtnLabel(value interface{}) *TableControl {
	a.Set("subAddBtnLabel", value)
	return a
}

func (a *TableControl) SubmitOnChange(value interface{}) *TableControl {
	a.Set("submitOnChange", value)
	return a
}

func (a *TableControl) TableClassName(value interface{}) *TableControl {
	a.Set("tableClassName", value)
	return a
}

func (a *TableControl) TableLayout(value interface{}) *TableControl {
	a.Set("tableLayout", value)
	return a
}

func (a *TableControl) TestIdBuilder(value interface{}) *TableControl {
	a.Set("testIdBuilder", value)
	return a
}

func (a *TableControl) Testid(value interface{}) *TableControl {
	a.Set("testid", value)
	return a
}

func (a *TableControl) Title(value interface{}) *TableControl {
	a.Set("title", value)
	return a
}

func (a *TableControl) ToolbarClassName(value interface{}) *TableControl {
	a.Set("toolbarClassName", value)
	return a
}

func (a *TableControl) Type(value interface{}) *TableControl {
	a.Set("type", value)
	return a
}

func (a *TableControl) UpdateApi(value interface{}) *TableControl {
	a.Set("updateApi", value)
	return a
}

func (a *TableControl) UseMobileUI(value interface{}) *TableControl {
	a.Set("useMobileUI", value)
	return a
}

func (a *TableControl) ValidateApi(value interface{}) *TableControl {
	a.Set("validateApi", value)
	return a
}

func (a *TableControl) ValidateOnChange(value interface{}) *TableControl {
	a.Set("validateOnChange", value)
	return a
}

func (a *TableControl) ValidationErrors(value interface{}) *TableControl {
	a.Set("validationErrors", value)
	return a
}

func (a *TableControl) Validations(value interface{}) *TableControl {
	a.Set("validations", value)
	return a
}

func (a *TableControl) Value(value interface{}) *TableControl {
	a.Set("value", value)
	return a
}

func (a *TableControl) ValueField(value interface{}) *TableControl {
	a.Set("valueField", value)
	return a
}

func (a *TableControl) Visible(value interface{}) *TableControl {
	a.Set("visible", value)
	return a
}

func (a *TableControl) VisibleOn(value interface{}) *TableControl {
	a.Set("visibleOn", value)
	return a
}

func (a *TableControl) Width(value interface{}) *TableControl {
	a.Set("width", value)
	return a
}
