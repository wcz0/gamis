package renderers


/**
 * 复选框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/checkboxes

 * @author wcz0
 * @version 6.2.2
 */
type CheckboxesControl struct {
	*BaseRenderer
}

func NewCheckboxesControl() *CheckboxesControl {
    a := &CheckboxesControl{
        BaseRenderer: NewBaseRenderer(),
    }

func (a *CheckboxesControl) Set(name string, value interface{}) *CheckboxesControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}

    a.Set("type", "checkboxes")
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *CheckboxesControl) Horizontal(value interface{}) *CheckboxesControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *CheckboxesControl) Hidden(value interface{}) *CheckboxesControl {
    a.Set("hidden", value)
    return a
}

/**
 * 组件样式
 */
func (a *CheckboxesControl) Style(value interface{}) *CheckboxesControl {
    a.Set("style", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *CheckboxesControl) Descriptionclassname(value interface{}) *CheckboxesControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *CheckboxesControl) Inline(value interface{}) *CheckboxesControl {
    a.Set("inline", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *CheckboxesControl) Deferapi(value interface{}) *CheckboxesControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *CheckboxesControl) Addapi(value interface{}) *CheckboxesControl {
    a.Set("addApi", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *CheckboxesControl) Labelwidth(value interface{}) *CheckboxesControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *CheckboxesControl) Creatable(value interface{}) *CheckboxesControl {
    a.Set("creatable", value)
    return a
}

/**
 * 配置 input className
 */
func (a *CheckboxesControl) Inputclassname(value interface{}) *CheckboxesControl {
    a.Set("inputClassName", value)
    return a
}

/**
 */
func (a *CheckboxesControl) Testidbuilder(value interface{}) *CheckboxesControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *CheckboxesControl) Source(value interface{}) *CheckboxesControl {
    a.Set("source", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *CheckboxesControl) Selectfirst(value interface{}) *CheckboxesControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *CheckboxesControl) Multiple(value interface{}) *CheckboxesControl {
    a.Set("multiple", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *CheckboxesControl) Validateapi(value interface{}) *CheckboxesControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *CheckboxesControl) Staticlabelclassname(value interface{}) *CheckboxesControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *CheckboxesControl) Width(value interface{}) *CheckboxesControl {
    a.Set("width", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *CheckboxesControl) Addcontrols(value interface{}) *CheckboxesControl {
    a.Set("addControls", value)
    return a
}

/**
 * 占位符
 */
func (a *CheckboxesControl) Placeholder(value interface{}) *CheckboxesControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *CheckboxesControl) Clearvalueonhidden(value interface{}) *CheckboxesControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *CheckboxesControl) Labelremark(value interface{}) *CheckboxesControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 表单项类型
 */
func (a *CheckboxesControl) Type(value interface{}) *CheckboxesControl {
    a.Set("type", value)
    return a
}

/**
 * 全选/不选文案
 */
func (a *CheckboxesControl) Checkalltext(value interface{}) *CheckboxesControl {
    a.Set("checkAllText", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *CheckboxesControl) Deferfield(value interface{}) *CheckboxesControl {
    a.Set("deferField", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *CheckboxesControl) Editable(value interface{}) *CheckboxesControl {
    a.Set("editable", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *CheckboxesControl) Staticplaceholder(value interface{}) *CheckboxesControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *CheckboxesControl) Editdialog(value interface{}) *CheckboxesControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 配置 label className
 */
func (a *CheckboxesControl) Labelclassname(value interface{}) *CheckboxesControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *CheckboxesControl) Staticon(value interface{}) *CheckboxesControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否开启全选功能
 */
func (a *CheckboxesControl) Checkall(value interface{}) *CheckboxesControl {
    a.Set("checkAll", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *CheckboxesControl) Resetvalue(value interface{}) *CheckboxesControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *CheckboxesControl) Adddialog(value interface{}) *CheckboxesControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *CheckboxesControl) Editapi(value interface{}) *CheckboxesControl {
    a.Set("editApi", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *CheckboxesControl) Value(value interface{}) *CheckboxesControl {
    a.Set("value", value)
    return a
}

/**
 */
func (a *CheckboxesControl) Row(value interface{}) *CheckboxesControl {
    a.Set("row", value)
    return a
}

/**
 * 选项集合
 */
func (a *CheckboxesControl) Options(value interface{}) *CheckboxesControl {
    a.Set("options", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *CheckboxesControl) Joinvalues(value interface{}) *CheckboxesControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *CheckboxesControl) Valuesnowrap(value interface{}) *CheckboxesControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *CheckboxesControl) Static(value interface{}) *CheckboxesControl {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *CheckboxesControl) Usemobileui(value interface{}) *CheckboxesControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 描述标题
 */
func (a *CheckboxesControl) Labelalign(value interface{}) *CheckboxesControl {
    a.Set("labelAlign", value)
    return a
}

/**
 */
func (a *CheckboxesControl) Desc(value interface{}) *CheckboxesControl {
    a.Set("desc", value)
    return a
}

/**
 * 是否为必填
 */
func (a *CheckboxesControl) Required(value interface{}) *CheckboxesControl {
    a.Set("required", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *CheckboxesControl) Hint(value interface{}) *CheckboxesControl {
    a.Set("hint", value)
    return a
}

/**
 * 是否只读
 */
func (a *CheckboxesControl) Readonly(value interface{}) *CheckboxesControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *CheckboxesControl) Classname(value interface{}) *CheckboxesControl {
    a.Set("className", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *CheckboxesControl) Extraname(value interface{}) *CheckboxesControl {
    a.Set("extraName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *CheckboxesControl) Remark(value interface{}) *CheckboxesControl {
    a.Set("remark", value)
    return a
}

/**
 * 每行显示多少个
 */
func (a *CheckboxesControl) Columnscount(value interface{}) *CheckboxesControl {
    a.Set("columnsCount", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *CheckboxesControl) Extractvalue(value interface{}) *CheckboxesControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *CheckboxesControl) Clearable(value interface{}) *CheckboxesControl {
    a.Set("clearable", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *CheckboxesControl) Editcontrols(value interface{}) *CheckboxesControl {
    a.Set("editControls", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *CheckboxesControl) Disabledon(value interface{}) *CheckboxesControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *CheckboxesControl) Editorsetting(value interface{}) *CheckboxesControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *CheckboxesControl) Initfetch(value interface{}) *CheckboxesControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *CheckboxesControl) Size(value interface{}) *CheckboxesControl {
    a.Set("size", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *CheckboxesControl) Mode(value interface{}) *CheckboxesControl {
    a.Set("mode", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *CheckboxesControl) Validateonchange(value interface{}) *CheckboxesControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *CheckboxesControl) Description(value interface{}) *CheckboxesControl {
    a.Set("description", value)
    return a
}

/**
 */
func (a *CheckboxesControl) Staticschema(value interface{}) *CheckboxesControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否默认全选
 */
func (a *CheckboxesControl) Defaultcheckall(value interface{}) *CheckboxesControl {
    a.Set("defaultCheckAll", value)
    return a
}

/**
 * 自定义选项展示
 */
func (a *CheckboxesControl) Menutpl(value interface{}) *CheckboxesControl {
    a.Set("menuTpl", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *CheckboxesControl) Deleteapi(value interface{}) *CheckboxesControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *CheckboxesControl) Deleteconfirmtext(value interface{}) *CheckboxesControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *CheckboxesControl) Name(value interface{}) *CheckboxesControl {
    a.Set("name", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *CheckboxesControl) Staticclassname(value interface{}) *CheckboxesControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 分割符
 */
func (a *CheckboxesControl) Delimiter(value interface{}) *CheckboxesControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 新增文字
 */
func (a *CheckboxesControl) Createbtnlabel(value interface{}) *CheckboxesControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *CheckboxesControl) Submitonchange(value interface{}) *CheckboxesControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 是否禁用
 */
func (a *CheckboxesControl) Disabled(value interface{}) *CheckboxesControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示
 */
func (a *CheckboxesControl) Visible(value interface{}) *CheckboxesControl {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *CheckboxesControl) Staticinputclassname(value interface{}) *CheckboxesControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否可删除
 */
func (a *CheckboxesControl) Removable(value interface{}) *CheckboxesControl {
    a.Set("removable", value)
    return a
}

/**
 */
func (a *CheckboxesControl) Validations(value interface{}) *CheckboxesControl {
    a.Set("validations", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *CheckboxesControl) Autofill(value interface{}) *CheckboxesControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *CheckboxesControl) Validationerrors(value interface{}) *CheckboxesControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *CheckboxesControl) Id(value interface{}) *CheckboxesControl {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *CheckboxesControl) Onevent(value interface{}) *CheckboxesControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *CheckboxesControl) Initfetchon(value interface{}) *CheckboxesControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 描述标题
 */
func (a *CheckboxesControl) Label(value interface{}) *CheckboxesControl {
    a.Set("label", value)
    return a
}

/**
 * 只读条件
 */
func (a *CheckboxesControl) Readonlyon(value interface{}) *CheckboxesControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 */
func (a *CheckboxesControl) Initautofill(value interface{}) *CheckboxesControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *CheckboxesControl) Hiddenon(value interface{}) *CheckboxesControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *CheckboxesControl) Visibleon(value interface{}) *CheckboxesControl {
    a.Set("visibleOn", value)
    return a
}
