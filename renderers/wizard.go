package renderers


/**
 * 表单向导 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/wizard

 * @author wcz0
 * @version 6.2.2
 */
type Wizard struct {
	*BaseRenderer
}

func NewWizard() *Wizard {
    a := &Wizard{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "wizard")
    return a
}
/**
 * 是否用panel包裹
 */
func (a *Wizard) WrapWithPanel(value interface{}) *Wizard {
    a.Set("wrapWithPanel", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Wizard) VisibleOn(value interface{}) *Wizard {
    a.Set("visibleOn", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Wizard) EditorSetting(value interface{}) *Wizard {
    a.Set("editorSetting", value)
    return a
}

/**
 * 指定为表单向导
 */
func (a *Wizard) Type(value interface{}) *Wizard {
    a.Set("type", value)
    return a
}

/**
 * 配置按钮 className
 */
func (a *Wizard) ActionClassName(value interface{}) *Wizard {
    a.Set("actionClassName", value)
    return a
}

/**
 */
func (a *Wizard) Name(value interface{}) *Wizard {
    a.Set("name", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Wizard) Disabled(value interface{}) *Wizard {
    a.Set("disabled", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Wizard) HiddenOn(value interface{}) *Wizard {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Wizard) StaticOn(value interface{}) *Wizard {
    a.Set("staticOn", value)
    return a
}

/**
 * Wizard 用来获取初始数据的 api。
 */
func (a *Wizard) InitApi(value interface{}) *Wizard {
    a.Set("initApi", value)
    return a
}

/**
 * 底部操作栏的css类
 */
func (a *Wizard) FooterClassName(value interface{}) *Wizard {
    a.Set("footerClassName", value)
    return a
}

/**
 * 默认表单提交自己会通过发送 api 保存数据，但是也可以设定另外一个 form 的 name 值，或者另外一个 `CRUD` 模型的 name 值。 如果 target 目标是一个 `Form` ，则目标 `Form` 会重新触发 `initApi` 和 `schemaApi`，api 可以拿到当前 form 数据。如果目标是一个 `CRUD` 模型，则目标模型会重新触发搜索，参数为当前 Form 数据。
 */
func (a *Wizard) Target(value interface{}) *Wizard {
    a.Set("target", value)
    return a
}

/**
 */
func (a *Wizard) Steps(value interface{}) *Wizard {
    a.Set("steps", value)
    return a
}

/**
 * 表单区域css类
 */
func (a *Wizard) BodyClassName(value interface{}) *Wizard {
    a.Set("bodyClassName", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Wizard) Id(value interface{}) *Wizard {
    a.Set("id", value)
    return a
}

/**
 * 组件样式
 */
func (a *Wizard) Style(value interface{}) *Wizard {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Wizard) UseMobileUI(value interface{}) *Wizard {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 完成按钮的文字描述
 */
func (a *Wizard) ActionFinishLabel(value interface{}) *Wizard {
    a.Set("actionFinishLabel", value)
    return a
}

/**
 * 上一步按钮的文字描述
 */
func (a *Wizard) ActionPrevLabel(value interface{}) *Wizard {
    a.Set("actionPrevLabel", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Wizard) ClassName(value interface{}) *Wizard {
    a.Set("className", value)
    return a
}

/**
 * 是否显示
 */
func (a *Wizard) Visible(value interface{}) *Wizard {
    a.Set("visible", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Wizard) StaticInputClassName(value interface{}) *Wizard {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Wizard) Reload(value interface{}) *Wizard {
    a.Set("reload", value)
    return a
}

/**
 * 是否将底部按钮固定在底部。
 */
func (a *Wizard) AffixFooter(value interface{}) *Wizard {
    a.Set("affixFooter", value)
    return a
}

/**
 * step + body区域css类
 */
func (a *Wizard) StepClassName(value interface{}) *Wizard {
    a.Set("stepClassName", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Wizard) Hidden(value interface{}) *Wizard {
    a.Set("hidden", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Wizard) StaticLabelClassName(value interface{}) *Wizard {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 * Wizard 用来保存数据的 api。 [详情](https://baidu.github.io/amis/docs/api#wizard)
 */
func (a *Wizard) Api(value interface{}) *Wizard {
    a.Set("api", value)
    return a
}

/**
 * 保存完后，可以指定跳转地址，支持相对路径和组内绝对路径，同时可以通过 $xxx 使用变量
 */
func (a *Wizard) Redirect(value interface{}) *Wizard {
    a.Set("redirect", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Wizard) StaticClassName(value interface{}) *Wizard {
    a.Set("staticClassName", value)
    return a
}

/**
 * 下一步并且保存按钮的文字描述
 */
func (a *Wizard) ActionNextSaveLabel(value interface{}) *Wizard {
    a.Set("actionNextSaveLabel", value)
    return a
}

/**
 * 步骤条区域css类
 */
func (a *Wizard) StepsClassName(value interface{}) *Wizard {
    a.Set("stepsClassName", value)
    return a
}

/**
 * 是否为只读模式。
 */
func (a *Wizard) ReadOnly(value interface{}) *Wizard {
    a.Set("readOnly", value)
    return a
}

/**
 */
func (a *Wizard) LoadingConfig(value interface{}) *Wizard {
    a.Set("loadingConfig", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Wizard) Static(value interface{}) *Wizard {
    a.Set("static", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Wizard) StaticPlaceholder(value interface{}) *Wizard {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 下一步按钮的文字描述
 */
func (a *Wizard) ActionNextLabel(value interface{}) *Wizard {
    a.Set("actionNextLabel", value)
    return a
}

/**
 * 展示模式
 * 可选值: vertical | horizontal
 */
func (a *Wizard) Mode(value interface{}) *Wizard {
    a.Set("mode", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Wizard) DisabledOn(value interface{}) *Wizard {
    a.Set("disabledOn", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Wizard) OnEvent(value interface{}) *Wizard {
    a.Set("onEvent", value)
    return a
}

/**
 */
func (a *Wizard) StaticSchema(value interface{}) *Wizard {
    a.Set("staticSchema", value)
    return a
}

/**
 * 是否合并后再提交
 */
func (a *Wizard) BulkSubmit(value interface{}) *Wizard {
    a.Set("bulkSubmit", value)
    return a
}

/**
 */
func (a *Wizard) StartStep(value interface{}) *Wizard {
    a.Set("startStep", value)
    return a
}
