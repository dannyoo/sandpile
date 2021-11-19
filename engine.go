package main

// coppies a matrix
func copy(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	output := make([][]int, rows)
	for i := range output {
		output[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			output[i][j] = matrix[i][j]
		}
	}
	return output
}

func SMP(checkerboard [][]int, numProcs int) {
	//rows per proc
	rpp := len(checkerboard) / numProcs
	notStable := true
	for notStable {
		notStable = false
		var topChannels = []chan []int{}
		var bottomChannels = []chan []int{}

		for i := 0; i < numProcs; i++ {
			if i == 0 {
				belowChannel := make(chan []int)
				bottomChannels = append(bottomChannels, belowChannel)
				go spsp1(checkerboard[:rpp], belowChannel, true)
				// spsp1
			} else if i == numProcs-1 { // last
				aboveChannel := make(chan []int)
				topChannels = append(topChannels, aboveChannel)
				go spsp1(checkerboard[rpp:], aboveChannel, false)
				//spsp1
			} else {
				aboveChannel := make(chan []int)
				belowChannel := make(chan []int)
				topChannels = append(topChannels, aboveChannel)
				bottomChannels = append(bottomChannels, belowChannel)
				go spsp2(checkerboard[i*rpp:(i+1)*rpp], aboveChannel, belowChannel)
			}
		}

		// check if coins went over on the top channel
		for i, topSlice := range topChannels {
			slice := <-topSlice
			if notStable {
				addToA(checkerboard[(i+1)*rpp-1], slice)
			} else {
				notStable = addToA(checkerboard[(i+1)*rpp-1], slice)
			}
		}
		// check if coins went over on the bottom channel
		for i, bottomSlice := range bottomChannels {
			slice := <-bottomSlice
			if notStable {
				addToA(checkerboard[((1+i)*rpp)+1], slice)
			} else {
				notStable = addToA(checkerboard[((1+i)*rpp)+1], slice)
			}
		}

	}
}

// add elements from second slice to first slice
// returns whether any coins crossed borders
func addToA(a []int, b []int) (notStable bool) {
	for i, x := range b {
		a[i] += x
		if x > 0 {
			notStable = true
		}
	}
	return
}

// single proc sandpile for 2 channel
func spsp2(checkerboard [][]int, aboveChannel chan []int, belowChannel chan []int) {
	aboveSlice := make([]int, len(checkerboard[0]))
	belowSlice := make([]int, len(checkerboard[0]))

	// we no longer have symmetrical board
	rows := len(checkerboard)
	cols := len(checkerboard[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			pile := checkerboard[i][j]
			if pile >= 4 {
				//coins per neighbor
				cpn := (pile - (pile % 4)) / 4
				if i+1 < rows { // S
					checkerboard[i+1][j] += cpn
				} else {
					belowSlice[j] = cpn
				}
				if j+1 < cols { //E
					checkerboard[i][j+1] += cpn
				}
				if i-1 > 0 { // N
					checkerboard[i-1][j] += cpn
				} else {
					aboveSlice[j] = cpn
				}
				if j-1 > 0 { // W
					checkerboard[i][j-1] += cpn
				}
				// update self
				checkerboard[i][j] = pile - (pile - (pile % 4))
			}
		}
	}

	aboveChannel <- aboveSlice
	belowChannel <- belowSlice

}

// single proc sandpile for 1 channel
func spsp1(checkerboard [][]int, channel chan []int, isFirst bool) {
	slice := make([]int, len(checkerboard[0]))

	// we no longer have symmetrical board
	rows := len(checkerboard)
	cols := len(checkerboard[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			pile := checkerboard[i][j]
			if pile >= 4 {
				//coins per neighbor
				cpn := (pile - (pile % 4)) / 4
				if i+1 < rows { // S
					checkerboard[i+1][j] += cpn
				} else if isFirst {
					slice[j] = cpn
				}
				if j+1 < cols { //E
					checkerboard[i][j+1] += cpn
				}
				if i-1 > 0 { // N
					checkerboard[i-1][j] += cpn
				} else if !isFirst {
					slice[j] = cpn
				}
				if j-1 > 0 { // W
					checkerboard[i][j-1] += cpn
				}
				// update self
				checkerboard[i][j] = pile - (pile - (pile % 4))
			}
		}
	}

	channel <- slice
}

func sandpile(checkerboard [][]int) {
	notStable := true
	for notStable {
		notStable = false
		notStable = topple(checkerboard, notStable)
	}

}

func topple(checkerboard [][]int, notStable bool) bool {
	size := len(checkerboard)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			pile := checkerboard[i][j]
			if pile >= 4 {
				//coins per neighbor
				cpn := (pile - (pile % 4)) / 4
				if i+1 < size {
					checkerboard[i+1][j] += cpn
				}
				if j+1 < size {
					checkerboard[i][j+1] += cpn
				}
				if i-1 > 0 {
					checkerboard[i-1][j] += cpn
				}
				if j-1 > 0 {
					checkerboard[i][j-1] += cpn
				}

				// update self
				checkerboard[i][j] = pile - (pile - (pile % 4))
				notStable = true
			}
		}
	}
	return notStable
}
