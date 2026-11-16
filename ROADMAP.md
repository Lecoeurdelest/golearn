# Learn Go in 7 days

Main source: [A Tour of Go](https://go.dev/tour/welcome/1) (~100 pages, 7 modules).
Rule of thumb: **read the Tour to understand → retype the code in this repo to remember it**. Don't just run it on the Playground.

- Time: ~2–3 hours/day (days 5 and 7 are heavier, reserve 3–4 hours).
- One `day-N/` directory per day, one subdirectory with a `main.go` per exercise.
- Each day ends with one commit (`git commit -m "day N: ..."`).

## Setup (15 minutes, done on Day 1)

```sh
go version                      # go1.27.1 is already installed
cd golearn && go mod init golearn
go get golang.org/x/tour@latest # the Pic, WordCount, Readers and Binary Trees exercises import this package
mkdir -p day-1/hello && cd day-1/hello
# create main.go, then:
go run .
gofmt -l . ; go vet ./...       # run after every exercise
```

Target layout:

```
golearn/
├── go.mod
├── ROADMAP.md
├── day-1/ hello/  vars/  ...
├── day-2/ sqrt/   fizzbuzz/ ...
└── day-7/ linkcheck/
```

To run the Tour offline: `go install golang.org/x/website/tour@latest && $(go env GOPATH)/bin/tour`.

> **Note:** when run locally, `pic.Show` (Exercise Pic, Images) doesn't display an image; it prints a single `IMAGE:<base64>` line.
> To see the image, paste the code into the [Playground](https://go.dev/tour/moretypes/18); otherwise it's enough that the code runs without errors.

---

## Day 1 — Basics: packages, variables, types, functions

**Tour:** `welcome` (the whole module) + `basics/1 → 17` (~20 pages)

| Learn | Pages |
|---|---|
| Packages, imports, exported names (capitalized) | basics/1–3 |
| Functions, multiple results, named returns | basics/4–7 |
| `var`, `:=`, zero values, type conversion, type inference | basics/8–14 |
| Constants | basics/15–16 |

**Do in the repo:**
1. `day-1/hello` — print "Hello, 世界" with `fmt`.
2. `day-1/vars` — declare every basic type and print each type's zero value with `%v` and `%T`.
3. `day-1/swap` — a `swap(a, b string) (string, string)` function and one function that uses named returns.
4. `day-1/consts` — type conversion (the 3-4-5 hypotenuse) and the `1 << 100` constant.

Step-by-step guide and starter code: [day-1/README.md](day-1/README.md).

**Self-check (answer without looking back):**
- Why does `fmt.println` fail while `fmt.Println` doesn't?
- Where can `:=` be used, and where can't it?
- What are the zero values of `int`, `string`, `bool` and `*int`?

---

## Day 2 — Flow control

**Tour:** `flowcontrol/1 → 14` (14 pages)

| Learn | Pages |
|---|---|
| `for` (3 forms: full, `while`-style, infinite) | 1–4 |
| `if` with an init statement, `else` | 5–7 |
| **Exercise: Loops and Functions** (Newton's sqrt) | 8 |
| `switch` (no `break` needed, switch with no condition) | 9–11 |
| `defer`, LIFO order | 12–13 |

**Do in the repo:**
1. `day-2/sqrt` — the Tour exercise: implement `Sqrt(x)` yourself with Newton's method and compare it with `math.Sqrt`.
2. `day-2/fizzbuzz` — use a `switch` with no condition.
3. `day-2/defer` — print 5 lines with `defer` in a loop, and guess the order before running it.

**Self-check:**
- Go has no `while`. Rewrite a `while` loop using `for`.
- When does `defer` evaluate its **arguments**, and when does it run the function?

---

## Day 3 — Pointers, structs, arrays, slices

**Tour:** `moretypes/1 → 18` (18 pages)

| Learn | Pages |
|---|---|
| Pointers (no pointer arithmetic) | 1 |
| Structs, fields, pointer-to-struct, literals | 2–5 |
| Array vs slice, a slice is a *view* of an array | 6–8 |
| Slice literals, len/cap, nil slices, `make`, `append` | 9–15 |
| `range` | 16–17 |
| **Exercise: Slices** (`Pic`) | 18 |

**Do in the repo:**
1. `day-3/pointers` — functions that take `int` vs `*int`, and show the difference.
2. `day-3/slices` — create a slice, cut `s[1:3]`, modify an element and observe the original slice; print `len`/`cap` after each `append`.
3. `day-3/pic` — Exercise `Pic`.

**Self-check (this is where beginners trip up the most):**
- Two slices point to the same array: if you change one, does the other change?
- When does `append` allocate a new array?
- Why doesn't changing `v` in `for i, v := range s` change `s`?

---

## Day 4 — Maps, closures, and methods

**Tour:** `moretypes/19 → 27` + `methods/1 → 8` (~17 pages)

| Learn | Pages |
|---|---|
| Maps, literals, insert/delete/check existence (`v, ok := m[k]`) | moretypes/19–22 |
| **Exercise: Maps** (`WordCount`) | moretypes/23 |
| Functions as values, closures | moretypes/24–25 |
| **Exercise: Fibonacci closure** | moretypes/26 |
| Methods on structs, pointer receiver vs value receiver | methods/1–8 |

**Do in the repo:**
1. `day-4/wordcount` — Exercise `WordCount`.
2. `day-4/fib` — Exercise Fibonacci closure.
3. `day-4/shapes` — structs `Rect` and `Circle` with an `Area()` method; add one method that resizes the shape using a pointer receiver.

**Self-check:**
- What does reading a missing key from a map return? What happens when you write to a `nil` map?
- When is a pointer receiver required?

---

## Day 5 — Interfaces, errors, Readers (the heaviest day)

**Tour:** `methods/9 → 26` (18 pages, many exercises)

| Learn | Pages |
|---|---|
| Interfaces, implicit implementation, interface values, nil interfaces | 9–13 |
| `interface{}` / `any`, type assertions, type switches | 14–16 |
| `Stringer` + **Exercise** | 17–18 |
| `error` + **Exercise: Errors** | 19–20 |
| `io.Reader` + **Exercise: Readers**, **rot13Reader** | 21–23 |
| `image` + **Exercise: Images** (optional) | 24–25 |

**Do in the repo:**
1. `day-5/stringer` — `IPAddr` with `String()`.
2. `day-5/errors` — `Sqrt` returns `ErrNegativeSqrt`; add one function that uses `errors.Is` / `errors.As`.
3. `day-5/readers` — `MyReader` (a stream of `'A'` characters) and `rot13Reader`.
4. Skip Images if you run out of time — it doesn't affect the later days.

**Self-check:**
- Does an interface holding a nil pointer `== nil`? Why is this a trap?
- What does the Go idiom "accept interfaces, return structs" mean?

---

## Day 6 — Generics and basic concurrency

**Tour:** `generics/1 → 3` + `concurrency/1 → 8` (11 pages)

| Learn | Pages |
|---|---|
| Type parameters, generic types | generics/1–2 |
| Goroutines | concurrency/1 |
| Channels, buffered channels, `range` + `close` | 2–4 |
| `select`, default case | 5–6 |
| **Exercise: Equivalent Binary Trees** (goroutines + channels are enough) | 7–8 |

**Do in the repo:**
1. `day-6/generics` — `Map[T, U any]`, `Filter[T any]`, and `Sum` with the constraint `~int | ~float64`.
2. `day-6/sum` — split a slice in two, sum each half in its own goroutine, and send the results over a channel (exactly the Tour's example, retyped by hand).
3. `day-6/fanin` — 2 producers + `select` to receive from both, with a timeout (`time.After`).
4. `day-6/trees` — Exercise Binary Trees (`Walk`, `Same`).

**Self-check:**
- Who should `close` a channel — the sender or the receiver?
- What happens when you send on a channel that nobody receives from?
- Run with `go run -race .` to see what a data race is (create one on purpose).

---

## Day 7 — Advanced concurrency + a small project + review

**Tour:** `concurrency/9 → 11` (3 pages)

| Learn | Pages |
|---|---|
| `sync.Mutex` | 9 |
| **Exercise: Web Crawler** | 10 |
| Where to go from here | 11 |

**Do in the repo (morning — finish the Tour):**
1. `day-7/mutex` — a thread-safe counter with `sync.Mutex` (the Tour's `SafeCounter` example); try it with and without the mutex using `go run -race .`.
2. `day-7/crawler` — Exercise Web Crawler using `sync.Mutex` or channels. This is the hardest exercise in the Tour, so set aside dedicated time.

**Capstone project (afternoon, ~2 hours): `day-7/linkcheck`**
A CLI that takes a list of URLs (flag or file), checks them in parallel with `net/http`, and prints a table of `URL | status | time`.
It reuses: structs, slices, maps, errors, goroutines, channels, mutex/WaitGroup, `defer resp.Body.Close()`.
Add a `linkcheck_test.go` file using `httptest.NewServer` — this is your first time touching `go test`.

**End-of-week review:** reread the "Self-check" sections from all 7 days, mark the ones you can't answer yet, and go back to exactly those Tour pages.

---

## After this week (optional)

| Goal | Resources |
|---|---|
| Syntax and idioms | [Effective Go](https://go.dev/doc/effective_go), [Go by Example](https://gobyexample.com) |
| Testing | `go test`, table-driven tests, the `testing` package |
| Commonly used standard library | `net/http`, `encoding/json`, `context`, `flag`, `os`, `bufio` |
| Next project | A small REST API (todo) using `net/http` + `encoding/json` |

## Tips for keeping pace

- Retype the code, don't copy it from the Playground.
- Run `gofmt` and `go vet` on every exercise — build the habit early.
- Stuck for more than 20 minutes on an exercise: look at the official solution at [golang/website/tour/solutions](https://github.com/golang/website/tree/master/_content/tour/solutions), but rewrite it yourself once you understand it.
- Falling a day behind is fine: Days 4 and 5 are the ones that tend to pile up, so use the morning of Day 7 to catch up.
