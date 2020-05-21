package Actions

import "fmt"

type ApplicationFunc struct {
	Name            string
	Args            []string
	Description     string

}


var AvailableFunctions  map[string]ApplicationFunc =  map[string]ApplicationFunc{
	"sumx": {
		Name: "sumx",
		Args: []string{"sint", "sint"},
		Description: "Sum two integers together",
	},
	"send2Ip": {
		Name: "send2Ip",
		Args: []string{"none", "sstring"},
		Description: "Send json data to an ip address",
	},
	"save2File": {
		Name: "save2File",
		Args: []string{"none", "sstring"},
		Description: "Store file in the data folder",
	},
	
}


func  (A ApplicationFunc) CheckTypes(data interface{}, index int) {

}


func (A ApplicationFunc) CheckCompatibility(data interface{}) error{
	
	if len(data.([]interface{})) != len(A.Args) {
		return fmt.Errorf("The lenght of supplied args (%d) is not the same as the function (%d)\n",len(data.([]interface{})), len(A.Args) )
	}
	
	return nil
}




		///////////////////////////// RUN FUNCTIONS /////////////////////////////

func (A ApplicationFunc) Run2Args(data interface{}, data2 interface{}) (interface{}, error) {
	
	switch A.Name {
	case "sumx":
		return Sumx(data, data2), nil
	case "send2IP":
		return Sumx(data, data2), nil
	case "save2File":
		return Sumx(data, data2), nil
	default:
	
	}
	return nil, fmt.Errorf("Error in Run2Args()\n")
}


func (A ApplicationFunc) Run3Args(data interface{}, data2 interface{}, data3 interface{}) (interface{}, error) {
	
	switch A.Name {
	
	default:

	}
	return nil, fmt.Errorf("Error in Run3Args()\n")
}

func (A ApplicationFunc) Run4Args(data interface{}, data2 interface{}, data3 interface{}, data4 interface{}) (interface{}, error) {
	
	switch A.Name {
	
	default:

	}
	return nil, fmt.Errorf("Error in Run4Args()\n")
}