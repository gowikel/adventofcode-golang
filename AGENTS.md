# AGENTS

## Repo Purpose
This repository contains my Advent of Code solutions in Go, aiming to cover all years over time.

## Key Map
- `year2023/`, `year2024/`: Solutions by year/day.
- `internal/`: Shared utilities used across solutions.
- `internal/conf/`: Runtime configuration and CLI options.
- `internal/constants/`: Shared constants.
- `internal/runner/`: Execution orchestration for solutions.
- `internal/summary/`: Solution result summaries and reporting.
- `internal/utils/`: Common helpers.
- `inputs/`: Problem inputs.
- `main.go`: CLI entry point.

## Essential Workflows
- Build: `go build`
- Run: `go run .`
- Test (all): `go test ./...`
- Test (module): `go test ./year2023/day01` (example)

## No-Solution Policy
Advent of Code is about challenging the developer. The agent must guide only: explain concepts, clarify problem logic, and suggest possible algorithms or approaches. It must not provide complete solutions or write code for the challenges.

## Notes
- Use the Go CLI directly; the `Taskfile.yml` is not in use.
