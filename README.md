# golearn

Learn Go in 7 days, following [A Tour of Go](https://go.dev/tour/welcome/1).

Rule of thumb: read the Tour to understand, then retype the code in this repo to remember it. Don't just run it on the Playground.

## Roadmap

| Day | Topic |
|---|---|
| Day 1 | Basics: packages, variables, types, functions |
| Day 2 | Flow control |
| Day 3 | Pointers, structs, arrays, slices |
| Day 4 | Maps, closures, methods |
| Day 5 | Interfaces, errors, Readers |
| Day 6 | Generics and basic concurrency |
| Day 7 | Advanced concurrency, a small project, review |

See [ROADMAP.md](ROADMAP.md) for the details of each day.

## Structure

```
golearn/
  go.mod
  go.sum
  README.md
  ROADMAP.md
  day-1/
    README.md      guide and checklist for the day
    hello/main.go
    vars/main.go
    swap/main.go
    consts/main.go
  day-2/ ...
```

Each day has its own `day-N/` directory. Each exercise is a subdirectory with its own `main.go`, which is a standalone program.

## Requirements

- Go 1.27 or later (check with `go version`)

## Running

Run from the repo root:

```sh
go run ./day-1/hello
go run ./day-1/vars
```

Note the leading `./`. If you write `go run day-1/hello`, Go treats it as an import path, not a directory.

After each exercise, check formatting and common mistakes:

```sh
gofmt -l .
go vet ./...
```

## How to work through each exercise

1. Read the matching Tour page.
2. Open the exercise's `main.go` and work through each `TODO` in order.
3. Run it and compare against the expected-output block at the top of the file (headed `Kết quả mong đợi`).
4. At the end of the day, commit: `git commit -m "day N: ..."`.
