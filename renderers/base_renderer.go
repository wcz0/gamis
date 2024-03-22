package renderers

import (
	"encoding/json"
)

type BaseRenderer struct {
	Type       string
	AmisSchema map[string]interface{}
}

func NewBaseRenderer() *BaseRenderer {
	return &BaseRenderer{
		AmisSchema: make(map[string]interface{}),
	}
}

func (b *BaseRenderer) Set(name string, value interface{}) *BaseRenderer {
	if name == "map" {
		if v, ok := value.([]interface{}); ok && isArrayOfArrays(v) {
			value = mapOfArrays(v)
		}
	}

	b.AmisSchema[name] = value
	return b
}

func (b *BaseRenderer) ToJson() (string, error) {
	bytes, err := json.Marshal(b.AmisSchema)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (b *BaseRenderer) ToArray() map[string]interface{} {
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
	m := make(map[string]interface{})
	for _, item := range v {
		if arr, ok := item.([]interface{}); ok {
			for i, a := range arr {
				m[string(i)] = a
			}
		}
	}
	return m
}