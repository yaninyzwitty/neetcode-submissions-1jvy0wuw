func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	currRow := 0
	goingDown := false

	for _, ch := range s {
		rows[currRow] += string(ch)

		if currRow == 0 || currRow == numRows - 1 {
			goingDown = !goingDown
		}

		if goingDown {
			currRow++
		} else {
			currRow--
		}
	}

	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row)
	}	
	return result.String()
}
