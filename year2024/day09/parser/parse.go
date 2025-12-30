package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type MemoryMap struct {
	FileInfo  map[int]int // FileId -> Size
	MemoryMap []int       // FileId + Free Blocks tuples
	TotalSize int
}

func Parse(input io.Reader) (*MemoryMap, error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)

	result := &MemoryMap{
		FileInfo:  make(map[int]int),
		MemoryMap: make([]int, 0),
	}

	fileId := -1
	isFile := false

	for scanner.Scan() {
		value := scanner.Text()
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("parse: unable to convert to number: %w", err)
		}

		isFile = !isFile

		if isFile {
			fileId++
			result.FileInfo[fileId] = num
			result.MemoryMap = append(result.MemoryMap, fileId)
			result.TotalSize += num
		} else {
			result.MemoryMap = append(result.MemoryMap, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("parse: failed to scan: %w", err)
	}

	return result, nil
}
