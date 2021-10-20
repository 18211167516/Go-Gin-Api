package tool

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	//"fmt"
)

type String = string
type Int = int

func JsonToStruct(data []byte, s interface{}) error {
	err := json.Unmarshal(data, s)
	if err != nil {
		//err = fmt.Sprintf("Json marshaling failed：%s", err)
		return err
	}

	return nil
}

func ToUpper(str string) string {
	return strings.ToUpper(str)
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func StructToJson(v interface{}) string {
	data, _ := json.Marshal(v)

	return string(data)
}

func StringToInt(str string) int {
	variable, _ := strconv.Atoi(str)
	return variable
}

func IntToString(i int) string {
	return strconv.Itoa(i)
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
