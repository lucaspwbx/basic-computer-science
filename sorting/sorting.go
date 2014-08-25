package sorting

func Bubblesort(list []int) []int {
	for i := len(list) - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}
	return list
}
