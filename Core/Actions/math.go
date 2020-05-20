package Actions

import (
	"errors"
)

func Sum(in interface{}) (int, error){
	if _, ok := in.(int); ok != true  {
		return 0, errors.New("Not an int type")
	}
	
	return in.(int)+132, nil
}

func Times(in interface{}) (int, error){
	if _, ok := in.(int); !ok {
		return 0, errors.New("Not an int type")
	}
	
	return in.(int)*3, nil
}
