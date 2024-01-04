package main_test

import (
	"backjoon/main"
	"github.com/go-playground/assert"
	"testing"
)

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
