package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	inputs := strings.Split(scanner.Text(), " ")
	w, _ := strconv.Atoi(inputs[0])
	h, _ := strconv.Atoi(inputs[1])

	fmt.Println(BuildBlocks(w, h) / 1000000007)
}

func BuildBlocks(w int, h int) int {
	answer := 1
	basicBlock := []int{1, 2, 3}
	block := make([]bool, w)
	for i := 0; i < w; i++ {
		block[i] = true
	}

	stacking(block, make([]bool, w), basicBlock, 0, w, h, &answer)

	return answer
}

func stacking(block []bool, upBlock []bool, basicBlock []int, start int, w int, h int, answer *int) {
	if h == 0 {
		return
	} else if start >= w {
		return
	}

	for i := start; i < w; i++ {
		if block[i] {
			for j := 0; j < len(basicBlock); j++ {
				if i+basicBlock[j] <= w && block[i+basicBlock[j]-1] {
					*answer += 1
					blockBuild(upBlock, true, i, i+basicBlock[j])

					//	위로 쌓기
					stacking(upBlock, make([]bool, len(block)), basicBlock, 0, w, h-1, answer)

					// 	가로로 이동
					stacking(block, upBlock, basicBlock, i+basicBlock[j], w, h, answer)

					blockBuild(upBlock, false, i, i+basicBlock[j])
				}
			}
		}
	}
}

func blockBuild(block []bool, flag bool, start int, end int) []bool {
	for i := start; i < end; i++ {
		block[i] = flag
	}
	return block
}
