package Socks

import (
	"SmartPolice-Interface/Core/Utils"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
	"math/rand"
	"os"
)

var  SimDevices map[string]interface{}  = map[string]interface{}{
	"LightBulb" : map[string]interface{}{
		"Manufacturer" : "Phillips",
		"Type" : "Hue",
		"Year" : 2017,
		"Data" : map[string]interface{}{
			"Red" : 0,
			"Green" : 0,
			"Blue" : 0,
		},
		
	},
	"LightSwitch" : map[string]interface{}{
		"Manufacturer" : "Nest",
		"Type" : "Gen",
		"Year" : 2018,
		"Data" : map[string]interface{}{
			"On" : false,
		},
	},
	
	"CodePad" : map[string]interface{}{
		"Manufacturer" : "Securitas",
		"Type" : "Securi",
		"Year" : 2018,
		"Data" : map[string]interface{}{
			"1" : 0,
			"2" : 0,
			"3" : 0,
			"4" : 0,
		},
	},
	
	"TempHumid" : map[string]interface{}{
		"Manufacturer" : "Arduino",
		"Type" : "Securi",
		"Year" : 2018,
		"Data" : map[string]interface{}{
			"Temperature" : 0,
			"Humid" : 0,
		},
	},
	
}

var  Simulator []DeviceList

type DeviceList struct {
	Name  string
	ID    string
	Data  map[string]interface{}
	Connection net.Conn
}



func (D *DeviceList) New() {
	rand.Seed(time.Now().UnixNano())
	pool := []string{"LightBulb", "LightSwitch", "CodePad", "TempHumid"}
	name := pool[Utils.RandInt(0, len(pool)-1)]
	
	D.Name = name
	D.Data = map[string] interface{}{
		name : SimDevices[name].(map[string]interface{}),
	}
	
	
	
}

func (D *DeviceList) RandomValues() {
	switch D.Name {
	case "LightBulb":
		//fmt.Println(D.Data[D.Name])
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["Red"] = Utils.RandInt(0, 255)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["Green"] = Utils.RandInt(0, 255)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["Blue"] = Utils.RandInt(0, 255)
		//fmt.Println(D.Name, D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{}))
		break
	case "LightSwitch":
		//fmt.Println(D.Name)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["On"] = Utils.RandBool()
		//fmt.Println(D.Name, D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{}))

		break
	case "CodePad":
		//fmt.Println(D.Name)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["1"] = Utils.RandInt(0, 9)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["2"] = Utils.RandInt(0, 9)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["3"] = Utils.RandInt(0, 9)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["4"] = Utils.RandInt(0, 9)
		//fmt.Println(D.Name, D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{}))
		break
	case "TempHumid":
		//fmt.Println(D.Name)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["Temperature"] = Utils.RandInt(0, 35)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["Humidity"] = Utils.RandInt(50, 90)
		//fmt.Println(D.Name, D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{}))
		break
	}
}



func (D *DeviceList) SendData (wg *sync.WaitGroup) {
	time.Sleep(time.Second * 3)
	
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:1201")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	var Success []int
	
	for len(Success) < 20 {
		bytes, _ := json.MarshalIndent(D.Data[D.Name].(map[string]interface{})["Data"], "", "\t")
		
		_, err := conn.Write(bytes)
		//fmt.Println(len(Success))
		if err != nil {
			panic(32)
		}
		Success = append(Success, 1)
	}
	
}





func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

