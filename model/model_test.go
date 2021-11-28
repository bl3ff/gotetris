package model

import (
	"testing"
)

func TestInit(t *testing.T) {
	m := GenerateMatrix(10, 5)
	m.Print()
}

func TestUpdate(t *testing.T) {
	b := GenerateBlockByType(O, Pos{X: 2, Y: 0})
	m := GenerateMatrix(10, 5)
	for i := 0; i < m.Cols; i++ {
		if i != 2 && i != 3 {
			m.Matrix[9][i] = 3
		}
	}
	m.Print()
	for b.CanMove(m, Down, 1) {
		b.Move(Down, 1)
	}
	m.StoreBlock(b)
	m.Print()
	m.Update()
	m.Print()
}

func TestUpdate1(t *testing.T) {
	b := GenerateBlockByType(O, Pos{X: 2, Y: 0})
	m := GenerateMatrix(10, 5)
	for i := 0; i < m.Cols; i++ {
		if i != 2 && i != 3 {
			m.Matrix[9][i] = 3
		}
	}
	m.Print()
	b.Move(Left, 1)
	for b.CanMove(m, Down, 1) {
		b.Move(Down, 1)
	}
	m.StoreBlock(b)
	m.Print()
	m.Update()
	m.Print()
}

func fillRow(m *TetrisMatrix, row int) {
	for i := 0; i < m.Cols; i++ {
		m.Matrix[row][i] = 3
	}
}

func TestUpdate2(t *testing.T) {
	m := GenerateMatrix(10, 5)

	fillRow(m, 7)
	fillRow(m, 8)
	fillRow(m, 9)

	m.Print()
	m.Update()
	m.Print()
}

func TestCanMove(t *testing.T) {
	m := GenerateMatrix(4, 4)
	b := GenerateBlockByType(O, Pos{X: 1, Y: 0})
	if !b.CanMove(m, Down, 1) {
		t.Fatalf("block should go down")
	}
	if !b.CanMove(m, Right, 1) {
		t.Fatalf("block should go right")
	}
	if !b.CanMove(m, Left, 1) {
		t.Fatalf("block should go Left")
	}
	b.Move(Down, 1)
	if !b.CanMove(m, Down, 1) {
		t.Fatalf("block should go down")
	}
	if !b.CanMove(m, Right, 1) {
		t.Fatalf("block should go right")
	}
	if !b.CanMove(m, Left, 1) {
		t.Fatalf("block should go Left")
	}
	b.Move(Down, 1)
	if b.CanMove(m, Down, 1) {
		t.Fatalf("block should not go down")
	}
	if !b.CanMove(m, Right, 1) {
		t.Fatalf("block should go right")
	}
	if !b.CanMove(m, Left, 1) {
		t.Fatalf("block should go Left")
	}
	b.Move(Left, 1)
	if b.CanMove(m, Down, 1) {
		t.Fatalf("block should not go down")
	}
	if !b.CanMove(m, Right, 1) {
		t.Fatalf("block should go right")
	}
	if b.CanMove(m, Left, 1) {
		t.Fatalf("block should not go Left")
	}
	b.Move(Right, 1)
	if b.CanMove(m, Down, 1) {
		t.Fatalf("block should not go down")
	}
	if !b.CanMove(m, Right, 1) {
		t.Fatalf("block should go right")
	}
	if !b.CanMove(m, Left, 1) {
		t.Fatalf("block should go Left")
	}
	b.Move(Right, 1)
	if b.CanMove(m, Down, 1) {
		t.Fatalf("block should not go down")
	}
	if b.CanMove(m, Right, 1) {
		t.Fatalf("block should not go right")
	}
	if !b.CanMove(m, Left, 1) {
		t.Fatalf("block should go Left")
	}
}

func TestCanMoveI(t *testing.T) {
	m := GenerateMatrix(4, 4)
	b := GenerateBlockByType(I, Pos{0,1})
	if b.CanMove(m, Down, 1) {
		t.Fatalf("block should not go down")
	}
	if !b.CanMove(m, Right, 1) {
		t.Fatalf("block should go right")
	}
	if b.CanMove(m, Left, 1) {
		t.Fatalf("block should not go Left")
	}

	if !b.CanRotate(m, ClockWise) {
		t.Fatalf("block should Rotate CW")
	}
	if b.CanRotate(m, CounterClockWise) {
		t.Fatalf("block shouldn't Rotate CCW")
	}

}

func TestCanRotate(t *testing.T) {
	m := GenerateMatrix(4, 4)
	b := GenerateBlockByType(S, Pos{1, 0})

	if b.CanRotate(m, ClockWise) {
		t.Fatalf("block shouldn't Rotate CW")
	}
	if b.CanRotate(m, CounterClockWise) {
		t.Fatalf("block shouldn't Rotate CCW")
	}
	b.Move(Down, 1)
	b.Rotate(CounterClockWise)
	if !b.CanMove(m, Down, 1) {
		t.Fatalf("block should go down")
	}
	if !b.CanMove(m, Right, 1) {
		t.Fatalf("block should go right")
	}
	if !b.CanMove(m, Left, 1) {
		t.Fatalf("block should go Left")
	}
	b.Rotate(ClockWise)
	b.Rotate(ClockWise)
	if !b.CanMove(m, Down, 1) {
		t.Fatalf("block should go down")
	}
	if !b.CanMove(m, Right, 1) {
		t.Fatalf("block should go right")
	}
	if !b.CanMove(m, Left, 1) {
		t.Fatalf("block should go Left")
	}
}

func clearBlock(b Block, m *TetrisMatrix) {
	m.Matrix[b.Center.Y][b.Center.X] = EmptyColor
	for _, p := range b.Pieces {
		m.Matrix[p.Y][p.X] = EmptyColor
	}
}

func TestRotateS(t *testing.T) {
	m := GenerateMatrix(10, 10)
	b := GenerateBlockByType(S, Pos{4, 0})

	b.Move(Down, 3)
	c := b.Center
	m.StoreBlock(b)
	m.Print()
	clearBlock(b, m)
	b.Rotate(ClockWise)
	d := b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d = b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d = b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}

}

func TestRotateI(t *testing.T) {
	m := GenerateMatrix(10, 10)
	b := GenerateBlockByType(I, Pos{5, 1})

	b.Move(Down, 3)
	c := b.Center
	m.StoreBlock(b)
	m.Print()

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d := b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d = b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d = b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}
}

func TestRotateT(t *testing.T) {
	m := GenerateMatrix(10, 10)
	b := GenerateBlockByType(T, Pos{5, 1})

	b.Move(Down, 3)
	c := b.Center
	m.StoreBlock(b)
	m.Print()

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d := b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d = b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d = b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}
}

func TestRotateJ(t *testing.T) {
	m := GenerateMatrix(10, 10)
	b := GenerateBlockByType(J, Pos{5, 1})

	b.Move(Down, 3)
	c := b.Center
	m.StoreBlock(b)
	m.Print()

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d := b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d = b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}

	clearBlock(b, m)
	b.Rotate(ClockWise)
	d = b.Center
	m.StoreBlock(b)
	m.Print()
	if c != d {
		t.Fatal("Invalid center after rotation")
	}
}
