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

	var button widget.Clickable

	theme := material.NewTheme()

	for {
		eventType := w.Event()

		switch typ := eventType.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ)

			buttonLayout := layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}

			// Rigid element to hold the button
			rigidElement := layout.Rigid(
				func(gtx layout.Context) layout.Dimensions {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),
					}

					return margins.Layout(gtx,
						func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(theme, &button, "Click me!")
							return btn.Layout(gtx)
						},
					)
				},
			)

			buttonLayout.Layout(gtx, rigidElement)
			typ.Frame(&ops)

		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}
