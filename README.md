[![Go Reference](https://pkg.go.dev/badge/go.eryndalor.dev/adventofcode-golang.svg)](https://pkg.go.dev/go.eryndalor.dev/adventofcode-golang) [![Go Report Card](https://goreportcard.com/badge/go.eryndalor.dev/adventofcode-golang)](https://goreportcard.com/report/go.eryndalor.dev/adventofcode-golang)

# Advent of Code

This repo contains my solutions to the Advent of Code for 2023 in Golang.

## Setup

Just install go and compile the binary with the normal `go build` or `go run .`.

The binary accepts the following actions:

### Solve

Intended to execute a given solution. It has two arguments, year and day, to locate
the intended solution, plus one flag, --example, to pass the example input.

```bash
aoc solve 2023 5 --example
```

## Taskfile

To make some task easier, there is a `Taskfile.yml`. Simply install [this](https://taskfile.dev/)
and run the desired task:

**Test all the files**

```bash
task test-all -- # Other params to go test
```

**Test only one module**

```bash
task test -- go.eryndalor.dev/adventofcode-golang/year2023/day01

# This will also work
task test -- ./year2023/day01
```

**Test only one function**
```bash
task test -- -run TestIsIncluded_WithStep_InsideRangeTarget_Valid
```

**Run the binary**

```bash
task run -- solve 2023 1 --example
```

**Format code**

```bash
task fmt
```

## Dependencies

- [Taskfile](https://taskfile.dev/) as task manager
- [pkgsite](https://github.com/golang/pkgsite) to generate the docs

## Authors

- Pedro Jose Piquero Plaza - pedropiqueroplaza@proton.me

## License

This project is licensed under the [MIT License](LICENSE.md) - see the file for details

The puzzles and its inputs are property of [AoC](https://adventofcode.com/)