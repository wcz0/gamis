package renderers


/**
 * Image 图片上传控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/image

 * @author wcz0
 * @version 6.2.2
 */
type ImageControl struct {
	*BaseRenderer
}

func NewImageControl() *ImageControl {
    a := &ImageControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-image")
    return a
}


func (a *ImageControl) Set(name string, value interface{}) *ImageControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否隐藏表达式
 */
func (a *ImageControl) Hiddenon(value interface{}) *ImageControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *ImageControl) Inline(value interface{}) *ImageControl {
    a.Set("inline", value)
    return a
}

/**
 */
func (a *ImageControl) Validations(value interface{}) *ImageControl {
    a.Set("validations", value)
    return a
}

/**
 * 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
 */
func (a *ImageControl) Joinvalues(value interface{}) *ImageControl {
    a.Set("joinValues", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *ImageControl) Labelremark(value interface{}) *ImageControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 图片上传完毕是否进入裁剪模式
 */
func (a *ImageControl) Dropcrop(value interface{}) *ImageControl {
    a.Set("dropCrop", value)
    return a
}

/**
 * 固定尺寸的 CSS类名
 */
func (a *ImageControl) Fixedsizeclassname(value interface{}) *ImageControl {
    a.Set("fixedSizeClassName", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *ImageControl) Staticplaceholder(value interface{}) *ImageControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 描述标题
 */
func (a *ImageControl) Labelalign(value interface{}) *ImageControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 是否可拖拽排序
 */
func (a *ImageControl) Draggable(value interface{}) *ImageControl {
    a.Set("draggable", value)
    return a
}

/**
 * 可拖拽排序的提示信息。
 */
func (a *ImageControl) Draggabletip(value interface{}) *ImageControl {
    a.Set("draggableTip", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *ImageControl) Labelwidth(value interface{}) *ImageControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 选择图片按钮的 CSS 类名
 */
func (a *ImageControl) Btnclassname(value interface{}) *ImageControl {
    a.Set("btnClassName", value)
    return a
}

/**
 * 裁剪后的质量
 */
func (a *ImageControl) Cropquality(value interface{}) *ImageControl {
    a.Set("cropQuality", value)
    return a
}

/**
 * 默认展示图片的链接
 */
func (a *ImageControl) Src(value interface{}) *ImageControl {
    a.Set("src", value)
    return a
}

/**
 * 配置接收的图片类型建议直接填写文件后缀 如：.txt,.csv多个类型用逗号隔开。
 */
func (a *ImageControl) Accept(value interface{}) *ImageControl {
    a.Set("accept", value)
    return a
}

/**
 * 是否为多选
 */
func (a *ImageControl) Multiple(value interface{}) *ImageControl {
    a.Set("multiple", value)
    return a
}

/**
 * 只读条件
 */
func (a *ImageControl) Readonlyon(value interface{}) *ImageControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 */
func (a *ImageControl) Desc(value interface{}) *ImageControl {
    a.Set("desc", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *ImageControl) Horizontal(value interface{}) *ImageControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 是否为必填
 */
func (a *ImageControl) Required(value interface{}) *ImageControl {
    a.Set("required", value)
    return a
}

/**
 */
func (a *ImageControl) Compressoptions(value interface{}) *ImageControl {
    a.Set("compressOptions", value)
    return a
}

/**
 * 初始化时是否打开裁剪模式
 */
func (a *ImageControl) Initcrop(value interface{}) *ImageControl {
    a.Set("initCrop", value)
    return a
}

/**
 * 默认占位图图片地址
 */
func (a *ImageControl) Frameimage(value interface{}) *ImageControl {
    a.Set("frameImage", value)
    return a
}

/**
 * 描述标题
 */
func (a *ImageControl) Label(value interface{}) *ImageControl {
    a.Set("label", value)
    return a
}

/**
 * 占位符
 */
func (a *ImageControl) Placeholder(value interface{}) *ImageControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 上传后把其他字段同步到表单内部。
 */
func (a *ImageControl) Autofill(value interface{}) *ImageControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 上传按钮的 CSS 类名
 */
func (a *ImageControl) Btnuploadclassname(value interface{}) *ImageControl {
    a.Set("btnUploadClassName", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *ImageControl) Submitonchange(value interface{}) *ImageControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 初始化时是否把其他字段同步到表单内部。
 */
func (a *ImageControl) Initautofill(value interface{}) *ImageControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 上传按钮文案
 */
func (a *ImageControl) Uploadbtntext(value interface{}) *ImageControl {
    a.Set("uploadBtnText", value)
    return a
}

/**
 * 默认为 false, 开启后，允许用户输入压缩选项。
 */
func (a *ImageControl) Showcompressoptions(value interface{}) *ImageControl {
    a.Set("showCompressOptions", value)
    return a
}

/**
 * 可配置移动端的拍照功能，比如配置 `camera` 移动端只能拍照，等
 */
func (a *ImageControl) Capture(value interface{}) *ImageControl {
    a.Set("capture", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *ImageControl) Classname(value interface{}) *ImageControl {
    a.Set("className", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *ImageControl) Disabledon(value interface{}) *ImageControl {
    a.Set("disabledOn", value)
    return a
}

/**
 */
func (a *ImageControl) Staticschema(value interface{}) *ImageControl {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *ImageControl) Row(value interface{}) *ImageControl {
    a.Set("row", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *ImageControl) Clearvalueonhidden(value interface{}) *ImageControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 默认展示图片的类名
 */
func (a *ImageControl) Imageclassname(value interface{}) *ImageControl {
    a.Set("imageClassName", value)
    return a
}

/**
 * 是否允许二次裁剪。
 */
func (a *ImageControl) Recropable(value interface{}) *ImageControl {
    a.Set("reCropable", value)
    return a
}

/**
 * 默认没有限制，当设置后，文件大小大于此值将不允许上传。
 */
func (a *ImageControl) Maxsize(value interface{}) *ImageControl {
    a.Set("maxSize", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *ImageControl) Visibleon(value interface{}) *ImageControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *ImageControl) Staticclassname(value interface{}) *ImageControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *ImageControl) Name(value interface{}) *ImageControl {
    a.Set("name", value)
    return a
}

/**
 * 是否只读
 */
func (a *ImageControl) Readonly(value interface{}) *ImageControl {
    a.Set("readOnly", value)
    return a
}

/**
 * 缩路图展示模式
 * 可选值: w-full | h-full | contain | cover
 */
func (a *ImageControl) Thumbmode(value interface{}) *ImageControl {
    a.Set("thumbMode", value)
    return a
}

/**
 */
func (a *ImageControl) Crop(value interface{}) *ImageControl {
    a.Set("crop", value)
    return a
}

/**
 * 是否禁用
 */
func (a *ImageControl) Disabled(value interface{}) *ImageControl {
    a.Set("disabled", value)
    return a
}

/**
 * 是否显示
 */
func (a *ImageControl) Visible(value interface{}) *ImageControl {
    a.Set("visible", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *ImageControl) Static(value interface{}) *ImageControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *ImageControl) Staticinputclassname(value interface{}) *ImageControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 清除时设置的值
 */
func (a *ImageControl) Resetvalue(value interface{}) *ImageControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *ImageControl) Width(value interface{}) *ImageControl {
    a.Set("width", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *ImageControl) Staticon(value interface{}) *ImageControl {
    a.Set("staticOn", value)
    return a
}

/**
 * 是否隐藏上传按钮
 */
func (a *ImageControl) Hideuploadbutton(value interface{}) *ImageControl {
    a.Set("hideUploadButton", value)
    return a
}

/**
 * 默认 `/api/upload` 如果想自己存储，请设置此选项。
 */
func (a *ImageControl) Receiver(value interface{}) *ImageControl {
    a.Set("receiver", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *ImageControl) Extractvalue(value interface{}) *ImageControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 裁剪后的图片类型
 */
func (a *ImageControl) Cropformat(value interface{}) *ImageControl {
    a.Set("cropFormat", value)
    return a
}

/**
 * 配置 label className
 */
func (a *ImageControl) Labelclassname(value interface{}) *ImageControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *ImageControl) Value(value interface{}) *ImageControl {
    a.Set("value", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *ImageControl) Validateapi(value interface{}) *ImageControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 配置 input className
 */
func (a *ImageControl) Inputclassname(value interface{}) *ImageControl {
    a.Set("inputClassName", value)
    return a
}

/**
 */
func (a *ImageControl) Compress(value interface{}) *ImageControl {
    a.Set("compress", value)
    return a
}

/**
 * 最多的个数
 */
func (a *ImageControl) Maxlength(value interface{}) *ImageControl {
    a.Set("maxLength", value)
    return a
}

/**
 * 缩路图展示比率。
 * 可选值: 1:1 | 4:3 | 16:9
 */
func (a *ImageControl) Thumbratio(value interface{}) *ImageControl {
    a.Set("thumbRatio", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *ImageControl) Onevent(value interface{}) *ImageControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 组件样式
 */
func (a *ImageControl) Style(value interface{}) *ImageControl {
    a.Set("style", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *ImageControl) Hint(value interface{}) *ImageControl {
    a.Set("hint", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *ImageControl) Description(value interface{}) *ImageControl {
    a.Set("description", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *ImageControl) Validationerrors(value interface{}) *ImageControl {
    a.Set("validationErrors", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *ImageControl) Id(value interface{}) *ImageControl {
    a.Set("id", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *ImageControl) Usemobileui(value interface{}) *ImageControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *ImageControl) Size(value interface{}) *ImageControl {
    a.Set("size", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *ImageControl) Validateonchange(value interface{}) *ImageControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 指定为图片上传控件
 */
func (a *ImageControl) Type(value interface{}) *ImageControl {
    a.Set("type", value)
    return a
}

/**
 * 默认都是通过用户选择图片后上传返回图片地址，如果开启此选项，则可以允许用户图片地址。
 */
func (a *ImageControl) Allowinput(value interface{}) *ImageControl {
    a.Set("allowInput", value)
    return a
}

/**
 * 分割符
 */
func (a *ImageControl) Delimiter(value interface{}) *ImageControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 是否开启固定尺寸
 */
func (a *ImageControl) Fixedsize(value interface{}) *ImageControl {
    a.Set("fixedSize", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *ImageControl) Staticlabelclassname(value interface{}) *ImageControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *ImageControl) Extraname(value interface{}) *ImageControl {
    a.Set("extraName", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *ImageControl) Descriptionclassname(value interface{}) *ImageControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *ImageControl) Mode(value interface{}) *ImageControl {
    a.Set("mode", value)
    return a
}

/**
 * 是否自动开始上传
 */
func (a *ImageControl) Autoupload(value interface{}) *ImageControl {
    a.Set("autoUpload", value)
    return a
}

/**
 * 限制图片大小，超出不让上传。
 */
func (a *ImageControl) Limit(value interface{}) *ImageControl {
    a.Set("limit", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *ImageControl) Hidden(value interface{}) *ImageControl {
    a.Set("hidden", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *ImageControl) Editorsetting(value interface{}) *ImageControl {
    a.Set("editorSetting", value)
    return a
}

/**
 */
func (a *ImageControl) Testidbuilder(value interface{}) *ImageControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *ImageControl) Remark(value interface{}) *ImageControl {
    a.Set("remark", value)
    return a
}
