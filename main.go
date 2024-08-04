package main

import (
	"gioui.org/app"
	"gioui.org/unit"
)

func main() {
	go func() {
		// Create new window
		w := new(app.Window)
		w.Option(app.Title("Calculator"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))

		// Listen for events in the window
		for {
			w.Event()
		}
	}()
	app.Main()
}
