package Actions

import (
	"errors"
	"math/rand"
	"time"
)

func Sum(in interface{}) (int, error){
	rand.Seed(time.Now().Unix())
	if _, ok := in.(int); ok != true  {
		return 0, errors.New("Not an int type")
	}
	
	return in.(int)+rand.Intn(100+333), nil
}

func Times(in interface{}) (int, error){
	if _, ok := in.(int); !ok {
		return 0, errors.New("Not an int type")
	}
	
	return in.(int)*3, nil
}
