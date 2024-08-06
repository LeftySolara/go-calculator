package main

import (
	"log"
	"os"
	"strconv"

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

	theme := material.NewTheme()

	for {
		eventType := w.Event()

		switch typ := eventType.(type) {

		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ)

			var numberButtons [10]widget.Clickable
			var rigidElements []layout.FlexChild

			flexbox := layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceAround,
			}

			for i := 0; i <= 9; i++ {
				button := material.Button(theme, &numberButtons[i], strconv.Itoa(i))
				rigidElement := layout.Rigid(button.Layout)
				rigidElements = append(rigidElements, rigidElement)
			}

			flexbox.Layout(gtx, rigidElements...)

			typ.Frame(gtx.Ops)

		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}
