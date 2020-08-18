package tool

import (
	"encoding/json"
	"reflect"
	"strconv"
	//"fmt"
)

func JsonToStruct(data []byte, s interface{}) error {
	err := json.Unmarshal(data, s)
	if err != nil {
		//err = fmt.Sprintf("Json marshaling failed：%s", err)
		return err
	}

	return nil
}

func StringToInt(str string) int {
	variable, _ := strconv.Atoi(str)
	return variable
}

func StructToMap(obj interface{}) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
