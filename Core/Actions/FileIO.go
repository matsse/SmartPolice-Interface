package Actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)





func SaveToFile(in interface{}, name string) (error){
	if _, err := os.Stat("../DataOutput/"+name+".json"); err == nil  {
	
	}

	
	file, err := os.OpenFile("../DataOutput/"+name+".json", os.O_RDWR, 0755)
	defer file.Close()
	
	if x, _ := file.Stat(); x.Size() < 10 {
		in = GenerateOutput(in)
		
		x, jerr := json.MarshalIndent(in, "", "\t")
		if jerr != nil {
			return jerr
		}
		
		_, err = file.Write(x)
		if err != nil {
			panic(err)
		}
	} else {
		
		var temp map[string] interface{}
		buffer, _ := ioutil.ReadAll(file)
		
		error2 := json.Unmarshal(buffer, &temp)
		
		
		
		if error2 != nil {
			panic(error2)
		}
		output := AppendOutput(in, temp)
		x, jerr := json.MarshalIndent(output, "", "\t")

		if jerr != nil {
			return jerr
		}
		file.Close()
		file, _ = os.Create("../DataOutput/"+name+".json")
		
		_, writeErr := file.Write(x)
		if writeErr != nil {
			panic(writeErr)
		}
		
	}
	
	return nil
}





func GenerateOutput(in interface{}) map[string]interface{}  {
	var temp map[string]interface{} = map[string]interface{}{}
	
	temp["data"] =  map[string]interface{}{
		"entry1" : in,
	}
	
	return temp
}

func AppendOutput(in interface{},  file map[string]interface{}) map[string]interface{}  {
	count := len(file["data"].(map[string]interface{}))
	entry := fmt.Sprintf("entry%d", count+1)
	
	file["data"].(map[string]interface{})[entry] = in
	
	return file
}