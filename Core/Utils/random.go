package Utils

import (
	"fmt"
	"time"
	"math/rand"
)


var (
	letterRunes string = "abcdefghijklmnopqrstuvwxyz"
)


func RandCoords(min, max float32) float32{
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float32() * (max - min)
}


func RandInt(min, max int ) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}


func RandByte() []byte  {
	rand.Seed(time.Now().UnixNano())
	token := make([]byte, 8)
	rand.Read(token)
	return token
}

func RandBool() bool {
	return rand.Float32() < 0.5
}


func RandomSuffix() string {
	output := ""
	for i := 0; i < 8; i++ {
		index :=  0 + rand.Intn(26-0)
		

		output += string(letterRunes[index])
		fmt.Println(output)
	}
	
	return output
	
}