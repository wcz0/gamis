package renderers


/**
 * Dialog 弹框渲染器。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/dialog

 * @author wcz0
 * @version 6.2.2
 */
type Dialog struct {
	*BaseRenderer
}

func NewDialog() *Dialog {
    a := &Dialog{
        BaseRenderer: NewBaseRenderer(),
    }

    a.Set("type", "dialog")
    return a
}


func (a *Dialog) Set(name string, value interface{}) *Dialog {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 组件样式
 */
func (a *Dialog) Style(value interface{}) *Dialog {
    a.Set("style", value)
    return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *Dialog) Usemobileui(value interface{}) *Dialog {
    a.Set("useMobileUI", value)
    return a
}

/**
 * 是否支持点其它区域关闭 Dialog
 */
func (a *Dialog) Closeonoutside(value interface{}) *Dialog {
    a.Set("closeOnOutside", value)
    return a
}

/**
 */
func (a *Dialog) Name(value interface{}) *Dialog {
    a.Set("name", value)
    return a
}

/**
 * 影响自动生成的按钮，如果自己配置了按钮这个配置无效。
 */
func (a *Dialog) Confirm(value interface{}) *Dialog {
    a.Set("confirm", value)
    return a
}

/**
 * 是否禁用
 */
func (a *Dialog) Disabled(value interface{}) *Dialog {
    a.Set("disabled", value)
    return a
}

/**
 * 是否支持按 ESC 关闭 Dialog
 */
func (a *Dialog) Closeonesc(value interface{}) *Dialog {
    a.Set("closeOnEsc", value)
    return a
}

/**
 */
func (a *Dialog) Headerclassname(value interface{}) *Dialog {
    a.Set("headerClassName", value)
    return a
}

/**
 * 是否静态展示表达式
 */
func (a *Dialog) Staticon(value interface{}) *Dialog {
    a.Set("staticOn", value)
    return a
}

/**
 * 静态展示空值占位
 */
func (a *Dialog) Staticplaceholder(value interface{}) *Dialog {
    a.Set("staticPlaceholder", value)
    return a
}

/**
 * 请通过配置 title 设置标题
 */
func (a *Dialog) Title(value interface{}) *Dialog {
    a.Set("title", value)
    return a
}

/**
 * 是否显示错误信息
 */
func (a *Dialog) Showerrormsg(value interface{}) *Dialog {
    a.Set("showErrorMsg", value)
    return a
}

/**
 */
func (a *Dialog) Testid(value interface{}) *Dialog {
    a.Set("testid", value)
    return a
}

/**
 * 内容区域
 */
func (a *Dialog) Body(value interface{}) *Dialog {
    a.Set("body", value)
    return a
}

/**
 * 静态展示表单项类名
 */
func (a *Dialog) Staticclassname(value interface{}) *Dialog {
    a.Set("staticClassName", value)
    return a
}

/**
 */
func (a *Dialog) Staticschema(value interface{}) *Dialog {
    a.Set("staticSchema", value)
    return a
}

/**
 */
func (a *Dialog) Type(value interface{}) *Dialog {
    a.Set("type", value)
    return a
}

/**
 * 弹窗参数说明，值格式为 JSONSchema。
 */
func (a *Dialog) Inputparams(value interface{}) *Dialog {
    a.Set("inputParams", value)
    return a
}

/**
 * 默认不用填写，自动会创建确认和取消按钮。
 */
func (a *Dialog) Actions(value interface{}) *Dialog {
    a.Set("actions", value)
    return a
}

/**
 * Dialog 大小
 * 可选值: xs | sm | md | lg | xl | full
 */
func (a *Dialog) Size(value interface{}) *Dialog {
    a.Set("size", value)
    return a
}

/**
 * 是否隐藏
 */
func (a *Dialog) Hidden(value interface{}) *Dialog {
    a.Set("hidden", value)
    return a
}

/**
 */
func (a *Dialog) Footer(value interface{}) *Dialog {
    a.Set("footer", value)
    return a
}

/**
 * 是否显示关闭按钮
 */
func (a *Dialog) Showclosebutton(value interface{}) *Dialog {
    a.Set("showCloseButton", value)
    return a
}

/**
 * Dialog 高度
 */
func (a *Dialog) Height(value interface{}) *Dialog {
    a.Set("height", value)
    return a
}

/**
 * Dialog 宽度
 */
func (a *Dialog) Width(value interface{}) *Dialog {
    a.Set("width", value)
    return a
}

/**
 * 是否显示表达式
 */
func (a *Dialog) Visibleon(value interface{}) *Dialog {
    a.Set("visibleOn", value)
    return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *Dialog) Staticinputclassname(value interface{}) *Dialog {
    a.Set("staticInputClassName", value)
    return a
}

/**
 */
func (a *Dialog) Testidbuilder(value interface{}) *Dialog {
    a.Set("testIdBuilder", value)
    return a
}

/**
 * 配置 Body 容器 className
 */
func (a *Dialog) Bodyclassname(value interface{}) *Dialog {
    a.Set("bodyClassName", value)
    return a
}

/**
 * 弹框类型 confirm 确认弹框
 */
func (a *Dialog) Dialogtype(value interface{}) *Dialog {
    a.Set("dialogType", value)
    return a
}

/**
 * 事件动作配置
 */
func (a *Dialog) Onevent(value interface{}) *Dialog {
    a.Set("onEvent", value)
    return a
}

/**
 * 是否静态展示
 */
func (a *Dialog) Static(value interface{}) *Dialog {
    a.Set("static", value)
    return a
}

/**
 * 是否显示
 */
func (a *Dialog) Visible(value interface{}) *Dialog {
    a.Set("visible", value)
    return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *Dialog) Editorsetting(value interface{}) *Dialog {
    a.Set("editorSetting", value)
    return a
}

/**
 * 是否显示蒙层
 */
func (a *Dialog) Overlay(value interface{}) *Dialog {
    a.Set("overlay", value)
    return a
}

/**
 * 数据映射
 */
func (a *Dialog) Data(value interface{}) *Dialog {
    a.Set("data", value)
    return a
}

/**
 * 是否禁用表达式
 */
func (a *Dialog) Disabledon(value interface{}) *Dialog {
    a.Set("disabledOn", value)
    return a
}

/**
 * 是否隐藏表达式
 */
func (a *Dialog) Hiddenon(value interface{}) *Dialog {
    a.Set("hiddenOn", value)
    return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *Dialog) Staticlabelclassname(value interface{}) *Dialog {
    a.Set("staticLabelClassName", value)
    return a
}

/**
 */
func (a *Dialog) Header(value interface{}) *Dialog {
    a.Set("header", value)
    return a
}

/**
 * 是否显示 spinner
 */
func (a *Dialog) Showloading(value interface{}) *Dialog {
    a.Set("showLoading", value)
    return a
}

/**
 * 可拖拽
 */
func (a *Dialog) Draggable(value interface{}) *Dialog {
    a.Set("draggable", value)
    return a
}

/**
 * 容器 css 类名
 */
func (a *Dialog) Classname(value interface{}) *Dialog {
    a.Set("className", value)
    return a
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (a *Dialog) Id(value interface{}) *Dialog {
    a.Set("id", value)
    return a
}
