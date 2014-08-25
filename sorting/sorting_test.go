package sorting

import "testing"

func TestBubblesort(t *testing.T) {
	list := []int{54, 26, 93}
	expected := []int{26, 54, 93}
	result := BubbleSort(list)

	for k, v := range result {
		if v != expected[k] {
			t.Errorf("Error sorting")
		}
	}
}

func TestSelectionSort(t *testing.T) {
	list := []int{54, 26, 93}
	expected := []int{26, 54, 93}
	result := SelectionSort(list)

	for k, v := range result {
		if v != expected[k] {
			t.Errorf("Error sorting")
		}
	}
}

func TestInsertionSort(t *testing.T) {
	list := []int{54, 26, 93}
	expected := []int{26, 54, 93}
	result := InsertionSort(list)

	for k, v := range result {
		if v != expected[k] {
			t.Errorf("Error sorting")
		}
	}
}
