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
	basicBlock := []int64{1, 3, 5}

	ground := int64(0)
	for i := 0; i < w; i++ {
		ground = ground << 1
		ground += 1
	}

	checkBlock(ground, 0, basicBlock, 0, w, h, &answer)

	return answer
}

func checkBlock(ground int64, block int64, basicBlock []int64, start int, w int, h int, answer *int) {
	if h == 0 {
		return
	}
	if start >= w {
		return
	}
	if block < 0 {
		return
	}

	for i := 0; i < len(basicBlock); i++ {
		if w >= start+i+1 {
			data := block<<i + 1
			data = data | basicBlock[i]

			if IsContain(ground, data) {
				*answer += 1
				stacking(data, basicBlock, w, h-1, answer)
				moving(ground, data, basicBlock, start+i+1, w, h, answer)
			}
		}
	}
	moving(ground, block<<1, basicBlock, start+1, w, h, answer)
}

func IsContain(ground int64, data int64) bool {
	groundBlock := fmt.Sprintf("%010b", ground)
	dataBlock := fmt.Sprintf("%010b", data)

	for i := 0; i < len(dataBlock); i++ {
		if dataBlock[i] == '1' {
			ln, rn := blockJump(dataBlock, i)
			if enableBlock(groundBlock, dataBlock, ln, rn-1) {
				i += rn - 1 - ln
			} else {
				return false
			}
		}
	}
	return true
}

func blockJump(dataBlock string, index int) (int, int) {
	for i := index + 1; i < len(dataBlock); i++ {
		if dataBlock[i] != '1' {
			return index, i
		}
	}
	return index, len(dataBlock)
}

func enableBlock(groundBlock string, dataBlock string, start int, end int) bool {
	if groundBlock[start] == dataBlock[start] && groundBlock[end] == dataBlock[end] {
		return true
	}
	return false
}

func stacking(data int64, basicBlock []int64, w int, h int, answer *int) {
	checkBlock(data, 0, basicBlock, 0, w, h, answer)
}

func moving(ground int64, data int64, basicBlock []int64, start int, w int, h int, answer *int) {
	checkBlock(ground, data, basicBlock, start, w, h, answer)
}
