package Utils

import (
	"SmartPolice-Interface/Core/Actions"
	"fmt"
	"strings"
)
var GlobalTemp map[string] interface{}

// !!NOT IN USE YET!! Alpha code of declaring inline arguments for Actions

const  (
	HCSelf   string   = "self"
	HCRef    string   = "sref"            // Declaration of a value from another data entry (name of entry is reference)
	HCString string   = "sstring"
	HCInt    string   = "sint"
	HCFloat  string   = "sfloat64"
	
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
		if arg[0] == HCSelf {
			temp = append(temp, in)
			continue
		} else if arg[0] == HCRef {
			// Get Reference from temporary global repository G
			variable := GlobalTemp["Globals"].(map[string]interface{})[arg[1]]
			Actions.AvailableFunctions[functionName].CheckTypes(variable, i)
			//fmt.Println("xxxx" , arg[1])
			temp = append(temp, variable)
			
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
