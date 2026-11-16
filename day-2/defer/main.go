// Bài 3: defer và thứ tự LIFO
// Đọc trước: https://go.dev/tour/flowcontrol/12  →  flowcontrol/13
// Hướng dẫn: ../README.md
//
// Kết quả mong đợi (TODO 1 và TODO 2):
//   start
//   end loop
//   4
//   3
//   2
//   1
//   0
//
// Chạy (từ thư mục gốc repo):  go run ./day-2/defer

package main

// TODO 0: import "fmt"

func main() {
	// TODO 1: in "start".

	// TODO 2: viết vòng for i từ 0 đến 4. Trong mỗi vòng, dùng defer để in i:
	//             defer fmt.Println(i)
	//         Sau vòng lặp in "end loop".
	//         ĐOÁN trước khi chạy: 5 dòng số in ra theo thứ tự nào, và in trước hay sau "end loop"?

	// TODO 3: defer đánh giá THAM SỐ ngay lúc gặp lệnh defer, nhưng chạy HÀM khi hàm bao ngoài kết thúc.
	//         Thử đoán output của đoạn sau, rồi gõ vào chạy để kiểm tra:
	//             x := 1
	//             defer fmt.Println("deferred x =", x)
	//             x = 100
	//             fmt.Println("x =", x)

	// TODO 4 (tự hỏi): vì sao Go thiết kế defer chạy theo thứ tự ngược (LIFO)?
	//         Gợi ý: nghĩ về việc mở file rồi mở kết nối, và đóng chúng theo thứ tự nào là an toàn.
}
