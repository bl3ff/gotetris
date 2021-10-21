package view

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"tetris/control"
	"tetris/model"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

type TetrisView struct {
	Win    *screen.Window
	Engine *control.EngineT
	s      size.Event
}

func NewTetrisView(w *screen.Window) *TetrisView {
	fmt.Println("Create New View")
	view := &TetrisView{
		Win:    w,
		Engine: control.GenerateEngine(),
	}
	return view
}

func (tw *TetrisView) Start() {
	fmt.Println("Starting View...")
	tw.Engine.Start(tw.drawAll)

	for {
		e := (*tw.Win).NextEvent()
		switch e := e.(type) {
		case mouse.Event:
			tw.sendMouseEvent(&e)
		case lifecycle.Event:
			if e.To == lifecycle.StageDead {
				return
			}

		case key.Event:
			if e.Code == key.CodeEscape {
				return
			}
			tw.sendKeyEvent(&e)
		case paint.Event:
			tw.drawAll()
		case size.Event:
			tw.s = e
			tw.drawAll()
		case error:
			log.Print(e)
		}
	}

}

func BlockColorToRGBA(bc model.BlockColor) color.RGBA {
	var rc color.RGBA
	switch bc {
	case model.Blue:
		rc = color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}
	case model.Red:
		rc = color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
	case model.Orange:
		rc = color.RGBA{R: 0xFF, G: 0x80, B: 0x00, A: 0xFF}
	case model.Yellow:
		rc = color.RGBA{R: 0xFF, G: 0xFF, B: 0x00, A: 0xFF}
	case model.Green:
		rc = color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}
	case model.Pink:
		rc = color.RGBA{R: 0xFF, G: 0x99, B: 0xCC, A: 0xFF}
	case model.Purple:
		rc = color.RGBA{R: 0x99, G: 0x33, B: 0xFF, A: 0xFF}
	default:
		rc = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
	}
	return rc
}

func (tw *TetrisView) sendKeyEvent(e *key.Event) {
	if e.Direction == key.DirPress {
		switch e.Code {
		case key.CodeLeftArrow:
			tw.Engine.SendKey(control.Key{DirKey: map[model.Direction]bool{model.Left: true}})
		case key.CodeDownArrow:
			tw.Engine.SendKey(control.Key{DirKey: map[model.Direction]bool{model.Down: true}})
		case key.CodeRightArrow:
			tw.Engine.SendKey(control.Key{DirKey: map[model.Direction]bool{model.Right: true}})
		case key.CodeUpArrow:
			tw.Engine.SendKey(control.Key{DirKey: map[model.Direction]bool{model.Up: true}})
		default:
			//log.Printf("key not supported: %s", e.Code)
		}
	}
}

func (tw *TetrisView) sendMouseEvent(e *mouse.Event) {
	if e.Direction == mouse.DirRelease {
		switch e.Button {
		case mouse.ButtonLeft:
			tw.Engine.SendClick(control.Click{Right: false, Left: true})
		case mouse.ButtonRight:
			tw.Engine.SendClick(control.Click{Right: true, Left: false})
		default:
			//log.Printf("mouse button not supported: %d", e.Button)
		}
	}
}

func (tw *TetrisView) drawAll() {
	tw.drawTMatrix()
	(*tw.Win).Publish()
	tw.drawCurrentBlock()
	(*tw.Win).Publish()
}

func (tw *TetrisView) drawTMatrix() {
	m := tw.Engine.TMatrix
	if m == nil {
		return
	}

	m.RLock()
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			//if m.Matrix[i][j] != model.EmptyColor {
			tw.drawSingleRect(model.Pos{X: j, Y: i}, BlockColorToRGBA(m.Matrix[i][j]), model.BlockSize)
			//}
		}
	}
	m.RUnlock()
}

func (tw *TetrisView) drawSingleRect(pos model.Pos, c color.RGBA, sizePx int) {
	realpos := model.Pos{X: pos.X * sizePx, Y: pos.Y * sizePx}
	(*tw.Win).Fill(image.Rect(realpos.X,
		realpos.Y,
		realpos.X+sizePx,
		realpos.Y+sizePx), c, screen.Over)
}

func (tw *TetrisView) drawCurrentBlock() {
	b := tw.Engine.CurrentBlock
	c := BlockColorToRGBA(b.Color)

	tw.drawSingleRect(b.Center, c, model.BlockSize)
	for _, p := range b.Pieces {
		tw.drawSingleRect(p, c, model.BlockSize)
	}
}
