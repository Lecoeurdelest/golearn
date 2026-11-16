# Day 1 — Basics: packages, variables, types, functions

**Time:** ~2–3 hours · **Tour:** [welcome](https://go.dev/tour/welcome/1) + [basics/1 → 17](https://go.dev/tour/basics/1)

How to work through each exercise: **read the Tour page → open `main.go` → work through each `TODO` → run it → compare with the expected output (headed `Kết quả mong đợi` in the file)**.
Don't copy from the Playground, type it by hand. If you're stuck for more than 20 minutes, open the *Hints* section (not the solution), and if you're still stuck, ask me.

---

## Step 0 — Initialize the module (once only)

Run this in the **repo root** (`golearn/`), not inside `day-1/`:

```sh
go version                       # check that Go is installed
go mod init golearn              # create go.mod
go get golang.org/x/tour@latest  # needed by the Tour exercises on later days
```

Then check that the skeleton runs (printing nothing is correct):

```sh
go run ./day-1/hello
```

<details>
<summary>Hint: why do we need go.mod?</summary>

Since Go 1.16, `go run` needs to know which module you're in. `go.mod` records the module name (`golearn`) and the libraries you use.
A single `go.mod` at the repo root is enough for all 7 days; each subdirectory (`day-1/hello`, …) is its own *package main*.

</details>

### Three commands to run after every exercise

```sh
go run ./day-1/<exercise-name>   # run
gofmt -l day-1                   # list files that aren't formatted correctly (printing nothing is good). Fix: gofmt -w day-1
go vet ./day-1/...               # catch common mistakes
```

---

## Exercise 1 — `hello` · print "Hello, 世界"

**Read first:** [basics/1](https://go.dev/tour/basics/1) (Packages) · [basics/2](https://go.dev/tour/basics/2) (Imports) · [basics/3](https://go.dev/tour/basics/3) (Exported names)
**File:** [hello/main.go](hello/main.go)

**Expected output:**

```
Hello, 世界
```

**Do:** the 3 TODOs in the file. TODO 3 is to deliberately break the code and then read the error message. This is the fastest way to learn.

<details>
<summary>Hints</summary>

- A runnable program needs `package main` and a `func main()`.
- `import "fmt"` or the grouped form `import ( "fmt" )`.
- The print function lives in the `fmt` package, and its name starts with a capital letter.
- The TODO 3 error will mention "unexported name" or "undefined". Answer for yourself: what does the capitalization of the first letter of a name decide?

</details>

---

## Exercise 2 — `vars` · variables, zero values, basic types

**Read first:** [basics/8](https://go.dev/tour/basics/8) → [basics/12](https://go.dev/tour/basics/12) (Variables → Zero values) · [basics/11](https://go.dev/tour/basics/11) (Basic types)
**File:** [vars/main.go](vars/main.go)

**Expected output** (the `TODO 2` part):

```
int      0
float64  0
bool     false
string   ""
*int     <nil>
```

**Do:** you'll use 3 ways of declaring variables: `var` at package level, `var` inside a function, and `:=`. The goal is to *see* the zero values for yourself, not to memorize them.

**`fmt.Printf` verb table** (used throughout the course):

| Verb | Prints |
|---|---|
| `%v` | the value in its default format |
| `%T` | the type of the value |
| `%q` | a string in double quotes |
| `\n` | a newline (`Printf` doesn't add one, `Println` does) |

<details>
<summary>Hints</summary>

- Declaring without assigning: `var i int` → `i` automatically gets its zero value.
- `fmt.Printf("%-8T %v\n", i, i)` left-aligns the type column. Use `%q` for strings so you can see the empty string.
- `:=` can only be used **inside a function**. Try putting it outside a function to read the error.
- A local variable that's declared but not used is a compile error, while a package-level variable isn't. Try both.

</details>

---

## Exercise 3 — `swap` · functions, multiple return values, named returns

**Read first:** [basics/4](https://go.dev/tour/basics/4) → [basics/7](https://go.dev/tour/basics/7)
**File:** [swap/main.go](swap/main.go)

**Expected output:**

```
world hello
7 10
```

**Do:** write `swap` (returns 2 values) and `split` (named returns + a *naked return*), then call both from `main`.

<details>
<summary>Hints</summary>

- Signature: `func swap(x, y string) (string, string)`: the return type goes **after** the parameter list, and multiple results go in parentheses.
- Receiving 2 values: `a, b := swap("hello", "world")`.
- Named returns: `func split(sum int) (x, y int)`. `x` and `y` are already declared, so assign to them directly and write a bare `return`.
- Formula: `x = sum * 4 / 9`, `y = sum - x`. With `17` this gives `7 10`. Watch out for integer division.
- Ask yourself: when does a naked return make code harder to read?

</details>

---

## Exercise 4 — `consts` · type conversion, type inference, constants

**Read first:** [basics/13](https://go.dev/tour/basics/13) → [basics/16](https://go.dev/tour/basics/16)
**File:** [consts/main.go](consts/main.go)

**Expected output:**

```
3 4 5
21
0.2
1.2676506002282295e+29
```

**Do:** the hypotenuse of a triangle (you need to convert between `int` ↔ `float64` ↔ `uint`, and Go does **not** convert automatically) and the huge constant `1 << 100`.

<details>
<summary>Hints</summary>

- A type conversion is written like a function call: `float64(x)`, `uint(f)`.
- `math.Sqrt` takes and returns a `float64`, so you have to convert `x*x + y*y` to `float64` first.
- An untyped constant doesn't get a fixed type until it's used, so `Big` (1 << 100) can be passed to a function that takes a `float64`, but passing it where an `int` is expected gives an overflow error. TODO 4 lets you see that error yourself.
- `const` can't be used with `:=`. You have to write `const Big = 1 << 100`.

</details>

---

## End of the day

Tick these off when you're done:

- [ ] `hello`, `vars`, `swap` and `consts` all run and match the expected output
- [ ] `gofmt -l day-1` prints nothing and `go vet ./day-1/...` reports no errors
- [ ] Read the whole Tour `basics` module through to the *Congratulations!* page

**Self-check** (answer without looking back):

1. Why does `fmt.println` fail while `fmt.Println` doesn't?
2. Where can `:=` be used, and where can't it?
3. What are the zero values of `int`, `string`, `bool` and `*int`?
4. With `x := 3` (type `int`), why is `var f float64 = x` a compile error while `var f float64 = 3` is fine?
5. How do `var x, y int` and `x, y := 1, 2` differ?

Then commit:

```sh
git add -A && git commit -m "day 1: basics"
```

After Day 1, go back to [ROADMAP.md](../ROADMAP.md) and open Day 2.
