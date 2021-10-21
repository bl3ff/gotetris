package control

import (
	"fmt"
	"log"
	"tetris/model"
	"time"
)

type EngineT struct {
	TMatrix      *model.TetrisMatrix
	CurrentBlock model.Block
	LoopDelay    int
	Points       int
	Lose         bool
	newGameChan  chan int
	reDrawChan   chan int
	commandChan  chan interface{}

	redrawView func()
}

var Engine EngineT

type Click struct {
	Right bool
	Left  bool
}

type Key struct {
	DirKey map[model.Direction]bool
}

func GenerateEngine() *EngineT {
	engine := &EngineT{
		LoopDelay:   LDelay,
		commandChan: make(chan interface{}, 20),
		newGameChan: make(chan int, 1),
		reDrawChan:  make(chan int, 1),
	}

	return engine
}

func (e *EngineT) InitNewGame() {
	fmt.Println("Init Game Engine")
	e.TMatrix = model.GenerateMatrix(model.Rows, model.Cols)
	e.CurrentBlock = model.GenerateRandomBlock(model.Pos{X: model.DefaultX, Y: model.DefaultY})
	e.Points = 0
}

func (e *EngineT) RestartGame() {
	fmt.Println("New Game Engine")
	e.TMatrix.Lock()
	e.TMatrix.Clean()
	e.TMatrix.Unlock()
	e.CurrentBlock = model.GenerateRandomBlock(model.Pos{X: model.DefaultX, Y: model.DefaultY})
	e.Points = 0
	e.Lose = false
}

// SendClick sends a click event from the user.
func (e *EngineT) SendClick(c Click) {
	e.commandChan <- &c
}

// SendKey sends a key event from the user.
func (e *EngineT) SendKey(k Key) {
	e.commandChan <- &k
}

func (e *EngineT) Start(redraw func()) {
	fmt.Println("Start Engine")
	e.InitNewGame()
	e.redrawView = redraw

	go e.loop()
}

func (e *EngineT) loop() {
	fmt.Println("Start Engine Loop")
	ticker := time.NewTicker(time.Duration(e.LoopDelay) * time.Millisecond)

	for {
		select {
		case <-e.newGameChan:
			e.RestartGame()
			e.redrawView()
			//e.TMatrix.Print()
		default:
		}

		e.TMatrix.Lock()

		if !e.canMoveDown() {
			fmt.Println("You lose and totalize ", e.Points)
			e.Lose = true
			time.Sleep(1 * time.Second)
			e.newGameChan <- 1
		}

		e.processInput()

		//update currentBlock
		if !e.Lose {
			e.moveBlock(model.Down)

			e.TMatrix.Unlock()
			e.redrawView()
			e.TMatrix.Lock()

			if !e.canMoveDown() {
				e.TMatrix.StoreBlock(e.CurrentBlock)
				e.Points += 100 * e.TMatrix.Update()
				//e.TMatrix.Print()
				e.CurrentBlock = model.GenerateRandomBlock(model.Pos{X: model.DefaultX, Y: model.DefaultY})
			}
		}

		e.TMatrix.Unlock()
		e.redrawView()
		<-ticker.C
	}

}

func (e *EngineT) canMoveDown() bool {
	return e.CurrentBlock.CanMove(e.TMatrix, model.Down, 1)
}

func (e *EngineT) moveBlock(d model.Direction) bool {
	if e.CurrentBlock.CanMove(e.TMatrix, d, 1) {
		e.CurrentBlock.Move(d, 1)
		return true
	}
	return false
}

func (e *EngineT) rotateBlock(r model.Rotation) bool {
	if e.CurrentBlock.CanRotate(e.TMatrix, r) {
		e.CurrentBlock.Rotate(r)
		return true
	}
	return false
}

func (e *EngineT) processInput() {
	for {
		select {
		case cmd := <-e.commandChan:

			switch cmd := cmd.(type) {
			case *Click:
				e.handleClick(cmd)
			case *Key:
				e.handleKey(cmd)
			default:
				log.Printf("Unhandled cmd type: %T", cmd)
			}

		default:
			return
		}
	}
}

func (e *EngineT) handleClick(c *Click) {
	if e.Lose {
		return
	}

	if c.Left {
		e.rotateBlock(model.CounterClockWise)
	}

	if c.Right {
		e.rotateBlock(model.ClockWise)
	}

}

func (e *EngineT) handleKey(k *Key) {
	if e.Lose {
		return
	}

	for k, v := range k.DirKey {
		if v {
			e.moveBlock(model.Direction(k))
			return
		}
	}

}
