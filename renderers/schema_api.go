package renderers


/**

 * @author wcz0
 * @version 6.2.2
 */
type SchemaApi struct {
	*BaseRenderer
}

func NewSchemaApi() *SchemaApi {
    a := &SchemaApi{
        BaseRenderer: NewBaseRenderer(),
    }

    return a
}


func (a *SchemaApi) Set(name string, value interface{}) *SchemaApi {
    if name == "map" {
        if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
            value = mapOfArrays(v)
        }
    }
    a.AmisSchema[name] = value
    return a
}
/**
 * 用来做接口返回的数据映射。
 */
func (a *SchemaApi) ResponseData(value interface{}) *SchemaApi {
    a.Set("responseData", value)
    return a
}

/**
 * 设置发送条件
 */
func (a *SchemaApi) SendOn(value interface{}) *SchemaApi {
    a.Set("sendOn", value)
    return a
}

/**
 * autoFill 是否显示自动填充错误提示
 */
func (a *SchemaApi) Silent(value interface{}) *SchemaApi {
    a.Set("silent", value)
    return a
}

/**
 * 发送体的格式
 * 可选值: json | form-data | form
 */
func (a *SchemaApi) DataType(value interface{}) *SchemaApi {
    a.Set("dataType", value)
    return a
}

/**
 * 如果是文件下载接口，请配置这个。
 */
func (a *SchemaApi) ResponseType(value interface{}) *SchemaApi {
    a.Set("responseType", value)
    return a
}

/**
 * 默认都是追加模式，如果想完全替换把这个配置成 true
 */
func (a *SchemaApi) ReplaceData(value interface{}) *SchemaApi {
    a.Set("replaceData", value)
    return a
}

/**
 * 强制将数据附加在 query，默认只有 api 地址中没有用变量的时候 crud 查询接口才会 自动附加数据到 query 部分，如果想强制附加请设置这个属性。 对于那种临时加了个变量但是又不想全部参数写一遍的时候配置很有用。
 */
func (a *SchemaApi) ForceAppendDataToQuery(value interface{}) *SchemaApi {
    a.Set("forceAppendDataToQuery", value)
    return a
}

/**
 * qs 配置项
 */
func (a *SchemaApi) QsOptions(value interface{}) *SchemaApi {
    a.Set("qsOptions", value)
    return a
}

/**
 * 文件下载时，指定文件名
 */
func (a *SchemaApi) DownloadFileName(value interface{}) *SchemaApi {
    a.Set("downloadFileName", value)
    return a
}

/**
 * API 发送类型
 * 可选值: get | post | put | delete | patch | jsonp | js
 */
func (a *SchemaApi) Method(value interface{}) *SchemaApi {
    a.Set("method", value)
    return a
}

/**
 * API 发送目标地址
 */
func (a *SchemaApi) Url(value interface{}) *SchemaApi {
    a.Set("url", value)
    return a
}

/**
 * 如果 method 为 get 的接口，设置了 data 信息。 默认 data 会自动附带在 query 里面发送给后端。如果想通过 body 发送给后端，那么请把这个配置成 false。但是，浏览器还不支持啊，设置了只是摆设。除非服务端支持 method-override
 */
func (a *SchemaApi) AttachDataToQuery(value interface{}) *SchemaApi {
    a.Set("attachDataToQuery", value)
    return a
}

/**
 * 是否自动刷新，当 url 中的取值结果变化时，自动刷新数据。
 */
func (a *SchemaApi) AutoRefresh(value interface{}) *SchemaApi {
    a.Set("autoRefresh", value)
    return a
}

/**
 * 当开启自动刷新的时候，默认是 api 的 url 来自动跟踪变量变化的。 如果你希望监控 url 外的变量，请配置 trackExpression。
 */
func (a *SchemaApi) TrackExpression(value interface{}) *SchemaApi {
    a.Set("trackExpression", value)
    return a
}

/**
 * 如果设置了值，同一个接口，相同参数，指定的时间（单位：ms）内请求将直接走缓存。
 */
func (a *SchemaApi) Cache(value interface{}) *SchemaApi {
    a.Set("cache", value)
    return a
}

/**
 * 用来控制携带数据. 当key 为 `&` 值为 `$$` 时, 将所有原始数据打平设置到 data 中. 当值为 $$ 将所有原始数据赋值到对应的 key 中. 当值为 $ 打头时, 将变量值设置到 key 中.
 */
func (a *SchemaApi) Data(value interface{}) *SchemaApi {
    a.Set("data", value)
    return a
}

/**
 * 默认数据映射中的key如果带点，或者带大括号，会转成对象比如：{   'a.b': '123' }经过数据映射后变成 {  a: {   b: '123  } }如果想要关闭此功能，请设置 convertKeyToPath 为 false
 */
func (a *SchemaApi) ConvertKeyToPath(value interface{}) *SchemaApi {
    a.Set("convertKeyToPath", value)
    return a
}

/**
 * 携带 headers，用法和 data 一样，可以用变量。
 */
func (a *SchemaApi) Headers(value interface{}) *SchemaApi {
    a.Set("headers", value)
    return a
}
