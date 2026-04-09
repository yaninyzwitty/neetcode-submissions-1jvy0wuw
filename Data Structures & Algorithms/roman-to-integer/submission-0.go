func romanToInt(s string) int {
	 values := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

	prev, total := 0, 0
	for i := len(s) -1; i >= 0; i-- {
		currentVal := values[s[i]]

		if currentVal < prev {
			total -= currentVal
		} else {
			total += currentVal
		}
		prev = currentVal
	}
	return total

	
}
