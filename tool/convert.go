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

func Rtrim(str string) string {
	return strings.TrimSuffix(str, "s")
}

//字符串转双驼峰写法
func CamelCase(str string, separator string) string {
	var text string
	for _, p := range strings.Split(str, separator) {
		switch length := len(p); length {
		case 0:
		case 1:
			text += strings.ToUpper(p[0:1])
		default:
			text += strings.ToUpper(p[0:1]) + p[1:]
		}
	}

	return text
}

//字符串转单驼峰写法
func SingleCamelCase(str string, separator string) string {
	temp := strings.Split(str, separator)
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

//首字母大写
func FirstUpper(s string) string {
    if s == "" {
        return ""
    }
    return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
  if s == "" {
      return ""
  }
  return strings.ToLower(s[:1]) + s[1:]
}

