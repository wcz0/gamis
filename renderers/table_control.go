package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
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

/**
 * 是否静态展示表达式
 */
func (a *TableControl) StaticOn(value interface{}) *TableControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 配置 label className
 */
func (a *TableControl) LabelClassName(value interface{}) *TableControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 取消按钮图标
 */
func (a *TableControl) CancelBtnIcon(value interface{}) *TableControl {
    a.Set("cancelBtnIcon", value)
    return a
}

/**
 * 底部总结行
 */
func (a *TableControl) AffixRow(value interface{}) *TableControl {
    a.Set("affixRow", value)
    return a
}

/**
 * 行样式表表达式
 */
func (a *TableControl) RowClassNameExpr(value interface{}) *TableControl {
    a.Set("rowClassNameExpr", value)
    return a
}

/**
 */
func (a *TableControl) Testid(value interface{}) *TableControl {
    a.Set("testid", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *TableControl) StaticPlaceholder(value interface{}) *TableControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 更新 API
 */
func (a *TableControl) UpdateApi(value interface{}) *TableControl {
    a.Set("updateApi", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *TableControl) Id(value interface{}) *TableControl {
    a.Set("id", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *TableControl) StaticClassName(value interface{}) *TableControl {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *TableControl) InitAutoFill(value interface{}) *TableControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 确认按钮图标
 */
func (a *TableControl) ConfirmBtnIcon(value interface{}) *TableControl {
    a.Set("confirmBtnIcon", value)
    return a
}

/**
 * 合并单元格配置，配置数字表示从左到右的多少列自动合并单元格。
 */
func (a *TableControl) CombineNum(value interface{}) *TableControl {
    a.Set("combineNum", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *TableControl) VisibleOn(value interface{}) *TableControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *TableControl) DescriptionClassName(value interface{}) *TableControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否开启 static 状态切换
 */
func (a *TableControl) EnableStaticTransform(value interface{}) *TableControl {
    a.Set("enableStaticTransform", value)
    return a
}

/**
 * 是否可以新增子项
 */
func (a *TableControl) ChildrenAddable(value interface{}) *TableControl {
    a.Set("childrenAddable", value)
    return a
}

/**
 * 是否显示表格操作栏新增按钮
 */
func (a *TableControl) ShowTableAddBtn(value interface{}) *TableControl {
    a.Set("showTableAddBtn", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *TableControl) Size(value interface{}) *TableControl {
    a.Set("size", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *TableControl) SubmitOnChange(value interface{}) *TableControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 */
func (a *TableControl) Desc(value interface{}) *TableControl {
    a.Set("desc", value)
    return a
}

/**
 * 可新增
 */
func (a *TableControl) Addable(value interface{}) *TableControl {
    a.Set("addable", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *TableControl) StaticLabelClassName(value interface{}) *TableControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *TableControl) LabelWidth(value interface{}) *TableControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 更新按钮名称
 */
func (a *TableControl) EditBtnLabel(value interface{}) *TableControl {
    a.Set("editBtnLabel", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *TableControl) HiddenOn(value interface{}) *TableControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *TableControl) ValidateOnChange(value interface{}) *TableControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 复制按钮图标
 */
func (a *TableControl) CopyBtnIcon(value interface{}) *TableControl {
    a.Set("copyBtnIcon", value)
    return a
}

/**
 * 限制最小个数
 */
func (a *TableControl) MinLength(value interface{}) *TableControl {
    a.Set("minLength", value)
    return a
}

/**
 * 值字段
 */
func (a *TableControl) ValueField(value interface{}) *TableControl {
    a.Set("valueField", value)
    return a
}

/**
 * 描述标题
 */
func (a *TableControl) LabelAlign(value interface{}) *TableControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *TableControl) Description(value interface{}) *TableControl {
    a.Set("description", value)
    return a
}

/**
 * 是否为必填
 */
func (a *TableControl) Required(value interface{}) *TableControl {
    a.Set("required", value)
    return a
}

/**
 * 可复制新增
 */
func (a *TableControl) Copyable(value interface{}) *TableControl {
    a.Set("copyable", value)
    return a
}

/**
 * 初始值，新增的时候
 */
func (a *TableControl) Scaffold(value interface{}) *TableControl {
    a.Set("scaffold", value)
    return a
}

/**
 * 是否可调整列宽
 */
func (a *TableControl) Resizable(value interface{}) *TableControl {
    a.Set("resizable", value)
    return a
}

/**
 * 描述标题
 */
func (a *TableControl) Label(value interface{}) *TableControl {
    a.Set("label", value)
    return a
}

/**
 */
func (a *TableControl) Row(value interface{}) *TableControl {
    a.Set("row", value)
    return a
}

/**
 * 底部外层 CSS 类名
 */
func (a *TableControl) FooterClassName(value interface{}) *TableControl {
    a.Set("footerClassName", value)
    return a
}

/**
 * 合并单元格配置，配置从第几列开始合并。
 */
func (a *TableControl) CombineFromIndex(value interface{}) *TableControl {
    a.Set("combineFromIndex", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *TableControl) DisabledOn(value interface{}) *TableControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 孩子新增按钮图标
 */
func (a *TableControl) SubAddBtnIcon(value interface{}) *TableControl {
    a.Set("subAddBtnIcon", value)
    return a
}

/**
 * 标题
 */
func (a *TableControl) Title(value interface{}) *TableControl {
    a.Set("title", value)
    return a
}

/**
 * 底部工具栏CSS样式类
 */
func (a *TableControl) ToolbarClassName(value interface{}) *TableControl {
    a.Set("toolbarClassName", value)
    return a
}

/**
 * 表格自动计算高度
 */
func (a *TableControl) AutoFillHeight(value interface{}) *TableControl {
    a.Set("autoFillHeight", value)
    return a
}

/**
 */
func (a *TableControl) TestIdBuilder(value interface{}) *TableControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 顶部总结行
 */
func (a *TableControl) PrefixRow(value interface{}) *TableControl {
    a.Set("prefixRow", value)
    return a
}

/**
 * 新增 API
 */
func (a *TableControl) AddApi(value interface{}) *TableControl {
    a.Set("addApi", value)
    return a
}

/**
 * 底部新增按钮配置
 */
func (a *TableControl) FooterAddBtn(value interface{}) *TableControl {
    a.Set("footerAddBtn", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *TableControl) ValidateApi(value interface{}) *TableControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *TableControl) Hidden(value interface{}) *TableControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *TableControl) Static(value interface{}) *TableControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *TableControl) StaticInputClassName(value interface{}) *TableControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *TableControl) Style(value interface{}) *TableControl {
    a.Set("style", value)
    return a
}

/**
 * 是否显示头部
 */
func (a *TableControl) ShowHeader(value interface{}) *TableControl {
    a.Set("showHeader", value)
    return a
}

/**
 * 表格 CSS 类名
 */
func (a *TableControl) TableClassName(value interface{}) *TableControl {
    a.Set("tableClassName", value)
    return a
}

/**
 */
func (a *TableControl) StaticSchema(value interface{}) *TableControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *TableControl) ValidationErrors(value interface{}) *TableControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 删除的 API
 */
func (a *TableControl) DeleteApi(value interface{}) *TableControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 数据源：绑定当前环境变量
 */
func (a *TableControl) Source(value interface{}) *TableControl {
    a.Set("source", value)
    return a
}

/**
 * 开启查询区域，会根据列元素的searchable属性值，自动生成查询条件表单
 */
func (a *TableControl) AutoGenerateFilter(value interface{}) *TableControl {
    a.Set("autoGenerateFilter", value)
    return a
}

/**
 * 是否只读
 */
func (a *TableControl) ReadOnly(value interface{}) *TableControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置 input className
 */
func (a *TableControl) InputClassName(value interface{}) *TableControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 展示列显示开关，自动即：列数量大于或等于5个时自动开启
 */
func (a *TableControl) ColumnsTogglable(value interface{}) *TableControl {
    a.Set("columnsTogglable", value)
    return a
}

/**
 */
func (a *TableControl) Validations(value interface{}) *TableControl {
    a.Set("validations", value)
    return a
}

/**
 * 确认按钮文字
 */
func (a *TableControl) ConfirmBtnLabel(value interface{}) *TableControl {
    a.Set("confirmBtnLabel", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *TableControl) ExtraName(value interface{}) *TableControl {
    a.Set("extraName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *TableControl) Inline(value interface{}) *TableControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否为确认的编辑模式。
 */
func (a *TableControl) NeedConfirm(value interface{}) *TableControl {
    a.Set("needConfirm", value)
    return a
}

/**
 * 是否可以拖拽排序
 */
func (a *TableControl) Draggable(value interface{}) *TableControl {
    a.Set("draggable", value)
    return a
}

/**
 * 取消按钮文字
 */
func (a *TableControl) CancelBtnLabel(value interface{}) *TableControl {
    a.Set("cancelBtnLabel", value)
    return a
}

/**
 * 限制最大个数
 */
func (a *TableControl) MaxLength(value interface{}) *TableControl {
    a.Set("maxLength", value)
    return a
}

/**
 * 是否固底
 */
func (a *TableControl) AffixFooter(value interface{}) *TableControl {
    a.Set("affixFooter", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *TableControl) Hint(value interface{}) *TableControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否显示复制按钮
 */
func (a *TableControl) CopyAddBtn(value interface{}) *TableControl {
    a.Set("copyAddBtn", value)
    return a
}

/**
 * 顶部外层 CSS 类名
 */
func (a *TableControl) HeaderClassName(value interface{}) *TableControl {
    a.Set("headerClassName", value)
    return a
}

/**
 * 只读条件
 */
func (a *TableControl) ReadOnlyOn(value interface{}) *TableControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *TableControl) ClearValueOnHidden(value interface{}) *TableControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *TableControl) Width(value interface{}) *TableControl {
    a.Set("width", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *TableControl) LabelRemark(value interface{}) *TableControl {
    a.Set("labelRemark", value)
    return a
}

/**
 */
func (a *TableControl) Type(value interface{}) *TableControl {
    a.Set("type", value)
    return a
}

/**
 * 可否编辑
 */
func (a *TableControl) Editable(value interface{}) *TableControl {
    a.Set("editable", value)
    return a
}

/**
 * 行角标
 */
func (a *TableControl) ItemBadge(value interface{}) *TableControl {
    a.Set("itemBadge", value)
    return a
}

/**
 * 是否可以访问父级数据，正常 combo 已经关联到数组成员，是不能访问父级数据的。
 */
func (a *TableControl) CanAccessSuperData(value interface{}) *TableControl {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * 是否禁用
 */
func (a *TableControl) Disabled(value interface{}) *TableControl {
    a.Set("disabled", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *TableControl) Name(value interface{}) *TableControl {
    a.Set("name", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *TableControl) OnEvent(value interface{}) *TableControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *TableControl) UseMobileUI(value interface{}) *TableControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 复制按钮文字
 */
func (a *TableControl) CopyBtnLabel(value interface{}) *TableControl {
    a.Set("copyBtnLabel", value)
    return a
}

/**
 * 是否显示底部新增按钮
 */
func (a *TableControl) ShowFooterAddBtn(value interface{}) *TableControl {
    a.Set("showFooterAddBtn", value)
    return a
}

/**
 * 是否开启底部展示功能，适合移动端展示
 */
func (a *TableControl) Footable(value interface{}) *TableControl {
    a.Set("footable", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *TableControl) ClassName(value interface{}) *TableControl {
    a.Set("className", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *TableControl) Horizontal(value interface{}) *TableControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 删除确认文字
 */
func (a *TableControl) DeleteConfirmText(value interface{}) *TableControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 是否显示底部
 */
func (a *TableControl) ShowFooter(value interface{}) *TableControl {
    a.Set("showFooter", value)
    return a
}

/**
 * 懒加载 API，当行数据中用 defer: true 标记了，则其孩子节点将会用这个 API 来拉取数据。
 */
func (a *TableControl) DeferApi(value interface{}) *TableControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *TableControl) Remark(value interface{}) *TableControl {
    a.Set("remark", value)
    return a
}

/**
 * 删除按钮文字
 */
func (a *TableControl) DeleteBtnLabel(value interface{}) *TableControl {
    a.Set("deleteBtnLabel", value)
    return a
}

/**
 * 是否固定表头
 */
func (a *TableControl) AffixHeader(value interface{}) *TableControl {
    a.Set("affixHeader", value)
    return a
}

/**
 * 复制的时候用来配置复制映射的数据。默认值是 {&:$$}，相当与复制整个行数据 通常有时候需要用来标记是复制过来的，也可能需要删掉一下主键字段。
 */
func (a *TableControl) CopyData(value interface{}) *TableControl {
    a.Set("copyData", value)
    return a
}

/**
 * 删除按钮图标
 */
func (a *TableControl) DeleteBtnIcon(value interface{}) *TableControl {
    a.Set("deleteBtnIcon", value)
    return a
}

/**
 * 新增按钮文字
 */
func (a *TableControl) AddBtnLabel(value interface{}) *TableControl {
    a.Set("addBtnLabel", value)
    return a
}

/**
 * 新增按钮图标
 */
func (a *TableControl) AddBtnIcon(value interface{}) *TableControl {
    a.Set("addBtnIcon", value)
    return a
}

/**
 * 是否显示序号
 */
func (a *TableControl) ShowIndex(value interface{}) *TableControl {
    a.Set("showIndex", value)
    return a
}

/**
 * 表格的列信息
 */
func (a *TableControl) Columns(value interface{}) *TableControl {
    a.Set("columns", value)
    return a
}

/**
 * table layout
 * 可选值: fixed | auto
 */
func (a *TableControl) TableLayout(value interface{}) *TableControl {
    a.Set("tableLayout", value)
    return a
}

/**
 * 是否显示
 */
func (a *TableControl) Visible(value interface{}) *TableControl {
    a.Set("visible", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *TableControl) Mode(value interface{}) *TableControl {
    a.Set("mode", value)
    return a
}

/**
 * 占位符
 */
func (a *TableControl) Placeholder(value interface{}) *TableControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *TableControl) EditorSetting(value interface{}) *TableControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *TableControl) AutoFill(value interface{}) *TableControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 更新按钮图标
 */
func (a *TableControl) EditBtnIcon(value interface{}) *TableControl {
    a.Set("editBtnIcon", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *TableControl) Value(value interface{}) *TableControl {
    a.Set("value", value)
    return a
}

/**
 * 孩子新增按钮文字
 */
func (a *TableControl) SubAddBtnLabel(value interface{}) *TableControl {
    a.Set("subAddBtnLabel", value)
    return a
}

/**
 * 可否删除
 */
func (a *TableControl) Removable(value interface{}) *TableControl {
    a.Set("removable", value)
    return a
}

/**
 * 分页个数，默认不分页
 */
func (a *TableControl) PerPage(value interface{}) *TableControl {
    a.Set("perPage", value)
    return a
}
