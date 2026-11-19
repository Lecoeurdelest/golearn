# Day 3 — Pointers, structs, arrays, slices

**Time:** ~2–3 hours · **Tour:** [moretypes/1 → 18](https://go.dev/tour/moretypes/1)

Same routine as Day 1 and Day 2: **read the Tour page, open `main.go`, work through each `TODO`, run it, compare with the expected output (headed `Expected output` in the file)**.
Type it by hand, don't copy from the Playground. If you're stuck for more than 20 minutes, open the *Hints* section (not the solution), and if you're still stuck, ask me.

This is the day beginners trip up the most, mostly on slices. Take your time with Exercise 2 and **guess before you run**.

---

## Step 0 — Check the module

Day 3 shares the `go.mod` in the **repo root** (`golearn/`) with the earlier days, so don't create a new one.
Check by running from the repo root:

```sh
go run ./day-3/pointers
```

It's correct if it prints nothing and shows no errors.

Exercise 3 imports `golang.org/x/tour/pic`, which you added to `go.mod` on Day 1. It's currently marked `// indirect`; once you import it, you can run `go mod tidy` to clean that up.

### Three commands to run after every exercise

```sh
go run ./day-3/<exercise-name>   # run
gofmt -l day-3                   # list files that aren't formatted correctly (printing nothing is good). Fix: gofmt -w day-3
go vet ./day-3/...               # catch common mistakes
```

---

## Exercise 1 — `pointers` · pointers and pointers to structs

**Read first:** [moretypes/1](https://go.dev/tour/moretypes/1) (Pointers) → [moretypes/5](https://go.dev/tour/moretypes/5) (Struct Literals)
**File:** [pointers/main.go](pointers/main.go)

**Expected output:**

```
after incByValue:   x = 1
after incByPointer: x = 2
*p = 2
after *p = 42: x = 42
p == &x: true
np = <nil>
{10 2}
```

**Do:** write one function that takes an `int` and one that takes a `*int`, and see for yourself which one changes the caller's variable. Then work with `&`, `*`, a nil pointer, and a pointer to a struct.

**Pointer cheat sheet:**

| Syntax | Meaning |
|---|---|
| `&x` | the address of `x` (a pointer to `x`) |
| `*p` | the value `p` points to (dereference) |
| `*int` | the type "pointer to an `int`" |
| `var p *int` | a nil pointer: its zero value is `nil` |

<details>
<summary>Hints</summary>

- Go passes everything **by value**. `incByValue(n int)` gets a copy of `x`, so `n++` only changes the copy.
- `incByPointer(n *int)` gets a copy of the *address*, but both addresses point to the same `x`, so `*n++` changes the caller's variable.
- Go has no pointer arithmetic, so you can't do `p + 1`.
- Dereferencing a nil pointer panics: `panic: runtime error: invalid memory address or nil pointer dereference`. Reading that message once is worth it.
- For a pointer to a struct, `pv.X` is shorthand for `(*pv).X`. Go does the dereference for you.

</details>

---

## Exercise 2 — `slices` · slices, len/cap, append

**Read first:** [moretypes/6](https://go.dev/tour/moretypes/6) → [moretypes/17](https://go.dev/tour/moretypes/17) (Arrays → Range continued)
**File:** [slices/main.go](slices/main.go)

**Expected output:** (the `cap` values after `append` can differ between Go versions)

```
s = [10 20 30 40 50] len=5 cap=5
t = [20 30] len=2 cap=4
s = [10 99 30 40 50]
t = [99 30]
len=1 cap=4
len=2 cap=4
len=3 cap=4
len=4 cap=4
len=5 cap=8
len=6 cap=8
len=7 cap=8
len=8 cap=8
len=9 cap=16
len=10 cap=16
[0 0 0 2] [0 0 0 2]
cap 3: [0 0 0 1] [0 0 0 2]
after range copy: [10 99 30 40 50]
after index loop: [20 198 60 80 100]
```

**Do:** build a slice, cut a sub-slice from it and modify it, watch `len` and `cap` change on every `append`, and find out what happens when two slices share one array. The last TODO is the classic `range` trap.

**Array vs slice:**

| | Array | Slice |
|---|---|---|
| Type | `[5]int` (the length is part of the type) | `[]int` |
| Size | fixed | grows with `append` |
| Assignment | copies all the elements | copies a small header (pointer, len, cap), not the data |

<details>
<summary>Hints</summary>

- `s[1:3]` includes index 1 and **excludes** index 3, so `len(t)` is 2. `cap(t)` counts from the start index to the end of the underlying array: `5 − 1 = 4`.
- `t` and `s` are two views of the **same** array, so `t[0] = 99` shows up as `s[1]`.
- The exact `cap` numbers depend on the Go version (this output is from Go 1.27.1). What matters is that `cap` jumps (4 → 8 → 16): that is the moment `append` allocated a bigger array and copied the elements over.
- `append(base, 1)` and `append(base, 2)` with `cap(base) == 4` both fit in the existing array, so they write into the **same** slot, and both results end in `2`. With `cap == 3` there is no room, so each `append` allocates its own array.
- In `for _, v := range s`, `v` is a **copy** of each element. To change `s`, write through the index: `s[i] *= 2`.
- `append` returns the new slice, so always assign it back: `a = append(a, i)`.

</details>

---

## Exercise 3 — `pic` · Exercise: Slices

**Read first:** [moretypes/18](https://go.dev/tour/moretypes/18)
**File:** [pic/main.go](pic/main.go)

**Expected output** (TODO 2, with the function `(x+y)/2`):

```
3 4
[[0 0 1 1] [0 1 1 2] [1 1 2 2]]
```

**Do:** implement `Pic(dx, dy int) [][]uint8`, which builds a slice of slices, then try a few different formulas. Once it works, hand it to `pic.Show`.

> **Note:** when run locally, `pic.Show` doesn't display an image; it prints a single `IMAGE:<base64>` line.
> To see the image, paste your code into the [Playground](https://go.dev/tour/moretypes/18); otherwise it's enough that the code runs without errors.

<details>
<summary>Hints</summary>

- Two levels of `make`: first `p := make([][]uint8, dy)`, then inside a loop `p[y] = make([]uint8, dx)`.
- `p[y][x]`: the **row** is `y` and the **column** is `x`. That's why `Pic(4, 3)` has `len(p) == 3` and `len(p[0]) == 4`.
- Convert the value with `uint8(...)`. For `x*y` and `x^y` the result can exceed 255 and wrap around, which is fine for a picture.
- Compare the formulas: `(x+y)/2` gives a smooth diagonal gradient, `x*y` gives curved bands, and `x^y` gives a fractal-looking pattern.

</details>

---

## End of the day

Tick these off when you're done:

- [ ] `pointers`, `slices` and `pic` all run and match the expected output
- [ ] `gofmt -l day-3` prints nothing and `go vet ./day-3/...` reports no errors
- [ ] Read the whole Tour `moretypes` module up to page 18 (the next pages are for Day 4)

**Self-check** (answer without looking back; this is where beginners trip up the most):

1. Two slices point to the same array: if you change one, does the other change?
2. When does `append` allocate a new array?
3. Why doesn't changing `v` in `for i, v := range s` change `s`?
4. What does `s[1:3]` give you, and what are its `len` and `cap`?
5. What's the zero value of a pointer, and what happens if you dereference it?
6. How does `pv.X` work when `pv` is a `*Vertex`?

Then commit:

```sh
git add -A && git commit -m "day 3: pointers and slices"
```

After Day 3, go back to [ROADMAP.md](../ROADMAP.md) and open Day 4.
