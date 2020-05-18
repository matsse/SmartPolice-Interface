package Core

import (
	"SmartPolice-Interface/Core/Actions"
	"fmt"
)

// Curated recipes  for handling data after acquisition





// Reads an input action command and dispatches the correct method
func ReadAction(method string, value interface{}, typ string, ) (interface{}, error) {
	switch method {
		case "convert2float64":
			if ok := Actions.Convert2float64[typ]; ok == nil  {
				fmt.Printf("The Convert2float64 is not supported on values of type %s\n", typ )
			} else {
				x, _ := ok(value, typ)
				
				return x, nil
				fmt.Println(x)
				
			}
			break
		case "convert2int":
			if ok := Actions.Convert2int[typ]; ok == nil  {
				fmt.Printf("The Convert2int is not supported on values of type %s\n", typ )
			} else {
				x, _ := ok(value, typ)
				
				return x, nil
				fmt.Println(x)
				
			}
			break
	}
	return nil, nil
}




// Runs an action
func RunAction() {

}

// Runs a chain of actions
func RunChain() {


}

