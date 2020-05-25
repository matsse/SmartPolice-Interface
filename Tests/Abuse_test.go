package Tests

import (
	"SmartPolice-Interface/Core/Actions"
	"SmartPolice-Interface/Core/Utils"
	"strings"
	"testing"
	"fmt"
)

func TestSyscall(t *testing.T) {
	fmt.Printf("\n\n\nAbused input Testing\n")
	
	var x  = map[string] interface{}{
		"Temperature" :             "key:temperature/type:float64",
		"Temperature.Action" :      "action:os.Remove(sstring#main.go)",
	}
	
	
	
	varInfo := strings.Split(x["Temperature"].(string), "/")
	
	identifier := strings.Split(varInfo[0], ":")[1]                 // Get test identifier key
	varType := strings.Split(varInfo[1], ":")[1]                    // Get test data type
	variable, _ := Actions.Convert2int[varType](343.3, varType)     // convert a float64 to int
	
	fmt.Println(variable, identifier, varType)
	
	ActInfo := strings.Split(x["Temperature.Action"].(string), ":")
	
	
	
	
	_ = Utils.AnalyzeAction(ActInfo[1], variable, varType)
	
	
	
}

