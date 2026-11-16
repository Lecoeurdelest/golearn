// Bài 2: FizzBuzz với switch không điều kiện
// Đọc trước: https://go.dev/tour/flowcontrol/9  →  flowcontrol/11
// Hướng dẫn: ../README.md
//
// Kết quả mong đợi (1 đến 15):
//   1
//   2
//   Fizz
//   4
//   Buzz
//   Fizz
//   7
//   8
//   Fizz
//   Buzz
//   11
//   Fizz
//   13
//   14
//   FizzBuzz
//
// Chạy (từ thư mục gốc repo):  go run ./day-2/fizzbuzz

package main

// TODO 0: import "fmt"

// TODO 1: trong main, dùng vòng for từ 1 đến 15 và in theo luật:
//           - chia hết cho 15  → in "FizzBuzz"
//           - chia hết cho 3   → in "Fizz"
//           - chia hết cho 5   → in "Buzz"
//           - còn lại          → in chính số đó
//         Dùng switch KHÔNG có điều kiện (flowcontrol/11):
//             switch {
//             case i%15 == 0:
//                 ...
//             }

// TODO 2: đổi thứ tự các case (đặt case i%3 == 0 lên trước i%15 == 0).
//         Chạy lại và xem số 15 in ra gì. Vì sao? Nhớ rằng switch dừng ở case đúng đầu tiên,
//         không cần break (flowcontrol/9). Sau đó hoàn tác.

// TODO 3 (nâng cấp): dùng if với câu lệnh khởi tạo (flowcontrol/6), ví dụ
//             if r := i % 3; r == 0 { ... }
//         để viết lại một nhánh. Biến r chỉ sống trong phạm vi if/else đó.

func main() {
}
