package Tests

import (
	"SmartPolice-Interface/Core/Actions"
	"strings"
	"testing"
	"SmartPolice-Interface/Core"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

func TestConversion(t *testing.T) {
	var x  = map[string] interface{}{
		
		// All variables that should be found
		"Temperature" :             "key:temperature/type:int",
		"Temperature.Action" :      "action:convert2float64",
		"Pressure" :                "key:pressure/type:float64",
		"Pressure.Action" :         "action:convert2int",
		// All variables that are missing
		
	}
	
	
	
	var testJson map[string] interface{} = map[string] interface{}{}
	
	//var result map[string] interface{} = map[string] interface{}{}
	
	file, err :=  os.Open("./files/Primitives.json")
	if err != nil {
		panic(err)
	}
	
	
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &testJson)
	
	
	// Int to float64
	value, returnf := Core.AnalyzeFormat( x["Temperature"].(string), testJson);
	if returnf == false {
		fmt.Println("Could not find key in file!")
	}
	
	dtype, returnf := Core.AnalyzeType( x["Temperature"].(string), testJson);
	if returnf == false {
		fmt.Println("Could not find key in file!")
	}
	
	
	action := strings.Split(x["Temperature.Action"].(string), ":")
	result, actErr := Core.ReadAction(action[1], value, dtype)
	if actErr != nil {
		panic(actErr)
	}
	
	 // Float64 to int
	 
	valuei, returni := Core.AnalyzeFormat( x["Pressure"].(string), testJson);
	if returni == false {
		fmt.Println("Could not find key in file!")
	}
	
	dtypei, returni := Core.AnalyzeType( x["Pressure"].(string), testJson);
	if returnf == false {
		fmt.Println("Could not find key in file!")
	}
	
	
	actioni := strings.Split(x["Pressure.Action"].(string), ":")
	resulti, actErri := Core.ReadAction(actioni[1], valuei, dtypei)
	if actErri != nil {
		panic(actErr)
	}
	
	fmt.Println("The converted ", dtype, "of", value, "is now", result)
	
	fmt.Println("The converted ", dtypei, "of", valuei, "is now", resulti)

	test2, _ :=  Actions.Convert2int[dtypei](value, "float64")
	
	
	test3, _ := Actions.Convert2float64["sint"](test2, "int")
	
	
	
	fmt.Println(50 + test2)
	
	fmt.Println(434.34 + test3)
	
}


func TestChain(t *testing.T) {
	var x  = map[string] interface{}{
		
		// All variables that should be found
		"Temperature" :             "key:temperature/type:float64",
		"Temperature.Action" :      "chain:convert2int.sum.times",
		// All variables that are missing
		
	}
	
	
	
	var testJson map[string] interface{} = map[string] interface{}{}
	
	//var result map[string] interface{} = map[string] interface{}{}
	
	file, err :=  os.Open("./files/Primitives.json")
	if err != nil {
		panic(err)
	}
	
	
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &testJson)
	
	
	value, returnf := Core.AnalyzeFormat( x["Temperature"].(string), testJson);
	if returnf == false {
		fmt.Println("Could not find key in file!")
	}
	
	dtype, returnf := Core.AnalyzeType( x["Temperature"].(string), testJson);
	if returnf == false {
		fmt.Println("Could not find key in file!")
	}
	
	action := strings.Split(x["Temperature.Action"].(string), ":")
	result, actErr := Core.ReadChain(action[1], value, dtype)
	if actErr != nil {
		panic(actErr)
	}
	
	fmt.Println(result)
	
	
	
	
}




