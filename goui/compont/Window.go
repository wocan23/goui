package compont

// 窗口
type Window interface {
	SetWidth(width int)
	SetHeight(height int)
	SetTitle(title string)
	GetWidth(width int) int
	GetHeight(height int)  int
	GetTitle(title string)  string
}

type TWindow struct{
	Width int
	Height int
	Title string
}

type WinWindow struct{
	TWindow
}

type MacWindow struct{
	TWindow
}

type LinuxWindow struct{
	TWindow
}

func (tWindow *TWindow)SetWidth(width int) {
	tWindow.Width = width
}

func (tWindow *TWindow)SetHeight(height int) {
	tWindow.Height = height
}

func (tWindow *TWindow)SetTitle(title string) {
	tWindow.Title = title
}

func (tWindow *TWindow)GetWidth(width int) int{
	return tWindow.Width
}

func (tWindow *TWindow)GetHeight(height int)  int{
	return tWindow.Height
}

func (tWindow *TWindow)GetTitle(title string)  string{
	return tWindow.Title
}





