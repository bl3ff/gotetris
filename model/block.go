package model

import (
	"math"
)

type BlockType int

const (
	// ****
	I = iota

	// *
	// ***
	J

	//   *
	// ***
	L

	// **
	// **
	O

	//  **
	// **
	S

	//  *
	// ***
	T

	// **
	//  **
	Z

	BlockTypeLen
)

type BlockColor int

const (
	EmptyColor = iota
	Red
	Green
	Blue
	Yellow
	Purple
	Orange
	Pink

	ColorLen
)

type Rotation float64

const (
	ClockWise        = math.Pi / 2
	CounterClockWise = -math.Pi / 2
)

type Direction int

const (
	Down = iota
	Left
	Right
)

type Pos struct {
	X, Y int
}

func (a *Pos) PosMinus(b Pos) {
	a.X -= b.X
	a.Y -= b.Y
}

func (a *Pos) PosPlus(b Pos) {
	a.X += b.X
	a.Y += b.Y
}

type Block struct {
	Center    Pos
	Pieces    []Pos
	Color     BlockColor
	TypeBlock BlockType
}

func (b Block) getCopy() Block {
	newblock := Block{}
	newblock.Center = b.Center
	newblock.Color = b.Color
	newblock.TypeBlock = b.TypeBlock
	newblock.Pieces = append(newblock.Pieces, b.Pieces...)
	return newblock
}

func (b Block) CanMove(m *TetrisMatrix, d Direction, step int) bool {
	target := b.getCopy()
	target.Move(d, step)
	return m.fit(target)
}

func (b Block) CanRotate(m *TetrisMatrix, r Rotation) bool {
	target := b.getCopy()
	target.Rotate(r)
	return m.fit(target)
}

func (b *Block) Rotate(r Rotation) {
	for i, p := range b.Pieces {
		//gets relatives coord
		p.PosMinus(b.Center)
		b.Pieces[i].X = p.X*int(math.Cos(float64(r))) - p.Y*int(math.Sin(float64(r)))
		b.Pieces[i].Y = p.X*int(math.Sin(float64(r))) + p.Y*int(math.Cos(float64(r)))
		b.Pieces[i].PosPlus(b.Center)
	}
}

func (b *Block) Move(d Direction, step int) {
	switch d {
	case Left:
		(*b).Center.X -= step
	case Right:
		(*b).Center.X += step
	default:
		(*b).Center.Y += step
	}
	b.UpdateBlockPieces(d, step)
}

func (b *Block) UpdateBlockPieces(d Direction, step int) {
	for i := range b.Pieces {
		switch d {
		case Left:
			b.Pieces[i].X -= step
		case Right:
			b.Pieces[i].X += step
		default:
			b.Pieces[i].Y += step
		}
	}
}

func GenerateBlockByType(t BlockType, at Pos) Block {
	var b Block
	b.Pieces = make([]Pos, 3)
	b.Center = at
	b.TypeBlock = t
	switch b.TypeBlock {
	case I:
		b.Color = Yellow
		b.Pieces[0].X = b.Center.X
		b.Pieces[0].Y = b.Center.Y - 1
		b.Pieces[1].X = b.Center.X
		b.Pieces[1].Y = b.Center.Y + 1
		b.Pieces[2].X = b.Center.X
		b.Pieces[2].Y = b.Center.Y + 2
	case J:
		b.Color = Red
		b.Pieces[0].X = b.Center.X
		b.Pieces[0].Y = b.Center.Y - 1
		b.Pieces[1].X = b.Center.X - 1
		b.Pieces[1].Y = b.Center.Y - 1
		b.Pieces[2].X = b.Center.X
		b.Pieces[2].Y = b.Center.Y + 1
	case L:
		b.Color = Blue
		b.Pieces[0].X = b.Center.X
		b.Pieces[0].Y = b.Center.Y - 1
		b.Pieces[1].X = b.Center.X + 1
		b.Pieces[1].Y = b.Center.Y - 1
		b.Pieces[2].X = b.Center.X
		b.Pieces[2].Y = b.Center.Y + 1
	case Z:
		b.Color = Purple
		b.Pieces[0].X = b.Center.X + 1
		b.Pieces[0].Y = b.Center.Y
		b.Pieces[1].X = b.Center.X
		b.Pieces[1].Y = b.Center.Y + 1
		b.Pieces[2].X = b.Center.X - 1
		b.Pieces[2].Y = b.Center.Y + 1
	case S:
		b.Color = Orange
		b.Pieces[0].X = b.Center.X - 1
		b.Pieces[0].Y = b.Center.Y
		b.Pieces[1].X = b.Center.X
		b.Pieces[1].Y = b.Center.Y + 1
		b.Pieces[2].X = b.Center.X + 1
		b.Pieces[2].Y = b.Center.Y + 1
	case T:
		b.Color = Pink
		b.Pieces[0].X = b.Center.X + 1
		b.Pieces[0].Y = b.Center.Y
		b.Pieces[1].X = b.Center.X - 1
		b.Pieces[1].Y = b.Center.Y
		b.Pieces[2].X = b.Center.X
		b.Pieces[2].Y = b.Center.Y + 1
	case O:
		b.Color = Green
		b.Pieces[0].X = b.Center.X + 1
		b.Pieces[0].Y = b.Center.Y
		b.Pieces[1].X = b.Center.X
		b.Pieces[1].Y = b.Center.Y + 1
		b.Pieces[2].X = b.Center.X + 1
		b.Pieces[2].Y = b.Center.Y + 1
	}
	return b
}
