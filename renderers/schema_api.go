package renderers

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

func (a *SchemaApi) AttachDataToQuery(value interface{}) *SchemaApi {
	a.Set("attachDataToQuery", value)
	return a
}

func (a *SchemaApi) AutoRefresh(value interface{}) *SchemaApi {
	a.Set("autoRefresh", value)
	return a
}

func (a *SchemaApi) Cache(value interface{}) *SchemaApi {
	a.Set("cache", value)
	return a
}

func (a *SchemaApi) ConvertKeyToPath(value interface{}) *SchemaApi {
	a.Set("convertKeyToPath", value)
	return a
}

func (a *SchemaApi) Data(value interface{}) *SchemaApi {
	a.Set("data", value)
	return a
}

func (a *SchemaApi) DataType(value interface{}) *SchemaApi {
	a.Set("dataType", value)
	return a
}

func (a *SchemaApi) DownloadFileName(value interface{}) *SchemaApi {
	a.Set("downloadFileName", value)
	return a
}

func (a *SchemaApi) ForceAppendDataToQuery(value interface{}) *SchemaApi {
	a.Set("forceAppendDataToQuery", value)
	return a
}

func (a *SchemaApi) Headers(value interface{}) *SchemaApi {
	a.Set("headers", value)
	return a
}

func (a *SchemaApi) Method(value interface{}) *SchemaApi {
	a.Set("method", value)
	return a
}

func (a *SchemaApi) QsOptions(value interface{}) *SchemaApi {
	a.Set("qsOptions", value)
	return a
}

func (a *SchemaApi) ReplaceData(value interface{}) *SchemaApi {
	a.Set("replaceData", value)
	return a
}

func (a *SchemaApi) ResponseData(value interface{}) *SchemaApi {
	a.Set("responseData", value)
	return a
}

func (a *SchemaApi) ResponseType(value interface{}) *SchemaApi {
	a.Set("responseType", value)
	return a
}

func (a *SchemaApi) SendOn(value interface{}) *SchemaApi {
	a.Set("sendOn", value)
	return a
}

func (a *SchemaApi) Silent(value interface{}) *SchemaApi {
	a.Set("silent", value)
	return a
}

func (a *SchemaApi) TrackExpression(value interface{}) *SchemaApi {
	a.Set("trackExpression", value)
	return a
}

func (a *SchemaApi) Url(value interface{}) *SchemaApi {
	a.Set("url", value)
	return a
}
