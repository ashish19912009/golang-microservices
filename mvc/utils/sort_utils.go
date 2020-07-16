package utils

//Input []int {9,8,7,6,5}
//Outut []int {5,6,7,8,9}

// BubbleSort Package
func BubbleSort(elements []int) {
	keepRunnning := true

	for keepRunnning {
		keepRunnning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunnning = true
			}
		}
	}
}
