// Bài 1: Loops and Functions (Newton's method)
// Đọc trước: https://go.dev/tour/flowcontrol/1  →  flowcontrol/8
// Hướng dẫn: ../README.md
//
// Kết quả mong đợi (số cuối có thể lệch vài chữ số thập phân, đó là bình thường):
//   Sqrt(2)  = 1.414213562373095
//   math.Sqrt = 1.4142135623730951
//   chênh lệch ~ 1e-16
//
// Chạy (từ thư mục gốc repo):  go run ./day-2/sqrt

package main

// TODO 0: import "fmt" và "math"
import (
	"fmt"
	"math"
)

// TODO 1: viết hàm Sqrt(x float64) float64 bằng phương pháp Newton.
//
//	Bắt đầu với z := 1.0, rồi lặp lại phép cập nhật:
//	    z -= (z*z - x) / (2 * z)
//	Bước đầu: lặp đúng 10 lần bằng vòng for đủ 3 phần (flowcontrol/1).
func Sqrt(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}
	z := 1.0
	diff := 1.0
	for diff > 1e-12 {
		next := z - (z*z-x)/(2*z)
		diff = math.Abs(next - z)
		z = next
	}
	return z
}

// TODO 2: trong main, in Sqrt(2), math.Sqrt(2) và hiệu của chúng.

// TODO 3 (nâng cấp): thay vòng lặp 10 lần bằng vòng "như while" (flowcontrol/3):
//         dừng khi z thay đổi rất nhỏ giữa hai lần lặp (ví dụ < 1e-12).
//         Đếm số lần lặp và in ra. Với x = 2 cần khoảng bao nhiêu lần?

// TODO 4 (thử phá): gọi Sqrt(-2). Chuyện gì xảy ra? Hàm có nên trả về lỗi không?
//         (Bài học về error sẽ đến ở Day 5, hiện tại chỉ cần quan sát.)

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(-2) - math.Sqrt(-2))
}
