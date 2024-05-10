package renderers


/**
 * Combo 组合输入框类型 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/combo

 * @author wcz0
 * @version 6.2.2
 */
type ComboControl struct {
	*BaseRenderer
}

func NewComboControl() *ComboControl {
    a := &ComboControl{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "combo")
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ComboControl) StaticOn(value interface{}) *ComboControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ComboControl) UseMobileUI(value interface{}) *ComboControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *ComboControl) Mode(value interface{}) *ComboControl {
    a.Set("mode", value)
    return a
}

/**
 * 确认删除时的提示
 */
func (a *ComboControl) DeleteConfirmText(value interface{}) *ComboControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 新增按钮CSS类名
 */
func (a *ComboControl) AddButtonClassName(value interface{}) *ComboControl {
    a.Set("addButtonClassName", value)
    return a
}

/**
 * Add at top
 */
func (a *ComboControl) Addattop(value interface{}) *ComboControl {
    a.Set("addattop", value)
    return a
}

/**
 * 描述标题
 */
func (a *ComboControl) LabelAlign(value interface{}) *ComboControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *ComboControl) Remark(value interface{}) *ComboControl {
    a.Set("remark", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ComboControl) StaticInputClassName(value interface{}) *ComboControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ComboControl) EditorSetting(value interface{}) *ComboControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *ComboControl) Size(value interface{}) *ComboControl {
    a.Set("size", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *ComboControl) LabelRemark(value interface{}) *ComboControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *ComboControl) ValidateApi(value interface{}) *ComboControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 当扁平化开启的时候，是否用分隔符的形式发送给后端，否则采用array的方式
 */
func (a *ComboControl) JoinValues(value interface{}) *ComboControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 是否可多选
 */
func (a *ComboControl) Multiple(value interface{}) *ComboControl {
    a.Set("multiple", value)
    return a
}

/**
 * 是否显示
 */
func (a *ComboControl) Visible(value interface{}) *ComboControl {
    a.Set("visible", value)
    return a
}

/**
 * 组件样式
 */
func (a *ComboControl) Style(value interface{}) *ComboControl {
    a.Set("style", value)
    return a
}

/**
 * 配置 label className
 */
func (a *ComboControl) LabelClassName(value interface{}) *ComboControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *ComboControl) Description(value interface{}) *ComboControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *ComboControl) Desc(value interface{}) *ComboControl {
    a.Set("desc", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *ComboControl) ClearValueOnHidden(value interface{}) *ComboControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 数组输入框的子项
 */
func (a *ComboControl) Items(value interface{}) *ComboControl {
    a.Set("items", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ComboControl) OnEvent(value interface{}) *ComboControl {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *ComboControl) TestIdBuilder(value interface{}) *ComboControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *ComboControl) Hint(value interface{}) *ComboControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否可拖拽排序
 */
func (a *ComboControl) Draggable(value interface{}) *ComboControl {
    a.Set("draggable", value)
    return a
}

/**
 * 限制最小个数
 */
func (a *ComboControl) MinLength(value interface{}) *ComboControl {
    a.Set("minLength", value)
    return a
}

/**
 * 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
 */
func (a *ComboControl) SubFormHorizontal(value interface{}) *ComboControl {
    a.Set("subFormHorizontal", value)
    return a
}

/**
 * 允许为空，如果子表单项里面配置验证器，且又是单条模式。可以允许用户选择清空（不填）。
 */
func (a *ComboControl) Nullable(value interface{}) *ComboControl {
    a.Set("nullable", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ComboControl) StaticPlaceholder(value interface{}) *ComboControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ComboControl) StaticClassName(value interface{}) *ComboControl {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *ComboControl) InitAutoFill(value interface{}) *ComboControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *ComboControl) HiddenOn(value interface{}) *ComboControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ComboControl) VisibleOn(value interface{}) *ComboControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *ComboControl) LabelWidth(value interface{}) *ComboControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *ComboControl) Horizontal(value interface{}) *ComboControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 指定为组合输入框类型
 */
func (a *ComboControl) Type(value interface{}) *ComboControl {
    a.Set("type", value)
    return a
}

/**
 * 限制最大个数
 */
func (a *ComboControl) MaxLength(value interface{}) *ComboControl {
    a.Set("maxLength", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ComboControl) Id(value interface{}) *ComboControl {
    a.Set("id", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *ComboControl) Value(value interface{}) *ComboControl {
    a.Set("value", value)
    return a
}

/**
 * 符合某类条件后才渲染的schema
 */
func (a *ComboControl) Conditions(value interface{}) *ComboControl {
    a.Set("conditions", value)
    return a
}

/**
 * 是否可以访问父级数据，正常 combo 已经关联到数组成员，是不能访问父级数据的。
 */
func (a *ComboControl) CanAccessSuperData(value interface{}) *ComboControl {
    a.Set("canAccessSuperData", value)
    return a
}

/**
 * 提示信息
 */
func (a *ComboControl) Messages(value interface{}) *ComboControl {
    a.Set("messages", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ComboControl) Static(value interface{}) *ComboControl {
    a.Set("static", value)
    return a
}

/**
 * 数据比较多，比较卡时，可以试试开启。
 */
func (a *ComboControl) LazyLoad(value interface{}) *ComboControl {
    a.Set("lazyLoad", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ComboControl) Disabled(value interface{}) *ComboControl {
    a.Set("disabled", value)
    return a
}

/**
 * 描述标题
 */
func (a *ComboControl) Label(value interface{}) *ComboControl {
    a.Set("label", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *ComboControl) Name(value interface{}) *ComboControl {
    a.Set("name", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *ComboControl) ValidateOnChange(value interface{}) *ComboControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *ComboControl) DescriptionClassName(value interface{}) *ComboControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *ComboControl) Inline(value interface{}) *ComboControl {
    a.Set("inline", value)
    return a
}

/**
 * 是否可切换条件，配合`conditions`使用
 */
func (a *ComboControl) TypeSwitchable(value interface{}) *ComboControl {
    a.Set("typeSwitchable", value)
    return a
}

/**
 * 内部单组表单项的类名
 */
func (a *ComboControl) FormClassName(value interface{}) *ComboControl {
    a.Set("formClassName", value)
    return a
}

/**
 * 新增按钮文字
 */
func (a *ComboControl) AddButtonText(value interface{}) *ComboControl {
    a.Set("addButtonText", value)
    return a
}

/**
 * Tabs 的展示模式。
 * 可选值:  | line | card | radio
 */
func (a *ComboControl) TabsStyle(value interface{}) *ComboControl {
    a.Set("tabsStyle", value)
    return a
}

/**
 * 配置 input className
 */
func (a *ComboControl) InputClassName(value interface{}) *ComboControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ComboControl) DisabledOn(value interface{}) *ComboControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *ComboControl) SubmitOnChange(value interface{}) *ComboControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 没有成员时显示。
 */
func (a *ComboControl) Placeholder(value interface{}) *ComboControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 是否可删除
 */
func (a *ComboControl) Removable(value interface{}) *ComboControl {
    a.Set("removable", value)
    return a
}

/**
 * 配置同步字段。只有 `strictMode` 为 `false` 时有效。 如果 Combo 层级比较深，底层的获取外层的数据可能不同步。 但是给 combo 配置这个属性就能同步下来。输入格式：`["os"]`
 */
func (a *ComboControl) SyncFields(value interface{}) *ComboControl {
    a.Set("syncFields", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ComboControl) Width(value interface{}) *ComboControl {
    a.Set("width", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ComboControl) ClassName(value interface{}) *ComboControl {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ComboControl) Hidden(value interface{}) *ComboControl {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ComboControl) StaticLabelClassName(value interface{}) *ComboControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 是否只读
 */
func (a *ComboControl) ReadOnly(value interface{}) *ComboControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *ComboControl) ValidationErrors(value interface{}) *ComboControl {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *ComboControl) Validations(value interface{}) *ComboControl {
    a.Set("validations", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *ComboControl) AutoFill(value interface{}) *ComboControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 单组表单项初始值。默认为 `{}`
 */
func (a *ComboControl) Scaffold(value interface{}) *ComboControl {
    a.Set("scaffold", value)
    return a
}

/**
 * 删除时调用的api
 */
func (a *ComboControl) DeleteApi(value interface{}) *ComboControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 是否可新增
 */
func (a *ComboControl) Addable(value interface{}) *ComboControl {
    a.Set("addable", value)
    return a
}

/**
 * 可拖拽排序的提示信息。
 */
func (a *ComboControl) DraggableTip(value interface{}) *ComboControl {
    a.Set("draggableTip", value)
    return a
}

/**
 * 采用 Tabs 展示方式？
 */
func (a *ComboControl) TabsMode(value interface{}) *ComboControl {
    a.Set("tabsMode", value)
    return a
}

/**
 * 只读条件
 */
func (a *ComboControl) ReadOnlyOn(value interface{}) *ComboControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 选项卡标题的生成模板。
 */
func (a *ComboControl) TabsLabelTpl(value interface{}) *ComboControl {
    a.Set("tabsLabelTpl", value)
    return a
}

/**
 * 严格模式，为了性能默认不开的。
 */
func (a *ComboControl) StrictMode(value interface{}) *ComboControl {
    a.Set("strictMode", value)
    return a
}

/**
 */
func (a *ComboControl) StaticSchema(value interface{}) *ComboControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *ComboControl) ExtraName(value interface{}) *ComboControl {
    a.Set("extraName", value)
    return a
}

/**
 * 是否为必填
 */
func (a *ComboControl) Required(value interface{}) *ComboControl {
    a.Set("required", value)
    return a
}

/**
 * 是否将结果扁平化(去掉name),只有当controls的length为1且multiple为true的时候才有效
 */
func (a *ComboControl) Flat(value interface{}) *ComboControl {
    a.Set("flat", value)
    return a
}

/**
 * 当扁平化开启并且joinValues为true时，用什么分隔符
 */
func (a *ComboControl) Delimiter(value interface{}) *ComboControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 子表单的模式。
 * 可选值: normal | horizontal | inline
 */
func (a *ComboControl) SubFormMode(value interface{}) *ComboControl {
    a.Set("subFormMode", value)
    return a
}

/**
 * 是否含有边框
 */
func (a *ComboControl) NoBorder(value interface{}) *ComboControl {
    a.Set("noBorder", value)
    return a
}

/**
 * 是否多行模式，默认一行展示完
 */
func (a *ComboControl) MultiLine(value interface{}) *ComboControl {
    a.Set("multiLine", value)
    return a
}

/**
 */
func (a *ComboControl) UpdatePristineAfterStoreDataReInit(value interface{}) *ComboControl {
    a.Set("updatePristineAfterStoreDataReInit", value)
    return a
}
