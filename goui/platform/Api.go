package platform

import "../compont"

type Api interface {
	CreateWindow(width int,height int,title string)*compont.Window
}
