package mypk2

const y = 2
func Sqrt100(x float64) float64 {
	z := 0.0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
func SumAndProduct(A, B int) (add int, Multiplied int) {
    add = A+B
    Multiplied = A*B
    test()
    return
}