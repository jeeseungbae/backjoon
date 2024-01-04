package main_test

import (
	"backjoon/main"
	"fmt"
	"github.com/go-playground/assert"
	"testing"
)

func TestBuildBlocks(t *testing.T) {
	w := []int{1, 3, 3, 4, 5}
	h := []int{3, 1, 2, 9, 5}
	result := []int{4, 13, 84, 132976888, 11676046}
	for i := 0; i < len(w); i++ {
		if main.BuildBlocks(w[i], h[i]) != result[i] {
			t.Errorf("Errors in case %d : Expected %d, Actual %d", i+1, main.BuildBlocks(w[i], h[i]), result[i])
		}
	}
}

func Test1(t *testing.T) {
	w := 1
	h := 3
	result := 4
	assert.Equal(t, result, main.BuildBlocks(w, h))
}

func Test2(t *testing.T) {
	w := 3
	h := 1
	result := 13
	assert.Equal(t, result, main.BuildBlocks(w, h))
}

func Test3(t *testing.T) {
	w := 3
	h := 2
	result := 84
	assert.Equal(t, result, main.BuildBlocks(w, h))
}

func Test4(t *testing.T) {
	w := 4
	h := 9
	result := 132976888
	assert.Equal(t, result, main.BuildBlocks(w, h))
}

func Test5(t *testing.T) {
	w := 5
	h := 5
	result := 11676046
	assert.Equal(t, result, main.BuildBlocks(w, h))
}

func TestGoodIsContain(t *testing.T) {
	ground := []int64{1, 3, 5, 7, 7}
	data := []int64{1, 3, 7, 5, 1}
	result := make([]bool, len(ground))
	for i := 0; i < len(ground); i++ {
		if main.IsContain(ground[i], data[i]) != !result[i] {
			t.Errorf("Errors in case %d : Expected %t, Actual %t", i+1, main.IsContain(ground[i], data[i]), !result[i])
		}
	}
}

func TestFailIsContain(t *testing.T) {
	ground := []int64{1, 3, 5, 1}
	data := []int64{2, 5, 3, 7}
	result := make([]bool, len(ground))
	for i := 0; i < len(ground); i++ {
		if main.IsContain(ground[i], data[i]) != result[i] {
			t.Errorf("Errors in case %d : Expected %t, Actual %t", i+1, main.IsContain(ground[i], data[i]), result[i])
		}
	}
}

func Test(t *testing.T) {
	var a int64 = 3
	str := fmt.Sprintf("%010b", a)
	print(str)
}
