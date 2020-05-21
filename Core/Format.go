package Core

import (
	"SmartPolice-Interface/Core/Utils"
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)



var Formats map[string] interface{}

type Format  map[string] interface{}

//
//type Format struct {
//	Entries map[string] interface{}     `json:"entries"`
//}



func FormatScreen() {

}



func LoadFormats() {
	file, err := os.Open("./Formats/Default.json")
	if err != nil {
		log.Fatal(err)
	}
	
	formats, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	
	
	json2.Unmarshal(formats, &Formats)
	
	
	fmt.Println(Formats)
	
	
	
}



func (F *Format)SetFormat() {
	var target string = ""
	var dtype string = ""
	fmt.Printf(`
	
	Specify the target of your data. It can be done with a key or a path
	
## PATH ##
	To specify a path, type the key value of the the path which leads to the right key.
	Denote each corresponding joints in the path with a period (.)
	
	Example:
		
		RootValue.Child1.Child2.Child3


## KEY ##
	To specify the target with a key will be reduce efficiency, as the program will recursively look for the value.
	However, after the value is found the first time, a path will be created to reduce time/complexity.


Target#`)
	fmt.Scan(&target)
	
	
	
	fmt.Printf(`
	
	Specify the type of your targetted data
	

	To specify the data type of your data entry, you can use the following options:
	
	## Primitives ##
	
		int = type:int
		uint16 = type:uint16
		float32 = type:float32
		float64 = type:float64
		string = type:string
		date = type:date
		interface = type:interface

	## Arrays ##
	
		int = type:[]int
		uint = type:[]uint16
		float32 = type:[]float32
		float64 = type:[]float64
		string = type:[]string
		interface = type:[]interface

Target#`)
	fmt.Scan(&dtype)
	
	
	fmt.Printf("%s/%s\n", target, dtype)
	
	
	fmt.Println()
	time.Sleep(time.Second * 10 )
	
}



func AnalyzeFormat(entry string, data map[string]interface{}) (interface{},  bool)  {
	format  := strings.Split(entry, "/")
	
	query := strings.Split(format[0], ":")
	
	if query[0] == "key" {
		if value, result := KeySearch(query[1], data); result == true {
			return value, true
		}
	} else if query[0] == "path" {
		
		if value, result := PathSearch(query[1], data); result == true {
			return value, true
		}
	}
	
	return nil, false
	
}



func AnalyzeType(entry string, data map[string]interface{})( string , bool)  {
	
	format  := strings.Split(entry, "/")
	
	query := strings.Split(format[0], ":")
	dataType := strings.Split(format[1], ":")
	var temp interface{}
	
	
	if query[0] == "key" {
		if value, result := KeySearch(query[1], data); result == true {
			temp = value
		} else {
			return "", false
		}
	} else if query[0] == "path" {
		if value, result := PathSearch(query[1], data); result == true {
			temp = value
		} else {
			return "", false
		}
	}
	
	
	
	//fmt.Println(query[1], temp)
	
	switch dataType[1] {
	case "string":
		return dataType[1],  Utils.String_Validation(temp)
		break
	case "int":
		return dataType[1], Utils.Integer_validation(temp)
		break
	case "int8":
		//fmt.Println(int8(temp.(float64)))
		return dataType[1], Utils.Integer8_validation(temp)
		break
	case "int16":
		return dataType[1], Utils.Integer16_validation(temp)
		break
	case "int32":
		return dataType[1], Utils.Integer32_validation(temp)
		break
	case "int64":
		return dataType[1], Utils.Integer64_validation(temp)
		break
	case "float32":
		return dataType[1], Utils.Float32_validation(temp)
		break
	case "float64":
		return dataType[1], Utils.Float64_validation(temp)
		break
	case "byte":
		return dataType[1], Utils.Byte_Validation(temp)
		break
		
	}
	
	
	return "", false

}


//func AnalyzeAction(formats string , data map[string]interface{}) (interface{}, bool) {
//	//format  := strings.Split(entry, "/")
//}



func PathSearch(formats string , data map[string]interface{}) (interface{}, bool)  {
	path := strings.Split(formats, ".")
	var value interface{}
	tData := data
	
	for loc := range path {
		value = tData[path[loc]]
		
		if reflect.TypeOf(value).String() == "map[string]interface {}" {
			tData = value.(map[string]interface{})
			return tData, true
		} else {
			fmt.Println("path value is ", value)
			
		}
	}
	return nil, false
}



func KeySearch(formats string , data map[string]interface{}) (interface{}, bool) {
	for key := range data {
		switch data[key].(type) {
		case map[string]interface{}:
			KeySearch(formats, data[key].(map[string]interface{}))
			break
		default:
			if key == formats {
				//fmt.Println("found", key, reflect.TypeOf(data[key]))
				return data[key], true
			}
			break
		}
	}
	return nil, false
}









