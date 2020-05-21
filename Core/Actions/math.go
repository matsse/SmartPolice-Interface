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




func Sumx(in1 interface{}, in2 interface{} ) (int) {
	if _, ok := in1.(int);  ok != true    {
		panic("arg1 is not an integer")
	}
	if _, ok := in2.(int);  ok != true    {
		panic("arg2 is not an integer")
	}
	
	return in1.(int) + in2.(int)
}