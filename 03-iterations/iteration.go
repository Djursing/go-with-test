package iterations

func Repeat(letter string, amount int) (word string) {
	for i := 0; i < amount; i++ {
		word += letter
	}

	return
}
