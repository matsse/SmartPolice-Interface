package Core

import (
	"fmt"
	"sync"
)




func ListenToAll() {
	var wg sync.WaitGroup
	
	for i := range Devices {
		wg.Add(1)
		fmt.Println(Devices[i])
		go Devices[i].FetchUplink(&wg)
	}
	wg.Wait()
}
