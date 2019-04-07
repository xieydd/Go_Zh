package bubblesort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	values := []int{4, 2, 7, 9, 5, 0}
	BubbleSort(values)
	if values[0] != 0 || values[1] != 2 || values[2] != 4 || values[3] != 5 || values[4] != 7 || values[5] != 9 {
		t.Error("BubbleSort failed. Got: ", values, "not Expected 0 2 4 5 7 9")
	}
}
