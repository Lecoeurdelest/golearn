// Exercise 3: Slices (Pic)
// Read first: https://go.dev/tour/moretypes/18
// Guide: ../README.md
//
// Expected output (TODO 2, with the function (x+y)/2):
//   3 4
//   [[0 0 1 1] [0 1 1 2] [1 1 2 2]]
//
// With pic.Show (TODO 4), running locally prints a single line "IMAGE:<base64>" instead of a picture.
// Paste your code into the Playground (https://go.dev/tour/moretypes/18) to see the actual image.
//
// Run (from the repo root):  go run ./day-3/pic

package main

// TODO 0: import "fmt" and "golang.org/x/tour/pic"

// TODO 1: write func Pic(dx, dy int) [][]uint8
//         It returns a slice of length dy, where each element is a slice of dx values of type uint8.
//         Build the outer slice with make([][]uint8, dy), then make each row with make([]uint8, dx).
//         Fill in each pixel with the value of (x+y)/2 (convert it with uint8(...)).

// TODO 2: in main, call p := Pic(4, 3) and print len(p), len(p[0]) and then p itself.
//         The first number is the number of rows (dy), the second is the number of columns (dx).

// TODO 3: change the value function and compare: x*y, then x^y. Which patterns do they make?

// TODO 4: once Pic works, replace the prints in main with pic.Show(Pic).

func main() {
}
