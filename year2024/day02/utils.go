package day02

import "strconv"

func convertToIntList(data []string) ([]int, error) {
	result := make([]int, 0, len(data))

	for _, s := range data {
		i, err := strconv.Atoi(s)
		if err != nil {
			return result, err
		}

		result = append(result, i)
	}

	return result, nil
}
