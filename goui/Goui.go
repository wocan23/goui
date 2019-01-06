package goui

import (
	"./platform"
	"./common"
)


func GetGoui() *Goui{

	osPlatform := common.GetPlatform()
	var api platform.Api
	switch osPlatform.Name {
	case "darwin": api = new(platform.Mac)
	case "unix": api = new(platform.Linux)
	case "windows": api = new(platform.Win)
	default:
		panic("unsupport platform:"+osPlatform.Name)
	}
	goui := Goui{Api:api}
	return &goui
}

type Goui struct{
	platform.Api
}

