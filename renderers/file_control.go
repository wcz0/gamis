package renderers


/**
 * File 文件上传控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/file

 * @author wcz0
 * @version 6.2.2
 */
type FileControl struct {
	*BaseRenderer
}

func NewFileControl() *FileControl {
    a := &FileControl{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "input-file")
    return a
}


func (a *FileControl) Set(name string, value interface{}) *FileControl {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 是否只读
 */
func (a *FileControl) ReadOnly(value interface{}) *FileControl {
    a.Set("readOnly", value)
    return a
}

/**
 */
func (a *FileControl) Desc(value interface{}) *FileControl {
    a.Set("desc", value)
    return a
}

/**
 * 初始化时是否把其他字段同步到表单内部。
 */
func (a *FileControl) InitAutoFill(value interface{}) *FileControl {
    a.Set("initAutoFill", value)
    return a
}

/**
 * 是否自动开始上传
 */
func (a *FileControl) AutoUpload(value interface{}) *FileControl {
    a.Set("autoUpload", value)
    return a
}

/**
 * 最多的个数
 */
func (a *FileControl) MaxLength(value interface{}) *FileControl {
    a.Set("maxLength", value)
    return a
}

/**
 * 说明文档内容配置
 */
func (a *FileControl) Documentation(value interface{}) *FileControl {
    a.Set("documentation", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容
 */
func (a *FileControl) Remark(value interface{}) *FileControl {
    a.Set("remark", value)
    return a
}

/**
 * 表单 control 是否为 inline 模式。
 */
func (a *FileControl) Inline(value interface{}) *FileControl {
    a.Set("inline", value)
    return a
}

/**
 */
func (a *FileControl) Validations(value interface{}) *FileControl {
    a.Set("validations", value)
    return a
}

/**
 * 远端校验表单项接口
 */
func (a *FileControl) ValidateApi(value interface{}) *FileControl {
    a.Set("validateApi", value)
    return a
}

/**
 * 控制 input 标签的 capture 属性，用于移动端拍照或录像。
 */
func (a *FileControl) Capture(value interface{}) *FileControl {
    a.Set("capture", value)
    return a
}

/**
 * 如果不希望 File 组件上传，可以配置 `asBlob` 或者 `asBase64`，采用这种方式后，组件不再自己上传了，而是直接把文件数据作为表单项的值，文件内容会在 Form 表单提交的接口里面一起带上。
 */
func (a *FileControl) AsBlob(value interface{}) *FileControl {
    a.Set("asBlob", value)
    return a
}

/**
 * 分块上传的并发数
 */
func (a *FileControl) Concurrency(value interface{}) *FileControl {
    a.Set("concurrency", value)
    return a
}

/**
 * 是否为多选
 */
func (a *FileControl) Multiple(value interface{}) *FileControl {
    a.Set("multiple", value)
    return a
}

/**
 * 1. 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交， 否则，整个选项对象都会作为该表单项的值提交。 2. 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来， 否则直接将以数组的形式提交值。
 */
func (a *FileControl) JoinValues(value interface{}) *FileControl {
    a.Set("joinValues", value)
    return a
}

/**
 */
func (a *FileControl) TestIdBuilder(value interface{}) *FileControl {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 描述标题
 */
func (a *FileControl) Label(value interface{}) *FileControl {
    a.Set("label", value)
    return a
}

/**
 * 上传按钮 CSS 类名
 */
func (a *FileControl) BtnUploadClassName(value interface{}) *FileControl {
    a.Set("btnUploadClassName", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *FileControl) Static(value interface{}) *FileControl {
    a.Set("static", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *FileControl) StaticClassName(value interface{}) *FileControl {
    a.Set("staticClassName", value)
    return a
}

/**
 * 组件样式
 */
func (a *FileControl) Style(value interface{}) *FileControl {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *FileControl) UseMobileUI(value interface{}) *FileControl {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 当修改完的时候是否提交表单。
 */
func (a *FileControl) SubmitOnChange(value interface{}) *FileControl {
    a.Set("submitOnChange", value)
    return a
}

/**
 * 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
 */
func (a *FileControl) ValidateOnChange(value interface{}) *FileControl {
    a.Set("validateOnChange", value)
    return a
}

/**
 * 配置 input className
 */
func (a *FileControl) InputClassName(value interface{}) *FileControl {
    a.Set("inputClassName", value)
    return a
}

/**
 * 上传后把其他字段同步到表单内部。
 */
func (a *FileControl) AutoFill(value interface{}) *FileControl {
    a.Set("autoFill", value)
    return a
}

/**
 * 分割符
 */
func (a *FileControl) Delimiter(value interface{}) *FileControl {
    a.Set("delimiter", value)
    return a
}

/**
 * 清除时设置的值
 */
func (a *FileControl) ResetValue(value interface{}) *FileControl {
    a.Set("resetValue", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *FileControl) ClassName(value interface{}) *FileControl {
    a.Set("className", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *FileControl) HiddenOn(value interface{}) *FileControl {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *FileControl) Id(value interface{}) *FileControl {
    a.Set("id", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *FileControl) EditorSetting(value interface{}) *FileControl {
    a.Set("editorSetting", value)
    return a
}

/**
 * 占位符
 */
func (a *FileControl) Placeholder(value interface{}) *FileControl {
    a.Set("placeholder", value)
    return a
}

/**
 * 接口返回的数据中，哪个用来当做值
 */
func (a *FileControl) ValueField(value interface{}) *FileControl {
    a.Set("valueField", value)
    return a
}

/**
 * 接口返回的数据中，哪个用来展示文件名
 */
func (a *FileControl) NameField(value interface{}) *FileControl {
    a.Set("nameField", value)
    return a
}

/**
 * 默认没有限制，当设置后，文件大小大于此值将不允许上传。
 */
func (a *FileControl) MaxSize(value interface{}) *FileControl {
    a.Set("maxSize", value)
    return a
}

/**
 * 在Table中调整宽度
 */
func (a *FileControl) Width(value interface{}) *FileControl {
    a.Set("width", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *FileControl) StaticPlaceholder(value interface{}) *FileControl {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 */
func (a *FileControl) StaticSchema(value interface{}) *FileControl {
    a.Set("staticSchema", value)
    return a
}

/**
 * 配置 label className
 */
func (a *FileControl) LabelClassName(value interface{}) *FileControl {
    a.Set("labelClassName", value)
    return a
}

/**
 * 是否为必填
 */
func (a *FileControl) Required(value interface{}) *FileControl {
    a.Set("required", value)
    return a
}

/**
 * 验证失败的提示信息
 */
func (a *FileControl) ValidationErrors(value interface{}) *FileControl {
    a.Set("validationErrors", value)
    return a
}

/**
 */
func (a *FileControl) Row(value interface{}) *FileControl {
    a.Set("row", value)
    return a
}

/**
 * 默认 `file`, 如果你不想自己存储，则可以忽略此属性。
 */
func (a *FileControl) FileField(value interface{}) *FileControl {
    a.Set("fileField", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *FileControl) VisibleOn(value interface{}) *FileControl {
    a.Set("visibleOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *FileControl) OnEvent(value interface{}) *FileControl {
    a.Set("onEvent", value)
    return a
}

/**
 * 只读条件
 */
func (a *FileControl) ReadOnlyOn(value interface{}) *FileControl {
    a.Set("readOnlyOn", value)
    return a
}

/**
 * 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
 */
func (a *FileControl) ClearValueOnHidden(value interface{}) *FileControl {
    a.Set("clearValueOnHidden", value)
    return a
}

/**
 * 分块大小，默认为 5M.
 */
func (a *FileControl) ChunkSize(value interface{}) *FileControl {
    a.Set("chunkSize", value)
    return a
}

/**
 * 接口返回的数据中哪个用来作为下载地址。
 */
func (a *FileControl) UrlField(value interface{}) *FileControl {
    a.Set("urlField", value)
    return a
}

/**
 * 是否显示
 */
func (a *FileControl) Visible(value interface{}) *FileControl {
    a.Set("visible", value)
    return a
}

/**
 * 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
 */
func (a *FileControl) LabelRemark(value interface{}) *FileControl {
    a.Set("labelRemark", value)
    return a
}

/**
 * 默认 `/api/upload/chunk` 想自己存储时才需要关注。
 */
func (a *FileControl) ChunkApi(value interface{}) *FileControl {
    a.Set("chunkApi", value)
    return a
}

/**
 * 默认 `/api/upload/startChunk` 想自己存储时才需要关注。
 */
func (a *FileControl) StartChunkApi(value interface{}) *FileControl {
    a.Set("startChunkApi", value)
    return a
}

/**
 * 说明文档链接配置
 */
func (a *FileControl) DocumentLink(value interface{}) *FileControl {
    a.Set("documentLink", value)
    return a
}

/**
 * 是否为拖拽上传
 */
func (a *FileControl) Drag(value interface{}) *FileControl {
    a.Set("drag", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *FileControl) DisabledOn(value interface{}) *FileControl {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *FileControl) StaticInputClassName(value interface{}) *FileControl {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
 */
func (a *FileControl) ExtraName(value interface{}) *FileControl {
    a.Set("extraName", value)
    return a
}

/**
 * 配置描述上的 className
 */
func (a *FileControl) DescriptionClassName(value interface{}) *FileControl {
    a.Set("descriptionClassName", value)
    return a
}

/**
 * 上传文件按钮说明
 */
func (a *FileControl) BtnLabel(value interface{}) *FileControl {
    a.Set("btnLabel", value)
    return a
}

/**
 * 表单项大小
 * 可选值: xs | sm | md | lg | full
 */
func (a *FileControl) Size(value interface{}) *FileControl {
    a.Set("size", value)
    return a
}

/**
 * 描述标题
 */
func (a *FileControl) LabelAlign(value interface{}) *FileControl {
    a.Set("labelAlign", value)
    return a
}

/**
 * 配置当前表单项展示模式
 * 可选值: normal | inline | horizontal
 */
func (a *FileControl) Mode(value interface{}) *FileControl {
    a.Set("mode", value)
    return a
}

/**
 * 默认为 'auto' amis 所在服务器，限制了文件上传大小不得超出10M，所以 amis 在用户选择大文件的时候，自动会改成分块上传模式。
 */
func (a *FileControl) UseChunk(value interface{}) *FileControl {
    a.Set("useChunk", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *FileControl) StaticLabelClassName(value interface{}) *FileControl {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
 */
func (a *FileControl) Name(value interface{}) *FileControl {
    a.Set("name", value)
    return a
}

/**
 * 当配置为水平布局的时候，用来配置具体的左右分配。
 */
func (a *FileControl) Horizontal(value interface{}) *FileControl {
    a.Set("horizontal", value)
    return a
}

/**
 * 指定为文件上传
 */
func (a *FileControl) Type(value interface{}) *FileControl {
    a.Set("type", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *FileControl) Hidden(value interface{}) *FileControl {
    a.Set("hidden", value)
    return a
}

/**
 * 输入提示，聚焦的时候显示
 */
func (a *FileControl) Hint(value interface{}) *FileControl {
    a.Set("hint", value)
    return a
}

/**
 * 描述内容，支持 Html 片段。
 */
func (a *FileControl) Description(value interface{}) *FileControl {
    a.Set("description", value)
    return a
}

/**
 * 如果上传的文件比较小可以设置此选项来简单的把文件 base64 的值给 form 一起提交，目前不支持多选。
 */
func (a *FileControl) AsBase64(value interface{}) *FileControl {
    a.Set("asBase64", value)
    return a
}

/**
 * 模板下载地址
 */
func (a *FileControl) TemplateUrl(value interface{}) *FileControl {
    a.Set("templateUrl", value)
    return a
}

/**
 * 默认 `/api/upload/file` 如果想自己存储，请设置此选项。
 */
func (a *FileControl) Receiver(value interface{}) *FileControl {
    a.Set("receiver", value)
    return a
}

/**
 * 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
 */
func (a *FileControl) ExtractValue(value interface{}) *FileControl {
    a.Set("extractValue", value)
    return a
}

/**
 * 是否禁用
 */
func (a *FileControl) Disabled(value interface{}) *FileControl {
    a.Set("disabled", value)
    return a
}

/**
 * 默认只支持纯文本，要支持其他类型，请配置此属性。建议直接填写文件后缀 如：.txt,.csv多个类型用逗号隔开。
 */
func (a *FileControl) Accept(value interface{}) *FileControl {
    a.Set("accept", value)
    return a
}

/**
 * 默认显示文件路径的时候会支持直接下载， 可以支持加前缀如：`http://xx.dom/filename=` ， 如果不希望这样，可以把当前配置项设置为 `false`。1.1.6 版本开始将支持变量 ${xxx} 来自己拼凑个下载地址，并且支持配置成 post.
 */
func (a *FileControl) DownloadUrl(value interface{}) *FileControl {
    a.Set("downloadUrl", value)
    return a
}

/**
 * 是否隐藏上传按钮
 */
func (a *FileControl) HideUploadButton(value interface{}) *FileControl {
    a.Set("hideUploadButton", value)
    return a
}

/**
 * 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
 */
func (a *FileControl) Value(value interface{}) *FileControl {
    a.Set("value", value)
    return a
}

/**
 * 默认 `/api/upload/finishChunkApi` 想自己存储时才需要关注。
 */
func (a *FileControl) FinishChunkApi(value interface{}) *FileControl {
    a.Set("finishChunkApi", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *FileControl) StaticOn(value interface{}) *FileControl {
    a.Set("staticOn", value)
    return a
}

/**
 * label自定义宽度，默认单位为px
 */
func (a *FileControl) LabelWidth(value interface{}) *FileControl {
    a.Set("labelWidth", value)
    return a
}

/**
 * 按钮 CSS 类名
 */
func (a *FileControl) BtnClassName(value interface{}) *FileControl {
    a.Set("btnClassName", value)
    return a
}

/**
 * 按钮状态文案配置。
 */
func (a *FileControl) StateTextMap(value interface{}) *FileControl {
    a.Set("stateTextMap", value)
    return a
}
