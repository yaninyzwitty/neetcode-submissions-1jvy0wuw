import "slices"
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		chars := []rune(str)
		slices.Sort(chars) // k log k, final complexity , 0(n * k log k ) from k log k 
		key := string(chars)
		m[key] = append(m[key], str)
	}

	var result [][]string
	for _, group := range m {
		result = append(result, group)
	}
	return result

}
