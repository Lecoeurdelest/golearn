# Day 2 — Flow control

**Time:** ~2–3 hours · **Tour:** [flowcontrol/1 → 14](https://go.dev/tour/flowcontrol/1)

Same routine as Day 1: **read the Tour page, open `main.go`, work through each `TODO`, run it, compare with the expected output (headed `Kết quả mong đợi` in the file)**.
Type it by hand, don't copy from the Playground. If you're stuck for more than 20 minutes, open the *Hints* section (not the solution), and if you're still stuck, ask me.

---

## Step 0 — Check the module

Day 2 shares the `go.mod` in the **repo root** (`golearn/`) with Day 1, so don't create a new one.
Check by running from the repo root:

```sh
go run ./day-2/sqrt
```

It's correct if it prints nothing and shows no errors. If it reports `go.mod file not found`, you need to move `go.mod` and `go.sum` from `day-1/` to the repo root:

```sh
git mv day-1/go.mod day-1/go.sum .
```

### Three commands to run after every exercise

```sh
go run ./day-2/<exercise-name>   # run
gofmt -l day-2                   # list files that aren't formatted correctly (printing nothing is good). Fix: gofmt -w day-2
go vet ./day-2/...               # catch common mistakes
```

---

## Exercise 1 — `sqrt` · loops and functions (Newton)

**Read first:** [flowcontrol/1](https://go.dev/tour/flowcontrol/1) → [flowcontrol/8](https://go.dev/tour/flowcontrol/8)
**File:** [sqrt/main.go](sqrt/main.go)

**Expected output:** `Sqrt(2)` is close to `math.Sqrt(2)` (about `1.41421356237`), with a difference on the order of `1e-16`.

**Do:** implement the square root yourself using Newton's method, then compare it with `math.Sqrt`. After that, upgrade it from "loop exactly 10 times" to "loop until it's accurate enough".

<details>
<summary>Hints</summary>

- The full 3-part `for` loop: `for i := 0; i < 10; i++ { ... }`.
- The "while-style" loop: drop the init and post statements and keep only the condition: `for condition { ... }`. Go has no `while` keyword.
- Newton update: `z -= (z*z - x) / (2 * z)`. With `x = 2`, it converges in just a few iterations.
- To check for a "small change", save the old value in a variable, subtract the new value from it, then take the absolute value with `math.Abs`.
- The printed value may differ from `math.Sqrt` in the last digits. Comparing floating-point numbers with `==` is a bad habit, so compare the difference against a small threshold instead.

</details>

---

## Exercise 2 — `fizzbuzz` · switch with no condition

**Read first:** [flowcontrol/9](https://go.dev/tour/flowcontrol/9) → [flowcontrol/11](https://go.dev/tour/flowcontrol/11) (and [flowcontrol/5](https://go.dev/tour/flowcontrol/5) → [7](https://go.dev/tour/flowcontrol/7) for `if`)
**File:** [fizzbuzz/main.go](fizzbuzz/main.go)

**Expected output** (1 to 15):

```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
```

**Do:** print 1 to 15 following the FizzBuzz rules, using a `switch` with no condition. TODO 2 is to swap the order of the `case`s to see that `switch` stops at the first matching branch.

<details>
<summary>Hints</summary>

- A `switch` with no condition is a tidy replacement for an `if / else if` chain: `switch { case i%15 == 0: ... }`.
- Go does **not** need a `break` at the end of each `case`; it stops automatically after the branch that ran.
- You must check divisibility by 15 **before** 3 and 5, otherwise `15` will fall into the `Fizz` branch.
- An `if` can have an init statement: `if r := i % 3; r == 0 { ... }`. The variable `r` only exists inside that if/else block.

</details>

---

## Exercise 3 — `defer` · LIFO execution order

**Read first:** [flowcontrol/12](https://go.dev/tour/flowcontrol/12) → [flowcontrol/13](https://go.dev/tour/flowcontrol/13)
**File:** [defer/main.go](defer/main.go)

**Expected output** (TODO 1 and 2):

```
start
end loop
4
3
2
1
0
```

**Do:** use `defer` in a loop, **guess** the output order before running it. Then try the code in TODO 3 to understand when `defer` evaluates its arguments.

<details>
<summary>Hints</summary>

- A `defer` statement postpones the **function call** until the surrounding function (here, `main`) returns.
- Multiple `defer`s are pushed onto a stack: the last one in runs first (LIFO).
- The **arguments** of a deferred function are evaluated immediately when the `defer` is reached, not when it runs. That's why TODO 3 prints `deferred x = 1` even though `x` is later changed to `100`.
- In practice `defer` is used most for cleanup: closing files, releasing locks (mutexes), closing connections.

</details>

---

## End of the day

Tick these off when you're done:

- [ ] `sqrt`, `fizzbuzz` and `defer` all run and match the expected output
- [ ] `gofmt -l day-2` prints nothing and `go vet ./day-2/...` reports no errors
- [ ] Read the whole Tour `flowcontrol` module through to the *Congratulations!* page

**Self-check** (answer without looking back):

1. Go has no `while`. Rewrite a `while` loop using `for`.
2. What are the 3 forms of the `for` loop? Which one is the infinite loop?
3. How long does a variable declared in an `if` init statement live?
4. Why doesn't `switch` in Go need `break`? When would you want to continue into the next case (the `fallthrough` keyword)?
5. When does `defer` evaluate its **arguments**, and when does it run the **function**?
6. In what order do three consecutive `defer` statements run?

Then commit:

```sh
git add -A && git commit -m "day 2: control flow"
```

After Day 2, go back to [ROADMAP.md](../ROADMAP.md) and open Day 3.
