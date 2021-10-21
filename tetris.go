package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"tetris/control"
	"tetris/model"
	"tetris/view"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

func processFlags() error {

	// Model package flags
	flag.IntVar(&model.Rows, "rows", 20, "the number of rows in the Labyrinth; must be odd; valid range: 9..99")
	flag.IntVar(&model.Cols, "cols", 15, "the number of columns in the Labyrinth; must be odd; valid range: 9..99")

	// Control/Engine flags
	flag.IntVar(&control.LDelay, "loopDelay", 500, "loop delay of the game engine, in milliseconds; valid range: 10..1500")
	//flag.Float64Var(&control.V, "v", model.BlockSize*2.0, "moving speed of Gopher and the Bulldogs in pixel/sec; valid range: 20..200")

	flag.Parse()

	if model.Rows < 9 || model.Rows > 99 {
		return fmt.Errorf("rows %d is outside of valid range", model.Rows)
	}

	if model.Cols < 9 || model.Cols > 99 {
		return fmt.Errorf("cols %d is outside of valid range", model.Cols)
	}

	model.TWidth = model.Cols * model.BlockSize
	model.THeight = model.Rows * model.BlockSize
	model.DefaultX = model.Cols / 2
	model.DefaultY = 2

	if control.LDelay < 10 || control.LDelay > 1500 {
		return fmt.Errorf("loopDelay %d is outside of valid range", control.LDelay)
	}

	return nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := processFlags(); err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	driver.Main(func(s screen.Screen) {
		fmt.Println("####### START TETRIS ########")
		fmt.Printf("# w=%d  h=%d #\n", model.TWidth, model.THeight)
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Title:  "Tetris GO!",
			Width:  model.TWidth,
			Height: model.THeight,
		})

		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()
		view := view.NewTetrisView(&w)
		view.Start()
	})
}
