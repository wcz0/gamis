package renderers

import (
	"encoding/json"
)

type BaseRenderer struct {
	Type       string `json:"type"`
	AmisSchema map[string]interface{}
}

func NewBaseRenderer() *BaseRenderer {
	return &BaseRenderer{
		AmisSchema: make(map[string]interface{}),
	}
}

func (b *BaseRenderer) ToJson() string {
	bytes, err := json.Marshal(b.AmisSchema)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (b *BaseRenderer) Set(name string, value interface{}) *BaseRenderer {
	// var obj  map[string]interface{}
	// err := json.Unmarshal([]byte(b.AmisSchema), &obj)
	// if err != nil {
	// 	panic(err)
	// }
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}
	// obj[name] = value
	// bytes, err := json.Marshal(obj)
	// if err != nil {
	// 	panic(err)
	// }
	// b.AmisSchema = string(bytes)
	b.AmisSchema[name] = value
	return b
}

func (b *BaseRenderer) ToArray() map[string]interface{} {
	// var obj  map[string]interface{}
	// err := json.Unmarshal([]byte(b.AmisSchema), &obj)
	// if err != nil {
	// 	panic(err)
	// }
	// return obj
	return b.AmisSchema
}

// isArrayOfArrays checks if a slice of interfaces contains only slices of interfaces.
func isArrayOfArrays(v []interface{}) bool {
	for _, item := range v {
		if _, ok := item.([]interface{}); !ok {
			return false
		}
	}
	return true
}

// mapOfArrays converts a slice of slices of interfaces to a map of interfaces.
func mapOfArrays(v []interface{}) map[string]interface{} {
	r := make(map[string]interface{})
	for _, v := range v {
		array := v.([]interface{})
		if len(array) >= 2 {
			key, ok1 := array[0].(string)
			if ok1 {
				r[key] = array[1]
			}
		}
	}
	return r
}

