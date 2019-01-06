package common

import (
	"runtime"
	"strings"
)

func GetPlatform() *Platform{
	arch := runtime.GOARCH
	os := runtime.GOOS
	var bit = 32
	if strings.HasSuffix(arch,"64"){
		bit = 64
	}
	platform := Platform{Name:os,Bit:bit}
	return &platform
}
