package renderers


/**
 * Location 选点组件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/location
 *

*/
type LocationControl struct {
	*BaseRenderer
}

func NewLocationControl() *LocationControl {
    a := &LocationControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "location-picker")
    return a
}
/**
 * 是否静态展示表达式
 */
func (a *LocationControl) StaticOn(value string) *LocationControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *LocationControl) DisabledOn(value string) *LocationControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否显示
 */
func (a *LocationControl) Visible(value string) *LocationControl {
    a.Set("visible", value)
    return a
}

/**
 * 选择地图类型
 * 可选值: baidu | gaode | tenxun
 */
func (a *LocationControl) Vendor(value string) *LocationControl {
    a.Set("vendor", value)
    return a
}

/**
 * 有的地图需要设置 ak 信息
 */
func (a *LocationControl) Ak(value string) *LocationControl {
    a.Set("ak", value)
    return a
}

/**
 * 是否自动选中当前地理位置
 */
func (a *LocationControl) AutoSelectCurrentLoc(value string) *LocationControl {
    a.Set("autoSelectCurrentLoc", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *LocationControl) VisibleOn(value string) *LocationControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 是否禁用
 */
func (a *LocationControl) Disabled(value string) *LocationControl {
    a.Set("disabled", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *LocationControl) ClearValueOnHidden(value string) *LocationControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *LocationControl) EditorSetting(value string) *LocationControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *LocationControl) Remark(value string) *LocationControl {
    a.Set("remark", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *LocationControl) ValidateOnChange(value string) *LocationControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *LocationControl) ClassName(value string) *LocationControl {
    a.Set("className", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *LocationControl) StaticLabelClassName(value string) *LocationControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *LocationControl) StaticInputClassName(value string) *LocationControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *LocationControl) ValidationErrors(value string) *LocationControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *LocationControl) Hidden(value string) *LocationControl {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *LocationControl) StaticPlaceholder(value string) *LocationControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 只读条件
 */
func (a *LocationControl) ReadOnlyOn(value string) *LocationControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 开启只读模式后的占位提示，默认为“点击获取位置信息” 备注：区分下现有的placeholder（“请选择位置”）
 */
func (a *LocationControl) GetLocationPlaceholder(value string) *LocationControl {
    a.Set("getLocationPlaceholder", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *LocationControl) HiddenOn(value string) *LocationControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 占位符
 */
func (a *LocationControl) Placeholder(value string) *LocationControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *LocationControl) Value(value string) *LocationControl {
    a.Set("value", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *LocationControl) Static(value string) *LocationControl {
    a.Set("static", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *LocationControl) ExtraName(value string) *LocationControl {
    a.Set("extraName", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *LocationControl) Inline(value string) *LocationControl {
    a.Set("inline", value)
    return a
}

/**
 * 组件样式
 */
func (a *LocationControl) Style(value string) *LocationControl {
    a.Set("style", value)
    return a
}

/**
 * 配置 label className
 */
func (a *LocationControl) LabelClassName(value string) *LocationControl {
    a.Set("labelClassName", value)
    return a
}

/**
 */
func (a *LocationControl) Desc(value string) *LocationControl {
    a.Set("desc", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *LocationControl) Mode(value string) *LocationControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否为必填
 */
func (a *LocationControl) Required(value string) *LocationControl {
    a.Set("required", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *LocationControl) UseMobileUI(value string) *LocationControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否只读
 */
func (a *LocationControl) ReadOnly(value string) *LocationControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 配置 input className
 */
func (a *LocationControl) InputClassName(value string) *LocationControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *LocationControl) LabelWidth(value string) *LocationControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *LocationControl) Description(value string) *LocationControl {
    a.Set("description", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *LocationControl) Horizontal(value string) *LocationControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否限制只能选中当前地理位置 备注：可用于充当定位组件，只允许选择当前位置
 */
func (a *LocationControl) OnlySelectCurrentLoc(value string) *LocationControl {
    a.Set("onlySelectCurrentLoc", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *LocationControl) StaticClassName(value string) *LocationControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 描述标题
 */
func (a *LocationControl) Label(value string) *LocationControl {
    a.Set("label", value)
    return a
}

/**
 * 描述标题
 */
func (a *LocationControl) LabelAlign(value string) *LocationControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *LocationControl) Id(value string) *LocationControl {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *LocationControl) OnEvent(value string) *LocationControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *LocationControl) Hint(value string) *LocationControl {
    a.Set("hint", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *LocationControl) Width(value string) *LocationControl {
    a.Set("width", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *LocationControl) SubmitOnChange(value string) *LocationControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *LocationControl) DescriptionClassName(value string) *LocationControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *LocationControl) ValidateApi(value string) *LocationControl {
    a.Set("validateApi", value)
    return a
}

/**
 */
func (a *LocationControl) StaticSchema(value string) *LocationControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *LocationControl) LabelRemark(value string) *LocationControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *LocationControl) Size(value string) *LocationControl {
    a.Set("size", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *LocationControl) Name(value string) *LocationControl {
    a.Set("name", value)
    return a
}

/**
 */
func (a *LocationControl) Validations(value string) *LocationControl {
    a.Set("validations", value)
    return a
}

/**
 * 表单项类型
 */
func (a *LocationControl) Type(value string) *LocationControl {
    a.Set("type", value)
    return a
}