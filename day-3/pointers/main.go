// Exercise 1: pointers (and pointers to structs)
// Read first: https://go.dev/tour/moretypes/1  →  moretypes/5
// Guide: ../README.md
//
// Expected output:
//   after incByValue:   x = 1
//   after incByPointer: x = 2
//   *p = 2
//   after *p = 42: x = 42
//   p == &x: true
//   np = <nil>
//   {10 2}
//
// Run (from the repo root):  go run ./day-3/pointers

package main

// TODO 0: import "fmt"

// TODO 1: write two functions:
//             func incByValue(n int)     // does n++
//             func incByPointer(n *int)  // does *n++
//         In main, declare x := 1, call each function in turn and print x after each call.
//         GUESS before running: which call changes x, and why?

// TODO 2: declare p := &x and print *p.
//         Then assign *p = 42, print x, and print the comparison p == &x.

// TODO 3: declare a pointer without a value: var np *int, and print np.
//         Then (as an experiment) add fmt.Println(*np), run it and read the panic message.
//         Comment that line out again before moving on.

// TODO 4: declare type Vertex struct{ X, Y int }.
//         Create v := Vertex{1, 2} and pv := &v, then set pv.X = 10 and print v.
//         Notice you wrote pv.X, not (*pv).X.

func main() {
}
