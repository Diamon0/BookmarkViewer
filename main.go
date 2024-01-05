package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Bookmark Viewer"),
			app.MinSize(unit.Dp(800), unit.Dp(600)),
			app.Size(unit.Dp(800), unit.Dp(600)),
		)
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var palette Palette = Palette{
	bg:      color.NRGBA{R: 0, G: 0, B: 0, A: 255},
	fg:      color.NRGBA{R: 51, G: 34, B: 69, A: 255},
	hl:      color.NRGBA{R: 124, G: 80, B: 171, A: 255},
	inter:   color.NRGBA{R: 114, G: 69, B: 163, A: 255},
	interhl: color.NRGBA{R: 144, G: 99, B: 193, A: 255},
	act:     color.NRGBA{R: 84, G: 39, B: 133, A: 255},
}

func run(w *app.Window) error {
	theme := material.NewTheme()
	theme.Bg = palette.bg
	theme.Fg = palette.fg
	var ops op.Ops
    settingsBtn := Button{}
    loadBtn := Button{}

    openSettings := func() {
        log.Println("Entered Settings!")
    }
    loadBookmarks := func() {
        log.Println("Loaded Bookmarks!")
    }

	for {
		switch e := w.NextEvent().(type) {
		case system.DestroyEvent:
			return e.Err

		case system.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := layout.NewContext(&ops, e)
            
            layout.Flex{}.Layout(gtx,
                FlexZone(gtx, image.Pt(gtx.Constraints.Max.X, 40), theme.Fg),
            )

			layout.Flex{}.Layout(gtx,
				settingsBtn.ButtonFlex(gtx, image.Pt(60, 40), image.Pt(0, 0), palette.inter, palette.interhl, palette.act, openSettings),
                loadBtn.ButtonFlex(gtx, image.Pt(60, 40), image.Pt(60, 0), palette.inter, palette.interhl, palette.act, loadBookmarks),
			)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
