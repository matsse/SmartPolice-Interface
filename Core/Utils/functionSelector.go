package Utils

import (
	"SmartPolice-Interface/Core/Actions"
)

var FindFunction map[int]  func( in []interface{}, name string) (interface{}, error) = map[int] func(in []interface{}, name string) (interface{}, error) {
	2 : func( in []interface{}, name string) (interface{}, error) {
		//fmt.Println(name)
		compatibility := Actions.AvailableFunctions[name].CheckCompatibility(in)
		if compatibility != nil {
			panic(compatibility)
		}

		return Actions.AvailableFunctions[name].Run2Args(in[0], in[1])
	},

	3 : func( in []interface{}, name string) (interface{}, error) {
		// TODO add functions that takes 3 args
		return nil, nil
	},
	4 : func( in []interface{}, name string) (interface{}, error) {
		//fmt.Println(name)
		compatibility := Actions.AvailableFunctions[name].CheckCompatibility(in)
		if compatibility != nil {
			panic(compatibility)
		}
		return Actions.AvailableFunctions[name].Run4Args(in[0], in[1], in[2], in[3])
	},
	
}
