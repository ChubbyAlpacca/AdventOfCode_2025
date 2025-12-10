# Advent of Code 2025 — Workspace

This repository contains my solutions and working files for Advent of Code 2025. Each puzzle day is placed in its own folder (`day1`, `day2`, ...). The solutions in this workspace are written in Go.

## Project layout

Top-level structure:

- `dayN/` — a folder per puzzle day (e.g. `day1/`, `day2/`). Each day folder typically contains:
  - `main.go` — the solution/runner for that day.
  - `go.mod` — module file indicating the Go version used.
  - `problem_statement.txt` — the puzzle text (for reference).
  - `data/` — input data files:
    - `test_input.txt` — the small example input shown in the puzzle statement.
    - `full_input.txt` — your personal puzzle input.

Example: `day1/` contains `main.go`, a `data/` directory with `test_input.txt` and `full_input.txt`, and `problem_statement.txt`.

## How each day is organised

- Solutions are runnable programs. They read input from a path defined near the top of `main.go` (often a constant like `inputFilePath`). By default many days point at `data/test_input.txt` for easy testing; switch it to `data/full_input.txt` to run your real puzzle input.
- Some days implement helper functions and small runners inside `main.go`. If you add more files for a day, keep them in the same `dayN` directory and update the `package`/`module` as needed.

## Running a day

Requirements:

- Go installed locally (use the Go version indicated in the day's `go.mod`, e.g. `go 1.24`).

Run from the day folder. Example for `day1`:

```bash
cd day1
go run .
```

or explicitly:

```bash
go run main.go
```

If you prefer to run the full input, open `main.go` and point the input constant to `data/full_input.txt`, or modify the code to accept a command-line flag.

## Suggested improvements (quick wins)

- Add a command-line flag for the input file (e.g. `-input data/full_input.txt`) so you don't edit the source to switch inputs.
- Return results from `solve`/helper functions rather than printing directly; this makes testing easier.
- Add small unit tests per day using Go's testing framework (`*_test.go`). Example tests should verify the sample input yields the expected result.
- For very large numeric ranges (some puzzle variants), avoid iterating every integer — enumerate only matching patterns and compute counts/sums mathematically.

## Troubleshooting

- `zsh: command not found: go` — install Go (https://golang.org/dl/), or make sure your `PATH` points to the Go binary.
- Build errors after editing — run `go build` in the day folder to see compile issues.
- If a file references `data/full_input.txt` but you don't have that file, create it or switch to `data/test_input.txt`.
