// Bài 4: Ép kiểu, type inference, constants
// Đọc trước: https://go.dev/tour/basics/13  →  basics/16
// Hướng dẫn: ../README.md
//
// Kết quả mong đợi:
//   3 4 5
//   21
//   0.2
//   1.2676506002282295e+29
//
// Chạy (từ thư mục gốc repo):  go run ./day-1/consts

package main

// TODO 0: import "fmt" và "math"

// TODO 1: khai báo hằng số khổng lồ (basics/16), cả hai là constant chứ không phải var:
//           Big   = 1 << 100
//           Small = Big >> 99
//         Gợi ý: dùng khối const ( ... ) cho gọn.

// TODO 2: viết 2 hàm:
//           needInt(x int) int          { return x*10 + 1 }
//           needFloat(x float64) float64 { return x * 0.1 }

func main() {
	// TODO 3: ép kiểu (basics/13). Cho x, y := 3, 4 (kiểu int).
	//         Tính cạnh huyền bằng math.Sqrt (nhận float64), rồi ép kết quả sang uint:
	//           f := math.Sqrt( ... )
	//           z := uint(f)
	//         In: x y z   → 3 4 5
	//         Chú ý: Go KHÔNG tự ép kiểu. Hãy xem điều gì xảy ra nếu bạn bỏ float64(...).

	// TODO 4: dùng hằng số (basics/16). In lần lượt:
	//           needInt(Small)    → 21
	//           needFloat(Small)  → 0.2
	//           needFloat(Big)    → 1.2676506002282295e+29

	// TODO 5 (thử phá): thêm fmt.Println(needInt(Big)), đọc lỗi (tràn int), rồi xóa dòng đó.
	//         Vì sao needFloat(Big) được mà needInt(Big) thì không?
}
