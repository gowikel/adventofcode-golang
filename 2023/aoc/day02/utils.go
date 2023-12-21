package day02

import "math"

func PowerCube(cube SetCubes) int {
	return cube.Red * cube.Green * cube.Blue
}

func MinimumPowerSet(cubes []SetCubes) SetCubes {
	result := SetCubes{
		Red:   math.MinInt,
		Green: math.MinInt,
		Blue:  math.MinInt,
	}

	for _, cube := range cubes {
		if cube.Red > result.Red {
			result.Red = cube.Red
		}

		if cube.Green > result.Green {
			result.Green = cube.Green
		}

		if cube.Blue > result.Blue {
			result.Blue = cube.Blue
		}
	}

	return result
}