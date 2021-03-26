package arrays_and_slices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numberArrays ...[]int) (sums []int) {
	for _, numbers := range numberArrays {
		sums = append(sums, Sum(numbers))
	}

	return
}

func SumTail(numbers []int) (sum int) {
	for i, number := range numbers {
		if i != 0 {
			sum += number
		}
	}

	return
}

func SumTails(numberArrays ...[]int) (sums []int) {
	for _, numbers := range numberArrays {
		sums = append(sums, SumTail(numbers))
	}

	return
}
