package Core

import (
	"SmartPolice-Interface/Core/Actions"
	"fmt"
	"strings"
)

// Curated recipes  for handling data after acquisition


func ReadChain(method string, value interface{}, typ string, ) (interface{}, error) {
	splitter := strings.Split(method, ".")
	var x interface{} = value
	var err error
	for i := range splitter {
		fmt.Println(splitter[i])
		x, err  = ReadAction(splitter[i], x, typ)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}


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
		case "sum":
				if x, err := Actions.Sum(value); err != nil {
					fmt.Println("x", x)
					return x, err
					
					
				} else {
					
					return x, nil
				}
			break
	
		case "times":
			if x, err := Actions.Times(value); err != nil {
				fmt.Println("x", x)
				return x, err
				
				
			} else {
				
				return x, nil
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

