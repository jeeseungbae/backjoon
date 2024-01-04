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

	checkBlock(block, make([]bool, w), basicBlock, 0, w, h, &answer)

	return answer
}

func checkBlock(block []bool, upBlock []bool, basicBlock []int, start int, w int, h int, answer *int) {
	if h == 0 {
		return
	}
	if start >= w {
		return
	}
	for i := 0; i < len(basicBlock); i++ {
		if block[start] && start+basicBlock[i]-1 < w && block[start+basicBlock[i]-1] {
			*answer += 1
			blockBuild(upBlock, true, start, start+basicBlock[i])

			stacking(upBlock, basicBlock, w, h, answer)

			moving(block, upBlock, basicBlock, start+basicBlock[i], w, h, answer)

			blockBuild(upBlock, false, start, start+basicBlock[i])
		}
	}
	// 이동
	moving(block, upBlock, basicBlock, start+1, w, h, answer)
}

func stacking(upBlock []bool, basicBlock []int, w int, h int, answer *int) {
	// 위로 쌓기
	checkBlock(upBlock, make([]bool, w), basicBlock, 0, w, h-1, answer)
}

func moving(block []bool, upBlock []bool, basicBlock []int, start int, w int, h int, answer *int) {
	// 쌓고 이동
	checkBlock(block, upBlock, basicBlock, start, w, h, answer)
}

func blockBuild(upBlock []bool, flag bool, start int, end int) []bool {
	for i := start; i < end; i++ {
		upBlock[i] = flag
	}
	return upBlock
}
