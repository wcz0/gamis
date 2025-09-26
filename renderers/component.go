package renderers

type Component struct {
	*BaseRenderer
}

func NewComponent(typeStr string) *Component {
	a := &Component{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", typeStr)
	return a
}

func (a *Component) Set(name string, value interface{}) *Component {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *Component) SetType(typeStr string) *Component {
	a.Set("type", typeStr)
	return a
}

func (a *Component) Accept(value interface{}) *Component {
	a.Set("accept", value)
	return a
}

func (a *Component) Accordion(value interface{}) *Component {
	a.Set("accordion", value)
	return a
}

func (a *Component) ActionClassName(value interface{}) *Component {
	a.Set("actionClassName", value)
	return a
}

func (a *Component) ActionFinishLabel(value interface{}) *Component {
	a.Set("actionFinishLabel", value)
	return a
}

func (a *Component) ActionNextLabel(value interface{}) *Component {
	a.Set("actionNextLabel", value)
	return a
}

func (a *Component) ActionNextSaveLabel(value interface{}) *Component {
	a.Set("actionNextSaveLabel", value)
	return a
}

func (a *Component) ActionPrevLabel(value interface{}) *Component {
	a.Set("actionPrevLabel", value)
	return a
}

func (a *Component) ActionType(value interface{}) *Component {
	a.Set("actionType", value)
	return a
}

func (a *Component) Actions(value interface{}) *Component {
	a.Set("actions", value)
	return a
}

func (a *Component) ActionsClassName(value interface{}) *Component {
	a.Set("actionsClassName", value)
	return a
}

func (a *Component) ActionsControlClassName(value interface{}) *Component {
	a.Set("actionsControlClassName", value)
	return a
}

func (a *Component) ActionsPosition(value interface{}) *Component {
	a.Set("actionsPosition", value)
	return a
}

func (a *Component) Active(value interface{}) *Component {
	a.Set("active", value)
	return a
}

func (a *Component) ActiveClassName(value interface{}) *Component {
	a.Set("activeClassName", value)
	return a
}

func (a *Component) ActiveItemSchema(value interface{}) *Component {
	a.Set("activeItemSchema", value)
	return a
}

func (a *Component) ActiveKey(value interface{}) *Component {
	a.Set("activeKey", value)
	return a
}

func (a *Component) ActiveLevel(value interface{}) *Component {
	a.Set("activeLevel", value)
	return a
}

func (a *Component) ActivePage(value interface{}) *Component {
	a.Set("activePage", value)
	return a
}

func (a *Component) AddApi(value interface{}) *Component {
	a.Set("addApi", value)
	return a
}

func (a *Component) AddBtnIcon(value interface{}) *Component {
	a.Set("addBtnIcon", value)
	return a
}

func (a *Component) AddBtnLabel(value interface{}) *Component {
	a.Set("addBtnLabel", value)
	return a
}

func (a *Component) AddBtnText(value interface{}) *Component {
	a.Set("addBtnText", value)
	return a
}

func (a *Component) AddBtnVisibleOn(value interface{}) *Component {
	a.Set("addBtnVisibleOn", value)
	return a
}

func (a *Component) AddButtonClassName(value interface{}) *Component {
	a.Set("addButtonClassName", value)
	return a
}

func (a *Component) AddButtonText(value interface{}) *Component {
	a.Set("addButtonText", value)
	return a
}

func (a *Component) AddControls(value interface{}) *Component {
	a.Set("addControls", value)
	return a
}

func (a *Component) AddDialog(value interface{}) *Component {
	a.Set("addDialog", value)
	return a
}

func (a *Component) AddGroupBtnVisibleOn(value interface{}) *Component {
	a.Set("addGroupBtnVisibleOn", value)
	return a
}

func (a *Component) AddOn(value interface{}) *Component {
	a.Set("addOn", value)
	return a
}

func (a *Component) Addable(value interface{}) *Component {
	a.Set("addable", value)
	return a
}

func (a *Component) Addattop(value interface{}) *Component {
	a.Set("addattop", value)
	return a
}

func (a *Component) AdvancedSettings(value interface{}) *Component {
	a.Set("advancedSettings", value)
	return a
}

func (a *Component) AffixFooter(value interface{}) *Component {
	a.Set("affixFooter", value)
	return a
}

func (a *Component) AffixHeader(value interface{}) *Component {
	a.Set("affixHeader", value)
	return a
}

func (a *Component) AffixRow(value interface{}) *Component {
	a.Set("affixRow", value)
	return a
}

func (a *Component) Ak(value interface{}) *Component {
	a.Set("ak", value)
	return a
}

func (a *Component) Align(value interface{}) *Component {
	a.Set("align", value)
	return a
}

func (a *Component) AlignContent(value interface{}) *Component {
	a.Set("alignContent", value)
	return a
}

func (a *Component) AlignItems(value interface{}) *Component {
	a.Set("alignItems", value)
	return a
}

func (a *Component) AllSheets(value interface{}) *Component {
	a.Set("allSheets", value)
	return a
}

func (a *Component) Allow(value interface{}) *Component {
	a.Set("allow", value)
	return a
}

func (a *Component) AllowCity(value interface{}) *Component {
	a.Set("allowCity", value)
	return a
}

func (a *Component) AllowClear(value interface{}) *Component {
	a.Set("allowClear", value)
	return a
}

func (a *Component) AllowCustomColor(value interface{}) *Component {
	a.Set("allowCustomColor", value)
	return a
}

func (a *Component) AllowDistrict(value interface{}) *Component {
	a.Set("allowDistrict", value)
	return a
}

func (a *Component) AllowFullscreen(value interface{}) *Component {
	a.Set("allowFullscreen", value)
	return a
}

func (a *Component) AllowInput(value interface{}) *Component {
	a.Set("allowInput", value)
	return a
}

func (a *Component) AllowStreet(value interface{}) *Component {
	a.Set("allowStreet", value)
	return a
}

func (a *Component) Alt(value interface{}) *Component {
	a.Set("alt", value)
	return a
}

func (a *Component) AlwaysShowArrow(value interface{}) *Component {
	a.Set("alwaysShowArrow", value)
	return a
}

func (a *Component) AlwaysShowPagination(value interface{}) *Component {
	a.Set("alwaysShowPagination", value)
	return a
}

func (a *Component) Animate(value interface{}) *Component {
	a.Set("animate", value)
	return a
}

func (a *Component) Animation(value interface{}) *Component {
	a.Set("animation", value)
	return a
}

func (a *Component) Api(value interface{}) *Component {
	a.Set("api", value)
	return a
}

func (a *Component) Args(value interface{}) *Component {
	a.Set("args", value)
	return a
}

func (a *Component) AsBase64(value interface{}) *Component {
	a.Set("asBase64", value)
	return a
}

func (a *Component) AsBlob(value interface{}) *Component {
	a.Set("asBlob", value)
	return a
}

func (a *Component) Aside(value interface{}) *Component {
	a.Set("aside", value)
	return a
}

func (a *Component) AsideClassName(value interface{}) *Component {
	a.Set("asideClassName", value)
	return a
}

func (a *Component) AsideMaxWidth(value interface{}) *Component {
	a.Set("asideMaxWidth", value)
	return a
}

func (a *Component) AsideMinWidth(value interface{}) *Component {
	a.Set("asideMinWidth", value)
	return a
}

func (a *Component) AsidePosition(value interface{}) *Component {
	a.Set("asidePosition", value)
	return a
}

func (a *Component) AsideResizor(value interface{}) *Component {
	a.Set("asideResizor", value)
	return a
}

func (a *Component) AsideSticky(value interface{}) *Component {
	a.Set("asideSticky", value)
	return a
}

func (a *Component) AspectRatio(value interface{}) *Component {
	a.Set("aspectRatio", value)
	return a
}

func (a *Component) AsyncApi(value interface{}) *Component {
	a.Set("asyncApi", value)
	return a
}

func (a *Component) AttachDataToQuery(value interface{}) *Component {
	a.Set("attachDataToQuery", value)
	return a
}

func (a *Component) Auto(value interface{}) *Component {
	a.Set("auto", value)
	return a
}

func (a *Component) AutoCancelParent(value interface{}) *Component {
	a.Set("autoCancelParent", value)
	return a
}

func (a *Component) AutoCheckChildren(value interface{}) *Component {
	a.Set("autoCheckChildren", value)
	return a
}

func (a *Component) AutoComplete(value interface{}) *Component {
	a.Set("autoComplete", value)
	return a
}

func (a *Component) AutoFill(value interface{}) *Component {
	a.Set("autoFill", value)
	return a
}

func (a *Component) AutoFillHeight(value interface{}) *Component {
	a.Set("autoFillHeight", value)
	return a
}

func (a *Component) AutoFocus(value interface{}) *Component {
	a.Set("autoFocus", value)
	return a
}

func (a *Component) AutoGenerateFilter(value interface{}) *Component {
	a.Set("autoGenerateFilter", value)
	return a
}

func (a *Component) AutoJumpToTopOnPagerChange(value interface{}) *Component {
	a.Set("autoJumpToTopOnPagerChange", value)
	return a
}

func (a *Component) AutoPlay(value interface{}) *Component {
	a.Set("autoPlay", value)
	return a
}

func (a *Component) AutoRefresh(value interface{}) *Component {
	a.Set("autoRefresh", value)
	return a
}

func (a *Component) AutoScroll(value interface{}) *Component {
	a.Set("autoScroll", value)
	return a
}

func (a *Component) AutoSelectCurrentLoc(value interface{}) *Component {
	a.Set("autoSelectCurrentLoc", value)
	return a
}

func (a *Component) AutoSet(value interface{}) *Component {
	a.Set("autoSet", value)
	return a
}

func (a *Component) AutoUpload(value interface{}) *Component {
	a.Set("autoUpload", value)
	return a
}

func (a *Component) Avatar(value interface{}) *Component {
	a.Set("avatar", value)
	return a
}

func (a *Component) BackgroundColor(value interface{}) *Component {
	a.Set("backgroundColor", value)
	return a
}

func (a *Component) Badge(value interface{}) *Component {
	a.Set("badge", value)
	return a
}

func (a *Component) Bcc(value interface{}) *Component {
	a.Set("bcc", value)
	return a
}

func (a *Component) BgColor(value interface{}) *Component {
	a.Set("bgColor", value)
	return a
}

func (a *Component) Big(value interface{}) *Component {
	a.Set("big", value)
	return a
}

func (a *Component) Blank(value interface{}) *Component {
	a.Set("blank", value)
	return a
}

func (a *Component) Block(value interface{}) *Component {
	a.Set("block", value)
	return a
}

func (a *Component) Body(value interface{}) *Component {
	a.Set("body", value)
	return a
}

func (a *Component) BodyClassName(value interface{}) *Component {
	a.Set("bodyClassName", value)
	return a
}

func (a *Component) BodyControlClassName(value interface{}) *Component {
	a.Set("bodyControlClassName", value)
	return a
}

func (a *Component) BodyWidth(value interface{}) *Component {
	a.Set("bodyWidth", value)
	return a
}

func (a *Component) Border(value interface{}) *Component {
	a.Set("border", value)
	return a
}

func (a *Component) BorderMode(value interface{}) *Component {
	a.Set("borderMode", value)
	return a
}

func (a *Component) Bordered(value interface{}) *Component {
	a.Set("bordered", value)
	return a
}

func (a *Component) Breakpoint(value interface{}) *Component {
	a.Set("breakpoint", value)
	return a
}

func (a *Component) BtnActiveClassName(value interface{}) *Component {
	a.Set("btnActiveClassName", value)
	return a
}

func (a *Component) BtnActiveLevel(value interface{}) *Component {
	a.Set("btnActiveLevel", value)
	return a
}

func (a *Component) BtnClassName(value interface{}) *Component {
	a.Set("btnClassName", value)
	return a
}

func (a *Component) BtnLabel(value interface{}) *Component {
	a.Set("btnLabel", value)
	return a
}

func (a *Component) BtnLevel(value interface{}) *Component {
	a.Set("btnLevel", value)
	return a
}

func (a *Component) BtnText(value interface{}) *Component {
	a.Set("btnText", value)
	return a
}

func (a *Component) BtnUploadClassName(value interface{}) *Component {
	a.Set("btnUploadClassName", value)
	return a
}

func (a *Component) BuilderMode(value interface{}) *Component {
	a.Set("builderMode", value)
	return a
}

func (a *Component) BulkActions(value interface{}) *Component {
	a.Set("bulkActions", value)
	return a
}

func (a *Component) BulkSubmit(value interface{}) *Component {
	a.Set("bulkSubmit", value)
	return a
}

func (a *Component) Buttons(value interface{}) *Component {
	a.Set("buttons", value)
	return a
}

func (a *Component) Cache(value interface{}) *Component {
	a.Set("cache", value)
	return a
}

func (a *Component) CanAccessSuperData(value interface{}) *Component {
	a.Set("canAccessSuperData", value)
	return a
}

func (a *Component) CanRetryStatusCode(value interface{}) *Component {
	a.Set("canRetryStatusCode", value)
	return a
}

func (a *Component) CancelBtnIcon(value interface{}) *Component {
	a.Set("cancelBtnIcon", value)
	return a
}

func (a *Component) CancelBtnLabel(value interface{}) *Component {
	a.Set("cancelBtnLabel", value)
	return a
}

func (a *Component) Caption(value interface{}) *Component {
	a.Set("caption", value)
	return a
}

func (a *Component) Capture(value interface{}) *Component {
	a.Set("capture", value)
	return a
}

func (a *Component) Card(value interface{}) *Component {
	a.Set("card", value)
	return a
}

func (a *Component) CardSchema(value interface{}) *Component {
	a.Set("cardSchema", value)
	return a
}

func (a *Component) Cascade(value interface{}) *Component {
	a.Set("cascade", value)
	return a
}

func (a *Component) Cc(value interface{}) *Component {
	a.Set("cc", value)
	return a
}

func (a *Component) Center(value interface{}) *Component {
	a.Set("center", value)
	return a
}

func (a *Component) Char(value interface{}) *Component {
	a.Set("char", value)
	return a
}

func (a *Component) CharClassName(value interface{}) *Component {
	a.Set("charClassName", value)
	return a
}

func (a *Component) ChartTheme(value interface{}) *Component {
	a.Set("chartTheme", value)
	return a
}

func (a *Component) ChartValueField(value interface{}) *Component {
	a.Set("chartValueField", value)
	return a
}

func (a *Component) CheckAll(value interface{}) *Component {
	a.Set("checkAll", value)
	return a
}

func (a *Component) CheckAllLabel(value interface{}) *Component {
	a.Set("checkAllLabel", value)
	return a
}

func (a *Component) CheckAllText(value interface{}) *Component {
	a.Set("checkAllText", value)
	return a
}

func (a *Component) CheckApi(value interface{}) *Component {
	a.Set("checkApi", value)
	return a
}

func (a *Component) CheckInterval(value interface{}) *Component {
	a.Set("checkInterval", value)
	return a
}

func (a *Component) CheckOnItemClick(value interface{}) *Component {
	a.Set("checkOnItemClick", value)
	return a
}

func (a *Component) Checkable(value interface{}) *Component {
	a.Set("checkable", value)
	return a
}

func (a *Component) Checked(value interface{}) *Component {
	a.Set("checked", value)
	return a
}

func (a *Component) Children(value interface{}) *Component {
	a.Set("children", value)
	return a
}

func (a *Component) ChildrenAddable(value interface{}) *Component {
	a.Set("childrenAddable", value)
	return a
}

func (a *Component) ChildrenColumnName(value interface{}) *Component {
	a.Set("childrenColumnName", value)
	return a
}

func (a *Component) ChunkApi(value interface{}) *Component {
	a.Set("chunkApi", value)
	return a
}

func (a *Component) ChunkSize(value interface{}) *Component {
	a.Set("chunkSize", value)
	return a
}

func (a *Component) ClassName(value interface{}) *Component {
	a.Set("className", value)
	return a
}

func (a *Component) ClassNameExpr(value interface{}) *Component {
	a.Set("classNameExpr", value)
	return a
}

func (a *Component) ClearAfterSubmit(value interface{}) *Component {
	a.Set("clearAfterSubmit", value)
	return a
}

func (a *Component) ClearAndSubmit(value interface{}) *Component {
	a.Set("clearAndSubmit", value)
	return a
}

func (a *Component) ClearBtnIcon(value interface{}) *Component {
	a.Set("clearBtnIcon", value)
	return a
}

func (a *Component) ClearBtnLabel(value interface{}) *Component {
	a.Set("clearBtnLabel", value)
	return a
}

func (a *Component) ClearPersistDataAfterSubmit(value interface{}) *Component {
	a.Set("clearPersistDataAfterSubmit", value)
	return a
}

func (a *Component) ClearValueOnEmpty(value interface{}) *Component {
	a.Set("clearValueOnEmpty", value)
	return a
}

func (a *Component) ClearValueOnHidden(value interface{}) *Component {
	a.Set("clearValueOnHidden", value)
	return a
}

func (a *Component) ClearValueOnSourceChange(value interface{}) *Component {
	a.Set("clearValueOnSourceChange", value)
	return a
}

func (a *Component) Clearable(value interface{}) *Component {
	a.Set("clearable", value)
	return a
}

func (a *Component) ClickAction(value interface{}) *Component {
	a.Set("clickAction", value)
	return a
}

func (a *Component) Closable(value interface{}) *Component {
	a.Set("closable", value)
	return a
}

func (a *Component) Close(value interface{}) *Component {
	a.Set("close", value)
	return a
}

func (a *Component) CloseButton(value interface{}) *Component {
	a.Set("closeButton", value)
	return a
}

func (a *Component) CloseButtonClassName(value interface{}) *Component {
	a.Set("closeButtonClassName", value)
	return a
}

func (a *Component) CloseIcon(value interface{}) *Component {
	a.Set("closeIcon", value)
	return a
}

func (a *Component) CloseOnClick(value interface{}) *Component {
	a.Set("closeOnClick", value)
	return a
}

func (a *Component) CloseOnEsc(value interface{}) *Component {
	a.Set("closeOnEsc", value)
	return a
}

func (a *Component) CloseOnOutside(value interface{}) *Component {
	a.Set("closeOnOutside", value)
	return a
}

func (a *Component) CloseOnSelect(value interface{}) *Component {
	a.Set("closeOnSelect", value)
	return a
}

func (a *Component) CodeSize(value interface{}) *Component {
	a.Set("codeSize", value)
	return a
}

func (a *Component) ColSpanExpr(value interface{}) *Component {
	a.Set("colSpanExpr", value)
	return a
}

func (a *Component) Collapsable(value interface{}) *Component {
	a.Set("collapsable", value)
	return a
}

func (a *Component) CollapseBtnLabel(value interface{}) *Component {
	a.Set("collapseBtnLabel", value)
	return a
}

func (a *Component) CollapseButton(value interface{}) *Component {
	a.Set("collapseButton", value)
	return a
}

func (a *Component) CollapseButtonText(value interface{}) *Component {
	a.Set("collapseButtonText", value)
	return a
}

func (a *Component) CollapseHeader(value interface{}) *Component {
	a.Set("collapseHeader", value)
	return a
}

func (a *Component) CollapseOnExceed(value interface{}) *Component {
	a.Set("collapseOnExceed", value)
	return a
}

func (a *Component) CollapseTitle(value interface{}) *Component {
	a.Set("collapseTitle", value)
	return a
}

func (a *Component) Collapsed(value interface{}) *Component {
	a.Set("collapsed", value)
	return a
}

func (a *Component) Collapsible(value interface{}) *Component {
	a.Set("collapsible", value)
	return a
}

func (a *Component) Color(value interface{}) *Component {
	a.Set("color", value)
	return a
}

func (a *Component) Colors(value interface{}) *Component {
	a.Set("colors", value)
	return a
}

func (a *Component) Cols(value interface{}) *Component {
	a.Set("cols", value)
	return a
}

func (a *Component) Column(value interface{}) *Component {
	a.Set("column", value)
	return a
}

func (a *Component) ColumnClassName(value interface{}) *Component {
	a.Set("columnClassName", value)
	return a
}

func (a *Component) ColumnCount(value interface{}) *Component {
	a.Set("columnCount", value)
	return a
}

func (a *Component) ColumnNum(value interface{}) *Component {
	a.Set("columnNum", value)
	return a
}

func (a *Component) ColumnWidth(value interface{}) *Component {
	a.Set("columnWidth", value)
	return a
}

func (a *Component) Columns(value interface{}) *Component {
	a.Set("columns", value)
	return a
}

func (a *Component) ColumnsCount(value interface{}) *Component {
	a.Set("columnsCount", value)
	return a
}

func (a *Component) ColumnsNum(value interface{}) *Component {
	a.Set("columnsNum", value)
	return a
}

func (a *Component) ColumnsTogglable(value interface{}) *Component {
	a.Set("columnsTogglable", value)
	return a
}

func (a *Component) CombineFromIndex(value interface{}) *Component {
	a.Set("combineFromIndex", value)
	return a
}

func (a *Component) CombineNum(value interface{}) *Component {
	a.Set("combineNum", value)
	return a
}

func (a *Component) ComponentId(value interface{}) *Component {
	a.Set("componentId", value)
	return a
}

func (a *Component) ComponentName(value interface{}) *Component {
	a.Set("componentName", value)
	return a
}

func (a *Component) Compress(value interface{}) *Component {
	a.Set("compress", value)
	return a
}

func (a *Component) CompressOptions(value interface{}) *Component {
	a.Set("compressOptions", value)
	return a
}

func (a *Component) ConcatDataFields(value interface{}) *Component {
	a.Set("concatDataFields", value)
	return a
}

func (a *Component) Concurrency(value interface{}) *Component {
	a.Set("concurrency", value)
	return a
}

func (a *Component) Condition(value interface{}) *Component {
	a.Set("condition", value)
	return a
}

func (a *Component) Conditions(value interface{}) *Component {
	a.Set("conditions", value)
	return a
}

func (a *Component) Config(value interface{}) *Component {
	a.Set("config", value)
	return a
}

func (a *Component) Confirm(value interface{}) *Component {
	a.Set("confirm", value)
	return a
}

func (a *Component) ConfirmBtnIcon(value interface{}) *Component {
	a.Set("confirmBtnIcon", value)
	return a
}

func (a *Component) ConfirmBtnLabel(value interface{}) *Component {
	a.Set("confirmBtnLabel", value)
	return a
}

func (a *Component) ConfirmText(value interface{}) *Component {
	a.Set("confirmText", value)
	return a
}

func (a *Component) ConfirmTitle(value interface{}) *Component {
	a.Set("confirmTitle", value)
	return a
}

func (a *Component) Conjunction(value interface{}) *Component {
	a.Set("conjunction", value)
	return a
}

func (a *Component) Connector(value interface{}) *Component {
	a.Set("connector", value)
	return a
}

func (a *Component) Content(value interface{}) *Component {
	a.Set("content", value)
	return a
}

func (a *Component) ContentClassName(value interface{}) *Component {
	a.Set("contentClassName", value)
	return a
}

func (a *Component) ContentStyle(value interface{}) *Component {
	a.Set("contentStyle", value)
	return a
}

func (a *Component) Controls(value interface{}) *Component {
	a.Set("controls", value)
	return a
}

func (a *Component) ControlsTheme(value interface{}) *Component {
	a.Set("controlsTheme", value)
	return a
}

func (a *Component) ConvertKeyToPath(value interface{}) *Component {
	a.Set("convertKeyToPath", value)
	return a
}

func (a *Component) Copy(value interface{}) *Component {
	a.Set("copy", value)
	return a
}

func (a *Component) CopyAddBtn(value interface{}) *Component {
	a.Set("copyAddBtn", value)
	return a
}

func (a *Component) CopyBtnIcon(value interface{}) *Component {
	a.Set("copyBtnIcon", value)
	return a
}

func (a *Component) CopyBtnLabel(value interface{}) *Component {
	a.Set("copyBtnLabel", value)
	return a
}

func (a *Component) CopyData(value interface{}) *Component {
	a.Set("copyData", value)
	return a
}

func (a *Component) Copyable(value interface{}) *Component {
	a.Set("copyable", value)
	return a
}

func (a *Component) Count(value interface{}) *Component {
	a.Set("count", value)
	return a
}

func (a *Component) CountDown(value interface{}) *Component {
	a.Set("countDown", value)
	return a
}

func (a *Component) CountDownTpl(value interface{}) *Component {
	a.Set("countDownTpl", value)
	return a
}

func (a *Component) Creatable(value interface{}) *Component {
	a.Set("creatable", value)
	return a
}

func (a *Component) CreateBtnLabel(value interface{}) *Component {
	a.Set("createBtnLabel", value)
	return a
}

func (a *Component) Crop(value interface{}) *Component {
	a.Set("crop", value)
	return a
}

func (a *Component) CropFormat(value interface{}) *Component {
	a.Set("cropFormat", value)
	return a
}

func (a *Component) CropQuality(value interface{}) *Component {
	a.Set("cropQuality", value)
	return a
}

func (a *Component) CrossOrigin(value interface{}) *Component {
	a.Set("crossOrigin", value)
	return a
}

func (a *Component) Css(value interface{}) *Component {
	a.Set("css", value)
	return a
}

func (a *Component) CssVars(value interface{}) *Component {
	a.Set("cssVars", value)
	return a
}

func (a *Component) CurrentIndex(value interface{}) *Component {
	a.Set("currentIndex", value)
	return a
}

func (a *Component) Data(value interface{}) *Component {
	a.Set("data", value)
	return a
}

func (a *Component) DataFilter(value interface{}) *Component {
	a.Set("dataFilter", value)
	return a
}

func (a *Component) DataMergeMode(value interface{}) *Component {
	a.Set("dataMergeMode", value)
	return a
}

func (a *Component) DataProvider(value interface{}) *Component {
	a.Set("dataProvider", value)
	return a
}

func (a *Component) DataType(value interface{}) *Component {
	a.Set("dataType", value)
	return a
}

func (a *Component) Debug(value interface{}) *Component {
	a.Set("debug", value)
	return a
}

func (a *Component) DebugConfig(value interface{}) *Component {
	a.Set("debugConfig", value)
	return a
}

func (a *Component) DefaultAvatar(value interface{}) *Component {
	a.Set("defaultAvatar", value)
	return a
}

func (a *Component) DefaultCheckAll(value interface{}) *Component {
	a.Set("defaultCheckAll", value)
	return a
}

func (a *Component) DefaultCollapsed(value interface{}) *Component {
	a.Set("defaultCollapsed", value)
	return a
}

func (a *Component) DefaultColor(value interface{}) *Component {
	a.Set("defaultColor", value)
	return a
}

func (a *Component) DefaultImage(value interface{}) *Component {
	a.Set("defaultImage", value)
	return a
}

func (a *Component) DefaultKey(value interface{}) *Component {
	a.Set("defaultKey", value)
	return a
}

func (a *Component) DefaultOpenLevel(value interface{}) *Component {
	a.Set("defaultOpenLevel", value)
	return a
}

func (a *Component) DefaultParams(value interface{}) *Component {
	a.Set("defaultParams", value)
	return a
}

func (a *Component) DefaultValue(value interface{}) *Component {
	a.Set("defaultValue", value)
	return a
}

func (a *Component) Defer(value interface{}) *Component {
	a.Set("defer", value)
	return a
}

func (a *Component) DeferApi(value interface{}) *Component {
	a.Set("deferApi", value)
	return a
}

func (a *Component) DeferField(value interface{}) *Component {
	a.Set("deferField", value)
	return a
}

func (a *Component) Definitions(value interface{}) *Component {
	a.Set("definitions", value)
	return a
}

func (a *Component) Delay(value interface{}) *Component {
	a.Set("delay", value)
	return a
}

func (a *Component) DeleteApi(value interface{}) *Component {
	a.Set("deleteApi", value)
	return a
}

func (a *Component) DeleteBtnIcon(value interface{}) *Component {
	a.Set("deleteBtnIcon", value)
	return a
}

func (a *Component) DeleteBtnLabel(value interface{}) *Component {
	a.Set("deleteBtnLabel", value)
	return a
}

func (a *Component) DeleteConfirmText(value interface{}) *Component {
	a.Set("deleteConfirmText", value)
	return a
}

func (a *Component) Delimiter(value interface{}) *Component {
	a.Set("delimiter", value)
	return a
}

func (a *Component) Desc(value interface{}) *Component {
	a.Set("desc", value)
	return a
}

func (a *Component) Description(value interface{}) *Component {
	a.Set("description", value)
	return a
}

func (a *Component) DescriptionClassName(value interface{}) *Component {
	a.Set("descriptionClassName", value)
	return a
}

func (a *Component) Detail(value interface{}) *Component {
	a.Set("detail", value)
	return a
}

func (a *Component) DetailClassName(value interface{}) *Component {
	a.Set("detailClassName", value)
	return a
}

func (a *Component) DetailCollapsedText(value interface{}) *Component {
	a.Set("detailCollapsedText", value)
	return a
}

func (a *Component) DetailExpandedText(value interface{}) *Component {
	a.Set("detailExpandedText", value)
	return a
}

func (a *Component) Dialog(value interface{}) *Component {
	a.Set("dialog", value)
	return a
}

func (a *Component) DialogType(value interface{}) *Component {
	a.Set("dialogType", value)
	return a
}

func (a *Component) DiffValue(value interface{}) *Component {
	a.Set("diffValue", value)
	return a
}

func (a *Component) Direction(value interface{}) *Component {
	a.Set("direction", value)
	return a
}

func (a *Component) DisableColor(value interface{}) *Component {
	a.Set("disableColor", value)
	return a
}

func (a *Component) DisableDataMapping(value interface{}) *Component {
	a.Set("disableDataMapping", value)
	return a
}

func (a *Component) DisableOn(value interface{}) *Component {
	a.Set("disableOn", value)
	return a
}

func (a *Component) Disabled(value interface{}) *Component {
	a.Set("disabled", value)
	return a
}

func (a *Component) DisabledDate(value interface{}) *Component {
	a.Set("disabledDate", value)
	return a
}

func (a *Component) DisabledOn(value interface{}) *Component {
	a.Set("disabledOn", value)
	return a
}

func (a *Component) DisabledOnAction(value interface{}) *Component {
	a.Set("disabledOnAction", value)
	return a
}

func (a *Component) DisabledTip(value interface{}) *Component {
	a.Set("disabledTip", value)
	return a
}

func (a *Component) DisabledTypes(value interface{}) *Component {
	a.Set("disabledTypes", value)
	return a
}

func (a *Component) DisplayDataTypes(value interface{}) *Component {
	a.Set("displayDataTypes", value)
	return a
}

func (a *Component) DisplayFormat(value interface{}) *Component {
	a.Set("displayFormat", value)
	return a
}

func (a *Component) DisplayMode(value interface{}) *Component {
	a.Set("displayMode", value)
	return a
}

func (a *Component) DisplayTimeZone(value interface{}) *Component {
	a.Set("displayTimeZone", value)
	return a
}

func (a *Component) DivideLine(value interface{}) *Component {
	a.Set("divideLine", value)
	return a
}

func (a *Component) Divider(value interface{}) *Component {
	a.Set("divider", value)
	return a
}

func (a *Component) DocumentLink(value interface{}) *Component {
	a.Set("documentLink", value)
	return a
}

func (a *Component) Documentation(value interface{}) *Component {
	a.Set("documentation", value)
	return a
}

func (a *Component) DotSize(value interface{}) *Component {
	a.Set("dotSize", value)
	return a
}

func (a *Component) DownloadFileName(value interface{}) *Component {
	a.Set("downloadFileName", value)
	return a
}

func (a *Component) DownloadUrl(value interface{}) *Component {
	a.Set("downloadUrl", value)
	return a
}

func (a *Component) Drag(value interface{}) *Component {
	a.Set("drag", value)
	return a
}

func (a *Component) DragOnSameLevel(value interface{}) *Component {
	a.Set("dragOnSameLevel", value)
	return a
}

func (a *Component) Draggable(value interface{}) *Component {
	a.Set("draggable", value)
	return a
}

func (a *Component) DraggableConfig(value interface{}) *Component {
	a.Set("draggableConfig", value)
	return a
}

func (a *Component) DraggableOn(value interface{}) *Component {
	a.Set("draggableOn", value)
	return a
}

func (a *Component) DraggableTip(value interface{}) *Component {
	a.Set("draggableTip", value)
	return a
}

func (a *Component) Drawer(value interface{}) *Component {
	a.Set("drawer", value)
	return a
}

func (a *Component) DropCrop(value interface{}) *Component {
	a.Set("dropCrop", value)
	return a
}

func (a *Component) Dropdown(value interface{}) *Component {
	a.Set("dropdown", value)
	return a
}

func (a *Component) DropdownClassName(value interface{}) *Component {
	a.Set("dropdownClassName", value)
	return a
}

func (a *Component) DropdownItemClassName(value interface{}) *Component {
	a.Set("dropdownItemClassName", value)
	return a
}

func (a *Component) Duration(value interface{}) *Component {
	a.Set("duration", value)
	return a
}

func (a *Component) EbmedCancelIcon(value interface{}) *Component {
	a.Set("ebmedCancelIcon", value)
	return a
}

func (a *Component) EbmedCancelLabel(value interface{}) *Component {
	a.Set("ebmedCancelLabel", value)
	return a
}

func (a *Component) EditApi(value interface{}) *Component {
	a.Set("editApi", value)
	return a
}

func (a *Component) EditBtnIcon(value interface{}) *Component {
	a.Set("editBtnIcon", value)
	return a
}

func (a *Component) EditBtnLabel(value interface{}) *Component {
	a.Set("editBtnLabel", value)
	return a
}

func (a *Component) EditControls(value interface{}) *Component {
	a.Set("editControls", value)
	return a
}

func (a *Component) EditDialog(value interface{}) *Component {
	a.Set("editDialog", value)
	return a
}

func (a *Component) Editable(value interface{}) *Component {
	a.Set("editable", value)
	return a
}

func (a *Component) EditorDidMount(value interface{}) *Component {
	a.Set("editorDidMount", value)
	return a
}

func (a *Component) EditorSetting(value interface{}) *Component {
	a.Set("editorSetting", value)
	return a
}

func (a *Component) EditorTheme(value interface{}) *Component {
	a.Set("editorTheme", value)
	return a
}

func (a *Component) EllipsisThreshold(value interface{}) *Component {
	a.Set("ellipsisThreshold", value)
	return a
}

func (a *Component) Embed(value interface{}) *Component {
	a.Set("embed", value)
	return a
}

func (a *Component) EmbedBtnIcon(value interface{}) *Component {
	a.Set("embedBtnIcon", value)
	return a
}

func (a *Component) EmbedBtnLabel(value interface{}) *Component {
	a.Set("embedBtnLabel", value)
	return a
}

func (a *Component) EmbedConfirmIcon(value interface{}) *Component {
	a.Set("embedConfirmIcon", value)
	return a
}

func (a *Component) EmbedConfirmLabel(value interface{}) *Component {
	a.Set("embedConfirmLabel", value)
	return a
}

func (a *Component) Emebed(value interface{}) *Component {
	a.Set("emebed", value)
	return a
}

func (a *Component) Enable(value interface{}) *Component {
	a.Set("enable", value)
	return a
}

func (a *Component) EnableAdvancedSetting(value interface{}) *Component {
	a.Set("enableAdvancedSetting", value)
	return a
}

func (a *Component) EnableBatchAdd(value interface{}) *Component {
	a.Set("enableBatchAdd", value)
	return a
}

func (a *Component) EnableBulkActions(value interface{}) *Component {
	a.Set("enableBulkActions", value)
	return a
}

func (a *Component) EnableBulkActionsOn(value interface{}) *Component {
	a.Set("enableBulkActionsOn", value)
	return a
}

func (a *Component) EnableClipboard(value interface{}) *Component {
	a.Set("enableClipboard", value)
	return a
}

func (a *Component) EnableDefaultIcon(value interface{}) *Component {
	a.Set("enableDefaultIcon", value)
	return a
}

func (a *Component) EnableFieldSetStyle(value interface{}) *Component {
	a.Set("enableFieldSetStyle", value)
	return a
}

func (a *Component) EnableNodePath(value interface{}) *Component {
	a.Set("enableNodePath", value)
	return a
}

func (a *Component) EnableStaticTransform(value interface{}) *Component {
	a.Set("enableStaticTransform", value)
	return a
}

func (a *Component) Encoding(value interface{}) *Component {
	a.Set("encoding", value)
	return a
}

func (a *Component) EndPlaceholder(value interface{}) *Component {
	a.Set("endPlaceholder", value)
	return a
}

func (a *Component) Enhance(value interface{}) *Component {
	a.Set("enhance", value)
	return a
}

func (a *Component) EnlargeAble(value interface{}) *Component {
	a.Set("enlargeAble", value)
	return a
}

func (a *Component) EnlargeWithGallary(value interface{}) *Component {
	a.Set("enlargeWithGallary", value)
	return a
}

func (a *Component) EnlargetWithImages(value interface{}) *Component {
	a.Set("enlargetWithImages", value)
	return a
}

func (a *Component) Enterable(value interface{}) *Component {
	a.Set("enterable", value)
	return a
}

func (a *Component) ErrorStatusCode(value interface{}) *Component {
	a.Set("errorStatusCode", value)
	return a
}

func (a *Component) Events(value interface{}) *Component {
	a.Set("events", value)
	return a
}

func (a *Component) Excavate(value interface{}) *Component {
	a.Set("excavate", value)
	return a
}

func (a *Component) ExcludeKeys(value interface{}) *Component {
	a.Set("excludeKeys", value)
	return a
}

func (a *Component) ExecOn(value interface{}) *Component {
	a.Set("execOn", value)
	return a
}

func (a *Component) ExpandConfig(value interface{}) *Component {
	a.Set("expandConfig", value)
	return a
}

func (a *Component) ExpandIcon(value interface{}) *Component {
	a.Set("expandIcon", value)
	return a
}

func (a *Component) ExpandIconPosition(value interface{}) *Component {
	a.Set("expandIconPosition", value)
	return a
}

func (a *Component) ExpandPosition(value interface{}) *Component {
	a.Set("expandPosition", value)
	return a
}

func (a *Component) Expandable(value interface{}) *Component {
	a.Set("expandable", value)
	return a
}

func (a *Component) ExpandableOn(value interface{}) *Component {
	a.Set("expandableOn", value)
	return a
}

func (a *Component) ExpandedRowClassNameExpr(value interface{}) *Component {
	a.Set("expandedRowClassNameExpr", value)
	return a
}

func (a *Component) ExpandedRowKeys(value interface{}) *Component {
	a.Set("expandedRowKeys", value)
	return a
}

func (a *Component) ExpandedRowKeysExpr(value interface{}) *Component {
	a.Set("expandedRowKeysExpr", value)
	return a
}

func (a *Component) ExpendButton(value interface{}) *Component {
	a.Set("expendButton", value)
	return a
}

func (a *Component) ExpendButtonText(value interface{}) *Component {
	a.Set("expendButtonText", value)
	return a
}

func (a *Component) Expression(value interface{}) *Component {
	a.Set("expression", value)
	return a
}

func (a *Component) ExtraName(value interface{}) *Component {
	a.Set("extraName", value)
	return a
}

func (a *Component) ExtractValue(value interface{}) *Component {
	a.Set("extractValue", value)
	return a
}

func (a *Component) EyeBorderColor(value interface{}) *Component {
	a.Set("eyeBorderColor", value)
	return a
}

func (a *Component) EyeBorderSize(value interface{}) *Component {
	a.Set("eyeBorderSize", value)
	return a
}

func (a *Component) EyeInnerColor(value interface{}) *Component {
	a.Set("eyeInnerColor", value)
	return a
}

func (a *Component) EyeType(value interface{}) *Component {
	a.Set("eyeType", value)
	return a
}

func (a *Component) FalseValue(value interface{}) *Component {
	a.Set("falseValue", value)
	return a
}

func (a *Component) Feedback(value interface{}) *Component {
	a.Set("feedback", value)
	return a
}

func (a *Component) FetchFailed(value interface{}) *Component {
	a.Set("fetchFailed", value)
	return a
}

func (a *Component) FetchOn(value interface{}) *Component {
	a.Set("fetchOn", value)
	return a
}

func (a *Component) FetchSuccess(value interface{}) *Component {
	a.Set("fetchSuccess", value)
	return a
}

func (a *Component) FieldSet(value interface{}) *Component {
	a.Set("fieldSet", value)
	return a
}

func (a *Component) Fields(value interface{}) *Component {
	a.Set("fields", value)
	return a
}

func (a *Component) FileField(value interface{}) *Component {
	a.Set("fileField", value)
	return a
}

func (a *Component) Filter(value interface{}) *Component {
	a.Set("filter", value)
	return a
}

func (a *Component) FilterDefaultVisible(value interface{}) *Component {
	a.Set("filterDefaultVisible", value)
	return a
}

func (a *Component) FilterTogglable(value interface{}) *Component {
	a.Set("filterTogglable", value)
	return a
}

func (a *Component) Filterable(value interface{}) *Component {
	a.Set("filterable", value)
	return a
}

func (a *Component) FinishChunkApi(value interface{}) *Component {
	a.Set("finishChunkApi", value)
	return a
}

func (a *Component) FinishStatusCode(value interface{}) *Component {
	a.Set("finishStatusCode", value)
	return a
}

func (a *Component) FinishedField(value interface{}) *Component {
	a.Set("finishedField", value)
	return a
}

func (a *Component) Fit(value interface{}) *Component {
	a.Set("fit", value)
	return a
}

func (a *Component) Fixed(value interface{}) *Component {
	a.Set("fixed", value)
	return a
}

func (a *Component) FixedSize(value interface{}) *Component {
	a.Set("fixedSize", value)
	return a
}

func (a *Component) FixedSizeClassName(value interface{}) *Component {
	a.Set("fixedSizeClassName", value)
	return a
}

func (a *Component) Flat(value interface{}) *Component {
	a.Set("flat", value)
	return a
}

func (a *Component) Font(value interface{}) *Component {
	a.Set("font", value)
	return a
}

func (a *Component) FontStyle(value interface{}) *Component {
	a.Set("fontStyle", value)
	return a
}

func (a *Component) Footable(value interface{}) *Component {
	a.Set("footable", value)
	return a
}

func (a *Component) Footer(value interface{}) *Component {
	a.Set("footer", value)
	return a
}

func (a *Component) FooterAddBtn(value interface{}) *Component {
	a.Set("footerAddBtn", value)
	return a
}

func (a *Component) FooterClassName(value interface{}) *Component {
	a.Set("footerClassName", value)
	return a
}

func (a *Component) FooterToolbar(value interface{}) *Component {
	a.Set("footerToolbar", value)
	return a
}

func (a *Component) FooterToolbarClassName(value interface{}) *Component {
	a.Set("footerToolbarClassName", value)
	return a
}

func (a *Component) FooterWrapClassName(value interface{}) *Component {
	a.Set("footerWrapClassName", value)
	return a
}

func (a *Component) ForceAppendDataToQuery(value interface{}) *Component {
	a.Set("forceAppendDataToQuery", value)
	return a
}

func (a *Component) ForegroundColor(value interface{}) *Component {
	a.Set("foregroundColor", value)
	return a
}

func (a *Component) Form(value interface{}) *Component {
	a.Set("form", value)
	return a
}

func (a *Component) FormClassName(value interface{}) *Component {
	a.Set("formClassName", value)
	return a
}

func (a *Component) Format(value interface{}) *Component {
	a.Set("format", value)
	return a
}

func (a *Component) Formula(value interface{}) *Component {
	a.Set("formula", value)
	return a
}

func (a *Component) FormulaForIf(value interface{}) *Component {
	a.Set("formulaForIf", value)
	return a
}

func (a *Component) FrameImage(value interface{}) *Component {
	a.Set("frameImage", value)
	return a
}

func (a *Component) Frames(value interface{}) *Component {
	a.Set("frames", value)
	return a
}

func (a *Component) FramesClassName(value interface{}) *Component {
	a.Set("framesClassName", value)
	return a
}

func (a *Component) FromNow(value interface{}) *Component {
	a.Set("fromNow", value)
	return a
}

func (a *Component) FullThumbMode(value interface{}) *Component {
	a.Set("fullThumbMode", value)
	return a
}

func (a *Component) Funcs(value interface{}) *Component {
	a.Set("funcs", value)
	return a
}

func (a *Component) Gap(value interface{}) *Component {
	a.Set("gap", value)
	return a
}

func (a *Component) GapDegree(value interface{}) *Component {
	a.Set("gapDegree", value)
	return a
}

func (a *Component) GapPosition(value interface{}) *Component {
	a.Set("gapPosition", value)
	return a
}

func (a *Component) GapRow(value interface{}) *Component {
	a.Set("gapRow", value)
	return a
}

func (a *Component) GetLocationPlaceholder(value interface{}) *Component {
	a.Set("getLocationPlaceholder", value)
	return a
}

func (a *Component) Grids(value interface{}) *Component {
	a.Set("grids", value)
	return a
}

func (a *Component) Gutter(value interface{}) *Component {
	a.Set("gutter", value)
	return a
}

func (a *Component) Half(value interface{}) *Component {
	a.Set("half", value)
	return a
}

func (a *Component) HasNext(value interface{}) *Component {
	a.Set("hasNext", value)
	return a
}

func (a *Component) Hash(value interface{}) *Component {
	a.Set("hash", value)
	return a
}

func (a *Component) Header(value interface{}) *Component {
	a.Set("header", value)
	return a
}

func (a *Component) HeaderAlign(value interface{}) *Component {
	a.Set("headerAlign", value)
	return a
}

func (a *Component) HeaderClassName(value interface{}) *Component {
	a.Set("headerClassName", value)
	return a
}

func (a *Component) HeaderControlClassName(value interface{}) *Component {
	a.Set("headerControlClassName", value)
	return a
}

func (a *Component) HeaderPosition(value interface{}) *Component {
	a.Set("headerPosition", value)
	return a
}

func (a *Component) HeaderToolbar(value interface{}) *Component {
	a.Set("headerToolbar", value)
	return a
}

func (a *Component) HeaderToolbarClassName(value interface{}) *Component {
	a.Set("headerToolbarClassName", value)
	return a
}

func (a *Component) Headers(value interface{}) *Component {
	a.Set("headers", value)
	return a
}

func (a *Component) HeadingClassName(value interface{}) *Component {
	a.Set("headingClassName", value)
	return a
}

func (a *Component) Height(value interface{}) *Component {
	a.Set("height", value)
	return a
}

func (a *Component) HeightAuto(value interface{}) *Component {
	a.Set("heightAuto", value)
	return a
}

func (a *Component) Hidden(value interface{}) *Component {
	a.Set("hidden", value)
	return a
}

func (a *Component) HiddenOn(value interface{}) *Component {
	a.Set("hiddenOn", value)
	return a
}

func (a *Component) HideCaret(value interface{}) *Component {
	a.Set("hideCaret", value)
	return a
}

func (a *Component) HideCheckToggler(value interface{}) *Component {
	a.Set("hideCheckToggler", value)
	return a
}

func (a *Component) HideDot(value interface{}) *Component {
	a.Set("hideDot", value)
	return a
}

func (a *Component) HideHeader(value interface{}) *Component {
	a.Set("hideHeader", value)
	return a
}

func (a *Component) HideNodePathLabel(value interface{}) *Component {
	a.Set("hideNodePathLabel", value)
	return a
}

func (a *Component) HideQuickSaveBtn(value interface{}) *Component {
	a.Set("hideQuickSaveBtn", value)
	return a
}

func (a *Component) HideRoot(value interface{}) *Component {
	a.Set("hideRoot", value)
	return a
}

func (a *Component) HideUploadButton(value interface{}) *Component {
	a.Set("hideUploadButton", value)
	return a
}

func (a *Component) HideViewControl(value interface{}) *Component {
	a.Set("hideViewControl", value)
	return a
}

func (a *Component) HighlightTxt(value interface{}) *Component {
	a.Set("highlightTxt", value)
	return a
}

func (a *Component) Hint(value interface{}) *Component {
	a.Set("hint", value)
	return a
}

func (a *Component) Horizontal(value interface{}) *Component {
	a.Set("horizontal", value)
	return a
}

func (a *Component) HotKey(value interface{}) *Component {
	a.Set("hotKey", value)
	return a
}

func (a *Component) HoverMode(value interface{}) *Component {
	a.Set("hoverMode", value)
	return a
}

func (a *Component) Href(value interface{}) *Component {
	a.Set("href", value)
	return a
}

func (a *Component) Html(value interface{}) *Component {
	a.Set("html", value)
	return a
}

func (a *Component) HtmlTarget(value interface{}) *Component {
	a.Set("htmlTarget", value)
	return a
}

func (a *Component) Icon(value interface{}) *Component {
	a.Set("icon", value)
	return a
}

func (a *Component) IconClassName(value interface{}) *Component {
	a.Set("iconClassName", value)
	return a
}

func (a *Component) IconOnly(value interface{}) *Component {
	a.Set("iconOnly", value)
	return a
}

func (a *Component) IconPosition(value interface{}) *Component {
	a.Set("iconPosition", value)
	return a
}

func (a *Component) IconRatio(value interface{}) *Component {
	a.Set("iconRatio", value)
	return a
}

func (a *Component) IconStyle(value interface{}) *Component {
	a.Set("iconStyle", value)
	return a
}

func (a *Component) Icons(value interface{}) *Component {
	a.Set("icons", value)
	return a
}

func (a *Component) Id(value interface{}) *Component {
	a.Set("id", value)
	return a
}

func (a *Component) If(value interface{}) *Component {
	a.Set("if", value)
	return a
}

func (a *Component) IgnoreConfirm(value interface{}) *Component {
	a.Set("ignoreConfirm", value)
	return a
}

func (a *Component) IgnoreError(value interface{}) *Component {
	a.Set("ignoreError", value)
	return a
}

func (a *Component) Image(value interface{}) *Component {
	a.Set("image", value)
	return a
}

func (a *Component) ImageCaption(value interface{}) *Component {
	a.Set("imageCaption", value)
	return a
}

func (a *Component) ImageClassName(value interface{}) *Component {
	a.Set("imageClassName", value)
	return a
}

func (a *Component) ImageGallaryClassName(value interface{}) *Component {
	a.Set("imageGallaryClassName", value)
	return a
}

func (a *Component) ImageMode(value interface{}) *Component {
	a.Set("imageMode", value)
	return a
}

func (a *Component) ImageSettings(value interface{}) *Component {
	a.Set("imageSettings", value)
	return a
}

func (a *Component) InTag(value interface{}) *Component {
	a.Set("inTag", value)
	return a
}

func (a *Component) InactiveColor(value interface{}) *Component {
	a.Set("inactiveColor", value)
	return a
}

func (a *Component) IncludeEmpty(value interface{}) *Component {
	a.Set("includeEmpty", value)
	return a
}

func (a *Component) IndentSize(value interface{}) *Component {
	a.Set("indentSize", value)
	return a
}

func (a *Component) IndexBarOffset(value interface{}) *Component {
	a.Set("indexBarOffset", value)
	return a
}

func (a *Component) IndexField(value interface{}) *Component {
	a.Set("indexField", value)
	return a
}

func (a *Component) IndexKeyName(value interface{}) *Component {
	a.Set("indexKeyName", value)
	return a
}

func (a *Component) Inherit(value interface{}) *Component {
	a.Set("inherit", value)
	return a
}

func (a *Component) InitApi(value interface{}) *Component {
	a.Set("initApi", value)
	return a
}

func (a *Component) InitAsyncApi(value interface{}) *Component {
	a.Set("initAsyncApi", value)
	return a
}

func (a *Component) InitAutoFill(value interface{}) *Component {
	a.Set("initAutoFill", value)
	return a
}

func (a *Component) InitCheckInterval(value interface{}) *Component {
	a.Set("initCheckInterval", value)
	return a
}

func (a *Component) InitCrop(value interface{}) *Component {
	a.Set("initCrop", value)
	return a
}

func (a *Component) InitFetch(value interface{}) *Component {
	a.Set("initFetch", value)
	return a
}

func (a *Component) InitFetchOn(value interface{}) *Component {
	a.Set("initFetchOn", value)
	return a
}

func (a *Component) InitFetchSchema(value interface{}) *Component {
	a.Set("initFetchSchema", value)
	return a
}

func (a *Component) InitFetchSchemaOn(value interface{}) *Component {
	a.Set("initFetchSchemaOn", value)
	return a
}

func (a *Component) InitFinishedField(value interface{}) *Component {
	a.Set("initFinishedField", value)
	return a
}

func (a *Component) InitSet(value interface{}) *Component {
	a.Set("initSet", value)
	return a
}

func (a *Component) InitialStatusCode(value interface{}) *Component {
	a.Set("initialStatusCode", value)
	return a
}

func (a *Component) InitiallyOpen(value interface{}) *Component {
	a.Set("initiallyOpen", value)
	return a
}

func (a *Component) Inline(value interface{}) *Component {
	a.Set("inline", value)
	return a
}

func (a *Component) InnerClassName(value interface{}) *Component {
	a.Set("innerClassName", value)
	return a
}

func (a *Component) InnerStyle(value interface{}) *Component {
	a.Set("innerStyle", value)
	return a
}

func (a *Component) InputClassName(value interface{}) *Component {
	a.Set("inputClassName", value)
	return a
}

func (a *Component) InputControlClassName(value interface{}) *Component {
	a.Set("inputControlClassName", value)
	return a
}

func (a *Component) InputForbid(value interface{}) *Component {
	a.Set("inputForbid", value)
	return a
}

func (a *Component) InputFormat(value interface{}) *Component {
	a.Set("inputFormat", value)
	return a
}

func (a *Component) InputName(value interface{}) *Component {
	a.Set("inputName", value)
	return a
}

func (a *Component) InputParams(value interface{}) *Component {
	a.Set("inputParams", value)
	return a
}

func (a *Component) InsertKeys(value interface{}) *Component {
	a.Set("insertKeys", value)
	return a
}

func (a *Component) Interval(value interface{}) *Component {
	a.Set("interval", value)
	return a
}

func (a *Component) InvalidSizeMessage(value interface{}) *Component {
	a.Set("invalidSizeMessage", value)
	return a
}

func (a *Component) InvalidTypeMessage(value interface{}) *Component {
	a.Set("invalidTypeMessage", value)
	return a
}

func (a *Component) IsEndDate(value interface{}) *Component {
	a.Set("isEndDate", value)
	return a
}

func (a *Component) IsLive(value interface{}) *Component {
	a.Set("isLive", value)
	return a
}

func (a *Component) IsolateScope(value interface{}) *Component {
	a.Set("isolateScope", value)
	return a
}

func (a *Component) ItemAction(value interface{}) *Component {
	a.Set("itemAction", value)
	return a
}

func (a *Component) ItemActions(value interface{}) *Component {
	a.Set("itemActions", value)
	return a
}

func (a *Component) ItemBadge(value interface{}) *Component {
	a.Set("itemBadge", value)
	return a
}

func (a *Component) ItemCheckableOn(value interface{}) *Component {
	a.Set("itemCheckableOn", value)
	return a
}

func (a *Component) ItemClassName(value interface{}) *Component {
	a.Set("itemClassName", value)
	return a
}

func (a *Component) ItemClearable(value interface{}) *Component {
	a.Set("itemClearable", value)
	return a
}

func (a *Component) ItemDraggableOn(value interface{}) *Component {
	a.Set("itemDraggableOn", value)
	return a
}

func (a *Component) ItemHeight(value interface{}) *Component {
	a.Set("itemHeight", value)
	return a
}

func (a *Component) ItemKeyName(value interface{}) *Component {
	a.Set("itemKeyName", value)
	return a
}

func (a *Component) ItemSchema(value interface{}) *Component {
	a.Set("itemSchema", value)
	return a
}

func (a *Component) ItemTitleSchema(value interface{}) *Component {
	a.Set("itemTitleSchema", value)
	return a
}

func (a *Component) ItemWidth(value interface{}) *Component {
	a.Set("itemWidth", value)
	return a
}

func (a *Component) Items(value interface{}) *Component {
	a.Set("items", value)
	return a
}

func (a *Component) ItemsClassName(value interface{}) *Component {
	a.Set("itemsClassName", value)
	return a
}

func (a *Component) JoinValues(value interface{}) *Component {
	a.Set("joinValues", value)
	return a
}

func (a *Component) JumpBufferDuration(value interface{}) *Component {
	a.Set("jumpBufferDuration", value)
	return a
}

func (a *Component) JumpFrame(value interface{}) *Component {
	a.Set("jumpFrame", value)
	return a
}

func (a *Component) Jumpable(value interface{}) *Component {
	a.Set("jumpable", value)
	return a
}

func (a *Component) JumpableOn(value interface{}) *Component {
	a.Set("jumpableOn", value)
	return a
}

func (a *Component) Justify(value interface{}) *Component {
	a.Set("justify", value)
	return a
}

func (a *Component) KeepItemSelectionOnPageChange(value interface{}) *Component {
	a.Set("keepItemSelectionOnPageChange", value)
	return a
}

func (a *Component) Key(value interface{}) *Component {
	a.Set("key", value)
	return a
}

func (a *Component) KeyField(value interface{}) *Component {
	a.Set("keyField", value)
	return a
}

func (a *Component) KeyPlaceholder(value interface{}) *Component {
	a.Set("keyPlaceholder", value)
	return a
}

func (a *Component) Keyboard(value interface{}) *Component {
	a.Set("keyboard", value)
	return a
}

func (a *Component) KilobitSeparator(value interface{}) *Component {
	a.Set("kilobitSeparator", value)
	return a
}

func (a *Component) Label(value interface{}) *Component {
	a.Set("label", value)
	return a
}

func (a *Component) LabelAlign(value interface{}) *Component {
	a.Set("labelAlign", value)
	return a
}

func (a *Component) LabelClassName(value interface{}) *Component {
	a.Set("labelClassName", value)
	return a
}

func (a *Component) LabelField(value interface{}) *Component {
	a.Set("labelField", value)
	return a
}

func (a *Component) LabelMap(value interface{}) *Component {
	a.Set("labelMap", value)
	return a
}

func (a *Component) LabelMaxLength(value interface{}) *Component {
	a.Set("labelMaxLength", value)
	return a
}

func (a *Component) LabelPlacement(value interface{}) *Component {
	a.Set("labelPlacement", value)
	return a
}

func (a *Component) LabelRemark(value interface{}) *Component {
	a.Set("labelRemark", value)
	return a
}

func (a *Component) LabelStyle(value interface{}) *Component {
	a.Set("labelStyle", value)
	return a
}

func (a *Component) LabelTpl(value interface{}) *Component {
	a.Set("labelTpl", value)
	return a
}

func (a *Component) LabelWidth(value interface{}) *Component {
	a.Set("labelWidth", value)
	return a
}

func (a *Component) Language(value interface{}) *Component {
	a.Set("language", value)
	return a
}

func (a *Component) LargeMode(value interface{}) *Component {
	a.Set("largeMode", value)
	return a
}

func (a *Component) Layout(value interface{}) *Component {
	a.Set("layout", value)
	return a
}

func (a *Component) LazyLoad(value interface{}) *Component {
	a.Set("lazyLoad", value)
	return a
}

func (a *Component) LazyRenderAfter(value interface{}) *Component {
	a.Set("lazyRenderAfter", value)
	return a
}

func (a *Component) Left(value interface{}) *Component {
	a.Set("left", value)
	return a
}

func (a *Component) LeftMode(value interface{}) *Component {
	a.Set("leftMode", value)
	return a
}

func (a *Component) LeftOptions(value interface{}) *Component {
	a.Set("leftOptions", value)
	return a
}

func (a *Component) Length(value interface{}) *Component {
	a.Set("length", value)
	return a
}

func (a *Component) Level(value interface{}) *Component {
	a.Set("level", value)
	return a
}

func (a *Component) LevelExpand(value interface{}) *Component {
	a.Set("levelExpand", value)
	return a
}

func (a *Component) Lg(value interface{}) *Component {
	a.Set("lg", value)
	return a
}

func (a *Component) Limit(value interface{}) *Component {
	a.Set("limit", value)
	return a
}

func (a *Component) LineColor(value interface{}) *Component {
	a.Set("lineColor", value)
	return a
}

func (a *Component) LineHeight(value interface{}) *Component {
	a.Set("lineHeight", value)
	return a
}

func (a *Component) LineStyle(value interface{}) *Component {
	a.Set("lineStyle", value)
	return a
}

func (a *Component) Link(value interface{}) *Component {
	a.Set("link", value)
	return a
}

func (a *Component) LinkClassName(value interface{}) *Component {
	a.Set("linkClassName", value)
	return a
}

func (a *Component) Links(value interface{}) *Component {
	a.Set("links", value)
	return a
}

func (a *Component) LinksClassName(value interface{}) *Component {
	a.Set("linksClassName", value)
	return a
}

func (a *Component) ListClassName(value interface{}) *Component {
	a.Set("listClassName", value)
	return a
}

func (a *Component) ListItem(value interface{}) *Component {
	a.Set("listItem", value)
	return a
}

func (a *Component) LoadBaiduMap(value interface{}) *Component {
	a.Set("loadBaiduMap", value)
	return a
}

func (a *Component) LoadDataOnce(value interface{}) *Component {
	a.Set("loadDataOnce", value)
	return a
}

func (a *Component) LoadDataOnceFetchOnFilter(value interface{}) *Component {
	a.Set("loadDataOnceFetchOnFilter", value)
	return a
}

func (a *Component) LoadMoreProps(value interface{}) *Component {
	a.Set("loadMoreProps", value)
	return a
}

func (a *Component) LoadType(value interface{}) *Component {
	a.Set("loadType", value)
	return a
}

func (a *Component) Loaded(value interface{}) *Component {
	a.Set("loaded", value)
	return a
}

func (a *Component) Loading(value interface{}) *Component {
	a.Set("loading", value)
	return a
}

func (a *Component) LoadingClassName(value interface{}) *Component {
	a.Set("loadingClassName", value)
	return a
}

func (a *Component) LoadingConfig(value interface{}) *Component {
	a.Set("loadingConfig", value)
	return a
}

func (a *Component) LoadingOn(value interface{}) *Component {
	a.Set("loadingOn", value)
	return a
}

func (a *Component) LoadingStatusCode(value interface{}) *Component {
	a.Set("loadingStatusCode", value)
	return a
}

func (a *Component) Loop(value interface{}) *Component {
	a.Set("loop", value)
	return a
}

func (a *Component) Map(value interface{}) *Component {
	a.Set("map", value)
	return a
}

func (a *Component) MapName(value interface{}) *Component {
	a.Set("mapName", value)
	return a
}

func (a *Component) MapURL(value interface{}) *Component {
	a.Set("mapURL", value)
	return a
}

func (a *Component) Marks(value interface{}) *Component {
	a.Set("marks", value)
	return a
}

func (a *Component) MaskColor(value interface{}) *Component {
	a.Set("maskColor", value)
	return a
}

func (a *Component) MasonryLayout(value interface{}) *Component {
	a.Set("masonryLayout", value)
	return a
}

func (a *Component) MatchFunc(value interface{}) *Component {
	a.Set("matchFunc", value)
	return a
}

func (a *Component) Max(value interface{}) *Component {
	a.Set("max", value)
	return a
}

func (a *Component) MaxButtons(value interface{}) *Component {
	a.Set("maxButtons", value)
	return a
}

func (a *Component) MaxDate(value interface{}) *Component {
	a.Set("maxDate", value)
	return a
}

func (a *Component) MaxDuration(value interface{}) *Component {
	a.Set("maxDuration", value)
	return a
}

func (a *Component) MaxHeight(value interface{}) *Component {
	a.Set("maxHeight", value)
	return a
}

func (a *Component) MaxKeepItemSelectionLength(value interface{}) *Component {
	a.Set("maxKeepItemSelectionLength", value)
	return a
}

func (a *Component) MaxLength(value interface{}) *Component {
	a.Set("maxLength", value)
	return a
}

func (a *Component) MaxRows(value interface{}) *Component {
	a.Set("maxRows", value)
	return a
}

func (a *Component) MaxSize(value interface{}) *Component {
	a.Set("maxSize", value)
	return a
}

func (a *Component) MaxTagCount(value interface{}) *Component {
	a.Set("maxTagCount", value)
	return a
}

func (a *Component) MaxTagLength(value interface{}) *Component {
	a.Set("maxTagLength", value)
	return a
}

func (a *Component) MaxVisibleCount(value interface{}) *Component {
	a.Set("maxVisibleCount", value)
	return a
}

func (a *Component) Md(value interface{}) *Component {
	a.Set("md", value)
	return a
}

func (a *Component) Media(value interface{}) *Component {
	a.Set("media", value)
	return a
}

func (a *Component) MenuClassName(value interface{}) *Component {
	a.Set("menuClassName", value)
	return a
}

func (a *Component) MenuTpl(value interface{}) *Component {
	a.Set("menuTpl", value)
	return a
}

func (a *Component) MergeData(value interface{}) *Component {
	a.Set("mergeData", value)
	return a
}

func (a *Component) Messages(value interface{}) *Component {
	a.Set("messages", value)
	return a
}

func (a *Component) Method(value interface{}) *Component {
	a.Set("method", value)
	return a
}

func (a *Component) Min(value interface{}) *Component {
	a.Set("min", value)
	return a
}

func (a *Component) MinDate(value interface{}) *Component {
	a.Set("minDate", value)
	return a
}

func (a *Component) MinDuration(value interface{}) *Component {
	a.Set("minDuration", value)
	return a
}

func (a *Component) MinLength(value interface{}) *Component {
	a.Set("minLength", value)
	return a
}

func (a *Component) MinRows(value interface{}) *Component {
	a.Set("minRows", value)
	return a
}

func (a *Component) Mini(value interface{}) *Component {
	a.Set("mini", value)
	return a
}

func (a *Component) MobileCSS(value interface{}) *Component {
	a.Set("mobileCSS", value)
	return a
}

func (a *Component) ModalMode(value interface{}) *Component {
	a.Set("modalMode", value)
	return a
}

func (a *Component) ModalSize(value interface{}) *Component {
	a.Set("modalSize", value)
	return a
}

func (a *Component) ModalTitle(value interface{}) *Component {
	a.Set("modalTitle", value)
	return a
}

func (a *Component) Mode(value interface{}) *Component {
	a.Set("mode", value)
	return a
}

func (a *Component) MosaicText(value interface{}) *Component {
	a.Set("mosaicText", value)
	return a
}

func (a *Component) MountOnEnter(value interface{}) *Component {
	a.Set("mountOnEnter", value)
	return a
}

func (a *Component) MouseEnterDelay(value interface{}) *Component {
	a.Set("mouseEnterDelay", value)
	return a
}

func (a *Component) MouseEvent(value interface{}) *Component {
	a.Set("mouseEvent", value)
	return a
}

func (a *Component) MouseLeaveDelay(value interface{}) *Component {
	a.Set("mouseLeaveDelay", value)
	return a
}

func (a *Component) MultiLine(value interface{}) *Component {
	a.Set("multiLine", value)
	return a
}

func (a *Component) Multiple(value interface{}) *Component {
	a.Set("multiple", value)
	return a
}

func (a *Component) Mutable(value interface{}) *Component {
	a.Set("mutable", value)
	return a
}

func (a *Component) Muted(value interface{}) *Component {
	a.Set("muted", value)
	return a
}

func (a *Component) Name(value interface{}) *Component {
	a.Set("name", value)
	return a
}

func (a *Component) NameField(value interface{}) *Component {
	a.Set("nameField", value)
	return a
}

func (a *Component) NativeAutoComplete(value interface{}) *Component {
	a.Set("nativeAutoComplete", value)
	return a
}

func (a *Component) NativeInputClassName(value interface{}) *Component {
	a.Set("nativeInputClassName", value)
	return a
}

func (a *Component) NeedConfirm(value interface{}) *Component {
	a.Set("needConfirm", value)
	return a
}

func (a *Component) NextCondition(value interface{}) *Component {
	a.Set("nextCondition", value)
	return a
}

func (a *Component) NoBorder(value interface{}) *Component {
	a.Set("noBorder", value)
	return a
}

func (a *Component) NodeBehavior(value interface{}) *Component {
	a.Set("nodeBehavior", value)
	return a
}

func (a *Component) Not(value interface{}) *Component {
	a.Set("not", value)
	return a
}

func (a *Component) Nullable(value interface{}) *Component {
	a.Set("nullable", value)
	return a
}

func (a *Component) OffText(value interface{}) *Component {
	a.Set("offText", value)
	return a
}

func (a *Component) Offset(value interface{}) *Component {
	a.Set("offset", value)
	return a
}

func (a *Component) OnClick(value interface{}) *Component {
	a.Set("onClick", value)
	return a
}

func (a *Component) OnError(value interface{}) *Component {
	a.Set("onError", value)
	return a
}

func (a *Component) OnEvent(value interface{}) *Component {
	a.Set("onEvent", value)
	return a
}

func (a *Component) OnMount(value interface{}) *Component {
	a.Set("onMount", value)
	return a
}

func (a *Component) OnText(value interface{}) *Component {
	a.Set("onText", value)
	return a
}

func (a *Component) OnUnmount(value interface{}) *Component {
	a.Set("onUnmount", value)
	return a
}

func (a *Component) OnUpdate(value interface{}) *Component {
	a.Set("onUpdate", value)
	return a
}

func (a *Component) OnlyChildren(value interface{}) *Component {
	a.Set("onlyChildren", value)
	return a
}

func (a *Component) OnlyLeaf(value interface{}) *Component {
	a.Set("onlyLeaf", value)
	return a
}

func (a *Component) OnlySelectCurrentLoc(value interface{}) *Component {
	a.Set("onlySelectCurrentLoc", value)
	return a
}

func (a *Component) Operation(value interface{}) *Component {
	a.Set("operation", value)
	return a
}

func (a *Component) OperationLabel(value interface{}) *Component {
	a.Set("operationLabel", value)
	return a
}

func (a *Component) Option(value interface{}) *Component {
	a.Set("option", value)
	return a
}

func (a *Component) OptionClassName(value interface{}) *Component {
	a.Set("optionClassName", value)
	return a
}

func (a *Component) OptionType(value interface{}) *Component {
	a.Set("optionType", value)
	return a
}

func (a *Component) Options(value interface{}) *Component {
	a.Set("options", value)
	return a
}

func (a *Component) OptionsTip(value interface{}) *Component {
	a.Set("optionsTip", value)
	return a
}

func (a *Component) OrderBy(value interface{}) *Component {
	a.Set("orderBy", value)
	return a
}

func (a *Component) OrderDir(value interface{}) *Component {
	a.Set("orderDir", value)
	return a
}

func (a *Component) OrderField(value interface{}) *Component {
	a.Set("orderField", value)
	return a
}

func (a *Component) OriginalSrc(value interface{}) *Component {
	a.Set("originalSrc", value)
	return a
}

func (a *Component) OutputName(value interface{}) *Component {
	a.Set("outputName", value)
	return a
}

func (a *Component) OutputVar(value interface{}) *Component {
	a.Set("outputVar", value)
	return a
}

func (a *Component) Overflow(value interface{}) *Component {
	a.Set("overflow", value)
	return a
}

func (a *Component) OverflowClassName(value interface{}) *Component {
	a.Set("overflowClassName", value)
	return a
}

func (a *Component) OverflowConfig(value interface{}) *Component {
	a.Set("overflowConfig", value)
	return a
}

func (a *Component) OverflowCount(value interface{}) *Component {
	a.Set("overflowCount", value)
	return a
}

func (a *Component) OverflowIndicator(value interface{}) *Component {
	a.Set("overflowIndicator", value)
	return a
}

func (a *Component) OverflowLabel(value interface{}) *Component {
	a.Set("overflowLabel", value)
	return a
}

func (a *Component) OverflowListClassName(value interface{}) *Component {
	a.Set("overflowListClassName", value)
	return a
}

func (a *Component) OverflowPopoverClassName(value interface{}) *Component {
	a.Set("overflowPopoverClassName", value)
	return a
}

func (a *Component) OverflowSuffix(value interface{}) *Component {
	a.Set("overflowSuffix", value)
	return a
}

func (a *Component) OverflowTagPopover(value interface{}) *Component {
	a.Set("overflowTagPopover", value)
	return a
}

func (a *Component) Overlay(value interface{}) *Component {
	a.Set("overlay", value)
	return a
}

func (a *Component) OverlayPlacement(value interface{}) *Component {
	a.Set("overlayPlacement", value)
	return a
}

func (a *Component) PageDirectionField(value interface{}) *Component {
	a.Set("pageDirectionField", value)
	return a
}

func (a *Component) PageField(value interface{}) *Component {
	a.Set("pageField", value)
	return a
}

func (a *Component) Pagination(value interface{}) *Component {
	a.Set("pagination", value)
	return a
}

func (a *Component) PanelClassName(value interface{}) *Component {
	a.Set("panelClassName", value)
	return a
}

func (a *Component) ParseMode(value interface{}) *Component {
	a.Set("parseMode", value)
	return a
}

func (a *Component) ParsePrimitiveQuery(value interface{}) *Component {
	a.Set("parsePrimitiveQuery", value)
	return a
}

func (a *Component) Partial(value interface{}) *Component {
	a.Set("partial", value)
	return a
}

func (a *Component) Parts(value interface{}) *Component {
	a.Set("parts", value)
	return a
}

func (a *Component) PathSeparator(value interface{}) *Component {
	a.Set("pathSeparator", value)
	return a
}

func (a *Component) PerPage(value interface{}) *Component {
	a.Set("perPage", value)
	return a
}

func (a *Component) PerPageAvailable(value interface{}) *Component {
	a.Set("perPageAvailable", value)
	return a
}

func (a *Component) PerPageField(value interface{}) *Component {
	a.Set("perPageField", value)
	return a
}

func (a *Component) PersistData(value interface{}) *Component {
	a.Set("persistData", value)
	return a
}

func (a *Component) PersistDataKeys(value interface{}) *Component {
	a.Set("persistDataKeys", value)
	return a
}

func (a *Component) PickerIcon(value interface{}) *Component {
	a.Set("pickerIcon", value)
	return a
}

func (a *Component) PickerSchema(value interface{}) *Component {
	a.Set("pickerSchema", value)
	return a
}

func (a *Component) PickerSize(value interface{}) *Component {
	a.Set("pickerSize", value)
	return a
}

func (a *Component) Placeholder(value interface{}) *Component {
	a.Set("placeholder", value)
	return a
}

func (a *Component) Placement(value interface{}) *Component {
	a.Set("placement", value)
	return a
}

func (a *Component) PlainText(value interface{}) *Component {
	a.Set("plainText", value)
	return a
}

func (a *Component) PlayerClassName(value interface{}) *Component {
	a.Set("playerClassName", value)
	return a
}

func (a *Component) PointSize(value interface{}) *Component {
	a.Set("pointSize", value)
	return a
}

func (a *Component) PointSizeRandom(value interface{}) *Component {
	a.Set("pointSizeRandom", value)
	return a
}

func (a *Component) PointType(value interface{}) *Component {
	a.Set("pointType", value)
	return a
}

func (a *Component) PopOver(value interface{}) *Component {
	a.Set("popOver", value)
	return a
}

func (a *Component) PopOverClassName(value interface{}) *Component {
	a.Set("popOverClassName", value)
	return a
}

func (a *Component) PopOverContainer(value interface{}) *Component {
	a.Set("popOverContainer", value)
	return a
}

func (a *Component) PopOverContainerSelector(value interface{}) *Component {
	a.Set("popOverContainerSelector", value)
	return a
}

func (a *Component) PopOverEnableOn(value interface{}) *Component {
	a.Set("popOverEnableOn", value)
	return a
}

func (a *Component) PopupClassName(value interface{}) *Component {
	a.Set("popupClassName", value)
	return a
}

func (a *Component) Position(value interface{}) *Component {
	a.Set("position", value)
	return a
}

func (a *Component) Poster(value interface{}) *Component {
	a.Set("poster", value)
	return a
}

func (a *Component) Precision(value interface{}) *Component {
	a.Set("precision", value)
	return a
}

func (a *Component) Prefix(value interface{}) *Component {
	a.Set("prefix", value)
	return a
}

func (a *Component) PrefixRow(value interface{}) *Component {
	a.Set("prefixRow", value)
	return a
}

func (a *Component) PresetColors(value interface{}) *Component {
	a.Set("presetColors", value)
	return a
}

func (a *Component) PreventDefault(value interface{}) *Component {
	a.Set("preventDefault", value)
	return a
}

func (a *Component) PreventEnterSubmit(value interface{}) *Component {
	a.Set("preventEnterSubmit", value)
	return a
}

func (a *Component) Primary(value interface{}) *Component {
	a.Set("primary", value)
	return a
}

func (a *Component) PrimaryField(value interface{}) *Component {
	a.Set("primaryField", value)
	return a
}

func (a *Component) ProgressClassName(value interface{}) *Component {
	a.Set("progressClassName", value)
	return a
}

func (a *Component) ProgressDot(value interface{}) *Component {
	a.Set("progressDot", value)
	return a
}

func (a *Component) PromptPageLeave(value interface{}) *Component {
	a.Set("promptPageLeave", value)
	return a
}

func (a *Component) PromptPageLeaveMessage(value interface{}) *Component {
	a.Set("promptPageLeaveMessage", value)
	return a
}

func (a *Component) Props(value interface{}) *Component {
	a.Set("props", value)
	return a
}

func (a *Component) PullRefresh(value interface{}) *Component {
	a.Set("pullRefresh", value)
	return a
}

func (a *Component) QrcodeClassName(value interface{}) *Component {
	a.Set("qrcodeClassName", value)
	return a
}

func (a *Component) QsOptions(value interface{}) *Component {
	a.Set("qsOptions", value)
	return a
}

func (a *Component) QuickEdit(value interface{}) *Component {
	a.Set("quickEdit", value)
	return a
}

func (a *Component) QuickEditOnUpdate(value interface{}) *Component {
	a.Set("quickEditOnUpdate", value)
	return a
}

func (a *Component) QuickSaveApi(value interface{}) *Component {
	a.Set("quickSaveApi", value)
	return a
}

func (a *Component) QuickSaveItemApi(value interface{}) *Component {
	a.Set("quickSaveItemApi", value)
	return a
}

func (a *Component) QuotesOnKeys(value interface{}) *Component {
	a.Set("quotesOnKeys", value)
	return a
}

func (a *Component) Ranges(value interface{}) *Component {
	a.Set("ranges", value)
	return a
}

func (a *Component) Rates(value interface{}) *Component {
	a.Set("rates", value)
	return a
}

func (a *Component) Raw(value interface{}) *Component {
	a.Set("raw", value)
	return a
}

func (a *Component) ReCropable(value interface{}) *Component {
	a.Set("reCropable", value)
	return a
}

func (a *Component) ReSubmitApi(value interface{}) *Component {
	a.Set("reSubmitApi", value)
	return a
}

func (a *Component) ReadOnly(value interface{}) *Component {
	a.Set("readOnly", value)
	return a
}

func (a *Component) ReadOnlyOn(value interface{}) *Component {
	a.Set("readOnlyOn", value)
	return a
}

func (a *Component) ReadyStatusCode(value interface{}) *Component {
	a.Set("readyStatusCode", value)
	return a
}

func (a *Component) Receiver(value interface{}) *Component {
	a.Set("receiver", value)
	return a
}

func (a *Component) Redirect(value interface{}) *Component {
	a.Set("redirect", value)
	return a
}

func (a *Component) Referrerpolicy(value interface{}) *Component {
	a.Set("referrerpolicy", value)
	return a
}

func (a *Component) Regions(value interface{}) *Component {
	a.Set("regions", value)
	return a
}

func (a *Component) Reload(value interface{}) *Component {
	a.Set("reload", value)
	return a
}

func (a *Component) Remark(value interface{}) *Component {
	a.Set("remark", value)
	return a
}

func (a *Component) RemarkLabel(value interface{}) *Component {
	a.Set("remarkLabel", value)
	return a
}

func (a *Component) Removable(value interface{}) *Component {
	a.Set("removable", value)
	return a
}

func (a *Component) ReplaceChartOption(value interface{}) *Component {
	a.Set("replaceChartOption", value)
	return a
}

func (a *Component) ReplaceData(value interface{}) *Component {
	a.Set("replaceData", value)
	return a
}

func (a *Component) RequireSelected(value interface{}) *Component {
	a.Set("requireSelected", value)
	return a
}

func (a *Component) Required(value interface{}) *Component {
	a.Set("required", value)
	return a
}

func (a *Component) RequiredOn(value interface{}) *Component {
	a.Set("requiredOn", value)
	return a
}

func (a *Component) ResetAfterSubmit(value interface{}) *Component {
	a.Set("resetAfterSubmit", value)
	return a
}

func (a *Component) ResetValue(value interface{}) *Component {
	a.Set("resetValue", value)
	return a
}

func (a *Component) Resizable(value interface{}) *Component {
	a.Set("resizable", value)
	return a
}

func (a *Component) ResponseData(value interface{}) *Component {
	a.Set("responseData", value)
	return a
}

func (a *Component) ResponseType(value interface{}) *Component {
	a.Set("responseType", value)
	return a
}

func (a *Component) ResultListModeFollowSelect(value interface{}) *Component {
	a.Set("resultListModeFollowSelect", value)
	return a
}

func (a *Component) ResultSearchPlaceholder(value interface{}) *Component {
	a.Set("resultSearchPlaceholder", value)
	return a
}

func (a *Component) ResultSearchable(value interface{}) *Component {
	a.Set("resultSearchable", value)
	return a
}

func (a *Component) ResultTitle(value interface{}) *Component {
	a.Set("resultTitle", value)
	return a
}

func (a *Component) RetryBtnClassName(value interface{}) *Component {
	a.Set("retryBtnClassName", value)
	return a
}

func (a *Component) RetryBtnText(value interface{}) *Component {
	a.Set("retryBtnText", value)
	return a
}

func (a *Component) Reverse(value interface{}) *Component {
	a.Set("reverse", value)
	return a
}

func (a *Component) Right(value interface{}) *Component {
	a.Set("right", value)
	return a
}

func (a *Component) RightIcon(value interface{}) *Component {
	a.Set("rightIcon", value)
	return a
}

func (a *Component) RightIconClassName(value interface{}) *Component {
	a.Set("rightIconClassName", value)
	return a
}

func (a *Component) RightMode(value interface{}) *Component {
	a.Set("rightMode", value)
	return a
}

func (a *Component) RootClose(value interface{}) *Component {
	a.Set("rootClose", value)
	return a
}

func (a *Component) RootCreatable(value interface{}) *Component {
	a.Set("rootCreatable", value)
	return a
}

func (a *Component) RootLabel(value interface{}) *Component {
	a.Set("rootLabel", value)
	return a
}

func (a *Component) RootTypeMutable(value interface{}) *Component {
	a.Set("rootTypeMutable", value)
	return a
}

func (a *Component) RootValue(value interface{}) *Component {
	a.Set("rootValue", value)
	return a
}

func (a *Component) Rotate(value interface{}) *Component {
	a.Set("rotate", value)
	return a
}

func (a *Component) Row(value interface{}) *Component {
	a.Set("row", value)
	return a
}

func (a *Component) RowClassNameExpr(value interface{}) *Component {
	a.Set("rowClassNameExpr", value)
	return a
}

func (a *Component) RowClick(value interface{}) *Component {
	a.Set("rowClick", value)
	return a
}

func (a *Component) RowHeight(value interface{}) *Component {
	a.Set("rowHeight", value)
	return a
}

func (a *Component) RowLabel(value interface{}) *Component {
	a.Set("rowLabel", value)
	return a
}

func (a *Component) RowSelection(value interface{}) *Component {
	a.Set("rowSelection", value)
	return a
}

func (a *Component) RowSpanExpr(value interface{}) *Component {
	a.Set("rowSpanExpr", value)
	return a
}

func (a *Component) Rows(value interface{}) *Component {
	a.Set("rows", value)
	return a
}

func (a *Component) Rules(value interface{}) *Component {
	a.Set("rules", value)
	return a
}

func (a *Component) Sandbox(value interface{}) *Component {
	a.Set("sandbox", value)
	return a
}

func (a *Component) SaveFailed(value interface{}) *Component {
	a.Set("saveFailed", value)
	return a
}

func (a *Component) SaveImmediately(value interface{}) *Component {
	a.Set("saveImmediately", value)
	return a
}

func (a *Component) SaveOrderApi(value interface{}) *Component {
	a.Set("saveOrderApi", value)
	return a
}

func (a *Component) SaveSuccess(value interface{}) *Component {
	a.Set("saveSuccess", value)
	return a
}

func (a *Component) Scaffold(value interface{}) *Component {
	a.Set("scaffold", value)
	return a
}

func (a *Component) ScheduleAction(value interface{}) *Component {
	a.Set("scheduleAction", value)
	return a
}

func (a *Component) ScheduleClassNames(value interface{}) *Component {
	a.Set("scheduleClassNames", value)
	return a
}

func (a *Component) Schedules(value interface{}) *Component {
	a.Set("schedules", value)
	return a
}

func (a *Component) SchemaApi(value interface{}) *Component {
	a.Set("schemaApi", value)
	return a
}

func (a *Component) ScopeLabel(value interface{}) *Component {
	a.Set("scopeLabel", value)
	return a
}

func (a *Component) Scrollable(value interface{}) *Component {
	a.Set("scrollable", value)
	return a
}

func (a *Component) SearchApi(value interface{}) *Component {
	a.Set("searchApi", value)
	return a
}

func (a *Component) SearchConfig(value interface{}) *Component {
	a.Set("searchConfig", value)
	return a
}

func (a *Component) SearchImediately(value interface{}) *Component {
	a.Set("searchImediately", value)
	return a
}

func (a *Component) SearchPlaceholder(value interface{}) *Component {
	a.Set("searchPlaceholder", value)
	return a
}

func (a *Component) SearchResultColumns(value interface{}) *Component {
	a.Set("searchResultColumns", value)
	return a
}

func (a *Component) SearchResultMode(value interface{}) *Component {
	a.Set("searchResultMode", value)
	return a
}

func (a *Component) Searchable(value interface{}) *Component {
	a.Set("searchable", value)
	return a
}

func (a *Component) Secondary(value interface{}) *Component {
	a.Set("secondary", value)
	return a
}

func (a *Component) SectionClassName(value interface{}) *Component {
	a.Set("sectionClassName", value)
	return a
}

func (a *Component) SelectFirst(value interface{}) *Component {
	a.Set("selectFirst", value)
	return a
}

func (a *Component) SelectMode(value interface{}) *Component {
	a.Set("selectMode", value)
	return a
}

func (a *Component) SelectTitle(value interface{}) *Component {
	a.Set("selectTitle", value)
	return a
}

func (a *Component) Selectable(value interface{}) *Component {
	a.Set("selectable", value)
	return a
}

func (a *Component) SelectedRowKeys(value interface{}) *Component {
	a.Set("selectedRowKeys", value)
	return a
}

func (a *Component) SelectedRowKeysExpr(value interface{}) *Component {
	a.Set("selectedRowKeysExpr", value)
	return a
}

func (a *Component) Selections(value interface{}) *Component {
	a.Set("selections", value)
	return a
}

func (a *Component) SendOn(value interface{}) *Component {
	a.Set("sendOn", value)
	return a
}

func (a *Component) Separator(value interface{}) *Component {
	a.Set("separator", value)
	return a
}

func (a *Component) SeparatorClassName(value interface{}) *Component {
	a.Set("separatorClassName", value)
	return a
}

func (a *Component) Shape(value interface{}) *Component {
	a.Set("shape", value)
	return a
}

func (a *Component) Shortcuts(value interface{}) *Component {
	a.Set("shortcuts", value)
	return a
}

func (a *Component) Show(value interface{}) *Component {
	a.Set("show", value)
	return a
}

func (a *Component) ShowANDOR(value interface{}) *Component {
	a.Set("showANDOR", value)
	return a
}

func (a *Component) ShowArrow(value interface{}) *Component {
	a.Set("showArrow", value)
	return a
}

func (a *Component) ShowAsPercent(value interface{}) *Component {
	a.Set("showAsPercent", value)
	return a
}

func (a *Component) ShowBadge(value interface{}) *Component {
	a.Set("showBadge", value)
	return a
}

func (a *Component) ShowBtnToolbar(value interface{}) *Component {
	a.Set("showBtnToolbar", value)
	return a
}

func (a *Component) ShowCloseButton(value interface{}) *Component {
	a.Set("showCloseButton", value)
	return a
}

func (a *Component) ShowCompressOptions(value interface{}) *Component {
	a.Set("showCompressOptions", value)
	return a
}

func (a *Component) ShowCounter(value interface{}) *Component {
	a.Set("showCounter", value)
	return a
}

func (a *Component) ShowDimensions(value interface{}) *Component {
	a.Set("showDimensions", value)
	return a
}

func (a *Component) ShowErrorModal(value interface{}) *Component {
	a.Set("showErrorModal", value)
	return a
}

func (a *Component) ShowErrorMsg(value interface{}) *Component {
	a.Set("showErrorMsg", value)
	return a
}

func (a *Component) ShowFooter(value interface{}) *Component {
	a.Set("showFooter", value)
	return a
}

func (a *Component) ShowFooterAddBtn(value interface{}) *Component {
	a.Set("showFooterAddBtn", value)
	return a
}

func (a *Component) ShowHeader(value interface{}) *Component {
	a.Set("showHeader", value)
	return a
}

func (a *Component) ShowIcon(value interface{}) *Component {
	a.Set("showIcon", value)
	return a
}

func (a *Component) ShowIndex(value interface{}) *Component {
	a.Set("showIndex", value)
	return a
}

func (a *Component) ShowIndexBar(value interface{}) *Component {
	a.Set("showIndexBar", value)
	return a
}

func (a *Component) ShowInput(value interface{}) *Component {
	a.Set("showInput", value)
	return a
}

func (a *Component) ShowInvalidMatch(value interface{}) *Component {
	a.Set("showInvalidMatch", value)
	return a
}

func (a *Component) ShowKey(value interface{}) *Component {
	a.Set("showKey", value)
	return a
}

func (a *Component) ShowLabel(value interface{}) *Component {
	a.Set("showLabel", value)
	return a
}

func (a *Component) ShowLoading(value interface{}) *Component {
	a.Set("showLoading", value)
	return a
}

func (a *Component) ShowOutline(value interface{}) *Component {
	a.Set("showOutline", value)
	return a
}

func (a *Component) ShowPageInput(value interface{}) *Component {
	a.Set("showPageInput", value)
	return a
}

func (a *Component) ShowPerPage(value interface{}) *Component {
	a.Set("showPerPage", value)
	return a
}

func (a *Component) ShowRootInfo(value interface{}) *Component {
	a.Set("showRootInfo", value)
	return a
}

func (a *Component) ShowSelection(value interface{}) *Component {
	a.Set("showSelection", value)
	return a
}

func (a *Component) ShowSteps(value interface{}) *Component {
	a.Set("showSteps", value)
	return a
}

func (a *Component) ShowTableAddBtn(value interface{}) *Component {
	a.Set("showTableAddBtn", value)
	return a
}

func (a *Component) ShowThresholdText(value interface{}) *Component {
	a.Set("showThresholdText", value)
	return a
}

func (a *Component) ShowTip(value interface{}) *Component {
	a.Set("showTip", value)
	return a
}

func (a *Component) ShowTipClassName(value interface{}) *Component {
	a.Set("showTipClassName", value)
	return a
}

func (a *Component) ShowToolbar(value interface{}) *Component {
	a.Set("showToolbar", value)
	return a
}

func (a *Component) ShowTooltipOnHighlight(value interface{}) *Component {
	a.Set("showTooltipOnHighlight", value)
	return a
}

func (a *Component) ShowValue(value interface{}) *Component {
	a.Set("showValue", value)
	return a
}

func (a *Component) SidePosition(value interface{}) *Component {
	a.Set("sidePosition", value)
	return a
}

func (a *Component) Silent(value interface{}) *Component {
	a.Set("silent", value)
	return a
}

func (a *Component) SilentPolling(value interface{}) *Component {
	a.Set("silentPolling", value)
	return a
}

func (a *Component) SingleSelectMode(value interface{}) *Component {
	a.Set("singleSelectMode", value)
	return a
}

func (a *Component) Size(value interface{}) *Component {
	a.Set("size", value)
	return a
}

func (a *Component) SkipRestOnCancel(value interface{}) *Component {
	a.Set("skipRestOnCancel", value)
	return a
}

func (a *Component) SkipRestOnConfirm(value interface{}) *Component {
	a.Set("skipRestOnConfirm", value)
	return a
}

func (a *Component) Sm(value interface{}) *Component {
	a.Set("sm", value)
	return a
}

func (a *Component) SortKeys(value interface{}) *Component {
	a.Set("sortKeys", value)
	return a
}

func (a *Component) SortType(value interface{}) *Component {
	a.Set("sortType", value)
	return a
}

func (a *Component) Sortable(value interface{}) *Component {
	a.Set("sortable", value)
	return a
}

func (a *Component) Sorter(value interface{}) *Component {
	a.Set("sorter", value)
	return a
}

func (a *Component) Source(value interface{}) *Component {
	a.Set("source", value)
	return a
}

func (a *Component) SpinnerClassName(value interface{}) *Component {
	a.Set("spinnerClassName", value)
	return a
}

func (a *Component) SpinnerWrapClassName(value interface{}) *Component {
	a.Set("spinnerWrapClassName", value)
	return a
}

func (a *Component) SplitPoster(value interface{}) *Component {
	a.Set("splitPoster", value)
	return a
}

func (a *Component) Square(value interface{}) *Component {
	a.Set("square", value)
	return a
}

func (a *Component) Src(value interface{}) *Component {
	a.Set("src", value)
	return a
}

func (a *Component) Stacked(value interface{}) *Component {
	a.Set("stacked", value)
	return a
}

func (a *Component) StartChunkApi(value interface{}) *Component {
	a.Set("startChunkApi", value)
	return a
}

func (a *Component) StartPlaceholder(value interface{}) *Component {
	a.Set("startPlaceholder", value)
	return a
}

func (a *Component) StartStep(value interface{}) *Component {
	a.Set("startStep", value)
	return a
}

func (a *Component) StateTextMap(value interface{}) *Component {
	a.Set("stateTextMap", value)
	return a
}

func (a *Component) Static(value interface{}) *Component {
	a.Set("static", value)
	return a
}

func (a *Component) StaticClassName(value interface{}) *Component {
	a.Set("staticClassName", value)
	return a
}

func (a *Component) StaticInputClassName(value interface{}) *Component {
	a.Set("staticInputClassName", value)
	return a
}

func (a *Component) StaticLabelClassName(value interface{}) *Component {
	a.Set("staticLabelClassName", value)
	return a
}

func (a *Component) StaticOn(value interface{}) *Component {
	a.Set("staticOn", value)
	return a
}

func (a *Component) StaticPlaceholder(value interface{}) *Component {
	a.Set("staticPlaceholder", value)
	return a
}

func (a *Component) StaticSchema(value interface{}) *Component {
	a.Set("staticSchema", value)
	return a
}

func (a *Component) Statistics(value interface{}) *Component {
	a.Set("statistics", value)
	return a
}

func (a *Component) Status(value interface{}) *Component {
	a.Set("status", value)
	return a
}

func (a *Component) StatusLabel(value interface{}) *Component {
	a.Set("statusLabel", value)
	return a
}

func (a *Component) StatusLabelMap(value interface{}) *Component {
	a.Set("statusLabelMap", value)
	return a
}

func (a *Component) StatusTextMap(value interface{}) *Component {
	a.Set("statusTextMap", value)
	return a
}

func (a *Component) Step(value interface{}) *Component {
	a.Set("step", value)
	return a
}

func (a *Component) StepClassName(value interface{}) *Component {
	a.Set("stepClassName", value)
	return a
}

func (a *Component) Steps(value interface{}) *Component {
	a.Set("steps", value)
	return a
}

func (a *Component) StepsClassName(value interface{}) *Component {
	a.Set("stepsClassName", value)
	return a
}

func (a *Component) Sticky(value interface{}) *Component {
	a.Set("sticky", value)
	return a
}

func (a *Component) StopAutoRefreshWhen(value interface{}) *Component {
	a.Set("stopAutoRefreshWhen", value)
	return a
}

func (a *Component) StopAutoRefreshWhenModalIsOpen(value interface{}) *Component {
	a.Set("stopAutoRefreshWhenModalIsOpen", value)
	return a
}

func (a *Component) StopOnNextFrame(value interface{}) *Component {
	a.Set("stopOnNextFrame", value)
	return a
}

func (a *Component) StopPropagation(value interface{}) *Component {
	a.Set("stopPropagation", value)
	return a
}

func (a *Component) StrictMode(value interface{}) *Component {
	a.Set("strictMode", value)
	return a
}

func (a *Component) Stripe(value interface{}) *Component {
	a.Set("stripe", value)
	return a
}

func (a *Component) StrokeWidth(value interface{}) *Component {
	a.Set("strokeWidth", value)
	return a
}

func (a *Component) Style(value interface{}) *Component {
	a.Set("style", value)
	return a
}

func (a *Component) SubAddBtnIcon(value interface{}) *Component {
	a.Set("subAddBtnIcon", value)
	return a
}

func (a *Component) SubAddBtnLabel(value interface{}) *Component {
	a.Set("subAddBtnLabel", value)
	return a
}

func (a *Component) SubFormHorizontal(value interface{}) *Component {
	a.Set("subFormHorizontal", value)
	return a
}

func (a *Component) SubFormMode(value interface{}) *Component {
	a.Set("subFormMode", value)
	return a
}

func (a *Component) SubTitle(value interface{}) *Component {
	a.Set("subTitle", value)
	return a
}

func (a *Component) Subject(value interface{}) *Component {
	a.Set("subject", value)
	return a
}

func (a *Component) SubmitApi(value interface{}) *Component {
	a.Set("submitApi", value)
	return a
}

func (a *Component) SubmitOnChange(value interface{}) *Component {
	a.Set("submitOnChange", value)
	return a
}

func (a *Component) SubmitOnDBClick(value interface{}) *Component {
	a.Set("submitOnDBClick", value)
	return a
}

func (a *Component) SubmitOnInit(value interface{}) *Component {
	a.Set("submitOnInit", value)
	return a
}

func (a *Component) SubmitText(value interface{}) *Component {
	a.Set("submitText", value)
	return a
}

func (a *Component) Suffix(value interface{}) *Component {
	a.Set("suffix", value)
	return a
}

func (a *Component) Svg(value interface{}) *Component {
	a.Set("svg", value)
	return a
}

func (a *Component) Swipeable(value interface{}) *Component {
	a.Set("swipeable", value)
	return a
}

func (a *Component) SyncFields(value interface{}) *Component {
	a.Set("syncFields", value)
	return a
}

func (a *Component) SyncLocation(value interface{}) *Component {
	a.Set("syncLocation", value)
	return a
}

func (a *Component) SyncResponse2Query(value interface{}) *Component {
	a.Set("syncResponse2Query", value)
	return a
}

func (a *Component) Tab(value interface{}) *Component {
	a.Set("tab", value)
	return a
}

func (a *Component) TabIndex(value interface{}) *Component {
	a.Set("tabIndex", value)
	return a
}

func (a *Component) TabSize(value interface{}) *Component {
	a.Set("tabSize", value)
	return a
}

func (a *Component) TableClassName(value interface{}) *Component {
	a.Set("tableClassName", value)
	return a
}

func (a *Component) TableLayout(value interface{}) *Component {
	a.Set("tableLayout", value)
	return a
}

func (a *Component) Tabs(value interface{}) *Component {
	a.Set("tabs", value)
	return a
}

func (a *Component) TabsClassName(value interface{}) *Component {
	a.Set("tabsClassName", value)
	return a
}

func (a *Component) TabsLabelTpl(value interface{}) *Component {
	a.Set("tabsLabelTpl", value)
	return a
}

func (a *Component) TabsMode(value interface{}) *Component {
	a.Set("tabsMode", value)
	return a
}

func (a *Component) TabsStyle(value interface{}) *Component {
	a.Set("tabsStyle", value)
	return a
}

func (a *Component) Tag(value interface{}) *Component {
	a.Set("tag", value)
	return a
}

func (a *Component) Target(value interface{}) *Component {
	a.Set("target", value)
	return a
}

func (a *Component) TaskNameLabel(value interface{}) *Component {
	a.Set("taskNameLabel", value)
	return a
}

func (a *Component) TemplateUrl(value interface{}) *Component {
	a.Set("templateUrl", value)
	return a
}

func (a *Component) Test(value interface{}) *Component {
	a.Set("test", value)
	return a
}

func (a *Component) TestIdBuilder(value interface{}) *Component {
	a.Set("testIdBuilder", value)
	return a
}

func (a *Component) Testid(value interface{}) *Component {
	a.Set("testid", value)
	return a
}

func (a *Component) Text(value interface{}) *Component {
	a.Set("text", value)
	return a
}

func (a *Component) TextClassName(value interface{}) *Component {
	a.Set("textClassName", value)
	return a
}

func (a *Component) TextPosition(value interface{}) *Component {
	a.Set("textPosition", value)
	return a
}

func (a *Component) Texts(value interface{}) *Component {
	a.Set("texts", value)
	return a
}

func (a *Component) ThemeColor(value interface{}) *Component {
	a.Set("themeColor", value)
	return a
}

func (a *Component) ThemeCss(value interface{}) *Component {
	a.Set("themeCss", value)
	return a
}

func (a *Component) Threshold(value interface{}) *Component {
	a.Set("threshold", value)
	return a
}

func (a *Component) ThumbClassName(value interface{}) *Component {
	a.Set("thumbClassName", value)
	return a
}

func (a *Component) ThumbMode(value interface{}) *Component {
	a.Set("thumbMode", value)
	return a
}

func (a *Component) ThumbRatio(value interface{}) *Component {
	a.Set("thumbRatio", value)
	return a
}

func (a *Component) Tiled(value interface{}) *Component {
	a.Set("tiled", value)
	return a
}

func (a *Component) Time(value interface{}) *Component {
	a.Set("time", value)
	return a
}

func (a *Component) TimeClassName(value interface{}) *Component {
	a.Set("timeClassName", value)
	return a
}

func (a *Component) TimeConstraints(value interface{}) *Component {
	a.Set("timeConstraints", value)
	return a
}

func (a *Component) TimeFormat(value interface{}) *Component {
	a.Set("timeFormat", value)
	return a
}

func (a *Component) Timeout(value interface{}) *Component {
	a.Set("timeout", value)
	return a
}

func (a *Component) Tip(value interface{}) *Component {
	a.Set("tip", value)
	return a
}

func (a *Component) TipPlacement(value interface{}) *Component {
	a.Set("tipPlacement", value)
	return a
}

func (a *Component) Title(value interface{}) *Component {
	a.Set("title", value)
	return a
}

func (a *Component) TitleClassName(value interface{}) *Component {
	a.Set("titleClassName", value)
	return a
}

func (a *Component) TitlePosition(value interface{}) *Component {
	a.Set("titlePosition", value)
	return a
}

func (a *Component) To(value interface{}) *Component {
	a.Set("to", value)
	return a
}

func (a *Component) Toast(value interface{}) *Component {
	a.Set("toast", value)
	return a
}

func (a *Component) TodayActiveStyle(value interface{}) *Component {
	a.Set("todayActiveStyle", value)
	return a
}

func (a *Component) Toggled(value interface{}) *Component {
	a.Set("toggled", value)
	return a
}

func (a *Component) Toolbar(value interface{}) *Component {
	a.Set("toolbar", value)
	return a
}

func (a *Component) ToolbarActions(value interface{}) *Component {
	a.Set("toolbarActions", value)
	return a
}

func (a *Component) ToolbarClassName(value interface{}) *Component {
	a.Set("toolbarClassName", value)
	return a
}

func (a *Component) ToolbarInline(value interface{}) *Component {
	a.Set("toolbarInline", value)
	return a
}

func (a *Component) ToolbarKeys(value interface{}) *Component {
	a.Set("toolbarKeys", value)
	return a
}

func (a *Component) Tooltip(value interface{}) *Component {
	a.Set("tooltip", value)
	return a
}

func (a *Component) TooltipArrowClassName(value interface{}) *Component {
	a.Set("tooltipArrowClassName", value)
	return a
}

func (a *Component) TooltipClassName(value interface{}) *Component {
	a.Set("tooltipClassName", value)
	return a
}

func (a *Component) TooltipPlacement(value interface{}) *Component {
	a.Set("tooltipPlacement", value)
	return a
}

func (a *Component) TooltipPosition(value interface{}) *Component {
	a.Set("tooltipPosition", value)
	return a
}

func (a *Component) TooltipStyle(value interface{}) *Component {
	a.Set("tooltipStyle", value)
	return a
}

func (a *Component) TooltipTheme(value interface{}) *Component {
	a.Set("tooltipTheme", value)
	return a
}

func (a *Component) TooltipVisible(value interface{}) *Component {
	a.Set("tooltipVisible", value)
	return a
}

func (a *Component) Total(value interface{}) *Component {
	a.Set("total", value)
	return a
}

func (a *Component) TotalField(value interface{}) *Component {
	a.Set("totalField", value)
	return a
}

func (a *Component) Tpl(value interface{}) *Component {
	a.Set("tpl", value)
	return a
}

func (a *Component) TrackExpression(value interface{}) *Component {
	a.Set("trackExpression", value)
	return a
}

func (a *Component) Transform(value interface{}) *Component {
	a.Set("transform", value)
	return a
}

func (a *Component) Trigger(value interface{}) *Component {
	a.Set("trigger", value)
	return a
}

func (a *Component) TrimContents(value interface{}) *Component {
	a.Set("trimContents", value)
	return a
}

func (a *Component) TrueValue(value interface{}) *Component {
	a.Set("trueValue", value)
	return a
}

func (a *Component) Type(value interface{}) *Component {
	a.Set("type", value)
	return a
}

func (a *Component) TypeSwitchable(value interface{}) *Component {
	a.Set("typeSwitchable", value)
	return a
}

func (a *Component) UnMountOnHidden(value interface{}) *Component {
	a.Set("unMountOnHidden", value)
	return a
}

func (a *Component) UndoBtnIcon(value interface{}) *Component {
	a.Set("undoBtnIcon", value)
	return a
}

func (a *Component) UndoBtnLabel(value interface{}) *Component {
	a.Set("undoBtnLabel", value)
	return a
}

func (a *Component) Unfolded(value interface{}) *Component {
	a.Set("unfolded", value)
	return a
}

func (a *Component) Unique(value interface{}) *Component {
	a.Set("unique", value)
	return a
}

func (a *Component) Unit(value interface{}) *Component {
	a.Set("unit", value)
	return a
}

func (a *Component) UnitOptions(value interface{}) *Component {
	a.Set("unitOptions", value)
	return a
}

func (a *Component) UnmountOnExit(value interface{}) *Component {
	a.Set("unmountOnExit", value)
	return a
}

func (a *Component) UpdateApi(value interface{}) *Component {
	a.Set("updateApi", value)
	return a
}

func (a *Component) UpdateFrequency(value interface{}) *Component {
	a.Set("updateFrequency", value)
	return a
}

func (a *Component) UpdatePristineAfterStoreDataReInit(value interface{}) *Component {
	a.Set("updatePristineAfterStoreDataReInit", value)
	return a
}

func (a *Component) UploadApi(value interface{}) *Component {
	a.Set("uploadApi", value)
	return a
}

func (a *Component) UploadBtnText(value interface{}) *Component {
	a.Set("uploadBtnText", value)
	return a
}

func (a *Component) UploadImageMaxNumber(value interface{}) *Component {
	a.Set("uploadImageMaxNumber", value)
	return a
}

func (a *Component) UploadImageMaxSize(value interface{}) *Component {
	a.Set("uploadImageMaxSize", value)
	return a
}

func (a *Component) UploadImageServer(value interface{}) *Component {
	a.Set("uploadImageServer", value)
	return a
}

func (a *Component) UploadVideoMaxNumber(value interface{}) *Component {
	a.Set("uploadVideoMaxNumber", value)
	return a
}

func (a *Component) UploadVideoMaxSize(value interface{}) *Component {
	a.Set("uploadVideoMaxSize", value)
	return a
}

func (a *Component) UploadVideoServer(value interface{}) *Component {
	a.Set("uploadVideoServer", value)
	return a
}

func (a *Component) Url(value interface{}) *Component {
	a.Set("url", value)
	return a
}

func (a *Component) UrlField(value interface{}) *Component {
	a.Set("urlField", value)
	return a
}

func (a *Component) UseCardLabel(value interface{}) *Component {
	a.Set("useCardLabel", value)
	return a
}

func (a *Component) UseChunk(value interface{}) *Component {
	a.Set("useChunk", value)
	return a
}

func (a *Component) UseMobileUI(value interface{}) *Component {
	a.Set("useMobileUI", value)
	return a
}

func (a *Component) Utc(value interface{}) *Component {
	a.Set("utc", value)
	return a
}

func (a *Component) VAlign(value interface{}) *Component {
	a.Set("vAlign", value)
	return a
}

func (a *Component) ValidateApi(value interface{}) *Component {
	a.Set("validateApi", value)
	return a
}

func (a *Component) ValidateOnChange(value interface{}) *Component {
	a.Set("validateOnChange", value)
	return a
}

func (a *Component) ValidationConfig(value interface{}) *Component {
	a.Set("validationConfig", value)
	return a
}

func (a *Component) ValidationErrors(value interface{}) *Component {
	a.Set("validationErrors", value)
	return a
}

func (a *Component) Validations(value interface{}) *Component {
	a.Set("validations", value)
	return a
}

func (a *Component) Valign(value interface{}) *Component {
	a.Set("valign", value)
	return a
}

func (a *Component) Value(value interface{}) *Component {
	a.Set("value", value)
	return a
}

func (a *Component) ValueField(value interface{}) *Component {
	a.Set("valueField", value)
	return a
}

func (a *Component) ValueFormat(value interface{}) *Component {
	a.Set("valueFormat", value)
	return a
}

func (a *Component) ValuePlaceholder(value interface{}) *Component {
	a.Set("valuePlaceholder", value)
	return a
}

func (a *Component) ValueTpl(value interface{}) *Component {
	a.Set("valueTpl", value)
	return a
}

func (a *Component) ValueType(value interface{}) *Component {
	a.Set("valueType", value)
	return a
}

func (a *Component) ValuesNoWrap(value interface{}) *Component {
	a.Set("valuesNoWrap", value)
	return a
}

func (a *Component) Vendor(value interface{}) *Component {
	a.Set("vendor", value)
	return a
}

func (a *Component) Vertical(value interface{}) *Component {
	a.Set("vertical", value)
	return a
}

func (a *Component) VideoReceiver(value interface{}) *Component {
	a.Set("videoReceiver", value)
	return a
}

func (a *Component) VideoType(value interface{}) *Component {
	a.Set("videoType", value)
	return a
}

func (a *Component) VirtualThreshold(value interface{}) *Component {
	a.Set("virtualThreshold", value)
	return a
}

func (a *Component) Visible(value interface{}) *Component {
	a.Set("visible", value)
	return a
}

func (a *Component) VisibleOn(value interface{}) *Component {
	a.Set("visibleOn", value)
	return a
}

func (a *Component) Width(value interface{}) *Component {
	a.Set("width", value)
	return a
}

func (a *Component) WithChildren(value interface{}) *Component {
	a.Set("withChildren", value)
	return a
}

func (a *Component) WordWrap(value interface{}) *Component {
	a.Set("wordWrap", value)
	return a
}

func (a *Component) Words(value interface{}) *Component {
	a.Set("words", value)
	return a
}

func (a *Component) Wrap(value interface{}) *Component {
	a.Set("wrap", value)
	return a
}

func (a *Component) WrapWithPanel(value interface{}) *Component {
	a.Set("wrapWithPanel", value)
	return a
}

func (a *Component) WrapperBody(value interface{}) *Component {
	a.Set("wrapperBody", value)
	return a
}

func (a *Component) WrapperComponent(value interface{}) *Component {
	a.Set("wrapperComponent", value)
	return a
}

func (a *Component) WrapperCustomStyle(value interface{}) *Component {
	a.Set("wrapperCustomStyle", value)
	return a
}

func (a *Component) Ws(value interface{}) *Component {
	a.Set("ws", value)
	return a
}

func (a *Component) X(value interface{}) *Component {
	a.Set("x", value)
	return a
}

func (a *Component) Xs(value interface{}) *Component {
	a.Set("xs", value)
	return a
}

func (a *Component) Y(value interface{}) *Component {
	a.Set("y", value)
	return a
}

func (a *Component) ZIndex(value interface{}) *Component {
	a.Set("zIndex", value)
	return a
}
