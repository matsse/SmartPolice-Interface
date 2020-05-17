package main

import (
	"SmartPolice-Interface/Core"
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
	
	fmt.Printf(`
	Welcome to the smart policing interface for data aqcuisition from IoT devices.
	This program lets you define custom formats to extract useful information from smart
	devices and control what gets prioritized.

	Disclaimer: this is a Work in progress, so some issues might arise from using its features!

	Please select one of the options:

	device)    Enters the device management interface
	format)    Enters the format management interface
	listen)    Listen for the data that is gathered
	options)   Adjust settings
	exit)      Exit the program

`)

prompt: {
	fmt.Printf(`
Root# `)
		
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
	
	fmt.Printf(`
	This interface lets you manage your devices. All devices are stored in json format for future use.

	Please select one of the options:
	new)       Create a new device - A prompt will ask for more info.
	edit)      Edits an existing device - Example: edit device1
	delete)    Deletes a device - Example: delete device1
	dat)       dump current data from all devices
	list)      List all devices
	exit)      goes back to precious prompt
`)
	
	devPrompt:{
		fmt.Printf(`
Devices# `)
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





func Listen() {

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
	path := "./Devices/"
	
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
		
		switch temp.DeviceType {
		case "TTN":
			temp.Source.TTN.LoadDevices()
			break
			
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
