package Utils

import "strconv"

var QuickConvert map[string]  func( in interface{}, typ string) (interface{}, error) =  map[string] func(in interface{}, typ string) (interface{}, error) {
	// When converted types must be reconverted
	"sint" : func(in interface{}, typ string) (interface{}, error) {
		switch typ {
			case "string":
				x, _ := strconv.Atoi(in.(string))
				return  x, nil
			break
		}
		return nil, nil
	},
	"sstring" : func(in interface{}, typ string) (interface{}, error) {
		switch typ {
		case "string":
			x, _ := strconv.Atoi(in.(string))
			return  x, nil
			break
		}
		return nil, nil
	},
}
