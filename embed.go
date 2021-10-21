package main

import (
	"embed"
	"go-api/global"
)

//go:embed static templates
var f embed.FS

func initEmbed() {
	global.FS = f //初始化FS
}
