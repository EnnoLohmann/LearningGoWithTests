package iteration

func Repeat(letter string, repeats int) (repeated string) {
	for i := 0; i < repeats; i++ {
		repeated += letter
	}
	return
}
