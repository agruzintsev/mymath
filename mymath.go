package mymath

import "math"

// Sqrt возвращает квадратный корень из x.
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Ceil возвращает наименьшее целое число, не меньшее x.
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Floor возвращает наибольшее целое число, не превышающее x.
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Pow возвращает x, возведенное в степень n.
func Pow(x float64, n float64) float64 {
	return math.Pow(x, n)
}

// Min возвращает меньшее из x или y.
func Min(x, y float64) float64 {
	return math.Min(x, y)
}

// Max возвращает большее из x или y.
func Max(x, y float64) float64 {
	return math.Max(x, y)
}

// Abs возвращает абсолютное значение x.
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Yn возвращает n-й порядок функции Бесселя второго рода при значении x.
func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
