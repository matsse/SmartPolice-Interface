package Core

import (
	"SmartPolice-Interface/Providers"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)


var (
	isTTN bool = false
)


var Devices []Device

type Device struct {
	DeviceName string                  `json:"device_name"`
	DeviceType string                  `json:"device_type"`
	Link       []string                `json:"link"`
	DeviceInfo DeviceInfo              `json:"deviceInfo"`
	Format     Format                  `json:"format"`
	Source     Provider                `json:"source"`
	Conf       Settings                `json:"settings"`
	Data       map[string] interface{} `json:"data"`
}

type DeviceInfo struct {
	Manufacturer string     `json:"manufacturer"`
	Longitude float32       `json:"longitude"`
	Latitude float32        `json:"latitude"`
	
}


type Settings struct {
	Duration    time.Duration   `json:"interval"`
	Running     bool            `json:"is_running"`
}

type Provider struct {
	TTN     Providers.TTN_Info  `json:"ttn,omitempty"`
	TS      Providers.TS_info   `json:"ts,omitempty"`
}







func (D *Device) SetDevice() {
	// Fetches the device name
	fmt.Println("What is the device name?")
	fmt.Scanf("%s\n", &D.DeviceName)
	
	// Fetches the device type
	fmt.Println("What type is this device?")
	fmt.Printf(`
	TTN)    TTN LoRA device
	TS)     ThingSpeak.com
	
Device-Type# `)
	fmt.Scanf("%s\n", &D.DeviceType)
	
	switch D.DeviceType {
	case "TTN":
		D.Source = Provider{}
		x := Providers.TTN_Info{}
		x.NewDevice()
		fmt.Println(x)
		D.Source.TTN = x
		fmt.Println(D.Source.TTN)
		D.DeviceInfo.Longitude, D.DeviceInfo.Latitude = D.Source.TTN.GetCoordinates()
		break
	case "TS":
		D.Source = Provider{}
		x := Providers.TS_info{}
		x.NewDevice()
		D.Source.TS = x
		break
		
	}
	
	Devices = append(Devices, *D)
	
	
	D.SaveDevice()
}


func (D Device) List() {
	
	
	for i := range Devices {
		fmt.Println(Devices[i])
	}
}






func (D *Device) SaveDevice() {
	
	fmt.Println("e")
	b, err := json.MarshalIndent(D, "", "\t")
	
	fmt.Println(D)
	fmt.Println(string(b))
	if err != nil {
		log.Printf("Could not load %s into a JSON file!\n", D.DeviceName)
		return
	}
	
	
	var s string = fmt.Sprintf("./Devices/Device-%s-%s.json", D.DeviceName, D.DeviceType )
	
	Output, err := os.Create(s)
	
	if err != nil {
		log.Printf("Could not create output file for device %s!\n", D.DeviceName)
		log.Println(err)
		return
	}
	
	Output.Write(b)
	Output.Close()
	
}


func (D *Device) Startup() {
	
	if D.DeviceType == "TTN" {
		
			
			D.Source.TTN.GetData(&D.Data, D.Format)
		
		//D.Source.TTN.WatchFeed()
		
	} else {
		fmt.Println("ERROR")
	}



}


func (D Device) FetchUplink(wg *sync.WaitGroup) {
	defer wg.Done()
	if D.DeviceType == "TTN" {
		fmt.Println(D.DeviceName)
		D.Source.TTN.WatchFeed()
	} else {
		fmt.Println("ERROR")
	}
}