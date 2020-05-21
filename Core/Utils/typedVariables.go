package Utils

import (
	"fmt"
	"strings"
)

// !!NOT IN USE YET!! Alpha code of declaring inline arguments for Actions

const (
	HCSelf      = "self:"
	HCRef       = "sref:"            // Declaration of a value from another data entry (name of entry is reference)
	HCString    = "sstring:"
	HCInt       = "sint:"
	HCFloat     = "sfloat64:"
	
)




func AnalyzeAction(method string, in interface{}, typ string) interface{}{
	
	
	function := strings.Split(method, "(")
	var functionName string
	
	if len(function) < 2 {
		fmt.Println("No arguments found or missing opening parenthesis '(' !")
	} else {
		functionName = function[0]
		function = strings.Split(function[1], ")")
		if len(function) <  1 {
			fmt.Println("No arguments found")
			panic("Improperly encoded action! Missing closing parenthesis ')' !")
		}
	}
	
	functions := strings.Split(function[0], ",")
	

	fmt.Println(functions)
	
	
	
	
	
	var temp  []interface{}
	for i := range functions {
		arg := strings.Split(functions[i], "#")
		if arg[0] == "self" {
			temp = append(temp, in)
			continue
		}
	
		
		x, _ := QuickConvert[arg[0]](arg[1], "string")
		temp = append(temp, x)
		
		
	}

	output, err := FindFunction[len(temp)](temp, functionName)
	
	if err != nil {
		panic("Something went wrong!")
	}
	
	
	fmt.Println(output)
	
	return nil
}
