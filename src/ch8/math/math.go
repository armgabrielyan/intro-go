package math

// Finds the average of a list of numbers
func Average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	total := float64(0)

	for _, n := range numbers {
		total += n
	}

	return total / float64(len(numbers))
}

// Finds the maximum of a list of numbers
func Max(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	m := numbers[0]

	for _, n := range numbers {
		if n > m {
			m = n
		}
	}

	return m
}

// Finds the minimum of a list of numbers
func Min(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	m := numbers[0]

	for _, n := range numbers {
		if n < m {
			m = n
		}
	}

	return m
}
