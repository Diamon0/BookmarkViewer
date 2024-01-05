package main

import (
	"image"
	"image/color"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type Palette struct {
	// Background Color
	bg color.NRGBA

	// Foreground Color
	fg color.NRGBA

	// Highlight Color
	hl color.NRGBA

	// Interactable Color
	inter color.NRGBA

	// Interactable Highlighted Color
	interhl color.NRGBA

	// Active Interactable Color
	act color.NRGBA
}

func ColorZone(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
    defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
    paint.ColorOp{Color: color}.Add(gtx.Ops)
    paint.PaintOp{}.Add(gtx.Ops)
    return layout.Dimensions{Size: size}
}

func FlexZone(gtx layout.Context, size image.Point, color color.NRGBA) layout.FlexChild {
    return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
        return ColorZone(gtx, size, color)
    })
}

type Button struct {
    Hover bool
    Pressed bool
    pressedOnce bool
}

func DrawButton(ops *op.Ops, size image.Point, color color.NRGBA) layout.Dimensions {
    defer clip.Rect{Max: size}.Push(ops).Pop()
    paint.ColorOp{Color: color}.Add(ops)
    paint.PaintOp{}.Add(ops)
    return layout.Dimensions{Size: size}
}

func (b *Button) ButtonFlex(gtx layout.Context, size, startPoint image.Point, color, hoverColor, pressColor color.NRGBA, action func()) layout.FlexChild {
    for _, e := range gtx.Events(b) {
        if e, ok := e.(pointer.Event); ok {
            switch e.Kind {
            case pointer.Press:
                b.Pressed = true
            case pointer.Release:
                b.Pressed = false
                b.pressedOnce = false
            case pointer.Enter:
                b.Hover = true
            case pointer.Leave:
                b.Hover = false

            }
        }
    }
    
    area := clip.Rect(image.Rect(startPoint.X, startPoint.Y, size.X + startPoint.X, size.Y + startPoint.Y)).Push(gtx.Ops)
    pointer.InputOp{
        Tag: b,
        Kinds: pointer.Press | pointer.Release | pointer.Enter | pointer.Leave,
    }.Add(gtx.Ops)
    area.Pop()

    return layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
        var col = color
        if b.Pressed && b.Hover && !b.pressedOnce {
            col = pressColor
            b.pressedOnce = true
            action()
        } else if b.Hover {
            col = hoverColor
        }
        return DrawButton(gtx.Ops, size, col)
    })
}
