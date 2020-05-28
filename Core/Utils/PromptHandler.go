package Utils

import (
	"SmartPolice-Interface/Core/Actions"
	"fmt"
	"log"
	"strings"
)

func CheckActionsHelp(Action string) bool{

	switch Action {
	case  "#help chains":
		fmt.Printf(ActionsHelp)
		return true
	case  "#help actions":
		fmt.Printf(ChainsHelp)
		return true
	case  "#help list":
		fmt.Printf("Not implemented")
		return true
	case  "#help args":
		fmt.Printf(ActTypessHelp)
		return true
	default:
		fmt.Println("Validating actions")
		ValidateActions(Action)
		}
		
		return false
}

func ValidateActions(Actions string) bool {
	
	// action:DecryptAES(self#,sref#key,sref#iv,sref#blocksize)
	
	split := strings.Split(Actions, ":")
	var key string = split[0]
	var rest string = split[1]
	var valid bool
	
	
	if !strings.Contains(key, "action") && !strings.Contains(key, "chain") {
		return false
	}

	// DecryptAES(self#,sref#key,sref#iv,sref#blocksize)
	functions := strings.Split(rest, ".")
	if len(functions) > 1 {
		valid = ChainsValid(functions)
	} else {
		valid = ActionValid(rest)
	}

	
	if !valid {
		
		return false
	}


	
	
	return true
}

func ChainsValid(chain []string) bool{
	for i := range chain {
		bChainvalid := ActionValid(chain[i])
		if bChainvalid == false {
			return false
		}
	}
	
	
	
	
	return true
}


func ActionValid(action string) bool{
	//DecryptAES(self#,sref#key,sref#iv,sref#blocksize)
	
	OpenParaen := strings.Split(action, "(")
	if Actions.AvailableFunctions[OpenParaen[0]].Name == ""  {
		log.Println("Action not found!")
		return false
	}
	// self#,sref#key,sref#iv,sref#blocksize)
	
	
	rest := strings.Split(OpenParaen[1], ")")	 // self#,sref#key,sref#iv,sref#blocksize
	args :=  strings.Split(rest[0], ",")
	
	if len(Actions.AvailableFunctions[OpenParaen[0]].Args) != len(args) {
		log.Println("Argument count mismatch!")
		return false
	}
	
	
	for i := range args {
		arg := strings.Split(rest[i], "#")
		fmt.Println(rest)
		if Actions.AvailableFunctions[OpenParaen[0]].Args[i] != arg[0] {
			fmt.Println(arg[0])
			log.Println("Argument mismatch!")
			return false
		}
		
	}
	
	
	

	
	
	
	
	
	return true
}