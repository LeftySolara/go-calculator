package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		// Create new window
		w := new(app.Window)
		w.Option(app.Title("Calculator"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))

		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {
	// Operations from the UI
	var ops op.Ops

	var numberButton widget.Clickable

	// Define the material design style
	theme := material.NewTheme()

	// Listen for events in the window
	for {
		// Detect the event type
		switch eventType := w.Event().(type) {

		// Sent when the application should re-render
		case app.FrameEvent:
			gtx := app.NewContext(&ops, eventType)

			// Flexbox layout
			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start (the top)
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				// We insert two rigid elements:
				// First one to hold a button...
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(35),
							Left:   unit.Dp(35),
						}

						return margins.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(theme, &numberButton, "0")
								return btn.Layout(gtx)
							},
						)
					},
				),
			)
			eventType.Frame(gtx.Ops)
		// Sent when the application should exit
		case app.DestroyEvent:
			return eventType.Err
		}
	}
}
