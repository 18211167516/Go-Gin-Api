package tool

import "go-api/global"

func Config(key string) interface{}{
	return global.VP.Get(key)
}