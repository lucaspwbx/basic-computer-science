package sorting

func BubbleSort(list []int) []int {
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

func SelectionSort(list []int) []int {
  var i, j, min, aux int
  for i = 0; i < len(list)-1; i++ {
    min = i
    for j = i+1; j < len(list); j++ {
      if list[j] < list[min] {
	min = j
      }
    }
    if i != min {
      aux = list[i]
      list[i] = list[min]
      list[min] = aux
    }
  }
  return list
}
