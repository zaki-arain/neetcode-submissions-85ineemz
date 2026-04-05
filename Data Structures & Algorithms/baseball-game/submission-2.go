func calPoints(operations []string) int {
	var record []int
	
	for _, op := range operations {
		if n, err := strconv.Atoi(op); err == nil {
			record = append(record, n)
			continue
		}

		switch op {
			case "+":
				u := record[len(record) - 1]
				pu := record[len(record) - 2]
				record = append(record, u + pu)

			case "D":
				record = append(record, record[len(record) - 1] * 2)

			case "C":
				record = record[:len(record) - 1]
		}
	}
	
    total := 0
    for _, n := range record {
        total += n
    }
    return total
}
