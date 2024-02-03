package templatefuncs

func Sub(a int, b int) int { return a - b }

func Span(arg1 int, args ...int) []int {
	start := 0
	end := 0
	if len(args) == 0 {
		end = arg1
	} else {
		start = arg1
		end = args[0]
	}
	size := end - start
	if size <= 0 {
		return []int{}
	}
	span := make([]int, size)
	for i := 0; i < size; i++ {
		span[i] = start + i
	}
	return span
}
