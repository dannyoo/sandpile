package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Sandpile")

	// Read in parameteres
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Error: Problem converting checkerboard width parameter to an integer.")
	}

	numCoins, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("Error: Problem converting number of coins to be placed on the pile parameter to an integer.")
	}

	placement := os.Args[3]
	if placement != "central" && placement != "random" {
		panic("placement param must be central or random")
	}

	//create checkerboard
	checkerboard := make([][]int, size)
	for i := 0; i < len(checkerboard); i++ {
		checkerboard[i] = make([]int, size)
	}

	// place coins
	if placement == "central" {
		checkerboard[size/2][size/2] = numCoins
	} else if placement == "random" {
		for i := 0; i <= numCoins; i++ {
			rand.Seed(time.Now().UnixNano())
			randR := rand.Intn(size)
			randC := rand.Intn(size)
			checkerboard[randR][randC]++
		}
	}
	//create board copy
	checkerboard2 := copy(checkerboard)
	// printMatrix(checkerboard)
	// fmt.Println("===")
	// printMatrix(checkerboard2)

	// topple logic for parallel and serial
	// stablize(checkerboard) // it's pointers underneath no duplicates
	start := time.Now()
	sandpile(checkerboard)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Serial took", elapsed.Seconds(), "seconds.")

	numProcs := runtime.NumCPU()
	// fmt.Println(numProcs)
	start = time.Now()
	SMP(checkerboard2, numProcs)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("Parallel took", elapsed.Seconds(), "seconds.")

	// Generate two drawings
	cellWidth := 2 // what was used for cell width
	list := make([][][]int, 0)
	list = append(list, checkerboard)
	list = append(list, checkerboard2)
	imgs := DrawGameBoards(list, cellWidth)


	// save the boards as pngs
	SavePNG(imgs[0], "serial.png")
	SavePNG(imgs[1], "parallel.png")

	fmt.Println("./sandpile", size, numCoins, placement)

}





