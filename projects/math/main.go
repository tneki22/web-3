// Функция для нахождения массы
func M() float64 {
	return p * v
}

// Функция для нахождения циклической частоты
func W() float64 {
	return math.Sqrt(k / M())
}

// Функция для нахождения периода колебаний
func T() float64 {
	return 6 / W()
}