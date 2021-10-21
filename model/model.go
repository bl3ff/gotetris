package model

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TetrisMatrix struct {
	sync.RWMutex
	Matrix [][]BlockColor
	Rows   int
	Cols   int
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func GenerateMatrix(rs int, cs int) *TetrisMatrix {
	fmt.Println("Generate Matrix")
	m := &TetrisMatrix{
		RWMutex: sync.RWMutex{},
		Matrix:  make([][]BlockColor, rs),
		Rows:    rs,
		Cols:    cs,
	}
	for i := range m.Matrix {
		m.Matrix[i] = make([]BlockColor, cs)
	}
	return m
}

func GenerateRandomBlock(at Pos) Block {
	return GenerateBlockByType(BlockType(r.Intn(BlockTypeLen-1)), at)
}

func (m *TetrisMatrix) Clean() {
	for i := range m.Matrix {
		m.Matrix[i] = make([]BlockColor, m.Cols)
	}
}

func (m *TetrisMatrix) fit(b Block) bool {
	for _, p := range b.Pieces {
		if !m.fitPos(p) {
			return false
		}
	}
	return m.fitPos(b.Center)
}

func (m *TetrisMatrix) fitPos(p Pos) bool {
	if p.X < 0 || p.X > m.Cols-1 {
		return false
	}

	if p.Y < 0 || p.Y > m.Rows-1 {
		return false
	}

	if m.Matrix[p.Y][p.X] != 0 {
		return false
	}

	return true
}

func (m *TetrisMatrix) StoreBlock(b Block) {
	m.Matrix[b.Center.Y][b.Center.X] = b.Color
	for _, p := range b.Pieces {
		m.Matrix[p.Y][p.X] = b.Color
	}
}

func (m *TetrisMatrix) Update() int {

	rowsRemoved := []int{}

	for i := 0; i < m.Rows; i++ {
		toRemove := true
		for j := 0; j < m.Cols; j++ {
			if m.Matrix[i][j] == EmptyColor {
				toRemove = false
				break
			}
		}
		if toRemove {
			rowsRemoved = append(rowsRemoved, i)
		}
	}

	for _, i := range rowsRemoved {
		m.removeRow(i)
	}
	return len(rowsRemoved)
}

func (m *TetrisMatrix) Print() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("%d ", m.Matrix[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("##############")
}

func (m *TetrisMatrix) removeRow(r int) {
	for i := r; i > 0; i-- {
		for j := 0; j < m.Cols; j++ {
			m.Matrix[i][j] = m.Matrix[i-1][j]
		}
	}

	// aggiorno prima riga
	for j := 0; j < m.Cols; j++ {
		m.Matrix[0][j] = EmptyColor
	}
}
