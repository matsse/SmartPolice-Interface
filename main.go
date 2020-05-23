package main

import (
	"SmartPolice-Interface/Core"
	"SmartPolice-Interface/Core/Utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func init() {
	fmt.Println("Checking for exisiting devices...")
	
	LoadDevices()
	StartupDevices()
	Core.LoadFormats()
}


var CurrentDevice *Core.Device


func exit ()  {
	
	for device := range Core.Devices {
		Core.Devices[device].SaveDevice()
	}
	
}

func main() {
	
	defer exit()
	
	
	var option string = ""
	
	fmt.Println(Utils.MainHome)

prompt: {
	fmt.Printf(Utils.MainPS)
		
		fmt.Scanf("%s\n", &option)
	
	
	switch option {
	case "device":
		DeviceScreen()
		break
	case "format":
		CurrentDevice.Format.SetFormat()
		break
	case "listen":
		Core.ListenToAll()
		break
	case "options":
		OptionScreen()
		break
	case "exit":
		return
		break
	}
	option = ""
	
	//end of option label
}
	
goto prompt

	
}


func DeviceScreen() {
	var devOption string = ""
	var devArg1 string = ""
	
	var tmpDevice Core.Device = Core.Device{}
	
	fmt.Println(Utils.DeviceHome)
	
	devPrompt:{
	fmt.Printf(Utils.DevicePS)
	
	
	
	
	fmt.Scanf("%s %v\n", &devOption, &devArg1 )
	
	switch devOption {
	case "new":
		tmpDevice.SetDevice()
		break
	case "edit":
		break
	case "dat":
		AllValues()
		break
	case "select":
		CurrentDevice = Select(devArg1)
		break
	case "delete":
		break
	case "list":
		tmpDevice.List()
		break
	case "exit":
		return
		break
	}
	
	devOption = ""
	
}


	
	
goto devPrompt



}



func OptionScreen() {

}




func Select(name string) *Core.Device {

	
	for i := range Core.Devices {
		if Core.Devices[i].DeviceName == name {
			defer fmt.Println("Selected", Core.Devices[i])
			return &Core.Devices[i]
		}
	}
	
	return nil
}



func LoadDevices() {
	var files []string
	//path := "./Devices/"
	path := "./Data/Devices/"
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	
	
	//open and read all files in the devices directory
	for _, file := range files {
		
		if file[len(file)-4:] != "json" {
			continue    // Error handling for non json files
		}
		
		var temp Core.Device = Core.Device{}
		temp.Data = map[string]interface{}{}
		file, errfile := os.Open(file)
		if errfile != nil {
			panic(errfile)
		}
		
		bytes, _ := ioutil.ReadAll(file)
		
		
		json.Unmarshal(bytes, &temp)
		
		//switch temp.DeviceType {
		//case "TTN":
		//	temp.Source.TTN.LoadDevices()
		//	break
		//
		//}
		
		if temp.DeviceType == "TTN" {
			temp.Source.TTN.LoadDevices()
		}
		
		Core.Devices = append(Core.Devices, temp)
		
		file.Close()
	}
}

func AllValues() {
	for i := range Core.Devices {
		fmt.Println(Core.Devices[i].DeviceName)
		for j := range Core.Devices[i].Data {
			fmt.Println(Core.Devices[i].Data[j])
		}
		
	}
}


func StartupDevices() {
	for i := range Core.Devices {
		go Core.Devices[i].Startup()
	}
}
