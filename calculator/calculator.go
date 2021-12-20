package calculator

func Addition(a, b int) int {
	result := a + b
	return result
}

func Pengurangan(a, b int) int {
	if a > 10 {
		return a
	}
	return a - b
}

func Pembagian(a, b int) int {
	return a / b
}

func Perkalian(a, b int) int {
	return a * b
}
