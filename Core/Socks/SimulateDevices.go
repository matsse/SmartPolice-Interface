package Socks

import (
	"SmartPolice-Interface/Core/Utils"
	"encoding/json"
	"math/rand"
	"net"
	"sync"
	"sync/atomic"
	"time"
)



var ByteCount uint64
var FailedWriteBytes uint64
var FailedDevice uint64
var FileBytes uint64
var  SimDevices map[string]interface{}  = map[string]interface{}{
	"LightBulb" : map[string]interface{}{
		"Data" : map[string]interface{}{
			"red" : 0,
			"green" : 0,
			"blue" : 0,
		},
		"format": map[string]interface{}{
			"GH-Hue": []string{
				"Red",
				"Green",
				"Blue",
			},
		},
	},
	"LightSwitch" : map[string]interface{}{
		"Data" : map[string]interface{}{
			"on" : false,
		},
		"format": map[string]interface{}{
			"LightSwitch": []string{
				"On",
			},
		},
	},
	
	"CodePad" : map[string]interface{}{
		"Data" : map[string]interface{}{
			"1" : 0,
			"2" : 0,
			"3" : 0,
			"4" : 0,
		},
		"format": map[string]interface{}{
			"Secu-Lock": []string{
				"1",
				"2",
				"3",
				"4",
			},
		},
	},
	
	"TempHumid" : map[string]interface{}{
		"Data" : map[string]interface{}{
			"temperature" : 0,
			"humidity" : 0,
		},
		"format": map[string]interface{}{
			"ArduinoTemp": []string{
				"Temperature",
				"Humidity",
			},
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
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["red"] = Utils.RandInt(0, 255)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["green"] = Utils.RandInt(0, 255)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["blue"] = Utils.RandInt(0, 255)
		//fmt.Println(D.Name, D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{}))
		break
	case "LightSwitch":
		//fmt.Println(D.Name)
		D.Data[D.Name].(map[string]interface{})["Data"].(map[string]interface{})["on"] = Utils.RandBool()
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

func (D *DeviceList) Authentication (conn net.Conn) bool  {
		//bytes, _ := json.MarshalIndent(D.Data[D.Name].(map[string]interface{})["format"], "", "\t")
		//_, err := conn.Write(bytes)
		//
		//if err != nil {
		//	return false
		//
		//	//continue
		//}
	return true
}

func (D *DeviceList) SendData (dg *sync.WaitGroup) {
	defer dg.Done()
	
	time.Sleep(time.Millisecond * time.Duration(Utils.RandInt(3000,     10000)))
	//time.Sleep(time.Second * 1)
	

	
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:1201")
	checkError(err, dg)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, dg)
	//var Success []int
	//defer conn.Close()
	defer closeConnection(conn)
	

	bytes, _ := json.Marshal(D.Data[D.Name].(map[string]interface{}))

	atomic.AddUint64(&FileBytes, uint64(len(bytes)))
		
		
		i, err := conn.Write(bytes)
		if err != nil {
			
			atomic.AddUint64(&FailedDevice, uint64(1))
			atomic.AddUint64(&FailedWriteBytes, uint64(1))
			return
			//time.Sleep(time.Second * 1)
			//goto reset
		
		}
		conn.CloseWrite()
		
		atomic.AddUint64(&ByteCount, uint64(i))
		//fmt.Println(len(Success))

		//conn.Write([]byte("EOC"))
		
		
		//Success = append(Success, 1)
		//time.Sleep(time.Second * 1)
	
	
}


func closeConnection(conn net.Conn) {
		
		if conn == nil{
			return
		}
		//if err := conn.Close(); err != nil {
		//	FailedDevice ++
		//}
}


func checkError(err error, dg *sync.WaitGroup) {
	if err != nil {
		//panic(err)
		atomic.AddUint64(&FailedDevice, 1)
		return
		//fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		//os.Exit(1)
	}
}

