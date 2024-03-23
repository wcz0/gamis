// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * 按钮组控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/button-group
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class ButtonGroupControl extends BaseRenderer
// {
//     public function __construct()
//     {
//         $this->set('type', 'button-group-select');


//     }

//     /**
//      * 添加时调用的接口
//      */
//     public function addApi($value = '')
//     {
//         return $this->set('addApi', $value);
//     }

//     /**
//      * 新增时的表单项。
//      */
//     public function addControls($value = '')
//     {
//         return $this->set('addControls', $value);
//     }

//     /**
//      * 控制新增弹框设置项
//      */
//     public function addDialog($value = '')
//     {
//         return $this->set('addDialog', $value);
//     }

//     /**
//      * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
//      */
//     public function autoFill($value = '')
//     {
//         return $this->set('autoFill', $value);
//     }

//     /**
//      *
//      */
//     public function btnActiveClassName($value = '')
//     {
//         return $this->set('btnActiveClassName', $value);
//     }

//     /**
//      * 按钮选中的样式级别
//      */
//     public function btnActiveLevel($value = '')
//     {
//         return $this->set('btnActiveLevel', $value);
//     }

//     /**
//      *
//      */
//     public function btnClassName($value = '')
//     {
//         return $this->set('btnClassName', $value);
//     }

//     /**
//      * 按钮样式级别
//      */
//     public function btnLevel($value = '')
//     {
//         return $this->set('btnLevel', $value);
//     }

//     /**
//      * 按钮集合
//      */
//     public function buttons($value = '')
//     {
//         return $this->set('buttons', $value);
//     }

//     /**
//      * 容器 css 类名
//      */
//     public function className($value = '')
//     {
//         return $this->set('className', $value);
//     }

//     /**
//      * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
//      */
//     public function clearValueOnHidden($value = true)
//     {
//         return $this->set('clearValueOnHidden', $value);
//     }

//     /**
//      * 是否可清除。
//      */
//     public function clearable($value = true)
//     {
//         return $this->set('clearable', $value);
//     }

//     /**
//      * 是否可以新增
//      */
//     public function creatable($value = true)
//     {
//         return $this->set('creatable', $value);
//     }

//     /**
//      * 新增文字
//      */
//     public function createBtnLabel($value = '')
//     {
//         return $this->set('createBtnLabel', $value);
//     }

//     /**
//      * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
//      */
//     public function deferApi($value = '')
//     {
//         return $this->set('deferApi', $value);
//     }

//     /**
//      * 懒加载字段
//      */
//     public function deferField($value = '')
//     {
//         return $this->set('deferField', $value);
//     }

//     /**
//      * 选项删除 API
//      */
//     public function deleteApi($value = '')
//     {
//         return $this->set('deleteApi', $value);
//     }

//     /**
//      * 选项删除提示文字。
//      */
//     public function deleteConfirmText($value = '')
//     {
//         return $this->set('deleteConfirmText', $value);
//     }

//     /**
//      * 分割符
//      */
//     public function delimiter($value = '')
//     {
//         return $this->set('delimiter', $value);
//     }

//     /**
//      *
//      */
//     public function desc($value = '')
//     {
//         return $this->set('desc', $value);
//     }

//     /**
//      * 描述内容，支持 Html 片段。
//      */
//     public function description($value = '')
//     {
//         return $this->set('description', $value);
//     }

//     /**
//      * 配置描述上的 className
//      */
//     public function descriptionClassName($value = '')
//     {
//         return $this->set('descriptionClassName', $value);
//     }

//     /**
//      * 是否为禁用状态。
//      */
//     public function disabled($value = true)
//     {
//         return $this->set('disabled', $value);
//     }

//     /**
//      * 通过 JS 表达式来配置当前表单项的禁用状态。
//      */
//     public function disabledOn($value = '')
//     {
//         return $this->set('disabledOn', $value);
//     }

//     /**
//      * 编辑时调用的 API
//      */
//     public function editApi($value = '')
//     {
//         return $this->set('editApi', $value);
//     }

//     /**
//      * 选项修改的表单项
//      */
//     public function editControls($value = '')
//     {
//         return $this->set('editControls', $value);
//     }

//     /**
//      * 控制编辑弹框设置项
//      */
//     public function editDialog($value = '')
//     {
//         return $this->set('editDialog', $value);
//     }

//     /**
//      * 是否可以编辑
//      */
//     public function editable($value = true)
//     {
//         return $this->set('editable', $value);
//     }

//     /**
//      * 编辑器配置，运行时可以忽略
//      */
//     public function editorSetting($value = '')
//     {
//         return $this->set('editorSetting', $value);
//     }

//     /**
//      * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
//      */
//     public function extraName($value = '')
//     {
//         return $this->set('extraName', $value);
//     }

//     /**
//      * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
//      */
//     public function extractValue($value = true)
//     {
//         return $this->set('extractValue', $value);
//     }

//     /**
//      * 是否隐藏
//      */
//     public function hidden($value = true)
//     {
//         return $this->set('hidden', $value);
//     }

//     /**
//      * 是否隐藏表达式
//      */
//     public function hiddenOn($value = '')
//     {
//         return $this->set('hiddenOn', $value);
//     }

//     /**
//      * 输入提示，聚焦的时候显示
//      */
//     public function hint($value = '')
//     {
//         return $this->set('hint', $value);
//     }

//     /**
//      * 当配置为水平布局的时候，用来配置具体的左右分配。
//      */
//     public function horizontal($value = '')
//     {
//         return $this->set('horizontal', $value);
//     }

//     /**
//      * 组件唯一 id，主要用于日志采集
//      */
//     public function id($value = '')
//     {
//         return $this->set('id', $value);
//     }

//     /**
//      *
//      */
//     public function initAutoFill($value = '')
//     {
//         return $this->set('initAutoFill', $value);
//     }

//     /**
//      * 配置 source 接口初始拉不拉取。
//      */
//     public function initFetch($value = true)
//     {
//         return $this->set('initFetch', $value);
//     }

//     /**
//      * 用表达式来配置 source 接口初始要不要拉取
//      */
//     public function initFetchOn($value = '')
//     {
//         return $this->set('initFetchOn', $value);
//     }

//     /**
//      * 表单 control 是否为 inline 模式。
//      */
//     public function inline($value = true)
//     {
//         return $this->set('inline', $value);
//     }

//     /**
//      * 配置 input className
//      */
//     public function inputClassName($value = '')
//     {
//         return $this->set('inputClassName', $value);
//     }

//     /**
//      * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
//      */
//     public function joinValues($value = true)
//     {
//         return $this->set('joinValues', $value);
//     }

//     /**
//      * 描述标题
//      */
//     public function label($value = '')
//     {
//         return $this->set('label', $value);
//     }

//     /**
//      * 描述标题
//      */
//     public function labelAlign($value = '')
//     {
//         return $this->set('labelAlign', $value);
//     }

//     /**
//      * 配置 label className
//      */
//     public function labelClassName($value = '')
//     {
//         return $this->set('labelClassName', $value);
//     }

//     /**
//      * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
//      */
//     public function labelRemark($value = '')
//     {
//         return $this->set('labelRemark', $value);
//     }

//     /**
//      * label自定义宽度，默认单位为px
//      */
//     public function labelWidth($value = '')
//     {
//         return $this->set('labelWidth', $value);
//     }

//     /**
//      * 配置当前表单项展示模式 可选值: normal | inline | horizontal
//      */
//     public function mode($value = '')
//     {
//         return $this->set('mode', $value);
//     }

//     /**
//      * 是否为多选模式
//      */
//     public function multiple($value = true)
//     {
//         return $this->set('multiple', $value);
//     }

//     /**
//      * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
//      */
//     public function name($value = '')
//     {
//         return $this->set('name', $value);
//     }

//     /**
//      * 事件动作配置
//      */
//     public function onEvent($value = '')
//     {
//         return $this->set('onEvent', $value);
//     }

//     /**
//      * 选项集合
//      */
//     public function options($value = '')
//     {
//         return $this->set('options', $value);
//     }

//     /**
//      * 占位符
//      */
//     public function placeholder($value = '')
//     {
//         return $this->set('placeholder', $value);
//     }

//     /**
//      * 是否只读
//      */
//     public function readOnly($value = true)
//     {
//         return $this->set('readOnly', $value);
//     }

//     /**
//      * 只读条件
//      */
//     public function readOnlyOn($value = '')
//     {
//         return $this->set('readOnlyOn', $value);
//     }

//     /**
//      * 显示一个小图标, 鼠标放上去的时候显示提示内容
//      */
//     public function remark($value = '')
//     {
//         return $this->set('remark', $value);
//     }

//     /**
//      * 是否可删除
//      */
//     public function removable($value = true)
//     {
//         return $this->set('removable', $value);
//     }

//     /**
//      * 是否为必填
//      */
//     public function required($value = true)
//     {
//         return $this->set('required', $value);
//     }

//     /**
//      * 点清除按钮时，将表单项设置成当前配置的值。
//      */
//     public function resetValue($value = '')
//     {
//         return $this->set('resetValue', $value);
//     }

//     /**
//      * 默认选择选项第一个值。
//      */
//     public function selectFirst($value = true)
//     {
//         return $this->set('selectFirst', $value);
//     }

//     /**
//      * 按钮大小 可选值: xs | sm | md | lg
//      */
//     public function size($value = '')
//     {
//         return $this->set('size', $value);
//     }

//     /**
//      * 可用来通过 API 拉取 options。
//      */
//     public function source($value = '')
//     {
//         return $this->set('source', $value);
//     }

//     /**
//      * 是否静态展示
//      */
//     public function static($value = true)
//     {
//         return $this->set('static', $value);
//     }

//     /**
//      * 静态展示表单项类名
//      */
//     public function staticClassName($value = '')
//     {
//         return $this->set('staticClassName', $value);
//     }

//     /**
//      * 静态展示表单项Value类名
//      */
//     public function staticInputClassName($value = '')
//     {
//         return $this->set('staticInputClassName', $value);
//     }

//     /**
//      * 静态展示表单项Label类名
//      */
//     public function staticLabelClassName($value = '')
//     {
//         return $this->set('staticLabelClassName', $value);
//     }

//     /**
//      * 是否静态展示表达式
//      */
//     public function staticOn($value = '')
//     {
//         return $this->set('staticOn', $value);
//     }

//     /**
//      * 静态展示空值占位
//      */
//     public function staticPlaceholder($value = '')
//     {
//         return $this->set('staticPlaceholder', $value);
//     }

//     /**
//      *
//      */
//     public function staticSchema($value = '')
//     {
//         return $this->set('staticSchema', $value);
//     }

//     /**
//      * 组件样式
//      */
//     public function style($value = '')
//     {
//         return $this->set('style', $value);
//     }

//     /**
//      * 当修改完的时候是否提交表单。
//      */
//     public function submitOnChange($value = true)
//     {
//         return $this->set('submitOnChange', $value);
//     }

//     /**
//      * 平铺展示？
//      */
//     public function tiled($value = true)
//     {
//         return $this->set('tiled', $value);
//     }

//     /**
//      *
//      */
//     public function type($value = 'button-group-select')
//     {
//         return $this->set('type', $value);
//     }

//     /**
//      * 可以组件级别用来关闭移动端样式
//      */
//     public function useMobileUI($value = true)
//     {
//         return $this->set('useMobileUI', $value);
//     }

//     /**
//      * 远端校验表单项接口
//      */
//     public function validateApi($value = '')
//     {
//         return $this->set('validateApi', $value);
//     }

//     /**
//      * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
//      */
//     public function validateOnChange($value = true)
//     {
//         return $this->set('validateOnChange', $value);
//     }

//     /**
//      * 验证失败的提示信息
//      */
//     public function validationErrors($value = '')
//     {
//         return $this->set('validationErrors', $value);
//     }

//     /**
//      *
//      */
//     public function validations($value = '')
//     {
//         return $this->set('validations', $value);
//     }

//     /**
//      * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
//      */
//     public function value($value = '')
//     {
//         return $this->set('value', $value);
//     }

//     /**
//      * 多选模式，值太多时是否避免折行
//      */
//     public function valuesNoWrap($value = true)
//     {
//         return $this->set('valuesNoWrap', $value);
//     }

//     /**
//      * 垂直展示？
//      */
//     public function vertical($value = true)
//     {
//         return $this->set('vertical', $value);
//     }

//     /**
//      * 是否显示
//      */
//     public function visible($value = true)
//     {
//         return $this->set('visible', $value);
//     }

//     /**
//      * 通过 JS 表达式来配置当前表单项是否显示
//      */
//     public function visibleOn($value = '')
//     {
//         return $this->set('visibleOn', $value);
//     }

//     /**
//      * 在Table中调整宽度
//      */
//     public function width($value = '')
//     {
//         return $this->set('width', $value);
//     }


// }

package renderers

type ButtonGroupControl struct {
	*BaseRenderer
}

func NewButtonGroupControl() *ButtonGroupControl {
	b := &ButtonGroupControl{
		BaseRenderer: NewBaseRenderer(),
	}
	b.Set("type", "button-group-select")
	return b
}

/**
 * 添加时调用的接口
 */
func (b *ButtonGroupControl) AddApi(value string) *ButtonGroupControl {
	b.Set("addApi", value)
	return b
}

/**
 * 新增时的表单项。
 */
func (b *ButtonGroupControl) AddControls(value string) *ButtonGroupControl {
	b.Set("addControls", value)
	return b
}

/**
 * 控制新增弹框设置项
 */
func (b *ButtonGroupControl) AddDialog(value string) *ButtonGroupControl {
	b.Set("addDialog", value)
	return b
}

/**
 * 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
 */
func (b *ButtonGroupControl) AutoFill(value string) *ButtonGroupControl {
	b.Set("autoFill", value)
	return b
}

func (b *ButtonGroupControl) BtnActiveClassName(value string) *ButtonGroupControl {
	b.Set("btnActiveClassName", value)
	return b
}

/**
 * 按钮选中的样式级别
 */
func (b *ButtonGroupControl) BtnActiveLevel(value string) *ButtonGroupControl {
	b.Set("btnActiveLevel", value)
	return b
}

func (b *ButtonGroupControl) BtnClassName(value string) *ButtonGroupControl {
	b.Set("btnClassName", value)
	return b
}

/**
 * 按钮样式级别
 */
func (b *ButtonGroupControl) BtnLevel(value string) *ButtonGroupControl {
	b.Set("btnLevel", value)
	return b
}

/**
 * 按钮集合
 */
func (b *ButtonGroupControl) Buttons(value string) *ButtonGroupControl {
	b.Set("buttons", value)
	return b
}

/**
 * 容器 css 类名
 */
func (b *ButtonGroupControl) ClassName(value string) *ButtonGroupControl {
	b.Set("className", value)
	return b
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (b *ButtonGroupControl) ClearValueOnHidden(value bool) *ButtonGroupControl {
	b.Set("clearValueOnHidden", value)
	return b
}

/**
 * 是否可清除。
 */
func (b *ButtonGroupControl) Clearable(value bool) *ButtonGroupControl {
	b.Set("clearable", value)
	return b
}

/**
 * 是否可以新增
 */
func (b *ButtonGroupControl) Creatable(value bool) *ButtonGroupControl {
	b.Set("creatable", value)
	return b
}

/**
 * 新增文字
 */
func (b *ButtonGroupControl) CreateBtnLabel(value string) *ButtonGroupControl {
	b.Set("createBtnLabel", value)
	return b
}

/**
 * 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
 */
func (b *ButtonGroupControl) DeferApi(value string) *ButtonGroupControl {
	b.Set("deferApi", value)
	return b
}

/**
 * 懒加载字段
 */
func (b *ButtonGroupControl) DeferField(value string) *ButtonGroupControl {
	b.Set("deferField", value)
	return b
}

/**
 * 选项删除 API
 */
func (b *ButtonGroupControl) DeleteApi(value string) *ButtonGroupControl {
	b.Set("deleteApi", value)
	return b
}

/**
 * 选项删除提示文字。
 */
func (b *ButtonGroupControl) DeleteConfirmText(value string) *ButtonGroupControl {
	b.Set("deleteConfirmText", value)
	return b
}

/**
 * 分割符
 */
func (b *ButtonGroupControl) Delimiter(value string) *ButtonGroupControl {
	b.Set("delimiter", value)
	return b
}

func (b *ButtonGroupControl) Desc(value string) *ButtonGroupControl {
	b.Set("desc", value)
	return b
}

/**
 * 描述内容，支持 Html 片段。
 */
func (b *ButtonGroupControl) Description(value string) *ButtonGroupControl {
	b.Set("description", value)
	return b
}

/**
 * 配置描述上的 className
 */
func (b *ButtonGroupControl) DescriptionClassName(value string) *ButtonGroupControl {
	b.Set("descriptionClassName", value)
	return b
}

/**
 * 是否为禁用状态。
 */
func (b *ButtonGroupControl) Disabled(value bool) *ButtonGroupControl {
	b.Set("disabled", value)
	return b
}

/**
 * 通过 JS 表达式来配置当前表单项的禁用状态。
 */
func (b *ButtonGroupControl) DisabledOn(value string) *ButtonGroupControl {
	b.Set("disabledOn", value)
	return b
}

/**
 * 编辑时调用的 API
 */
func (b *ButtonGroupControl) EditApi(value string) *ButtonGroupControl {
	b.Set("editApi", value)
	return b
}

/**
 * 选项修改的表单项
 */
func (b *ButtonGroupControl) EditControls(value string) *ButtonGroupControl {
	b.Set("editControls", value)
	return b
}

/**
 * 控制编辑弹框设置项
 */
func (b *ButtonGroupControl) EditDialog(value string) *ButtonGroupControl {
	b.Set("editDialog", value)
	return b
}

/**
 * 是否可以编辑
 */
func (b *ButtonGroupControl) Editable(value bool) *ButtonGroupControl {
	b.Set("editable", value)
	return b
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (b *ButtonGroupControl) EditorSetting(value string) *ButtonGroupControl {
	b.Set("editorSetting", value)
	return b
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (b *ButtonGroupControl) ExtraName(value string) *ButtonGroupControl {
	b.Set("extraName", value)
	return b
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (b *ButtonGroupControl) ExtractValue(value bool) *ButtonGroupControl {
	b.Set("extractValue", value)
	return b
}

/**
 * 是否隐藏
 */
func (b *ButtonGroupControl) Hidden(value bool) *ButtonGroupControl {
	b.Set("hidden", value)
	return b
}

/**
 * 是否隐藏表达式
 */
func (b *ButtonGroupControl) HiddenOn(value string) *ButtonGroupControl {
	b.Set("hiddenOn", value)
	return b
}

/**
 * 输入提示，聚焦的时候显示
 */
func (b *ButtonGroupControl) Hint(value string) *ButtonGroupControl {
	b.Set("hint", value)
	return b
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (b *ButtonGroupControl) Horizontal(value string) *ButtonGroupControl {
	b.Set("horizontal", value)
	return b
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (b *ButtonGroupControl) Id(value string) *ButtonGroupControl {
	b.Set("id", value)
	return b
}

func (b *ButtonGroupControl) InitAutoFill(value string) *ButtonGroupControl {
	b.Set("initAutoFill", value)
	return b
}

/**
 * 配置 source 接口初始拉不拉取。
 */
func (b *ButtonGroupControl) InitFetch(value bool) *ButtonGroupControl {
	b.Set("initFetch", value)
	return b
}

/**
 * 用表达式来配置 source 接口初始要不要拉取
 */
func (b *ButtonGroupControl) InitFetchOn(value string) *ButtonGroupControl {
	b.Set("initFetchOn", value)
	return b
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (b *ButtonGroupControl) Inline(value bool) *ButtonGroupControl {
	b.Set("inline", value)
	return b
}

/**
 * 配置 input className
 */
func (b *ButtonGroupControl) InputClassName(value string) *ButtonGroupControl {
	b.Set("inputClassName", value)
	return b
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (b *ButtonGroupControl) JoinValues(value bool) *ButtonGroupControl {
	b.Set("joinValues", value)
	return b
}

/**
 * 描述标题
 */
func (b *ButtonGroupControl) Label(value string) *ButtonGroupControl {
	b.Set("label", value)
	return b
}

/**
 * 描述标题
 */
func (b *ButtonGroupControl) LabelAlign(value string) *ButtonGroupControl {
	b.Set("labelAlign", value)
	return b
}

/**
 * 配置 label className
 */
func (b *ButtonGroupControl) LabelClassName(value string) *ButtonGroupControl {
	b.Set("labelClassName", value)
	return b
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (b *ButtonGroupControl) LabelRemark(value string) *ButtonGroupControl {
	b.Set("labelRemark", value)
	return b
}

/**
 * label自定义宽度，默认单位为px
 */
func (b *ButtonGroupControl) LabelWidth(value string) *ButtonGroupControl {
	b.Set("labelWidth", value)
	return b
}

/**
 * 配置当前表单项展示模式 可选值: normal | inline | horizontal
 */
func (b *ButtonGroupControl) Mode(value string) *ButtonGroupControl {
	b.Set("mode", value)
	return b
}

/**
 * 是否为多选模式
 */
func (b *ButtonGroupControl) Multiple(value bool) *ButtonGroupControl {
	b.Set("multiple", value)
	return b
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (b *ButtonGroupControl) Name(value string) *ButtonGroupControl {
	b.Set("name", value)
	return b
}

/**
 * 事件动作配置
 */
func (b *ButtonGroupControl) OnEvent(value string) *ButtonGroupControl {
	b.Set("onEvent", value)
	return b
}

/**
 * 选项集合
 */
func (b *ButtonGroupControl) Options(value string) *ButtonGroupControl {
	b.Set("options", value)
	return b
}

/**
 * 占位符
 */
func (b *ButtonGroupControl) Placeholder(value string) *ButtonGroupControl {
	b.Set("placeholder", value)
	return b
}

/**
 * 是否只读
 */
func (b *ButtonGroupControl) ReadOnly(value bool) *ButtonGroupControl {
	b.Set("readOnly", value)
	return b
}

/**
 * 只读条件
 */
func (b *ButtonGroupControl) ReadOnlyOn(value string) *ButtonGroupControl {
	b.Set("readOnlyOn", value)
	return b
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (b *ButtonGroupControl) Remark(value string) *ButtonGroupControl {
	b.Set("remark", value)
	return b
}

/**
 * 是否可删除
 */
func (b *ButtonGroupControl) Removable(value bool) *ButtonGroupControl {
	b.Set("removable", value)
	return b
}

/**
 * 是否为必填
 */
func (b *ButtonGroupControl) Required(value bool) *ButtonGroupControl {
	b.Set("required", value)
	return b
}

/**
 * 点清除按钮时，将表单项设置成当前配置的值。
 */
func (b *ButtonGroupControl) ResetValue(value string) *ButtonGroupControl {
	b.Set("resetValue", value)
	return b
}

/**
 * 默认选择选项第一个值。
 */
func (b *ButtonGroupControl) SelectFirst(value bool) *ButtonGroupControl {
	b.Set("selectFirst", value)
	return b
}

/**
 * 按钮大小 可选值: xs | sm | md | lg
 */
func (b *ButtonGroupControl) Size(value string) *ButtonGroupControl {
	b.Set("size", value)
	return b
}

/**
 * 可用来通过 API 拉取 options。
 */
func (b *ButtonGroupControl) Source(value string) *ButtonGroupControl {
	b.Set("source", value)
	return b
}

/**
 * 是否静态展示
 */
func (b *ButtonGroupControl) Static(value bool) *ButtonGroupControl {
	b.Set("static", value)
	return b
}

/**
 * 静态展示表单项类名
 */
func (b *ButtonGroupControl) StaticClassName(value string) *ButtonGroupControl {
	b.Set("staticClassName", value)
	return b
}

/**
 * 静态展示表单项Value类名
 */
func (b *ButtonGroupControl) StaticInputClassName(value string) *ButtonGroupControl {
	b.Set("staticInputClassName", value)
	return b
}

/**
 * 静态展示表单项Label类名
 */
func (b *ButtonGroupControl) StaticLabelClassName(value string) *ButtonGroupControl {
	b.Set("staticLabelClassName", value)
	return b
}

/**
 * 是否静态展示表达式
 */
func (b *ButtonGroupControl) StaticOn(value string) *ButtonGroupControl {
	b.Set("staticOn", value)
	return b
}

/**
 * 静态展示空值占位
 */
func (b *ButtonGroupControl) StaticPlaceholder(value string) *ButtonGroupControl {
	b.Set("staticPlaceholder", value)
	return b
}

func (b *ButtonGroupControl) StaticSchema(value string) *ButtonGroupControl {
	b.Set("staticSchema", value)
	return b
}

/**
 * 组件样式
 */
func (b *ButtonGroupControl) Style(value string) *ButtonGroupControl {
	b.Set("style", value)
	return b
}

/**
 * 当修改完的时候是否提交表单。
 */
func (b *ButtonGroupControl) SubmitOnChange(value bool) *ButtonGroupControl {
	b.Set("submitOnChange", value)
	return b
}

/**
 * 平铺展示？
 */
func (b *ButtonGroupControl) Tiled(value bool) *ButtonGroupControl {
	b.Set("tiled", value)
	return b
}

func (b *ButtonGroupControl) Type(value string) *ButtonGroupControl {
	b.Set("type", value)
	return b
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (b *ButtonGroupControl) UseMobileUI(value bool) *ButtonGroupControl {
	b.Set("useMobileUI", value)
	return b
}

/**
 * 远端校验表单项接口
 */
func (b *ButtonGroupControl) ValidateApi(value string) *ButtonGroupControl {
	b.Set("validateApi", value)
	return b
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (b *ButtonGroupControl) ValidateOnChange(value bool) *ButtonGroupControl {
	b.Set("validateOnChange", value)
	return b
}

/**
 * 验证失败的提示信息
 */
func (b *ButtonGroupControl) ValidationErrors(value string) *ButtonGroupControl {
	b.Set("validationErrors", value)
	return b
}

func (b *ButtonGroupControl) Validations(value string) *ButtonGroupControl {
	b.Set("validations", value)
	return b
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (b *ButtonGroupControl) Value(value string) *ButtonGroupControl {
	b.Set("value", value)
	return b
}

/**
 * 多选模式，值太多时是否避免折行
 */
func (b *ButtonGroupControl) ValuesNoWrap(value bool) *ButtonGroupControl {
	b.Set("valuesNoWrap", value)
	return b
}

/**
 * 垂直展示？
 */
func (b *ButtonGroupControl) Vertical(value bool) *ButtonGroupControl {
	b.Set("vertical", value)
	return b
}

/**
 * 是否显示
 */
func (b *ButtonGroupControl) Visible(value bool) *ButtonGroupControl {
	b.Set("visible", value)
	return b
}

/**
 * 通过 JS 表达式来配置当前表单项是否显示
 */
func (b *ButtonGroupControl) VisibleOn(value string) *ButtonGroupControl {
	b.Set("visibleOn", value)
	return b
}

/**
 * 在Table中调整宽度
 */
func (b *ButtonGroupControl) Width(value string) *ButtonGroupControl {
	b.Set("width", value)
	return b
}
