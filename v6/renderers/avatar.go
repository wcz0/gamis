package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type Avatar struct {
	*BaseRenderer
}

func NewAvatar() *Avatar {
    a := &Avatar{
        BaseRenderer: NewBaseRenderer(),
    }
    a.Set("type", "avatar")
    a.Set("crossOrigin", "anonymous")
    return a
}

/**
 */
func (a *Avatar) Testid(value interface{}) *Avatar {
    a.Set("testid", value)
    return a
}

/**
 * 默认头像
 */
func (a *Avatar) DefaultAvatar(value interface{}) *Avatar {
    a.Set("defaultAvatar", value)
    return a
}

/**
 * 大小
 */
func (a *Avatar) Size(value interface{}) *Avatar {
    a.Set("size", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Avatar) Hidden(value interface{}) *Avatar {
    a.Set("hidden", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Avatar) HiddenOn(value interface{}) *Avatar {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Avatar) Id(value interface{}) *Avatar {
    a.Set("id", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Avatar) OnEvent(value interface{}) *Avatar {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Avatar) Static(value interface{}) *Avatar {
    a.Set("static", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Avatar) UseMobileUI(value interface{}) *Avatar {
    a.Set("useMobileUI", value)
    return a
}

/**
 */
func (a *Avatar) TestIdBuilder(value interface{}) *Avatar {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 角标
 */
func (a *Avatar) Badge(value interface{}) *Avatar {
    a.Set("badge", value)
    return a
}

/**
 * 图片CORS属性
 * 可选值: anonymous | use-credentials | 
 */
func (a *Avatar) CrossOrigin(value interface{}) *Avatar {
    a.Set("crossOrigin", value)
    return a
}

/**
 * 图片加载失败的是否默认处理，字符串函数
 */
func (a *Avatar) OnError(value interface{}) *Avatar {
    a.Set("onError", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Avatar) DisabledOn(value interface{}) *Avatar {
    a.Set("disabledOn", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Avatar) StaticClassName(value interface{}) *Avatar {
    a.Set("staticClassName", value)
    return a
}

/**
 * 图片地址
 */
func (a *Avatar) Src(value interface{}) *Avatar {
    a.Set("src", value)
    return a
}

/**
 * 形状
 * 可选值: circle | square | rounded
 */
func (a *Avatar) Shape(value interface{}) *Avatar {
    a.Set("shape", value)
    return a
}

/**
 * 文本
 */
func (a *Avatar) Text(value interface{}) *Avatar {
    a.Set("text", value)
    return a
}

/**
 * 类名
 */
func (a *Avatar) ClassName(value interface{}) *Avatar {
    a.Set("className", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Avatar) StaticInputClassName(value interface{}) *Avatar {
    a.Set("staticInputClassName", value)
    return a
}

/**
 * 是否显示
 */
func (a *Avatar) Visible(value interface{}) *Avatar {
    a.Set("visible", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Avatar) VisibleOn(value interface{}) *Avatar {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Avatar) StaticLabelClassName(value interface{}) *Avatar {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Avatar) StaticSchema(value interface{}) *Avatar {
    a.Set("staticSchema", value)
    return a
}

/**
 * 自定义样式
 */
func (a *Avatar) Style(value interface{}) *Avatar {
    a.Set("style", value)
    return a
}

/**
 */
func (a *Avatar) Type(value interface{}) *Avatar {
    a.Set("type", value)
    return a
}

/**
 * 图片是否允许拖动
 */
func (a *Avatar) Draggable(value interface{}) *Avatar {
    a.Set("draggable", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Avatar) Disabled(value interface{}) *Avatar {
    a.Set("disabled", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Avatar) EditorSetting(value interface{}) *Avatar {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Avatar) StaticOn(value interface{}) *Avatar {
    a.Set("staticOn", value)
    return a
}

/**
 * 图片相对于容器的缩放方式
 * 可选值: fill | contain | cover | none | scale-down
 */
func (a *Avatar) Fit(value interface{}) *Avatar {
    a.Set("fit", value)
    return a
}

/**
 * 图片无法显示时的替换文字地址
 */
func (a *Avatar) Alt(value interface{}) *Avatar {
    a.Set("alt", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Avatar) StaticPlaceholder(value interface{}) *Avatar {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 图标
 */
func (a *Avatar) Icon(value interface{}) *Avatar {
    a.Set("icon", value)
    return a
}

/**
 * 字符类型距离左右两侧边界单位像素
 */
func (a *Avatar) Gap(value interface{}) *Avatar {
    a.Set("gap", value)
    return a
}
