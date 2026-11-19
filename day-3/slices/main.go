// Exercise 2: slices, len/cap, append
// Read first: https://go.dev/tour/moretypes/6  →  moretypes/17
// Guide: ../README.md
//
// Expected output (the cap values after append can differ between Go versions, that's fine):
//   s = [10 20 30 40 50] len=5 cap=5
//   t = [20 30] len=2 cap=4
//   s = [10 99 30 40 50]
//   t = [99 30]
//   len=1 cap=4
//   len=2 cap=4
//   len=3 cap=4
//   len=4 cap=4
//   len=5 cap=8
//   len=6 cap=8
//   len=7 cap=8
//   len=8 cap=8
//   len=9 cap=16
//   len=10 cap=16
//   [0 0 0 2] [0 0 0 2]
//   cap 3: [0 0 0 1] [0 0 0 2]
//   after range copy: [10 99 30 40 50]
//   after index loop: [20 198 60 80 100]
//
// Run (from the repo root):  go run ./day-3/slices

package main

// TODO 0: import "fmt"

func main() {
	// TODO 1: create s := []int{10, 20, 30, 40, 50} and print it with its length and capacity:
	//             fmt.Printf("s = %v len=%d cap=%d\n", s, len(s), cap(s))

	// TODO 2: cut t := s[1:3] and print it the same way (use the label "t").
	//         GUESS first: what are len(t) and cap(t)?
	//         Then set t[0] = 99 and print s and t with fmt.Println("s =", s) and fmt.Println("t =", t).
	//         Did the original slice s change? Why?

	// TODO 3: start from a nil slice (var a []int). In a loop of 10 iterations, append i to it
	//         and print len and cap after every append:
	//             fmt.Printf("len=%d cap=%d\n", len(a), cap(a))
	//         Watch where cap jumps: that is exactly when append allocates a new array.

	// TODO 4: base := make([]int, 3, 4)
	//         x := append(base, 1)
	//         y := append(base, 2)
	//         Print x and y. GUESS first: are they [0 0 0 1] [0 0 0 2]?
	//         Then change the capacity from 4 to 3 and print again with the label "cap 3:".

	// TODO 5: loop over s twice, doubling every element:
	//         (a) for _, v := range s { v *= 2 }   then print "after range copy:" and s
	//         (b) for i := range s { s[i] *= 2 }   then print "after index loop:" and s
	//         Why does only one of them change s?
}
