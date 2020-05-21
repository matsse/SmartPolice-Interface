package Utils

import (
	"SmartPolice-Interface/Core/Actions"
	"fmt"
)

var FindFunction map[int]  func( in []interface{}, name string) (interface{}, error) = map[int] func(in []interface{}, name string) (interface{}, error) {
	2 : func( in []interface{}, name string) (interface{}, error) {
		fmt.Println(name)
		x:= Actions.Sumx(in[0], in[1])
		return x, nil
	},
	3 : func( in []interface{}, name string) (interface{}, error) {
		// TODO add functions that takes 3 args
		return nil, nil
	},
	4 : func( in []interface{}, name string) (interface{}, error) {
		// TODO add functions that takes 4 args
		return nil, nil
	},
	
}
