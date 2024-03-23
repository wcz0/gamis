// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * AjaxAction
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class AjaxAction extends BaseRenderer
// {
//     public function __construct()
//     {
//         $this->set('type', 'button');
// $this->set('actionType', 'ajax');


//     }

//     /**
//      * 指定为发送 ajax 的行为。
//      */
//     public function actionType($value = 'ajax')
//     {
//         return $this->set('actionType', $value);
//     }

//     /**
//      * 激活状态时的类名
//      */
//     public function activeClassName($value = '')
//     {
//         return $this->set('activeClassName', $value);
//     }

//     /**
//      * 激活状态时的样式
//      */
//     public function activeLevel($value = '')
//     {
//         return $this->set('activeLevel', $value);
//     }

//     /**
//      * 配置 ajax 发送地址
//      */
//     public function api($value = '')
//     {
//         return $this->set('api', $value);
//     }

//     /**
//      * 角标
//      */
//     public function badge($value = '')
//     {
//         return $this->set('badge', $value);
//     }

//     /**
//      * 是否为块状展示，默认为内联。
//      */
//     public function block($value = true)
//     {
//         return $this->set('block', $value);
//     }

//     /**
//      * 子内容
//      */
//     public function body($value = '')
//     {
//         return $this->set('body', $value);
//     }

//     /**
//      * 容器 css 类名
//      */
//     public function className($value = '')
//     {
//         return $this->set('className', $value);
//     }

//     /**
//      * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
//      */
//     public function close($value = '')
//     {
//         return $this->set('close', $value);
//     }

//     /**
//      * 提示文字，配置了操作前会要求用户确认。
//      */
//     public function confirmText($value = '')
//     {
//         return $this->set('confirmText', $value);
//     }

//     /**
//      * 点击后的禁止倒计时（秒）
//      */
//     public function countDown($value = '')
//     {
//         return $this->set('countDown', $value);
//     }

//     /**
//      * 倒计时文字自定义
//      */
//     public function countDownTpl($value = '')
//     {
//         return $this->set('countDownTpl', $value);
//     }

//     /**
//      * 是否禁用
//      */
//     public function disabled($value = true)
//     {
//         return $this->set('disabled', $value);
//     }

//     /**
//      * 是否禁用表达式
//      */
//     public function disabledOn($value = '')
//     {
//         return $this->set('disabledOn', $value);
//     }

//     /**
//      * 禁用时的文案提示。
//      */
//     public function disabledTip($value = '')
//     {
//         return $this->set('disabledTip', $value);
//     }

//     /**
//      * 编辑器配置，运行时可以忽略
//      */
//     public function editorSetting($value = '')
//     {
//         return $this->set('editorSetting', $value);
//     }

//     /**
//      *
//      */
//     public function feedback($value = '')
//     {
//         return $this->set('feedback', $value);
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
//      * 键盘快捷键
//      */
//     public function hotKey($value = '')
//     {
//         return $this->set('hotKey', $value);
//     }

//     /**
//      * 按钮图标， iconfont 的类名
//      */
//     public function icon($value = '')
//     {
//         return $this->set('icon', $value);
//     }

//     /**
//      * icon 上的css 类名
//      */
//     public function iconClassName($value = '')
//     {
//         return $this->set('iconClassName', $value);
//     }

//     /**
//      * 主要用于用户行为跟踪里区分是哪个按钮
//      */
//     public function id($value = '')
//     {
//         return $this->set('id', $value);
//     }

//     /**
//      *
//      */
//     public function ignoreConfirm($value = true)
//     {
//         return $this->set('ignoreConfirm', $value);
//     }

//     /**
//      * 是否开启请求隔离, 主要用于隔离联动CRUD, Service的请求
//      */
//     public function isolateScope($value = true)
//     {
//         return $this->set('isolateScope', $value);
//     }

//     /**
//      * 按钮文字
//      */
//     public function label($value = '')
//     {
//         return $this->set('label', $value);
//     }

//     /**
//      * 按钮样式 可选值: info | success | warning | danger | link | primary | dark | light | secondary
//      */
//     public function level($value = '')
//     {
//         return $this->set('level', $value);
//     }

//     /**
//      * loading 上的css 类名
//      */
//     public function loadingClassName($value = '')
//     {
//         return $this->set('loadingClassName', $value);
//     }

//     /**
//      * 是否显示loading效果
//      */
//     public function loadingOn($value = '')
//     {
//         return $this->set('loadingOn', $value);
//     }

//     /**
//      * 是否将弹框中数据 merge 到父级作用域。
//      */
//     public function mergeData($value = true)
//     {
//         return $this->set('mergeData', $value);
//     }

//     /**
//      * 自定义事件处理函数
//      */
//     public function onClick($value = '')
//     {
//         return $this->set('onClick', $value);
//     }

//     /**
//      * 事件动作配置
//      */
//     public function onEvent($value = '')
//     {
//         return $this->set('onEvent', $value);
//     }

//     /**
//      *
//      */
//     public function primary($value = true)
//     {
//         return $this->set('primary', $value);
//     }

//     /**
//      *
//      */
//     public function redirect($value = '')
//     {
//         return $this->set('redirect', $value);
//     }

//     /**
//      *
//      */
//     public function reload($value = '')
//     {
//         return $this->set('reload', $value);
//     }

//     /**
//      * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
//      */
//     public function requireSelected($value = true)
//     {
//         return $this->set('requireSelected', $value);
//     }

//     /**
//      * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
//      */
//     public function required($value = '')
//     {
//         return $this->set('required', $value);
//     }

//     /**
//      * 右侧按钮图标， iconfont 的类名
//      */
//     public function rightIcon($value = '')
//     {
//         return $this->set('rightIcon', $value);
//     }

//     /**
//      * 右侧 icon 上的 css 类名
//      */
//     public function rightIconClassName($value = '')
//     {
//         return $this->set('rightIconClassName', $value);
//     }

//     /**
//      * 按钮大小 可选值: xs | sm | md | lg
//      */
//     public function size($value = '')
//     {
//         return $this->set('size', $value);
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
//      * 可以指定让谁来触发这个动作。
//      */
//     public function target($value = '')
//     {
//         return $this->set('target', $value);
//     }

//     /**
//      *
//      */
//     public function testid($value = '')
//     {
//         return $this->set('testid', $value);
//     }

//     /**
//      *
//      */
//     public function tooltip($value = '')
//     {
//         return $this->set('tooltip', $value);
//     }

//     /**
//      *  可选值: top | right | bottom | left
//      */
//     public function tooltipPlacement($value = '')
//     {
//         return $this->set('tooltipPlacement', $value);
//     }

//     /**
//      * 指定按钮类型，支持 button、submit或者reset三种类型。 可选值: button | submit | reset
//      */
//     public function type($value = '')
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
//      * 是否显示
//      */
//     public function visible($value = true)
//     {
//         return $this->set('visible', $value);
//     }

//     /**
//      * 是否显示表达式
//      */
//     public function visibleOn($value = '')
//     {
//         return $this->set('visibleOn', $value);
//     }


// }

package renderers

type AjaxAction struct {
	*BaseRenderer
}

func NewAjaxAction() *AjaxAction{
	a := &AjaxAction{
		BaseRenderer: NewBaseRenderer(),
	}
	a.Set("type", "button")
	a.Set("actionType", "ajax")
	return a
}

/**
 * 指定为发送 ajax 的行为。
 */
func (a *AjaxAction) ActionType(value string) *AjaxAction {
	a.Set("actionType", value)
	return a
}

/**
 * 激活状态时的类名
 */
func (a *AjaxAction) ActiveClassName(value string) *AjaxAction {
	a.Set("activeClassName", value)
	return a
}

/**
 * 激活状态时的样式
 */
func (a *AjaxAction) ActiveLevel(value string) *AjaxAction {
	a.Set("activeLevel", value)
	return a
}

/**
 * 配置 ajax 发送地址
 */
func (a *AjaxAction) Api(value string) *AjaxAction {
	a.Set("api", value)
	return a
}

/**
 * 角标
 */
func (a *AjaxAction) Badge(value string) *AjaxAction {
	a.Set("badge", value)
	return a
}

/**
 * 是否为块状展示，默认为内联。
 */
func (a *AjaxAction) Block(value bool) *AjaxAction {
	a.Set("block", value)
	return a
}

/**
 * 子内容
 */
func (a *AjaxAction) Body(value string) *AjaxAction {
	a.Set("body", value)
	return a
}

/**
 * 容器 css 类名
 */
func (a *AjaxAction) ClassName(value string) *AjaxAction {
	a.Set("className", value)
	return a
}

/**
 * 如果按钮在弹框中，可以配置这个动作完成后是否关闭弹窗，或者指定关闭目标弹框。
 */
func (a *AjaxAction) Close(value string) *AjaxAction {
	a.Set("close", value)
	return a
}

/**
 * 提示文字，配置了操作前会要求用户确认。
 */
func (a *AjaxAction) ConfirmText(value string) *AjaxAction {
	a.Set("confirmText", value)
	return a
}

/**
 * 点击后的禁止倒计时（秒）
 */
func (a *AjaxAction) CountDown(value string) *AjaxAction {
	a.Set("countDown", value)
	return a
}

/**
 * 倒计时文字自定义
 */
func (a *AjaxAction) CountDownTpl(value string) *AjaxAction {
	a.Set("countDownTpl", value)
	return a
}

/**
 * 是否禁用
 */
func (a *AjaxAction) Disabled(value bool) *AjaxAction {
	a.Set("disabled", value)
	return a
}

/**
 * 是否禁用表达式
 */
func (a *AjaxAction) DisabledOn(value string) *AjaxAction {
	a.Set("disabledOn", value)
	return a
}

/**
 * 禁用时的文案提示。
 */
func (a *AjaxAction) DisabledTip(value string) *AjaxAction {
	a.Set("disabledTip", value)
	return a
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (a *AjaxAction) EditorSetting(value string) *AjaxAction {
	a.Set("editorSetting", value)
	return a
}

/**
 *
 */
func (a *AjaxAction) Feedback(value string) *AjaxAction {
	a.Set("feedback", value)
	return a
}

/**
 * 是否隐藏
 */
func (a *AjaxAction) Hidden(value bool) *AjaxAction {
	a.Set("hidden", value)
	return a
}

/**
 * 是否隐藏表达式
 */
func (a *AjaxAction) HiddenOn(value string) *AjaxAction {
	a.Set("hiddenOn", value)
	return a
}

/**
 * 键盘快捷键
 */
func (a *AjaxAction) HotKey(value string) *AjaxAction {
	a.Set("hotKey", value)
	return a
}

/**
 * 按钮图标， iconfont 的类名
 */
func (a *AjaxAction) Icon(value string) *AjaxAction {
	a.Set("icon", value)
	return a
}

/**
 * icon 上的css 类名
 */
func (a *AjaxAction) IconClassName(value string) *AjaxAction {
	a.Set("iconClassName", value)
	return a
}

/**
 * 主要用于用户行为跟踪里区分是哪个按钮
 */
func (a *AjaxAction) Id(value string) *AjaxAction {
	a.Set("id", value)
	return a
}

/**
 *
 */
func (a *AjaxAction) IgnoreConfirm(value bool) *AjaxAction {
	a.Set("ignoreConfirm", value)
	return a
}

/**
 * 是否开启请求隔离, 主要用于隔离联动CRUD, Service的请求
 */
func (a *AjaxAction) IsolateScope(value bool) *AjaxAction {
	a.Set("isolateScope", value)
	return a
}

/**
 * 按钮文字
 */
func (a *AjaxAction) Label(value string) *AjaxAction {
	a.Set("label", value)
	return a
}

/**
 * 按钮样式 可选值: info | success | warning | danger | link | primary | dark | light | secondary
 */
func (a *AjaxAction) Level(value string) *AjaxAction {
	a.Set("level", value)
	return a
}

/**
 * loading 上的css 类名
 */
func (a *AjaxAction) LoadingClassName(value string) *AjaxAction {
	a.Set("loadingClassName", value)
	return a
}

/**
 * 是否显示loading效果
 */
func (a *AjaxAction) LoadingOn(value string) *AjaxAction {
	a.Set("loadingOn", value)
	return a
}

/**
 * 是否将弹框中数据 merge 到父级作用域。
 */
func (a *AjaxAction) MergeData(value bool) *AjaxAction {
	a.Set("mergeData", value)
	return a
}

/**
 * 自定义事件处理函数
 */
func (a *AjaxAction) OnClick(value string) *AjaxAction {
	a.Set("onClick", value)
	return a
}

/**
 * 事件动作配置
 */
func (a *AjaxAction) OnEvent(value string) *AjaxAction {
	a.Set("onEvent", value)
	return a
}

/**
 *
 */
func (a *AjaxAction) Primary(value bool) *AjaxAction {
	a.Set("primary", value)
	return a
}

/**
 *
 */
func (a *AjaxAction) Redirect(value string) *AjaxAction {
	a.Set("redirect", value)
	return a
}

/**
 *
 */
func (a *AjaxAction) Reload(value string) *AjaxAction {
	a.Set("reload", value)
	return a
}

/**
 * 当按钮时批量操作按钮时，默认必须有勾选元素才能可点击，如果此属性配置成 false，则没有点选成员也能点击。
 */
func (a *AjaxAction) RequireSelected(value bool) *AjaxAction {
	a.Set("requireSelected", value)
	return a
}

/**
 * 如果按钮在form中，配置此属性会要求用户把指定的字段通过验证后才会触发行为。
 */
func (a *AjaxAction) Required(value string) *AjaxAction {
	a.Set("required", value)
	return a
}

/**
 * 右侧按钮图标， iconfont 的类名
 */
func (a *AjaxAction) RightIcon(value string) *AjaxAction {
	a.Set("rightIcon", value)
	return a
}

/**
 * 右侧 icon 上的 css 类名
 */
func (a *AjaxAction) RightIconClassName(value string) *AjaxAction {
	a.Set("rightIconClassName", value)
	return a
}

/**
 * 按钮大小 可选值: xs | sm | md | lg
 */
func (a *AjaxAction) Size(value string) *AjaxAction {
	a.Set("size", value)
	return a
}

/**
 * 是否静态展示
 */
func (a *AjaxAction) Static(value bool) *AjaxAction {
	a.Set("static", value)
	return a
}

/**
 * 静态展示表单项类名
 */
func (a *AjaxAction) StaticClassName(value string) *AjaxAction {
	a.Set("staticClassName", value)
	return a
}

/**
 * 静态展示表单项Value类名
 */
func (a *AjaxAction) StaticInputClassName(value string) *AjaxAction {
	a.Set("staticInputClassName", value)
	return a
}

/**
 * 静态展示表单项Label类名
 */
func (a *AjaxAction) StaticLabelClassName(value string) *AjaxAction {
	a.Set("staticLabelClassName", value)
	return a
}

/**
 * 是否静态展示表达式
 */
func (a *AjaxAction) StaticOn(value string) *AjaxAction {
	a.Set("staticOn", value)
	return a
}

/**
 * 静态展示空值占位
 */
func (a *AjaxAction) StaticPlaceholder(value string) *AjaxAction {
	a.Set("staticPlaceholder", value)
	return a
}

/**
 *
 */
func (a *AjaxAction) StaticSchema(value string) *AjaxAction {
	a.Set("staticSchema", value)
	return a
}

/**
 * 组件样式
 */
func (a *AjaxAction) Style(value string) *AjaxAction {
	a.Set("style", value)
	return a
}

/**
 * 可以指定让谁来触发这个动作。
 */
func (a *AjaxAction) Target(value string) *AjaxAction {
	a.Set("target", value)
	return a
}

/**
 *
 */
func (a *AjaxAction) Testid(value string) *AjaxAction {
	a.Set("testid", value)
	return a
}

/**
 *
 */
func (a *AjaxAction) Tooltip(value string) *AjaxAction {
	a.Set("tooltip", value)
	return a
}

/**
 *  可选值: top | right | bottom | left
 */
func (a *AjaxAction) TooltipPlacement(value string) *AjaxAction {
	a.Set("tooltipPlacement", value)
	return a
}

/**
 * 指定按钮类型，支持 button、submit或者reset三种类型。 可选值: button | submit | reset
 */
func (a *AjaxAction) Type(value string) *AjaxAction {
	a.Set("type", value)
	return a
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (a *AjaxAction) UseMobileUI(value bool) *AjaxAction {
	a.Set("useMobileUI", value)
	return a
}

/**
 * 是否显示
 */
func (a *AjaxAction) Visible(value bool) *AjaxAction {
	a.Set("visible", value)
	return a
}

/**
 * 是否显示表达式
 */
func (a *AjaxAction) VisibleOn(value string) *AjaxAction {
	a.Set("visibleOn", value)
	return a
}
