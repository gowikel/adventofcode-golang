package day01

type Exercise struct{}

func (e Exercise) Part1(data string) int {
	numbers := ParseInput(data)
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func (e Exercise) Part2(data string) int {
	numbers := EnhancedParseInput(data)
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
