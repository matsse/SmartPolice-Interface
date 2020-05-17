package Core

import (
	"reflect"
)

// runs tests on variables


func Integer_validation(value interface{}) bool {
	if value.(float64) != 0.0 && int(value.(float64)) == 0 {
		return false
	}
	tmp := int(value.(float64))
	myType := reflect.TypeOf(tmp)
	if k := myType.Kind(); k == reflect.Int {
		return true
	}
	return false
}

func Integer8_validation(value interface{}) bool {
	if value.(float64) != 0.0 && int8(value.(float64)) == 0 {
		return false
	}
	tmp := int8(value.(float64))
	myType := reflect.TypeOf(tmp)
	if k := myType.Kind(); k == reflect.Int8 {
		return true
	}
	return false
}
func Integer16_validation(value interface{}) bool {
	if value.(float64) != 0.0 && int16(value.(float64)) == 0 {
		return false
	}
	
	tmp := int16(value.(float64))
	myType := reflect.TypeOf(tmp)
	if k := myType.Kind(); k == reflect.Int16 {
		return true
	}
	return false
}
func Integer32_validation(value interface{}) bool {
	if value.(float64) != 0.0 && int32(value.(float64)) == 0 {
		return false
	}
	
	tmp := int32(value.(float64))
	myType := reflect.TypeOf(tmp)
	if k := myType.Kind(); k == reflect.Int32 {
		return true
	}
	return false
}

func Integer64_validation(value interface{}) bool {
	if value.(float64) != 0.0 && int64(value.(float64)) == 0 {
		return false
	}
	
	tmp := int64(value.(float64))
	myType := reflect.TypeOf(tmp)
	if k := myType.Kind(); k == reflect.Int64 {
		return true
	}
	return false
}



// IEEE-754 float32 value
func Float32_validation(value interface{}) bool {
	tmp := float32(value.(float64))
	myType := reflect.TypeOf(tmp)
	if k := myType.Kind(); k == reflect.Float32 {
		return true
	}
	return false
}

// IEEE-754 float63 value
func Float64_validation(value interface{}) bool {
	if _, s := value.(float64); s == true {
		return true
	}
	return false
}


func String_Validation(value interface{}) bool {
	if _, s := value.(string); s == true {
		return true
	}
	
	
	return false
	
}
func Bool_Validation(value interface{}) bool {
	tmp := uint8(value.(float64))
	myType := reflect.TypeOf(tmp)
	if k := myType.Kind(); k == reflect.Uint8 {
		return true
	}
	return false
}


func Byte_Validation(value interface{}) bool {
	tmp := uint8(value.(float64))
	myType := reflect.TypeOf(tmp)
	if k := myType.Kind(); k == reflect.Uint8 {
		return true
	}
	return false
}
