package Actions





// Implementing a map of

var Convert2float64 map[string]  func( in interface{}, typ string) (float64, error) =  map[string] func(in interface{}, typ string) (float64, error) {
	"int" : func(in interface{}, typ string) (float64, error) {
		out := float64(in.(float64))
		return out, nil
	},
	"int8" : func(in interface{}, typ string) (float64, error) {
		out := float64(in.(int8))
		return out, nil
	},
	"int16" : func(in interface{}, typ string) (float64, error) {
		out := float64(in.(int16))
		return out, nil
	},
	"int32" : func(in interface{}, typ string) (float64, error) {
		out := float64(in.(int32))
		return out, nil
	},
	"int64" : func(in interface{}, typ string) (float64, error) {
		out := float64(in.(int64))
		return out, nil
	},
	"float32" : func(in interface{}, typ string) (float64, error) {
		out := float64(in.(float32))
		return out, nil
	},
}



var Convert2int map[string]  func( in interface{}, typ string) (int , error) =  map[string] func(in interface{}, typ string) (int , error) {

	"int8" : func(in interface{}, typ string) (int, error) {
		out := int(int8(in.(float64)))
		return out, nil
	},
	"int16" : func(in interface{}, typ string) (int, error) {
		out := int(int16(in.(float64)))
		return out, nil
	},
	"int32" : func(in interface{}, typ string) (int, error) {
		out := int(int32(in.(float64)))
		return out, nil
	},
	"int64" : func(in interface{}, typ string) (int, error) {
		out := int(int64(in.(float64)))
		return out, nil
	},
	"float32" : func(in interface{}, typ string) (int, error) {
		out := int(float32(in.(float64)))
		return out, nil
	},
	"float64" : func(in interface{}, typ string) (int, error) {
		out := int(in.(float64))
		return out, nil
	},
}




var Convert2bytes map[string]  func( in interface{}, typ string) ([]byte , error) =  map[string] func(in interface{}, typ string) ([]byte , error) {
	"string" : func(in interface{}, typ string) ([]byte, error) {
		out := []byte(in.(string))
		return out, nil
	},
}




