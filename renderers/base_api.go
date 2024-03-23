package renderers


/**

*/
type BaseApi struct {
	*BaseRenderer
}

func NewBaseApi() *BaseApi {
    a := &BaseApi{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}
/**
 * autoFill 是否显示自动填充错误提示
 */
func (a *BaseApi) Silent(value string) *BaseApi {
    a.Set("silent", value)
    return a
}

/**
 * 默认数据映射中的key如果带点，或者带大括号，会转成对象比如：{   'a.b': '123' }经过数据映射后变成 {  a: {   b: '123  } }如果想要关闭此功能，请设置 convertKeyToPath 为 false
 */
func (a *BaseApi) ConvertKeyToPath(value string) *BaseApi {
    a.Set("convertKeyToPath", value)
    return a
}

/**
 * 用来做接口返回的数据映射。
 */
func (a *BaseApi) ResponseData(value string) *BaseApi {
    a.Set("responseData", value)
    return a
}

/**
 * 默认都是追加模式，如果想完全替换把这个配置成 true
 */
func (a *BaseApi) ReplaceData(value string) *BaseApi {
    a.Set("replaceData", value)
    return a
}

/**
 * 强制将数据附加在 query，默认只有 api 地址中没有用变量的时候 crud 查询接口才会 自动附加数据到 query 部分，如果想强制附加请设置这个属性。 对于那种临时加了个变量但是又不想全部参数写一遍的时候配置很有用。
 */
func (a *BaseApi) ForceAppendDataToQuery(value string) *BaseApi {
    a.Set("forceAppendDataToQuery", value)
    return a
}

/**
 * qs 配置项
 */
func (a *BaseApi) QsOptions(value string) *BaseApi {
    a.Set("qsOptions", value)
    return a
}

/**
 * 设置发送条件
 */
func (a *BaseApi) SendOn(value string) *BaseApi {
    a.Set("sendOn", value)
    return a
}

/**
 * 当开启自动刷新的时候，默认是 api 的 url 来自动跟踪变量变化的。 如果你希望监控 url 外的变量，请配置 trackExpression。
 */
func (a *BaseApi) TrackExpression(value string) *BaseApi {
    a.Set("trackExpression", value)
    return a
}

/**
 * 提示信息
 */
func (a *BaseApi) Messages(value string) *BaseApi {
    a.Set("messages", value)
    return a
}

/**
 * API 发送类型
 * 可选值: get | post | put | delete | patch | jsonp | js
 */
func (a *BaseApi) Method(value string) *BaseApi {
    a.Set("method", value)
    return a
}

/**
 * 如果 method 为 get 的接口，设置了 data 信息。 默认 data 会自动附带在 query 里面发送给后端。如果想通过 body 发送给后端，那么请把这个配置成 false。但是，浏览器还不支持啊，设置了只是摆设。除非服务端支持 method-override
 */
func (a *BaseApi) AttachDataToQuery(value string) *BaseApi {
    a.Set("attachDataToQuery", value)
    return a
}

/**
 * 如果是文件下载接口，请配置这个。
 */
func (a *BaseApi) ResponseType(value string) *BaseApi {
    a.Set("responseType", value)
    return a
}

/**
 * 是否将两次返回的数据字段，做一个合并。配置返回对象中的字段名，支持配置多个。比如：同时返回 log 字段，第一次返回 {log: '1'}，第二次返回 {log: '2'}，合并后的结果是 {log: ['1', '2']]} 再比如：同时返回 items 字段，第一次返回 {items: [1, 2]}，第二次返回 {items: [3, 4]}，合并后的结果是 {items: [1, 2, 3, 4]}
 */
func (a *BaseApi) ConcatDataFields(value string) *BaseApi {
    a.Set("concatDataFields", value)
    return a
}

/**
 * 如果设置了值，同一个接口，相同参数，指定的时间（单位：ms）内请求将直接走缓存。
 */
func (a *BaseApi) Cache(value string) *BaseApi {
    a.Set("cache", value)
    return a
}

/**
 * API 发送目标地址
 */
func (a *BaseApi) Url(value string) *BaseApi {
    a.Set("url", value)
    return a
}

/**
 * 用来控制携带数据. 当key 为 `&` 值为 `$$` 时, 将所有原始数据打平设置到 data 中. 当值为 $$ 将所有原始数据赋值到对应的 key 中. 当值为 $ 打头时, 将变量值设置到 key 中.
 */
func (a *BaseApi) Data(value string) *BaseApi {
    a.Set("data", value)
    return a
}

/**
 * 发送体的格式
 * 可选值: json | form-data | form
 */
func (a *BaseApi) DataType(value string) *BaseApi {
    a.Set("dataType", value)
    return a
}

/**
 * 携带 headers，用法和 data 一样，可以用变量。
 */
func (a *BaseApi) Headers(value string) *BaseApi {
    a.Set("headers", value)
    return a
}

/**
 * 是否自动刷新，当 url 中的取值结果变化时，自动刷新数据。
 */
func (a *BaseApi) AutoRefresh(value string) *BaseApi {
    a.Set("autoRefresh", value)
    return a
}
