package renderers


/**
 * Picker 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/picker

 * @author wcz0
 * @version 6.2.2
 */
type PickerControl struct {
	*BaseRenderer
}

func NewPickerControl() *PickerControl {
    a := &PickerControl{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *PickerControl) Set(name string, value interface{}) *PickerControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "picker")
    return a
}

/**
 * 配置 label className
 */
func (a *PickerControl) Labelclassname(value interface{}) *PickerControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *PickerControl) Style(value interface{}) *PickerControl {
    a.Set("style", value)
    return a
}

/**
 * 内嵌模式，也就是说不弹框了。
 */
func (a *PickerControl) Embed(value interface{}) *PickerControl {
    a.Set("embed", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *PickerControl) Multiple(value interface{}) *PickerControl {
    a.Set("multiple", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *PickerControl) Validationerrors(value interface{}) *PickerControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *PickerControl) Hiddenon(value interface{}) *PickerControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *PickerControl) Deferfield(value interface{}) *PickerControl {
    a.Set("deferField", value)
    return a
}

/**
 * 新增文字
 */
func (a *PickerControl) Createbtnlabel(value interface{}) *PickerControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *PickerControl) Editapi(value interface{}) *PickerControl {
    a.Set("editApi", value)
    return a
}

/**
 * 是否只读
 */
func (a *PickerControl) Readonly(value interface{}) *PickerControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *PickerControl) Staticclassname(value interface{}) *PickerControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 开启最大标签展示数量的相关配置
 */
func (a *PickerControl) Overflowconfig(value interface{}) *PickerControl {
    a.Set("overflowConfig", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *PickerControl) Clearable(value interface{}) *PickerControl {
    a.Set("clearable", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *PickerControl) Creatable(value interface{}) *PickerControl {
    a.Set("creatable", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *PickerControl) Hint(value interface{}) *PickerControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *PickerControl) Hidden(value interface{}) *PickerControl {
    a.Set("hidden", value)
    return a
}

/**
 * 是否显示
 */
func (a *PickerControl) Visible(value interface{}) *PickerControl {
    a.Set("visible", value)
    return a
}

/**
 */
func (a *PickerControl) Staticschema(value interface{}) *PickerControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *PickerControl) Remark(value interface{}) *PickerControl {
    a.Set("remark", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *PickerControl) Initfetch(value interface{}) *PickerControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 是否为必填
 */
func (a *PickerControl) Required(value interface{}) *PickerControl {
    a.Set("required", value)
    return a
}

/**
 * 弹窗选择框详情。
 */
func (a *PickerControl) Pickerschema(value interface{}) *PickerControl {
    a.Set("pickerSchema", value)
    return a
}

/**
 * 是否可删除
 */
func (a *PickerControl) Removable(value interface{}) *PickerControl {
    a.Set("removable", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *PickerControl) Classname(value interface{}) *PickerControl {
    a.Set("className", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *PickerControl) Selectfirst(value interface{}) *PickerControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *PickerControl) Editable(value interface{}) *PickerControl {
    a.Set("editable", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *PickerControl) Mode(value interface{}) *PickerControl {
    a.Set("mode", value)
    return a
}

/**
 * 配置 input className
 */
func (a *PickerControl) Inputclassname(value interface{}) *PickerControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 选项集合
 */
func (a *PickerControl) Options(value interface{}) *PickerControl {
    a.Set("options", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *PickerControl) Initfetchon(value interface{}) *PickerControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *PickerControl) Joinvalues(value interface{}) *PickerControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *PickerControl) Addapi(value interface{}) *PickerControl {
    a.Set("addApi", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *PickerControl) Editcontrols(value interface{}) *PickerControl {
    a.Set("editControls", value)
    return a
}

/**
 */
func (a *PickerControl) Desc(value interface{}) *PickerControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *PickerControl) Descriptionclassname(value interface{}) *PickerControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *PickerControl) Visibleon(value interface{}) *PickerControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *PickerControl) Deleteconfirmtext(value interface{}) *PickerControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *PickerControl) Size(value interface{}) *PickerControl {
    a.Set("size", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *PickerControl) Extraname(value interface{}) *PickerControl {
    a.Set("extraName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *PickerControl) Inline(value interface{}) *PickerControl {
    a.Set("inline", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *PickerControl) Validateapi(value interface{}) *PickerControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *PickerControl) Staticinputclassname(value interface{}) *PickerControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *PickerControl) Labelremark(value interface{}) *PickerControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *PickerControl) Addcontrols(value interface{}) *PickerControl {
    a.Set("addControls", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *PickerControl) Labelwidth(value interface{}) *PickerControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *PickerControl) Disabledon(value interface{}) *PickerControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *PickerControl) Usemobileui(value interface{}) *PickerControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 可用来生成选中的值的描述文字
 */
func (a *PickerControl) Labeltpl(value interface{}) *PickerControl {
    a.Set("labelTpl", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *PickerControl) Valuesnowrap(value interface{}) *PickerControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *PickerControl) Deferapi(value interface{}) *PickerControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *PickerControl) Value(value interface{}) *PickerControl {
    a.Set("value", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *PickerControl) Autofill(value interface{}) *PickerControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 是否禁用
 */
func (a *PickerControl) Disabled(value interface{}) *PickerControl {
    a.Set("disabled", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *PickerControl) Onevent(value interface{}) *PickerControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *PickerControl) Staticlabelclassname(value interface{}) *PickerControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *PickerControl) Label(value interface{}) *PickerControl {
    a.Set("label", value)
    return a
}

/**
 * 占位符
 */
func (a *PickerControl) Placeholder(value interface{}) *PickerControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *PickerControl) Width(value interface{}) *PickerControl {
    a.Set("width", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *PickerControl) Submitonchange(value interface{}) *PickerControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *PickerControl) Horizontal(value interface{}) *PickerControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 建议用 labelTpl 选中一个字段名用来作为值的描述文字
 */
func (a *PickerControl) Labelfield(value interface{}) *PickerControl {
    a.Set("labelField", value)
    return a
}

/**
 * 选一个可以用来作为值的字段。
 */
func (a *PickerControl) Valuefield(value interface{}) *PickerControl {
    a.Set("valueField", value)
    return a
}

/**
 * 弹窗模式，dialog 或者 drawer
 * 可选值: dialog | drawer
 */
func (a *PickerControl) Modalmode(value interface{}) *PickerControl {
    a.Set("modalMode", value)
    return a
}

/**
 * 弹窗的标题，默认为情选择
 */
func (a *PickerControl) Modaltitle(value interface{}) *PickerControl {
    a.Set("modalTitle", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *PickerControl) Source(value interface{}) *PickerControl {
    a.Set("source", value)
    return a
}

/**
 * 分割符
 */
func (a *PickerControl) Delimiter(value interface{}) *PickerControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *PickerControl) Deleteapi(value interface{}) *PickerControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 只读条件
 */
func (a *PickerControl) Readonlyon(value interface{}) *PickerControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *PickerControl) Description(value interface{}) *PickerControl {
    a.Set("description", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *PickerControl) Editdialog(value interface{}) *PickerControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *PickerControl) Name(value interface{}) *PickerControl {
    a.Set("name", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *PickerControl) Clearvalueonhidden(value interface{}) *PickerControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 */
func (a *PickerControl) Row(value interface{}) *PickerControl {
    a.Set("row", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *PickerControl) Static(value interface{}) *PickerControl {
    a.Set("static", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *PickerControl) Staticon(value interface{}) *PickerControl {
    a.Set("staticOn", value)
    return a
}

/**
 */
func (a *PickerControl) Testidbuilder(value interface{}) *PickerControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *PickerControl) Id(value interface{}) *PickerControl {
    a.Set("id", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *PickerControl) Extractvalue(value interface{}) *PickerControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *PickerControl) Resetvalue(value interface{}) *PickerControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *PickerControl) Adddialog(value interface{}) *PickerControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 描述标题
 */
func (a *PickerControl) Labelalign(value interface{}) *PickerControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *PickerControl) Validateonchange(value interface{}) *PickerControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 */
func (a *PickerControl) Validations(value interface{}) *PickerControl {
    a.Set("validations", value)
    return a
}

/**
 */
func (a *PickerControl) Initautofill(value interface{}) *PickerControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *PickerControl) Staticplaceholder(value interface{}) *PickerControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *PickerControl) Editorsetting(value interface{}) *PickerControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 表单项类型
 */
func (a *PickerControl) Type(value interface{}) *PickerControl {
    a.Set("type", value)
    return a
}
