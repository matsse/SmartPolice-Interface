package Tests

import (
	"SmartPolice-Interface/Core"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)





func TestTypes(t *testing.T) {
	fmt.Printf("Structure testing \n\n\n")
	
	
	var x  = map[string] interface{}{
		
		// All variables that should be found
		"value1" :          "key:value1/type:int",
		"value2" :          "key:value2/type:int8",
		"value3" :          "key:value3/type:int16",
		"value4" :          "key:value4/type:int32",
		"value5" :          "key:value5/type:int64",
		"value6" :          "key:value6/type:float32",
		"value7" :          "key:value7/type:float64",
		"value8" :          "key:value8/type:string",
		"value9" :          "key:value9/type:byte",
		
		
		// All variables that are missing
		"value10" :         "key:value10/type:int8",
		"value11" :         "key:missing2/type:float32",
		"value12" :         "key:desire/type:[]int",
		"value13" :         "key:scarlet/type:[]string",
		
		
	}
	
	
	var testJson map[string] interface{} = map[string] interface{}{}
	
	var result map[string] interface{} = map[string] interface{}{}
	
	file, err :=  os.Open("./files/Types.json")
	if err != nil {
		panic(err)
	}
	
	
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &testJson)
	
	
	for queries := range x {
		if analyFormat := Core.AnalyzeType( x[queries].(string), testJson); analyFormat == true {
			result[queries] = "True"
		} else {
		
		}
	}
	
	fmt.Println(result)
	
	
	
	
}


func TestPrimitives(t *testing.T) {
	fmt.Printf("\n\n\nPrimitive testing\n")

	var x  = map[string] interface{}{
		
		// All variables that should be found
		"Temperature" :     "key:temperature/type:int",
		"Humidity" :        "key:humidity/type:int",
		"Dew Point" :       "key:dew point/type:float32",
		"LDR" :             "key:ldr/type:int",
		"Message" :         "key:message/type:string",
		"Latitude" :        "key:latitude/type:float32",
		"Colors" :          "key:colors/type:[]int",
		"Messages" :        "key:messages/type:[]string",
		
		
		// All variables that are missing
		
		"Missing" :         "key:missing/type:string",
		"Missing2" :        "key:missing2/type:float32",
		"Desire" :          "key:desire/type:[]int",
		"Scarlet" :        "key:scarlet/type:[]string",
		
		
	}
	
	
	var testJson map[string] interface{} = map[string] interface{}{}
	
	var result map[string] interface{} = map[string] interface{}{}
	
	file, err :=  os.Open("./files/Primitives.json")
	if err != nil {
		panic(err)
	}
	
	
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &testJson)
	
	for queries := range x {
		result[queries] = "missing"
	}
	
	
	
	for queries := range x {
		if analyFormat := Core.AnalyzeFormat( x[queries].(string), testJson); analyFormat == true {
			result[queries] = "found"
		}
	}
	
	fmt.Println(result)
	
	
}



func TestPaths(t *testing.T) {
	
	fmt.Printf("\n\n\nPath testing\n")
	
	var x  = map[string] interface{}{
		
		// All variables that should be found
		"Temperature" :     "path:layer1.temperature/type:int",
		"Humidity" :        "path:layer1.humidity/type:int",
		"Message" :         "path:message feed.Tom.Saturday.message/type:string",
		"Latitude" :        "path:Coordinates.new folder.continue.getting closer.Warmer.warmer.latitude/type:float32",
		"Messages" :        "path:message feed.Tom.Friday.unnamed folder.assorted messages.messages/type:[]string",
		
		
		// All variables that are missing
		
		"Missing" :         "key:missing/type:string",
		"Scarlet" :        "key:scarlet/type:[]string",
		
		
	}
	
	var testJson map[string] interface{} = map[string] interface{}{}
	
	var result map[string] interface{} = map[string] interface{}{}
	
	file, err :=  os.Open("./files/Paths.json")
	if err != nil {
		panic(err)
	}
	
	
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &testJson)
	
	for queries := range x {
		result[queries] = "missing"
	}
	
	
	for queries := range x {
		if analyFormat := Core.AnalyzeFormat( x[queries].(string), testJson); analyFormat == true {
			result[queries] = "found"
		}
	}
	
	fmt.Println(result)
	
}

