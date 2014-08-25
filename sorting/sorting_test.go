package sorting

import "testing"

func TestBubblesort(t *testing.T) {
	expected := []int{26, 54, 93}
	list := []int{54, 26, 93}
	result := Bubblesort(list)

	for k, v := range result {
		if v != expected[k] {
			t.Errorf("Error sorting")
		}
	}
}
