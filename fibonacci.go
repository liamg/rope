package rope

var fibonacci []int

func init() {
	a, b := 0, 1
	for i := 0; i < 48; i++ {
		a, b = b, a+b
		fibonacci = append(fibonacci, a)
	}
}
