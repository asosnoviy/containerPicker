package picker

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var picker *Picker

func init() {

	picker = NewPicker()
}

func TestPicker_Calculate(t *testing.T) {

	example := []int{
		2, 3, 5, 7, 10, 15,
	}

	answer := picker.Calculate(example, 21)
	assert.Equal(t, true, answer.Can)
	assert.Equal(t, 4, answer.Count)
	assert.Equal(t, []int{0, 1, 2, 4}, answer.Nums)
	fmt.Println(answer.Nums)
}

func TestPicker_CalculateZero(t *testing.T) {

	example := []int{
		2, 3, 5, 7, 10, 15,
	}

	answer := picker.Calculate(example, 0)
	assert.Equal(t, false, answer.Can)
	assert.Equal(t, 0, answer.Count)
	assert.Equal(t, []int{}, answer.Nums)

	answer = picker.Calculate(example, 1)
	assert.Equal(t, false, answer.Can)
	assert.Equal(t, 0, answer.Count)
	assert.Equal(t, []int{}, answer.Nums)

	answer = picker.Calculate(example, 2)
	assert.Equal(t, true, answer.Can)
	assert.Equal(t, 1, answer.Count)
	assert.Equal(t, []int{0}, answer.Nums)

}

func TestPicker_CalculateMore(t *testing.T) {

	example := []int{
		20, 20, 15, 20, 15, 10, 10, 20, 10, 20, 15, 10, 15, 10, 10, 20, 15, 20, 20, 15, 20, 15, 15, 10, 15, 20, 15, 15, 10, 10, 10, 10, 10, 10, 10, 15,
		10, 15, 100, 20, 15, 20, 26, 100, 15, 67, 15, 7, 10, 20, 100, 100, 20, 20, 20, 15, 20, 100, 100}

	answer := picker.Calculate(example, 100)
	assert.Equal(t, true, answer.Can)
	assert.Equal(t, 10, answer.Count)
	assert.Equal(t, []int{5, 6, 8, 11, 13, 14, 23, 28, 29, 30}, answer.Nums)

}

func TestPicker_CalculateMega(t *testing.T) {

	cntConteiners := 10000
	var example = make([]int, cntConteiners)
	for i := 0; i < cntConteiners; i++ {
		example[i] = rand.Intn(1000)
	}

	startTime := time.Now()
	answer := picker.Calculate(example, 100)
	fmt.Println(cntConteiners, "X", len(example))
	fmt.Println(answer.Can)
	fmt.Println(answer.Count)
	fmt.Println(answer.Nums)
	fmt.Println(time.Since(startTime))
}
