package iterations

func Repeat(letter string) (word string) {
	for i := 0; i < 5; i++ {
		word += letter
	}

	return
}
