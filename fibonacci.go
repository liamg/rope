package rope

var fibonacci []int

func init() {
	a, b := 0, 1
	for i := 0; i < 48; i++ {
		fibonacci = append(fibonacci, a)
		a, b = b, a+b
	}
}
