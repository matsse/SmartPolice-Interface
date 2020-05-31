package Tests

import (
	"SmartPolice-Interface/Core"
	"SmartPolice-Interface/Core/Socks"
	"SmartPolice-Interface/Core/Utils"
	"encoding/json"
	"log"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"math/rand"
	"os"
	"fmt"
)

var packets []Incoming

type Incoming struct {
	Conn    net.Conn
	ID      int
	bytes   int
	Auth    bool
	Format    Core.Format
}

var currentTest TestResult
var TestResults []TestResult
type TestResult struct {
	ServerBytes         uint64
	DeviceBytes         uint64
	ServerTime          time.Duration
	DeviceFailure       uint64
	SizeDiff            uint64
	SizePercent         float64
}



var wg sync.WaitGroup
var dg sync.WaitGroup
var TestType string = ""
var run_count = 1
var counter uint64
var TotalRuns int = 2000
var TestMode = ""


func Test_32(t *testing.T) {
	Core.LoadFormats()


	
	if os.Args[4] == "Time" || os.Args[4] == "Loss" {
		TestMode = os.Args[4]
	} else {
		log.Println("No test type is specified! Please append \"Loss\" or \"Time\"")
		os.Exit(1)
	}
	

	
	
	
	Core.TestingLock = 1

	
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	
	checkError(err)
	
	resend:
	start := time.Now()
	currentTest = TestResult{}
	wg.Add(1)
	go ListenerTest(&wg, listener)
	
	ClearValues()
	for i := 0; i < TotalRuns; i ++ {
		dg.Add(1)
		rand.Seed(time.Now().UnixNano() * int64(Utils.RandInt(0, 9999999999)))
		Socks.Simulator = append(Socks.Simulator, Socks.DeviceList{})
		Socks.Simulator[i].New()
		Socks.Simulator[i].RandomValues()
		
		go Socks.Simulator[i].SendData(&dg)
	}
	dg.Wait()
	
	//wg.Wait()
	
	
	
	
	currentTest.ServerTime = time.Since(start)
	//log.Printf("All devices generated %d bytes during run %d.", Socks.ByteCount, run_count)
	//log.Printf("Test run %d took %s and wrote %d bytes!", run_count, elapsed, counter)
	currentTest.DeviceBytes = Socks.ByteCount
	currentTest.DeviceFailure = Socks.FailedDevice
	Calculate_differences()
	TestResults = append(TestResults, currentTest)
	
	if run_count < 10 {
		time.Sleep(time.Millisecond * 10)
		run_count++
		goto resend
	} else {
	//log.Println("Total number of bytes", counter, "written in ", elapsed)
	}
}



func ListenerTest(wg *sync.WaitGroup, listener *net.TCPListener) {
	
	//Deadline := time.Now().Add(time.Second * 10)
	defer wg.Done()

	
	for {
		
		if len(packets) == TotalRuns-int(Socks.FailedDevice) /*|| time.Now().UnixNano() > Deadline.UnixNano()*/ {
		
			packets = nil
			return
		} else {
			//fmt.Println(time.Now().Unix(), Deadline.Unix())
			//fmt.Println(len(packets), dg, wg )
			
		}
		conn, err := listener.Accept()
		if err != nil {
		
		}
		// run as a goroutine
		go handleClient(conn)
	}
	
}

func handleClient(conn net.Conn) {
	
	defer conn.Close()
	
	if TestType == "TestConnnections" {
		return
	}
 

	tmp := Incoming{
		Conn: nil,
		ID:   0,
		
	}
	
	
	var buf []byte = make([]byte, 512)
	i, err := conn.Read(buf[0:])
	
	if err != nil {
		log.Fatal("The message was not recieved", err)
		return
	}
	
	
		authenticated := tmp.Authenticate(buf[:i])
	
		if authenticated == false {
			panic(0)
			return
		}

	
	packets = append(packets, tmp)
	
	atomic.AddUint64(&currentTest.ServerBytes, uint64(i))
	
	

	
	
	//total_bytes <-+ i
	//log.Println("Incoming from : ", string(buf[:i]), len(packets))
	
}

func CloseConnection(conn net.Conn) {

}


func (i *Incoming)Authenticate(message [] byte) bool {
	
	//fmt.Println(string(message))
	
	
	err := json.Unmarshal(message, &i.Format)
	if err != nil {
		panic(string(message))
	}
	//fmt.Println(i.Format["format"])
	for in := range i.Format["format"].(map[string]interface{}) {
		if Core.Formats[in] == nil {
			return false
		}
		
		for entry := range i.Format["format"].(map[string]interface{})[in].([]interface{}) {
			name := i.Format["format"].(map[string]interface{})[in].([]interface{})[entry].(string)
			format := Core.Formats[in].(map[string]interface{})["Target"].(map[string]interface{})[name]
			//fmt.Println("this", name, format.(string), i.Format["Data"].(map[string]interface{}))
			
			value, returnf := Core.AnalyzeFormat(format.(string) ,i.Format["Data"].(map[string]interface{}))
			if returnf == false {
				fmt.Println("Could not find location key in file!")
			}
			dtype, returnf := Core.AnalyzeType(format.(string) ,i.Format["Data"].(map[string]interface{}))
			if returnf == false {
				fmt.Println("Could not find type key in file!")
			}
			action := strings.Split(name+".Action", ":")
			if len(action) > 1 {
				result, actErr := Core.ReadChain(action[1], value, dtype)
				if actErr != nil {
					panic(actErr)
				}
				fmt.Println(result)
			}
			
		}
	}
	
	//log.Panicf("%v",i.Format["format"])
	return true
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	
}

func ClearValues() {
	Socks.ByteCount = 0
	Socks.FailedWriteBytes = 0
	Socks.FailedDevice = 0
	
}



func Calculate_differences() {
	//log.Printf("Test %d generagted %d bytes from %d bytes on the devices. %d devices have failed ", run_count, currentTest.ServerBytes, currentTest.DeviceBytes, Socks.FailedWriteBytes)
	//fmt.Println("FileBytes", Socks.FileBytes)
	
	
	if TestMode ==  "Time" {
		fmt.Printf("Test %d recieved %d bytes and took %s to complete\n", run_count, currentTest.ServerBytes, currentTest.ServerTime)
	} else if TestMode == "Loss" {
		currentTest.SizeDiff = currentTest.DeviceBytes - currentTest.ServerBytes
		currentTest.SizePercent = 100.0 * (float64(currentTest.SizeDiff) / float64(currentTest.DeviceBytes))
		fmt.Println("Bytes lost", currentTest.SizeDiff, " which is a loss of ", currentTest.SizePercent, "%" )
		fmt.Println("Device bytes ", currentTest.DeviceBytes, " Server bytes ", currentTest.ServerBytes)
	}

}


func Test_ByteRange(t *testing.T) {
	LightBulb, _ := json.Marshal(Socks.SimDevices["LightBulb"])
	LightSwitch, _ := json.Marshal(Socks.SimDevices["LightSwitch"])
	CodePad, _ := json.Marshal(Socks.SimDevices["CodePad"])
	TempHumid, _ := json.Marshal(Socks.SimDevices["TempHumid"])
	
	
	fmt.Printf("LightBulb has the size of: %d bytes \n", len(LightBulb))
	fmt.Printf("LightSwitch has the size of: %d bytes \n", len(LightSwitch))
	fmt.Printf("CodePad has the size of: %d bytes \n", len(CodePad))
	fmt.Printf("TempHumid has the size of: %d bytes \n", len(TempHumid))
	
}