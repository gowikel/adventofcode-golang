package day09

import (
	"github.com/gowikel/adventofcode-golang/internal/runner"
)

type Exercise struct{}

func (e Exercise) Part1(input string) (int, error) {
	data, err := parseFilePart1(input)
	if err != nil {
		return 0, err
	}

	memoryMap := data.MemoryMap
	fileInfo := data.FileInfo
	movedFiles := make([]int, 0, data.TotalSize)
	var result int

	leftIndex := 0
	rightIndex := len(memoryMap) - 1
	if rightIndex%2 != 0 {
		rightIndex--
	}

	for range fileInfo[memoryMap[leftIndex]] {
		movedFiles = append(movedFiles, memoryMap[leftIndex])
	}
	leftIndex++

	remainingRightSize := fileInfo[memoryMap[rightIndex]]
	spaceRemaining := memoryMap[leftIndex]

	for leftIndex < rightIndex {
		isLeftAFile := leftIndex%2 == 0
		isRightAFile := rightIndex%2 == 0

		if isLeftAFile {
			fileId := memoryMap[leftIndex]
			fileSize := fileInfo[fileId]

			for range fileSize {
				movedFiles = append(movedFiles, fileId)
			}
			leftIndex++
			spaceRemaining += memoryMap[leftIndex]
			continue
		}

		if isRightAFile {
			fileId := memoryMap[rightIndex]
			spaceToMove := remainingRightSize

			if spaceRemaining < spaceToMove {
				spaceToMove = spaceRemaining
			}

			for range spaceToMove {
				movedFiles = append(movedFiles, fileId)
			}

			remainingRightSize -= spaceToMove
			spaceRemaining -= spaceToMove

			if spaceRemaining == 0 {
				leftIndex++
			}
			if remainingRightSize == 0 {
				rightIndex -= 2
				remainingRightSize = fileInfo[memoryMap[rightIndex]]
			}
			continue
		}
	}

	for range remainingRightSize {
		movedFiles = append(movedFiles, memoryMap[rightIndex])
	}

	for idx, value := range movedFiles {
		result += idx * value
	}

	return result, nil
}

func (e Exercise) Part2(input string) (int, error) {
	return 0, runner.ErrPartNotImplemented
}
