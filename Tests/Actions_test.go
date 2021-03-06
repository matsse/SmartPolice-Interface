package Tests

import (
	"SmartPolice-Interface/Core/Actions"
	"SmartPolice-Interface/Core/Utils"
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




func TestHCTypes(t *testing.T) {
	fmt.Printf("\n\n\nHardcoded type testing\n")
	
	
	var x  = map[string] interface{}{
		"Temperature" :             "key:temperature/type:float64",
		"Temperature.Action" :      "action:sumx(self#,sint#32)",
	}
	
	

	varInfo := strings.Split(x["Temperature"].(string), "/")
	
	identifier := strings.Split(varInfo[0], ":")[1]                 // Get test identifier key
	varType := strings.Split(varInfo[1], ":")[1]                    // Get test data type
	variable, _ := Actions.Convert2int[varType](343.3, varType)     // convert a float64 to int
	
	fmt.Println(variable, identifier, varType)
	
	ActInfo := strings.Split(x["Temperature.Action"].(string), ":")
	
	
	
	
	_ = Utils.AnalyzeAction(ActInfo[1], variable, varType)
	
	
	
}

func Test4Input(t *testing.T) {
	fmt.Printf("\n\n\nDecryption Testing\n")
	

	Utils.GlobalTemp = map[string] interface{}{
		"Temperature" :             "key:temperature/type:string",
		"Temperature.Action" :      "action:DecryptAES(self#,sref#key,sref#iv,sref#blocksize)",
		
		"Globals" : map[string]  interface{} {
			// Global variables to use for encryption
			"key" : "12345678901234567890123456789012",
			"iv" : "1234567890123456",
			"blocksize" : 128,
		},
	}
	
	input := Actions.EncryptAES(`{
  "layer1" : {
    "temperature" : 32,
    "humidity" :  52
  },
}`, "12345678901234567890123456789012", "1234567890123456", 128)
	
	
	varInfo := strings.Split(Utils.GlobalTemp["Temperature"].(string), "/")
	_ = strings.Split(varInfo[0], ":")[1]                 // Get test identifier key
	varType := strings.Split(varInfo[1], ":")[1]                    // Get test data type
	//variable, _ := Actions.Convert2int[varType](343.3, varType)     // convert a float64 to int
	
	ActInfo := strings.Split(Utils.GlobalTemp["Temperature.Action"].(string), ":")
	
	
	
	_ = Utils.AnalyzeAction(ActInfo[1], input, varType)
	
	
}





