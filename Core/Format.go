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
	//file, err := os.Open("./Formats/Default.json")
	file, err := os.Open("./Data/Formats/Default.json")
	
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



func (F Format)SetFormat() {
	var target string = ""
	var dtype string = ""
	var dName string = ""
	var AddAct string = ""
	var AddT string = ""
	var Action string = ""
	var ActArg string = ""
	
	// Sets name of the format
	fmt.Printf(Utils.FormatName)
	fmt.Printf(Utils.FmtNamePS)
	fmt.Scan(&dName)
	
	if dName == "" {
		panic("Bad Name")
	}
	
	F[dName] = map[string] interface{}{
		"Target" : map[string] interface{}{},
		"Actions" : map[string] interface{}{},
	}
	
	
	EntryPrompt:
	// Sets the Target of the format
	fmt.Printf(Utils.TargetHelp)
	fmt.Printf(Utils.TargetPS)
	fmt.Scan(&target)
	
	
	// Sets the type of the target
	fmt.Printf(Utils.TypeHelp)
	fmt.Printf(Utils.TypePS)
	fmt.Scan(&dtype)
	
	// Asks if actions are required
	fmt.Printf(Utils.ActionsScreen)
	fmt.Printf(Utils.AddActionPS)
	fmt.Scan(&AddAct)
	
	if AddAct != "yes" && AddAct == "no" {
		goto AddMore
	}
	
	
	
	// Sets the actions for target
	fmt.Printf(Utils.ActionsMain)
	PromptAction:
		fmt.Printf(Utils.ActionsPS)
		fmt.Scan( &Action)
		fmt.Println(Action, ActArg)
		if ActArg != "" {
			Action = fmt.Sprintf("%s %s", Action, ActArg)
		}
	
		if Utils.CheckActionsHelp(Action) == false {
			goto PromptAction
		} else {
		
		}
	
	
	
	
	
	AddMore:
	fmt.Printf(Utils.AddMore)
	fmt.Printf(Utils.AdditonalT)
	fmt.Scan(&AddT)
	
	if AddT == "yes" {
		goto EntryPrompt
	} else {
	
	}
	
	

	
	
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









