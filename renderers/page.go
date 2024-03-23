// <?php

// namespace Slowlyo\AmisRenderers;

// /**
//  * amis Page 渲染器。详情请见：https://aisuda.bce.baidu.com/amis/zh-CN/components/page
//  *
//  * @author slowlyo
//  * @version 6.2.2
//  */
// class Page extends BaseRenderer
// {
//     public function __construct()
//     {
//         $this->set('type', 'page');


//     }

//     /**
//      * 边栏区域
//      */
//     public function aside($value = '')
//     {
//         return $this->set('aside', $value);
//     }

//     /**
//      * 边栏区 css 类名
//      */
//     public function asideClassName($value = '')
//     {
//         return $this->set('asideClassName', $value);
//     }

//     /**
//      * 边栏最小宽度
//      */
//     public function asideMaxWidth($value = '')
//     {
//         return $this->set('asideMaxWidth', $value);
//     }

//     /**
//      * 边栏最小宽度
//      */
//     public function asideMinWidth($value = '')
//     {
//         return $this->set('asideMinWidth', $value);
//     }

//     /**
//      * 边栏是否允许拖动
//      */
//     public function asideResizor($value = true)
//     {
//         return $this->set('asideResizor', $value);
//     }

//     /**
//      * 边栏内容是否粘住，即不跟随滚动。
//      */
//     public function asideSticky($value = true)
//     {
//         return $this->set('asideSticky', $value);
//     }

//     /**
//      * 内容区域
//      */
//     public function body($value = '')
//     {
//         return $this->set('body', $value);
//     }

//     /**
//      * 内容区 css 类名
//      */
//     public function bodyClassName($value = '')
//     {
//         return $this->set('bodyClassName', $value);
//     }

//     /**
//      * 配置容器 className
//      */
//     public function className($value = '')
//     {
//         return $this->set('className', $value);
//     }

//     /**
//      * 自定义页面级别样式表
//      */
//     public function css($value = '')
//     {
//         return $this->set('css', $value);
//     }

//     /**
//      * css 变量
//      */
//     public function cssVars($value = '')
//     {
//         return $this->set('cssVars', $value);
//     }

//     /**
//      * 页面级别的初始数据
//      */
//     public function data($value = '')
//     {
//         return $this->set('data', $value);
//     }

//     /**
//      *
//      */
//     public function definitions($value = '')
//     {
//         return $this->set('definitions', $value);
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
//      * 编辑器配置，运行时可以忽略
//      */
//     public function editorSetting($value = '')
//     {
//         return $this->set('editorSetting', $value);
//     }

//     /**
//      * 配置 header 容器 className
//      */
//     public function headerClassName($value = '')
//     {
//         return $this->set('headerClassName', $value);
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
//      * 组件唯一 id，主要用于日志采集
//      */
//     public function id($value = '')
//     {
//         return $this->set('id', $value);
//     }

//     /**
//      * 页面初始化的时候，可以设置一个 API 让其取拉取，发送数据会携带当前 data 数据（包含地址栏参数），获取得数据会合并到 data 中，供组件内使用。
//      */
//     public function initApi($value = '')
//     {
//         return $this->set('initApi', $value);
//     }

//     /**
//      * 是否默认就拉取？
//      */
//     public function initFetch($value = true)
//     {
//         return $this->set('initFetch', $value);
//     }

//     /**
//      * 是否默认就拉取表达式
//      */
//     public function initFetchOn($value = '')
//     {
//         return $this->set('initFetchOn', $value);
//     }

//     /**
//      * 配置轮询间隔，配置后 initApi 将轮询加载。
//      */
//     public function interval($value = '')
//     {
//         return $this->set('interval', $value);
//     }

//     /**
//      *
//      */
//     public function loadingConfig($value = '')
//     {
//         return $this->set('loadingConfig', $value);
//     }

//     /**
//      *
//      */
//     public function messages($value = '')
//     {
//         return $this->set('messages', $value);
//     }

//     /**
//      * 移动端下的样式表
//      */
//     public function mobileCSS($value = '')
//     {
//         return $this->set('mobileCSS', $value);
//     }

//     /**
//      *
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
//      * 下拉刷新配置
//      */
//     public function pullRefresh($value = '')
//     {
//         return $this->set('pullRefresh', $value);
//     }

//     /**
//      * 默认不设置自动感觉内容来决定要不要展示这些区域 如果配置了，以配置为主。
//      */
//     public function regions($value = '')
//     {
//         return $this->set('regions', $value);
//     }

//     /**
//      * 页面描述, 标题旁边会出现个小图标，放上去会显示这个属性配置的内容。
//      */
//     public function remark($value = '')
//     {
//         return $this->set('remark', $value);
//     }

//     /**
//      * 是否显示错误信息，默认是显示的。
//      */
//     public function showErrorMsg($value = true)
//     {
//         return $this->set('showErrorMsg', $value);
//     }

//     /**
//      * 是否要静默加载，也就是说不显示进度
//      */
//     public function silentPolling($value = true)
//     {
//         return $this->set('silentPolling', $value);
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
//      * 配置停止轮询的条件。
//      */
//     public function stopAutoRefreshWhen($value = '')
//     {
//         return $this->set('stopAutoRefreshWhen', $value);
//     }

//     /**
//      * 自定义样式
//      */
//     public function style($value = '')
//     {
//         return $this->set('style', $value);
//     }

//     /**
//      * 页面副标题
//      */
//     public function subTitle($value = '')
//     {
//         return $this->set('subTitle', $value);
//     }

//     /**
//      * 页面标题
//      */
//     public function title($value = '')
//     {
//         return $this->set('title', $value);
//     }

//     /**
//      * 页面顶部区域，当存在 title 时在右上角显示。
//      */
//     public function toolbar($value = '')
//     {
//         return $this->set('toolbar', $value);
//     }

//     /**
//      * 配置 toolbar 容器 className
//      */
//     public function toolbarClassName($value = '')
//     {
//         return $this->set('toolbarClassName', $value);
//     }

//     /**
//      * 指定为 page 渲染器。
//      */
//     public function type($value = 'page')
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

type Page struct {
	*BaseRenderer
}

func NewPage() *Page{
	p := &Page{
		BaseRenderer: NewBaseRenderer(),
	}
	p.Set("type", "page")
	return p
}

/**
 * 边栏区域
 */
func (p *Page) Aside(value string) *Page {
	p.Set("aside", value)
	return p
}

/**
 * 边栏区 css 类名
 */
func (p *Page) AsideClassName(value string) *Page {
	p.Set("asideClassName", value)
	return p
}

/**
 * 边栏最小宽度
 */
func (p *Page) AsideMaxWidth(value string) *Page {
	p.Set("asideMaxWidth", value)
	return p
}

/**
 * 边栏最小宽度
 */
func (p *Page) AsideMinWidth(value string) *Page {
	p.Set("asideMinWidth", value)
	return p
}

/**
 * 边栏是否允许拖动
 */
func (p *Page) AsideResizor(value bool) *Page {
	p.Set("asideResizor", value)
	return p
}

/**
 * 边栏内容是否粘住，即不跟随滚动。
 */
func (p *Page) AsideSticky(value bool) *Page {
	p.Set("asideSticky", value)
	return p
}

/**
 * 内容区域
 */
func (p *Page) Body(value string) *Page {
	p.Set("body", value)
	return p
}

/**
 * 内容区 css 类名
 */
func (p *Page) BodyClassName(value string) *Page {
	p.Set("bodyClassName", value)
	return p
}

/**
 * 配置容器 className
 */
func (p *Page) ClassName(value string) *Page {
	p.Set("className", value)
	return p
}

/**
 * 自定义页面级别样式表
 */
func (p *Page) Css(value string) *Page {
	p.Set("css", value)
	return p
}

/**
 * css 变量
 */
func (p *Page) CssVars(value string) *Page {
	p.Set("cssVars", value)
	return p
}

/**
 * 页面级别的初始数据
 */
func (p *Page) Data(value string) *Page {
	p.Set("data", value)
	return p
}

/**
 *
 */
func (p *Page) Definitions(value string) *Page {
	p.Set("definitions", value)
	return p
}

/**
 * 是否禁用
 */
func (p *Page) Disabled(value bool) *Page {
	p.Set("disabled", value)
	return p
}

/**
 * 是否禁用表达式
 */
func (p *Page) DisabledOn(value string) *Page {
	p.Set("disabledOn", value)
	return p
}

/**
 * 编辑器配置，运行时可以忽略
 */
func (p *Page) EditorSetting(value string) *Page {
	p.Set("editorSetting", value)
	return p
}

/**
 * 配置 header 容器 className
 */
func (p *Page) HeaderClassName(value string) *Page {
	p.Set("headerClassName", value)
	return p
}

/**
 * 是否隐藏
 */
func (p *Page) Hidden(value bool) *Page {
	p.Set("hidden", value)
	return p
}

/**
 * 是否隐藏表达式
 */
func (p *Page) HiddenOn(value string) *Page {
	p.Set("hiddenOn", value)
	return p
}

/**
 * 组件唯一 id，主要用于日志采集
 */
func (p *Page) Id(value string) *Page {
	p.Set("id", value)
	return p
}

/**
 * 页面初始化的时候，可以设置一个 API 让其取拉取，发送数据会携带当前 data 数据（包含地址栏参数），获取得数据会合并到 data 中，供组件内使用。
 */
func (p *Page) InitApi(value string) *Page {
	p.Set("initApi", value)
	return p
}

/**
 * 是否默认就拉取？
 */
func (p *Page) InitFetch(value bool) *Page {
	p.Set("initFetch", value)
	return p
}

/**
 * 是否默认就拉取表达式
 */
func (p *Page) InitFetchOn(value string) *Page {
	p.Set("initFetchOn", value)
	return p
}

/**
 * 配置轮询间隔，配置后 initApi 将轮询加载。
 */
func (p *Page) Interval(value string) *Page {
	p.Set("interval", value)
	return p
}

func (p *Page) LoadingConfig(value string) *Page {
	p.Set("loadingConfig", value)
	return p
}

func (p *Page) Messages(value string) *Page {
	p.Set("messages", value)
	return p
}

/**
 * 移动端下的样式表
 */
func (p *Page) MobileCSS(value string) *Page {
	p.Set("mobileCSS", value)
	return p
}

func (p *Page) Name(value string) *Page {
	p.Set("name", value)
	return p
}

/**
 * 事件动作配置
 */
func (p *Page) OnEvent(value string) *Page {
	p.Set("onEvent", value)
	return p
}

/**
 * 下拉刷新配置
 */
func (p *Page) PullRefresh(value string) *Page {
	p.Set("pullRefresh", value)
	return p
}

/**
 * 默认不设置自动感觉内容来决定要不要展示这些区域 如果配置了，以配置为主。
 */
func (p *Page) Regions(value string) *Page {
	p.Set("regions", value)
	return p
}

/**
 * 页面描述, 标题旁边会出现个小图标，放上去会显示这个属性配置的内容。
 */
func (p *Page) Remark(value string) *Page {
	p.Set("remark", value)
	return p
}

/**
 * 是否显示错误信息，默认是显示的。
 */
func (p *Page) ShowErrorMsg(value bool) *Page {
	p.Set("showErrorMsg", value)
	return p
}

/**
 * 是否要静默加载，也就是说不显示进度
 */
func (p *Page) SilentPolling(value bool) *Page {
	p.Set("silentPolling", value)
	return p
}

/**
 * 是否静态展示
 */
func (p *Page) Static(value bool) *Page {
	p.Set("static", value)
	return p
}

/**
 * 静态展示表单项类名
 */
func (p *Page) StaticClassName(value string) *Page {
	p.Set("staticClassName", value)
	return p
}

/**
 * 静态展示表单项Value类名
 */
func (p *Page) StaticInputClassName(value string) *Page {
	p.Set("staticInputClassName", value)
	return p
}

/**
 * 静态展示表单项Label类名
 */
func (p *Page) StaticLabelClassName(value string) *Page {
	p.Set("staticLabelClassName", value)
	return p
}

/**
 * 是否静态展示表达式
 */
func (p *Page) StaticOn(value string) *Page {
	p.Set("staticOn", value)
	return p
}

/**
 * 静态展示空值占位
 */
func (p *Page) StaticPlaceholder(value string) *Page {
	p.Set("staticPlaceholder", value)
	return p
}

func (p *Page) StaticSchema(value string) *Page {
	p.Set("staticSchema", value)
	return p
}

/**
 * 配置停止轮询的条件。
 */
func (p *Page) StopAutoRefreshWhen(value string) *Page {
	p.Set("stopAutoRefreshWhen", value)
	return p
}

/**
 * 自定义样式
 */
func (p *Page) Style(value string) *Page {
	p.Set("style", value)
	return p
}

/**
 * 页面副标题
 */
func (p *Page) SubTitle(value string) *Page {
	p.Set("subTitle", value)
	return p
}

/**
 * 页面标题
 */
func (p *Page) Title(value string) *Page {
	p.Set("title", value)
	return p
}

/**
 * 页面顶部区域，当存在 title 时在右上角显示。
 */
func (p *Page) Toolbar(value string) *Page {
	p.Set("toolbar", value)
	return p
}

/**
 * 配置 toolbar 容器 className
 */
func (p *Page) ToolbarClassName(value string) *Page {
	p.Set("toolbarClassName", value)
	return p
}

/**
 * 指定为 page 渲染器。
 */
func (p *Page) Type(value string) *Page {
	p.Set("type", value)
	return p
}

/**
 * 可以组件级别用来关闭移动端样式
 */
func (p *Page) UseMobileUI(value bool) *Page {
	p.Set("useMobileUI", value)
	return p
}

/**
 * 是否显示
 */
func (p *Page) Visible(value bool) *Page {
	p.Set("visible", value)
	return p
}

/**
 * 是否显示表达式
 */
func (p *Page) VisibleOn(value string) *Page {
	p.Set("visibleOn", value)
	return p
}

