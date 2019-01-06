package main

import (
	"../goui"
	"fmt"
)

func main(){
	mainFunc()
}


func mainFunc(){
	goui := goui.GetGoui()
	winow := goui.CreateWindow(800,600,"demo")
	fmt.Println(winow)
}
