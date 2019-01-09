package main

import (
	"fmt"
	"os"
	sciter "sciter-sdk"
	"sciter-sdk/window"
	"win"
)

type t_resolution struct {
	width  int
	height int
}

// Specifying  havily used
// Singltons to make them
// package wide available
var root *sciter.Element
var rootSelectorErr error
var w *window.Window
var windowErr error

// Preapare Scitre For Execution ///
func init() {

	// initlzigin window for downloaer
	// app with appropriate properties
	fmt.Println("init")
	res := getResolution()
	rect := sciter.NewRect(5, 5, res.width-100, res.height-100)
	w, windowErr = window.New(sciter.SW_TITLEBAR|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_GLASSY,
		rect)

	if windowErr != nil {
		fmt.Println("Can not create new window")
		return
	}
	// Loading main html file for app
	htloadErr := w.LoadFile("./simpleGuiApp.html")
	if htloadErr != nil {
		fmt.Println("Can not load html in the screen", htloadErr.Error())
		return
	}

	// Initializng  Selector at global level as we  are going to need
	// it mostly and as it is
	root, rootSelectorErr = w.GetRootElement()
	if rootSelectorErr != nil {
		fmt.Println("Can not select root element")
		return
	}

	// Set title of the appliaction window
	w.SetTitle("Simple Calc")

}
func getResolution() t_resolution {
	var l_resolution t_resolution
	hDC := win.GetDC(0)
	defer win.ReleaseDC(0, hDC)
	l_resolution.width = int(win.GetDeviceCaps(hDC, win.HORZRES))
	l_resolution.height = int(win.GetDeviceCaps(hDC, win.VERTRES))
	return l_resolution
}

// Preaprare Program for execution ///
func main() {

	fmt.Println("main")
	addbutton, _ := root.SelectById("add")

	addbutton.OnClick(func() {
		output := add()
		fmt.Println(output)
		os.Exit(3)
	})
	liExitbut, _ := root.SelectById("liExit")

	liExitbut.OnClick(func() {
		fmt.Println("Exits")
	})
	w.Show()
	w.Run()

}

//////////////////////////////////////////////////
/// Function of calc                           ///
//////////////////////////////////////////////////

func add() int {

	return 1 + 2
}
