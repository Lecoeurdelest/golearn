// Bài 2: Biến, zero value, kiểu cơ bản
// Đọc trước: https://go.dev/tour/basics/8  →  basics/12  và  basics/11
// Hướng dẫn: ../README.md
//
// Kết quả mong đợi (TODO 2):
//   int      0
//   float64  0
//   bool     false
//   string   ""
//   *int     <nil>
//
// Chạy (từ thư mục gốc repo):  go run ./day-1/vars

package main

// TODO 0: import "fmt"
import "fmt"

var c, python, java bool

// TODO 1: khai báo biến CẤP PACKAGE bằng var, ví dụ:  var c, python, java bool
//         (basics/8). Biến cấp package không dùng cũng không lỗi.

func main() {
	// TODO 2: khai báo 5 biến KHÔNG gán giá trị: int, float64, bool, string, *int
	//         rồi in kiểu và giá trị của từng biến bằng Printf với %T, %v (string dùng %q).
	//         Đây là zero value (basics/12).
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%-8T %v\n", a, a)
	fmt.Printf("%-8T %v\n", b, b)
	fmt.Printf("%-8T %v\n", c, c)
	fmt.Printf("%-8T %q\n", d, d) // string dùng %q
	fmt.Printf("%-8T %v\n", e, e)

	// TODO 3: khai báo i, j int = 1, 2 bằng var CÓ giá trị khởi tạo (basics/9).
	//         Rồi khai báo k := 3 bằng toán tử := (basics/10). In cả ba.

	// TODO 4: type inference (basics/14). Khai báo biến chỉ bằng giá trị, không ghi kiểu:
	//         a := 42        // kiểu gì?
	//         b := 3.14      // kiểu gì?
	//         c := 0.867 + 0.5i  // kiểu gì?
	//         In %T của từng biến. Đoán trước khi chạy.

	// TODO 5 (thử phá, làm lần lượt, mỗi lần hoàn tác):
	//         a) đặt `x := 1` NGOÀI hàm, ở cấp package → đọc lỗi
	//         b) khai báo một biến local trong main mà không dùng → đọc lỗi
	//         c) gán `i = "hello"` cho biến int → đọc lỗi
}
