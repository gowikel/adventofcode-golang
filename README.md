# Advent of Code

This repo contains my solutions to the Advent of Code for 2023 in Golang.

## Setup

Just install go and compile the binary with the normal `go build` or `go run .`.

The binary accepts three flags:

- **-year**: to specify the year to run. Defaults to the current year
- **-day**: to specify the day to run. Defaults to today if we are on December, 1 otherwise.
- **-run-example**: to pass the example input to the solver instead of the puzzle data

That's it! You don't need to pass the inputs, as they are embedded in the binary.

## Taskfile

To make some task easier, there is a `Taskfile.yml`. Simply install [this](https://taskfile.dev/)
and run the desired task:

**Test all the files**

```bash
task test-all -- # Other params to go test
```

**Test only one module**

```bash
task test PATH=year2023/day01 -- # Other params to go test
```

**Run the binary**

```bash
task run -- --year 2023 --day 1 --run-example
```

**Format code**

```bash
task fmt
```

## Dependencies

- [Taskfile](https://taskfile.dev/) as task manager
- [pkgsite](https://github.com/golang/pkgsite) to generate the docs

## Authors

- Pedro Jose Piquero Plaza - gowikel@gmail.com

## License

This project is licensed under the [MIT License](LICENSE.md) - see the file for details

The puzzles and its inputs are property of [AoC](https://adventofcode.com/)