package model

import (
	"testing"
)

func TestGetByBlockType(t *testing.T) {
	b := GenerateBlockByType(I, Pos{})

	if (b.Center != Pos{}) {
		t.Fatalf("Center isn't in start position")
	}

	if b.Color != Yellow {
		t.Fatalf("block hasn't the correct color")
	}

	if b.TypeBlock != I {
		t.Fatalf("block hasn't the correct Type")
	}

	if (b.Pieces[0] != Pos{0, -1} || b.Pieces[1] != Pos{0, 1} || b.Pieces[2] != Pos{0, 2}) {
		t.Fatalf("block hasn't the correct shape")
	}

}

func TestMove(t *testing.T) {
	b := GenerateBlockByType(I, Pos{})
	b.Move(Down, 2)

	if (b.Center != Pos{0, 2}) {
		t.Fatalf("block hasn't the correct position")
	}
	if (b.Pieces[0] != Pos{0, 1} &&
		b.Pieces[1] != Pos{0, 3} &&
		b.Pieces[2] != Pos{0, 4}) {
		t.Fatalf("block hasn't the correct position")
	}
}
