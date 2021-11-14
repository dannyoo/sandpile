package main

import (
	"canvas"
	"fmt"
	"image"
	"image/png"
	"os"
	"strconv"
)

func DrawGameBoards(boards [][][]int, cellWidth int) []image.Image {
	numGenerations := len(boards)
	imageList := make([]image.Image, numGenerations)
	for i := range boards {
		imageList[i] = DrawGameBoard(boards[i], cellWidth)
	}
	return imageList
}

func DrawGameBoard(board [][]int, cellWidth int) image.Image {
	height := len(board) * cellWidth
	width := len(board[0]) * cellWidth
	fmt.Println(len(board[0]))
	c := canvas.CreateNewPalettedCanvas(width, height, nil)

	// declare colors
	// darkGray := canvas.MakeColor(50, 50, 50)
	black := canvas.MakeColor(0, 0, 0)
	blue := canvas.MakeColor(85, 85, 85)
	red := canvas.MakeColor(170, 170, 170)
	green := canvas.MakeColor(255, 255, 255)


	// fill in colored squares
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				c.SetFillColor(black)
			} else if board[i][j] == 1 {
				c.SetFillColor(blue)
			} else if board[i][j] == 2 {
				c.SetFillColor(red)
			} else if board[i][j] == 3 {
				c.SetFillColor(green)
			} else {
				panic("Error: Out of range value " + strconv.Itoa(board[i][j]) + " in board when drawing board.")
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}

	return canvas.GetImage(c)
}

func DrawGridLines(pic canvas.Canvas, cellWidth int) {
	w, h := pic.Width(), pic.Height()
	// first, draw vertical lines
	for i := 1; i < w/cellWidth; i++ {
		y := i * cellWidth
		pic.MoveTo(0.0, float64(y))
		pic.LineTo(float64(w), float64(y))
	}
	// next, draw horizontal lines
	for j := 1; j < h/cellWidth; j++ {
		x := j * cellWidth
		pic.MoveTo(float64(x), 0.0)
		pic.LineTo(float64(x), float64(h))
	}
	pic.Stroke()
}

// saves png of image called "Prisoners.png" to current directory
func SavePNG(i image.Image, filename string) {
	// creates new file
	f, err := os.Create(filename)
	if err != nil {
		// Handle error
		panic("problem with creating png file")
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, i)
	if err != nil {
		// Handle error
		panic("problem with encoding image")
	}
}
