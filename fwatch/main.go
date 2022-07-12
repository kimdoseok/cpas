package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(IconFwatch)
	systray.SetTitle("Fwatch App")
	systray.SetTooltip("Fwatch configuration")
	mControl := systray.AddMenuItem("Setting", "Setting the app")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(IconQuit)
	mControl.SetIcon(IconControl)
	go func() {
		for {
			select {
			case <-mControl.ClickedCh:
				go EditConfig()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	// clean up here
	systray.Quit()
}

func EditConfig() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}

func loop() {
	g.SingleWindow().Layout(
		g.Label("Hello world from giu"),
		g.Row(
			g.Button("Click Me").OnClick(onClickMe),
			g.Button("I'm so cute").OnClick(onImSoCute),
		),
	)
}

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}
