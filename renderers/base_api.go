package renderers

type BaseApi struct {
	*BaseRenderer
}

func NewBaseApi() *BaseApi {
	a := &BaseApi{
		BaseRenderer: NewBaseRenderer(),
	}

	return a
}

func (a *BaseApi) Set(name string, value interface{}) *BaseApi {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	a.AmisSchema[name] = value
	return a
}

func (a *BaseApi) AttachDataToQuery(value interface{}) *BaseApi {
	a.Set("attachDataToQuery", value)
	return a
}

func (a *BaseApi) AutoRefresh(value interface{}) *BaseApi {
	a.Set("autoRefresh", value)
	return a
}

func (a *BaseApi) Cache(value interface{}) *BaseApi {
	a.Set("cache", value)
	return a
}

func (a *BaseApi) ConcatDataFields(value interface{}) *BaseApi {
	a.Set("concatDataFields", value)
	return a
}

func (a *BaseApi) ConvertKeyToPath(value interface{}) *BaseApi {
	a.Set("convertKeyToPath", value)
	return a
}

func (a *BaseApi) Data(value interface{}) *BaseApi {
	a.Set("data", value)
	return a
}

func (a *BaseApi) DataType(value interface{}) *BaseApi {
	a.Set("dataType", value)
	return a
}

func (a *BaseApi) ForceAppendDataToQuery(value interface{}) *BaseApi {
	a.Set("forceAppendDataToQuery", value)
	return a
}

func (a *BaseApi) Headers(value interface{}) *BaseApi {
	a.Set("headers", value)
	return a
}

func (a *BaseApi) Messages(value interface{}) *BaseApi {
	a.Set("messages", value)
	return a
}

func (a *BaseApi) Method(value interface{}) *BaseApi {
	a.Set("method", value)
	return a
}

func (a *BaseApi) QsOptions(value interface{}) *BaseApi {
	a.Set("qsOptions", value)
	return a
}

func (a *BaseApi) ReplaceData(value interface{}) *BaseApi {
	a.Set("replaceData", value)
	return a
}

func (a *BaseApi) ResponseData(value interface{}) *BaseApi {
	a.Set("responseData", value)
	return a
}

func (a *BaseApi) ResponseType(value interface{}) *BaseApi {
	a.Set("responseType", value)
	return a
}

func (a *BaseApi) SendOn(value interface{}) *BaseApi {
	a.Set("sendOn", value)
	return a
}

func (a *BaseApi) Silent(value interface{}) *BaseApi {
	a.Set("silent", value)
	return a
}

func (a *BaseApi) TrackExpression(value interface{}) *BaseApi {
	a.Set("trackExpression", value)
	return a
}

func (a *BaseApi) Url(value interface{}) *BaseApi {
	a.Set("url", value)
	return a
}
