package renderers


/**
 * Radio 单选框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/radios

 * @author wcz0
 * @version 6.2.2
 */
type RadiosControl struct {
	*BaseRenderer
}

func NewRadiosControl() *RadiosControl {
    a := &RadiosControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "radios")
    return a
}


func (a *RadiosControl) Set(name string, value interface{}) *RadiosControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否显示
 */
func (a *RadiosControl) Visible(value interface{}) *RadiosControl {
    a.Set("visible", value)
    return a
}

/**
 * 新增时的表单项。
 */
func (a *RadiosControl) AddControls(value interface{}) *RadiosControl {
    a.Set("addControls", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *RadiosControl) Name(value interface{}) *RadiosControl {
    a.Set("name", value)
    return a
}

/**
 */
func (a *RadiosControl) InitAutoFill(value interface{}) *RadiosControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *RadiosControl) DisabledOn(value interface{}) *RadiosControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 可用来通过 API 拉取 options。
 */
func (a *RadiosControl) Source(value interface{}) *RadiosControl {
    a.Set("source", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *RadiosControl) Inline(value interface{}) *RadiosControl {
    a.Set("inline", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *RadiosControl) OnEvent(value interface{}) *RadiosControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *RadiosControl) StaticOn(value interface{}) *RadiosControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (a *RadiosControl) ValuesNoWrap(value interface{}) *RadiosControl {
    a.Set("valuesNoWrap", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *RadiosControl) ExtractValue(value interface{}) *RadiosControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *RadiosControl) Value(value interface{}) *RadiosControl {
    a.Set("value", value)
    return a
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (a *RadiosControl) AutoFill(value interface{}) *RadiosControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *RadiosControl) DescriptionClassName(value interface{}) *RadiosControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 是否为必填
 */
func (a *RadiosControl) Required(value interface{}) *RadiosControl {
    a.Set("required", value)
    return a
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (a *RadiosControl) InitFetch(value interface{}) *RadiosControl {
    a.Set("initFetch", value)
    return a
}

/**
 * 编辑时调用的 API
 */
func (a *RadiosControl) EditApi(value interface{}) *RadiosControl {
    a.Set("editApi", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *RadiosControl) LabelWidth(value interface{}) *RadiosControl {
    a.Set("labelWidth", value)
    return a
}

/**
 */
func (a *RadiosControl) Desc(value interface{}) *RadiosControl {
    a.Set("desc", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *RadiosControl) HiddenOn(value interface{}) *RadiosControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *RadiosControl) StaticLabelClassName(value interface{}) *RadiosControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (a *RadiosControl) ResetValue(value interface{}) *RadiosControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 是否可以新增
 */
func (a *RadiosControl) Creatable(value interface{}) *RadiosControl {
    a.Set("creatable", value)
    return a
}

/**
 */
func (a *RadiosControl) Validations(value interface{}) *RadiosControl {
    a.Set("validations", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *RadiosControl) ValidateApi(value interface{}) *RadiosControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (a *RadiosControl) InitFetchOn(value interface{}) *RadiosControl {
    a.Set("initFetchOn", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *RadiosControl) ClearValueOnHidden(value interface{}) *RadiosControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 是否禁用
 */
func (a *RadiosControl) Disabled(value interface{}) *RadiosControl {
    a.Set("disabled", value)
    return a
}

/**
 */
func (a *RadiosControl) TestIdBuilder(value interface{}) *RadiosControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 是否可清除。
 */
func (a *RadiosControl) Clearable(value interface{}) *RadiosControl {
    a.Set("clearable", value)
    return a
}

/**
 * 描述标题
 */
func (a *RadiosControl) Label(value interface{}) *RadiosControl {
    a.Set("label", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *RadiosControl) Description(value interface{}) *RadiosControl {
    a.Set("description", value)
    return a
}

/**
 * 是否只读
 */
func (a *RadiosControl) ReadOnly(value interface{}) *RadiosControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 每行显示多少个
 */
func (a *RadiosControl) ColumnsCount(value interface{}) *RadiosControl {
    a.Set("columnsCount", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *RadiosControl) Width(value interface{}) *RadiosControl {
    a.Set("width", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *RadiosControl) ClassName(value interface{}) *RadiosControl {
    a.Set("className", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *RadiosControl) VisibleOn(value interface{}) *RadiosControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *RadiosControl) StaticInputClassName(value interface{}) *RadiosControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *RadiosControl) Remark(value interface{}) *RadiosControl {
    a.Set("remark", value)
    return a
}

/**
 * 选项修改的表单项
 */
func (a *RadiosControl) EditControls(value interface{}) *RadiosControl {
    a.Set("editControls", value)
    return a
}

/**
 * 控制编辑弹框设置项
 */
func (a *RadiosControl) EditDialog(value interface{}) *RadiosControl {
    a.Set("editDialog", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *RadiosControl) Mode(value interface{}) *RadiosControl {
    a.Set("mode", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *RadiosControl) Horizontal(value interface{}) *RadiosControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *RadiosControl) LabelRemark(value interface{}) *RadiosControl {
    a.Set("labelRemark", value)
    return a
}

/**
 */
func (a *RadiosControl) StaticSchema(value interface{}) *RadiosControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *RadiosControl) UseMobileUI(value interface{}) *RadiosControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否为多选模式
 */
func (a *RadiosControl) Multiple(value interface{}) *RadiosControl {
    a.Set("multiple", value)
    return a
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (a *RadiosControl) DeferApi(value interface{}) *RadiosControl {
    a.Set("deferApi", value)
    return a
}

/**
 * 添加时调用的接口
 */
func (a *RadiosControl) AddApi(value interface{}) *RadiosControl {
    a.Set("addApi", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *RadiosControl) Static(value interface{}) *RadiosControl {
    a.Set("static", value)
    return a
}

/**
 * 新增文字
 */
func (a *RadiosControl) CreateBtnLabel(value interface{}) *RadiosControl {
    a.Set("createBtnLabel", value)
    return a
}

/**
 * 配置 label className
 */
func (a *RadiosControl) LabelClassName(value interface{}) *RadiosControl {
    a.Set("labelClassName", value)
    return a
}

/**
 */
func (a *RadiosControl) Row(value interface{}) *RadiosControl {
    a.Set("row", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *RadiosControl) Hidden(value interface{}) *RadiosControl {
    a.Set("hidden", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *RadiosControl) ValidationErrors(value interface{}) *RadiosControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *RadiosControl) StaticPlaceholder(value interface{}) *RadiosControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 组件样式
 */
func (a *RadiosControl) Style(value interface{}) *RadiosControl {
    a.Set("style", value)
    return a
}

/**
 * 分割符
 */
func (a *RadiosControl) Delimiter(value interface{}) *RadiosControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 控制新增弹框设置项
 */
func (a *RadiosControl) AddDialog(value interface{}) *RadiosControl {
    a.Set("addDialog", value)
    return a
}

/**
 * 描述标题
 */
func (a *RadiosControl) LabelAlign(value interface{}) *RadiosControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 占位符
 */
func (a *RadiosControl) Placeholder(value interface{}) *RadiosControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *RadiosControl) EditorSetting(value interface{}) *RadiosControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *RadiosControl) JoinValues(value interface{}) *RadiosControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 懒加载字段
 */
func (a *RadiosControl) DeferField(value interface{}) *RadiosControl {
    a.Set("deferField", value)
    return a
}

/**
 * 是否可删除
 */
func (a *RadiosControl) Removable(value interface{}) *RadiosControl {
    a.Set("removable", value)
    return a
}

/**
 * 选项删除提示文字。
 */
func (a *RadiosControl) DeleteConfirmText(value interface{}) *RadiosControl {
    a.Set("deleteConfirmText", value)
    return a
}

/**
 * 选项集合
 */
func (a *RadiosControl) Options(value interface{}) *RadiosControl {
    a.Set("options", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *RadiosControl) SubmitOnChange(value interface{}) *RadiosControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 表单项类型
 */
func (a *RadiosControl) Type(value interface{}) *RadiosControl {
    a.Set("type", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *RadiosControl) ValidateOnChange(value interface{}) *RadiosControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置 input className
 */
func (a *RadiosControl) InputClassName(value interface{}) *RadiosControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *RadiosControl) Id(value interface{}) *RadiosControl {
    a.Set("id", value)
    return a
}

/**
 * 默认选择选项第一个值。
 */
func (a *RadiosControl) SelectFirst(value interface{}) *RadiosControl {
    a.Set("selectFirst", value)
    return a
}

/**
 * 是否可以编辑
 */
func (a *RadiosControl) Editable(value interface{}) *RadiosControl {
    a.Set("editable", value)
    return a
}

/**
 * 选项删除 API
 */
func (a *RadiosControl) DeleteApi(value interface{}) *RadiosControl {
    a.Set("deleteApi", value)
    return a
}

/**
 * 只读条件
 */
func (a *RadiosControl) ReadOnlyOn(value interface{}) *RadiosControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *RadiosControl) Size(value interface{}) *RadiosControl {
    a.Set("size", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *RadiosControl) ExtraName(value interface{}) *RadiosControl {
    a.Set("extraName", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *RadiosControl) Hint(value interface{}) *RadiosControl {
    a.Set("hint", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *RadiosControl) StaticClassName(value interface{}) *RadiosControl {
    a.Set("staticClassName", value)
    return a
}
