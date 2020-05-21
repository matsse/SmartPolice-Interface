package Actions

import (
	"encoding/json"
	"os"
)

func SaveToFile(in interface{}, name string) (error){
	
	
	if _, err := os.Stat("../DataOutput/"+name+".json"); err == nil  {
		//log.Fatal("File already exist!")
		//return err
		os.Remove("../DataOutput/"+name+".json")

	}
	file, err := os.Create("../DataOutput/"+name+".json")
	if err != nil {
		return err
	}
	defer file.Close()
	
	in = GenerateOutput(in)
	
	x, jerr := json.MarshalIndent(in, "", "\t")
	if jerr != nil {
		return err
	}
	
	
	file.Write(x)
	
	
	
	return nil
}




func GenerateOutput(in interface{}) map[string]interface{}  {
	var temp map[string]interface{} = map[string]interface{}{}
	
	temp["data"] =  map[string]interface{}{
		"entry1" : in,
	}
	
	
	
	
	return temp
	
	
}