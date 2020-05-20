package Actions

import (
	"encoding/json"
	"os"
)

func SaveToFile(in interface{}, name string) (error){
	
	
	if _, err := os.Stat("../DataOutput/"+name+".json"); err == nil  {
		return err
	}
	file, err := os.Create("../DataOutput/"+name+".json")
	if err != nil {
		return err
	}
	defer file.Close()
	
	x, jerr := json.Marshal(in)
	if jerr != nil {
		return err
	}
	
	
	file.Write(x)
	
	
	
	return nil
}
