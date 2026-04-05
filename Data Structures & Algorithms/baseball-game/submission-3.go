func calPoints(operations []string) int {
	var record []int
	total := 0
	
	for _, op := range operations {
		if n, err := strconv.Atoi(op); err == nil {
			record = append(record, n)
			total += n
			continue
		}

		switch op {
			case "+":
				u := record[len(record) - 1]
				pu := record[len(record) - 2]
				record = append(record, u + pu)
				total += record[len(record) - 1]

			case "D":
				record = append(record, record[len(record) - 1] * 2)
				total += record[len(record) - 1]

			case "C":
				total -= record[len(record) - 1]
				record = record[:len(record) - 1]
		}
	}

	return total
}
