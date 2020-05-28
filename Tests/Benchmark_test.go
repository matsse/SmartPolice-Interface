package Tests

import (
	"SmartPolice-Interface/Core"
	"SmartPolice-Interface/Core/Socks"
	"SmartPolice-Interface/Core/Utils"
	"log"
	"net"
	"sync"
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
}

var wg sync.WaitGroup
var TestType string = ""

func Test_32(t *testing.T) {
	fmt.Println(os.Args[3] )
	if os.Args[3] == "TestConnnections" {
	TestType = os.Args[3]
	}
	Core.TestingLock = 1
	start := time.Now()
	wg.Add(1)
	go ListenerTest(&wg)
	
	
	for i := 0; i < 2000; i ++ {
		wg.Add(1)
		rand.Seed(time.Now().UnixNano() * int64(Utils.RandInt(0, 1001034345101010101)))
		Socks.Simulator = append(Socks.Simulator, Socks.DeviceList{})
		Socks.Simulator[i].New()
		Socks.Simulator[i].RandomValues()
		
		go Socks.Simulator[i].SendData(&wg)
	}
	wg.Wait()
	
	
	
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}



func ListenerTest(wg *sync.WaitGroup) {
	defer wg.Done()
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	
	for {

		if len(packets) == 1999 {
			//break
			return
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
	
	packets = append(packets, tmp)
	log.Println("Incoming from : ", string(buf[:i]), len(packets))
	
	
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	
}
