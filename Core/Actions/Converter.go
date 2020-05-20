package Actions

import (
	"go/types"
)

// Implementing a map of

var Convert2float64 map[string]  func( in interface{}, typ string) (float64, error) =  map[string] func(in interface{}, typ string) (float64, error) {
	"int" : func(in interface{}, typ string) (float64, error) {
		return float64(int(in.(float64))), nil
	},
	"int8" : func(in interface{}, typ string) (float64, error) {
		return float64(int8(in.(float64))), nil
	},
	"int16" : func(in interface{}, typ string) (float64, error) {
		return float64(int16(in.(float64))), nil
	},
	"int32" : func(in interface{}, typ string) (float64, error) {
		return  float64(int32(in.(float64))), nil
	},
	"int64" : func(in interface{}, typ string) (float64, error) {
		return float64(int64(in.(float64))), nil
	},
	"float32" : func(in interface{}, typ string) (float64, error) {
		return  float64(float32(in.(float64))), nil
	},
	
	// When converted types must be reconverted
	"sint" : func(in interface{}, typ string) (float64, error) {

		return  float64(in.(int)), nil
	},
	"sint8" : func(in interface{}, typ string) (float64, error) {

		return float64(in.(int8)), nil
	},
	"sint16" : func(in interface{}, typ string) (float64, error) {
		return float64(in.(int16)), nil
	},
	"sint32" : func(in interface{}, typ string) (float64, error) {
		return float64(in.(int32)), nil
	},
	"sint64" : func(in interface{}, typ string) (float64, error) {
		return float64(in.(int64)), nil
	},
	"sfloat32" : func(in interface{}, typ string) (float64, error) {
		out := float64(in.(float32))
		return out, nil
	},
}



var Convert2int map[string]  func( in interface{}, typ string) (int , error) =  map[string] func(in interface{}, typ string) (int , error) {

	"int8" : func(in interface{}, typ string) (int, error) {
		return int(int8(in.(float64))), nil
	},
	"int16" : func(in interface{}, typ string) (int, error) {
		return int(int16(in.(float64))), nil
	},
	"int32" : func(in interface{}, typ string) (int, error) {
		return int(int32(in.(float64))), nil
	},
	"int64" : func(in interface{}, typ string) (int, error) {
		return int(int64(in.(float64))), nil
	},
	"float32" : func(in interface{}, typ string) (int, error) {
		return int(float32(in.(float64))), nil
	},
	"float64" : func(in interface{}, typ string) (int, error) {
		return int(in.(float64)), nil
	},
	
	// When converted types must be reconverted
	
	"sint" : func(in interface{}, typ string) (int, error) {
		return  int(in.(int)), nil
	},
	"sint8" : func(in interface{}, typ string) (int, error) {
		return int(in.(int8)), nil
	},
	"sint16" : func(in interface{}, typ string) (int, error) {
		return int(in.(int16)), nil
	},
	"sint32" : func(in interface{}, typ string) (int, error) {
		return int(in.(int32)), nil
	},
	"sint64" : func(in interface{}, typ string) (int, error) {
		return int(in.(int64)), nil
	},
	"sfloat32" : func(in interface{}, typ string) (int, error) {
		return int(in.(float32)), nil
	},
	
}




var Convert2bytes map[string]  func( in interface{}, typ string) ([]byte , error) =  map[string] func(in interface{}, typ string) ([]byte , error) {
	"string" : func(in interface{}, typ string) ([]byte, error) {
		out := []byte(in.(string))
		return out, nil
	},
}

var GetType map[string]  func() (types.Type , error) =  map[string] func() (types.Type , error) {
	"int" : func() (types.Type, error) {
		
		return types.Universe.Lookup("string").Type(), nil
	},
}






